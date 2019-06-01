package main

import (
    "sse.tongji.edu.cn/leon/web3/handler"

    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/labstack/gommon/log"
    "gopkg.in/mgo.v2"
)

func main() {
    e := echo.New()
    e.Logger.SetLevel(log.ERROR)
    e.Use(middleware.Logger())

    // Database connection
    db, err := mgo.Dial("127.0.0.1")
    if err != nil {
        e.Logger.Fatal(err)
    }

    // Create indices
    if err = db.Copy().DB("webCourse").C("film").EnsureIndex(mgo.Index{
        Key:    []string{"imdb"},
        Unique: true,
    }); err != nil {
        log.Fatal(err)
    }

    // Initialize handler
    h := &handler.Handler{DB: db}

    // Routes
    e.GET("/film", h.FetchFilm)
    e.GET("/statistics", h.FetchStatis)
    e.GET("/film/:index", h.FetchFilmByIndex)
    e.GET("/search", h.SearchFilm)

    // Start server
    e.Logger.Fatal(e.Start(":2333"))
}
