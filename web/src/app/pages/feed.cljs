(ns app.pages.feed
  (:require [app.helper :as helper]
            [app.router :as router]
            [reagent.core :as r]
            [ajax.core :refer [POST]]))

(defn feed-slug-valid? []
  (let [slug (:slug (router/get-route-params))]
    (contains? @app.state/feeds (keyword slug))))

(defn get-feed-from-slug []
  ((keyword (:slug (router/get-route-params))) @app.state/feeds))

(defn invalid-identifier []
  [:article.message.is-danger
   [:div.message-header
    [:p "Invalid identifier"]]
   [:div.message-body
    "The specified identifier " [:b (:slug (router/get-route-params))] " is invalid."]])

(defn delete-feed-link! [feed link]
  (let [url (str "/feeds/" (name (:identifier feed)) "/delete-link")]
    (POST url {:body    link
               :handler #(router/redirect-raw! (router/url-for :feed :slug (:identifier feed)))})))

(defn render-link-elem [feed link]
  [:article.message.is-info
   [:div.message-header
    [:p link]
    [:button.delete {:onClick #(delete-feed-link! feed link)}]]])

(defn on-submit-new-link! [input feed]
  (let [url (str "/feeds/" (name (:identifier feed)) "/add-link")]
    (POST url {:body    (:link @input)
               :handler #(router/redirect-raw! (router/url-for :feed :slug (:identifier feed)))})))

(defn render-feed-form []
  (let [feed (get-feed-from-slug) new-link-input (r/atom {})]
    [:div
     [:h2.title (:title feed)]
     (map #(render-link-elem feed %) (:links feed))
     [:div.field.has-addons
      [:div.control.w100
       [:input.input {:type        "text"
                      :placeholder "New link"
                      :name        "link"
                      :on-change   #(swap! new-link-input assoc :link (-> % .-target .-value))}]]
      [:div.control
       [:a.button.is-info {:onClick #(on-submit-new-link! new-link-input feed)} "Submit"]]]]))

(defn feed []
  (if (not (feed-slug-valid?))
    (invalid-identifier)
    (render-feed-form)))