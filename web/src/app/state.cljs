(ns app.state
  (:require [reagent.core :as r]
            [ajax.core :refer [GET]]))

(def active-page (r/atom {:active {:handler :home}}))

(def add-feed-input-default {:title "" :links []})
(def add-feed-input (r/atom add-feed-input-default))

(def feeds (r/atom {}))

(defn update-feeds! []
  (GET "/feeds" {:response-format :json
                 :keywords? true
                 :handler (fn [res] (reset! feeds res))}))