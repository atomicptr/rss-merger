(ns app.router
  (:require [bidi.bidi :as bidi]
            [pushy.core :as pushy]
            [app.state :as state]))

(def routes
  ["/" {""         :home
        "feed/add" :add-feed
        "feed/"    {[:slug] :feed}}])

(defn set-page! [match]
  (swap! state/active-page assoc :active match))

(def history
  (pushy/pushy set-page! (partial bidi/match-route routes)))

(defn start! []
  (pushy/start! history))

(def url-for (partial bidi/path-for routes))

(defn get-route-params [] (get-in @state/active-page [:active :route-params]))

(defn redirect-raw! [url]
  (set! js/window.location.href url))

(defn redirect! [location]
  (redirect-raw! (url-for location)))