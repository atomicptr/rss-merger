(ns app.core
  (:require [reagent.core :as r]
            [reagent.dom :as rdom]
            [app.state :as state]
            [app.router :as router]
            [app.components.header :refer [header]]
            [app.pages.pages :refer [pages]]))

(defn rss-merger-app []
  [:div
   [header]
   [:div.container.body
    [pages (get-in @state/active-page [:active :handler])]]])

(defn bootstrap-react []
  (rdom/render [rss-merger-app] (js/document.getElementById "app")))

(defn init []
  ; update feed once on start, this could probably be done better but rn the
  ; app is reloading the UI at some points anyway so probably won't matter
  (state/update-feeds!)
  (router/start!)
  (bootstrap-react))

(defn main! [] (init))

(defn reload! [] (init))
