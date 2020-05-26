(ns app.main
    (:require [reagent.core :as r]
              [reagent.dom :as rdom]))

(defn test-component []
    [:div "Testassss"])

(defn app-container []
    [:div#root
        [test-component]])

(defn bootstrap-react []
    (rdom/render [app-container] (js/document.getElementById "app")))

(defn reload! []
  (println "Code updated.")
  (bootstrap-react))

(defn main! []
    (println "yolo")
    (bootstrap-react))