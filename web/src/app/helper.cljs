(ns app.helper
  (:require
    [cljs.core.async :refer [go]]
    [cljs.core.async.interop :refer-macros [<p!]]))

(defn fetch-json [url func]
  (go (let [res (<p! (js/fetch url))]
        (func (<p! (.json res))))))