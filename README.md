# rss-merger

Simple tool to merge several RSS feeds into one.

## Why?

Because building an app for a few hours to save you a few minutes is the way to go, isn't it? :)

## Configuration

The application is purely configured through environment variables since I'm deploying this inside a container.

### Things you might want to change

* **RSS_MERGER_USERNAME** - HTTP Auth Username, recommended.
* **RSS_MERGER_PASSWORD** - HTTP Auth Password, recommended.
* **RSS_MERGER_SITELINK** - According to the RSS Spec having a channel url set is required, this var helps you with that.

### Other configurations

* **RSS_MERGER_STORAGEDIR** - Location for the storage files (yes this app just stores plain JSON as everything else
seems to be overkill). Default is [your configuration directory](https://golang.org/pkg/os/#UserConfigDir).
* **RSS_MERGER_CACHE_DURATION_IN_MINUTES** - Cache duration, default is 15 minutes.
* **RSS_MERGER_PORT** - Well, the port. Default is 8081.

## Tech

* The backend is built in Go mostly powered by Echo.
* The frontend is built in ClojureScript with Reagent (React).

## License

MIT