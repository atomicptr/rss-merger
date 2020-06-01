(ns app.components.header
  (:require [app.router :as router]
            [app.components.icons :as icons]))

(defn header []
  [:nav.navbar.is-link {:role "navigation"}
   [:div.navbar-brand
    [:a.navbar-item {:href (router/url-for :home)}
     [icons/rss]
     [:span "RSS Merger"]]]
   [:div.navbar-end
    [:div.navbar-item
     [:div.field
      [:a.button.is-light {:href (router/url-for :add-feed)}
       [:span.icon (icons/add)]
       [:span "Add feed"]]]]]])