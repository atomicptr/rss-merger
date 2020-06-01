package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/atomicptr/rss-merger/pkg/feed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func runApi(username, password string, port int) error {
	e := echo.New()
	e.HideBanner = true

	e.Use(simpleLoggerMiddleware)
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}))

	if username != "" && password != "" {
		e.Use(middleware.BasicAuth(func(user, pass string, context echo.Context) (bool, error) {
			return username == user && pass == password, nil
		}))
	}

	e.HTTPErrorHandler = errorHandler

	e.GET("/", serveFrontend)
	e.GET("/feeds", getFeeds)
	e.POST("/feeds", postNewFeed)
	e.GET("/feeds/:identifier", getFeedByIdentifier)
	e.DELETE("/feeds/:identifier", deleteFeed)
	e.POST("/feeds/:identifier/add-link", postAddLink)
	e.POST("/feeds/:identifier/delete-link", postDeleteLink)
	e.GET("/rss/:identifier", getRssFeed)

	e.Static("/", "public")

	return e.Start(fmt.Sprintf(":%d", port))
}

func serveFrontend(context echo.Context) error {
	return context.File("public/index.html")
}

func errorHandler(err error, context echo.Context) {
	code := http.StatusInternalServerError
	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
	}

	// push all requests that couldn't be found to the frontend
	if code == http.StatusNotFound {
		err = serveFrontend(context)
		if err != nil {
			log.Println(err)
		}
		return
	}

	_ = context.HTML(code, fmt.Sprintf("<h2>%d - %s</h2>", code, "Error"))
}

func getFeeds(context echo.Context) error {
	return context.JSON(http.StatusOK, feedStorage)
}

func postNewFeed(context echo.Context) error {
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return err
	}

	var f feed.Feed
	err = json.Unmarshal(body, &f)
	if err != nil {
		return err
	}

	if f.Title == "" {
		return errors.New("no title found")
	}

	identifier := createIdentifierFromTitle(f.Title)
	f.Identifier = identifier
	feedStorage[identifier] = f
	err = saveStorage(storageDir)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{"identifier": identifier})
}

func getFeedByIdentifier(context echo.Context) error {
	identifier := context.Param("identifier")
	f, ok := feedStorage[identifier]
	if !ok {
		return context.JSON(http.StatusNotFound, "not found")
	}
	return context.JSON(http.StatusOK, f)
}

func deleteFeed(context echo.Context) error {
	identifier := context.Param("identifier")
	_, ok := feedStorage[identifier]
	if !ok {
		return context.JSON(http.StatusNotFound, "not found")
	}

	delete(feedStorage, identifier)
	err := saveStorage(storageDir)
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, "ok")
}

func postAddLink(context echo.Context) error {
	identifier := context.Param("identifier")
	f, ok := feedStorage[identifier]
	if !ok {
		return context.JSON(http.StatusBadRequest, "not found")
	}

	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return err
	}
	link := string(body)
	if link == "" {
		return context.JSON(http.StatusBadRequest, "no link found")
	}
	_, err = url.Parse(link)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "link was not a valid url")
	}

	f.Links = append(f.Links, link)
	feedStorage[identifier] = f
	return saveStorage(storageDir)
}

func postDeleteLink(context echo.Context) error {
	identifier := context.Param("identifier")
	f, ok := feedStorage[identifier]
	if !ok {
		return context.JSON(http.StatusBadRequest, "not found")
	}

	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return err
	}
	link := string(body)

	linkIndex := -1

	for index, l := range f.Links {
		if l == link {
			linkIndex = index
			break
		}
	}

	if linkIndex == -1 {
		return errors.New("unknown link supplied")
	}

	f.Links = append(f.Links[:linkIndex], f.Links[linkIndex+1:]...)
	feedStorage[identifier] = f
	return saveStorage(storageDir)
}

func getRssFeed(context echo.Context) error {
	identifier := context.Param("identifier")
	str, err := createMergedFeed(identifier)
	if err != nil {
		return context.String(http.StatusInternalServerError, err.Error())
	}
	return context.XML(http.StatusOK, str)
}

func simpleLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		req := context.Request()
		res := context.Response()

		message := ""

		if err := next(context); err != nil {
			context.Error(err)
			message = err.Error()
		}

		log.Printf("%d - %s %s %s\n", res.Status, req.Method, req.URL.String(), message)
		return nil
	}
}
