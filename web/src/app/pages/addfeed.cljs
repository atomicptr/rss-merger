(ns app.pages.addfeed
  (:require [app.state :as state]
            [reagent.core :as r]
            [ajax.core :refer [POST]]
            [app.router :as router]))

(defn feed-input-to-json []
  (clj->js @app.state/add-feed-input))

(defn on-add-feed-submit []
  (POST "/feeds" {:body    (js/JSON.stringify (feed-input-to-json))
                  :handler #(router/redirect! :home)}))

(defn link-input [num]
  [:div.field.is-horizontal
   [:div.field-label.is-normal
    [:label.label (str "Link #" (+ num 1))]]
   [:div.field-body
    [:div.field
     [:p.control
      [:input.input {:type          "text"
                     :name          "link[]"
                     :placeholder   "Link"
                     :default-value (get-in @state/add-feed-input [:links num] "")
                     :on-change     #(swap! state/add-feed-input assoc-in [:links num] (-> % .-target .-value))}]]]]])

(defn add-feed []
  [:div
   [:div.field.is-horizontal
    [:div.field-label.is-normal
     [:label.label "Title"]]
    [:div.field-body
     [:div.field
      [:p.control
       [:input.input {:type          "text"
                      :name          "title"
                      :placeholder   "Title"
                      :default-value (:title @state/add-feed-input)
                      :on-change     #(swap! state/add-feed-input assoc :title (-> % .-target .-value))}]]]]]
   (map #(link-input %) (range (+ 1 (count (:links @state/add-feed-input)))))
   [:div.field.is-grouped.is-grouped-right
    [:p.control
     [:button.button.is-info {:onClick on-add-feed-submit} "Submit"]]]])