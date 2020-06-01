(ns app.pages.home
  (:require [app.router :as router]
            [app.state :as state]
            [app.components.icons :as icons]
            [ajax.core :refer [DELETE]]))

(defn open-feed-url! [feed]
  (->> (js/window.open (str "/rss/" (:identifier feed)) "_blank")
       (.focus)))

(defn delete-feed! [feed]
  (let [should-delete (js/window.confirm "Are you sure?")]
    (if (true? should-delete)
      (DELETE (str "/feeds/" (:identifier feed)) {:handler (router/redirect! :home)}))))

(defn feed-list-item [feed]
  [:div.column.is-one-third
   [:div.card
    [:div.card-content
     [:div.content
      [:h2.title (:title feed)]
      [:p (str "Contains " (count (:links feed)) " links.")]]]
    [:footer.card-footer
     [:a.card-footer-item {:onClick #(open-feed-url! feed)}
      [:span.icon (icons/open) " Open"]]
     [:a.card-footer-item {:href (router/url-for :feed :slug (:identifier feed))}
      [:span.icon (icons/edit) " Edit"]]
     [:a.card-footer-item {:onClick #(delete-feed! feed)}
      [:span.icon (icons/delete) " Delete"]]]]])

(defn home []
  [:div
   [:h2.title "Feeds"]
   (if (empty? @state/feeds)
     [:article.message.is-warning
      [:div.message-header "No feeds found"]
      [:div.message-body "No feeds were found, add some."]]
     [:div.feed-list.columns.is-multiline
      (map feed-list-item (vals @state/feeds))])])