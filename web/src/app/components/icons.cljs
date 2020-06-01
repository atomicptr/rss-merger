(ns app.components.icons
  (:require ["@fortawesome/react-fontawesome" :refer (FontAwesomeIcon)]
            ["@fortawesome/free-solid-svg-icons" :refer (faRss faEdit faTrash faPlus faShareSquare)]))

(defn render-font-aweseome [icon]
  [FontAwesomeIcon #js{:className "icon" :icon icon}])

(defn rss [] (render-font-aweseome faRss))

(defn edit [] (render-font-aweseome faEdit))

(defn delete [] (render-font-aweseome faTrash))

(defn open [] (render-font-aweseome faShareSquare))

(defn add [] (render-font-aweseome faPlus))