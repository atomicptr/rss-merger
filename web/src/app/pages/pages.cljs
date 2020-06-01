(ns app.pages.pages
  (:require [app.pages.home :refer [home]]
            [app.pages.feed :refer [feed]]
            [app.pages.addfeed :refer [add-feed]]))

(defn pages [active-page]
  (case active-page
    :home [home]
    :feed [feed]
    :add-feed [add-feed]
    [home]))