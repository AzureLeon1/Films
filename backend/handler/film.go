package handler

import (
    "sse.tongji.edu.cn/leon/web3/model"
    "net/http"
    "strconv"

    "github.com/labstack/echo"
    // "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Result struct {
	Films []*model.Film `bson:"films" json:"films"`
	FilmsForRank   []*FilmForRank `bson:"filmsForRank"   json:"filmsForRank"`
}

type FilmForRank struct {
	Rating    model.Rating   `bson:"rating" json:"rating"`
	Title   string   `bson:"title" json:"title"`
}

type FilmForAnalysis struct {
	Genres    []string   `bson:"genres" json:"genres"`
    Countries  []string   `bson:"countries" json:"countries"`
}

func (h *Handler) SearchFilm(c echo.Context) (err error) {

    standard := c.QueryParam("standard")
    keyWord := c.QueryParam("keyWord")

    // Retrieve posts from database
    films  := []*model.Film{}
    db := h.DB.Clone()

    if err = db.DB("webCourse").C("film").
        Find(bson.M{standard : &bson.RegEx{Pattern: keyWord, Options: "i"}}).
        All(&films); err != nil {
        return
    }
    defer db.Close()

    c.Response().Header().Set("Access-Control-Allow-Origin", "*")
    return c.JSON(http.StatusOK, films)
}

func (h *Handler) FetchFilm(c echo.Context) (err error) {
    page, _ := strconv.Atoi(c.QueryParam("page"))
    limit, _ := strconv.Atoi(c.QueryParam("limit"))

    // Defaults
    if page == 0 {
        page = 1
    }
    if limit == 0 {
        limit = 10
    }

    // Retrieve posts from database
    films := []*model.Film{}
    filmsForRank  := []*FilmForRank{}
    db := h.DB.Clone()

    if err = db.DB("webCourse").C("film").
		Find(nil).
        Skip((page - 1) * limit).
        Limit(limit).
        All(&films); err != nil {
        return
    }

    if err = db.DB("webCourse").C("film").
		// 只读取 title 和 rating 字段
        Find(bson.M{}).
        Select(bson.M{"title":1, "rating":1, "_id":0}).
        Sort("-rating.average").
        Limit(40).
        All(&filmsForRank); err != nil {
        return
    }
    defer db.Close()
    result := &Result{
        Films : films,
        FilmsForRank : filmsForRank,
    }

    c.Response().Header().Set("Access-Control-Allow-Origin", "*")
    return c.JSON(http.StatusOK, result)
}

func (h *Handler) FetchStatis(c echo.Context) (err error) {

    // Retrieve posts from database
    filmsForAnalysis  := []*FilmForAnalysis{}
    db := h.DB.Clone()


    if err = db.DB("webCourse").C("film").
		// 只读取 title 和 rating 字段
        Find(bson.M{}).
        Select(bson.M{"genres":1, "countries":1, "_id":0}).
        All(&filmsForAnalysis); err != nil {
        return
    }
    defer db.Close()

    c.Response().Header().Set("Access-Control-Allow-Origin", "*")
    return c.JSON(http.StatusOK, filmsForAnalysis)
}

func (h *Handler) FetchFilmByIndex(c echo.Context) (err error) {
    index, _ := strconv.Atoi(c.Param("index"))
    var limit = 10

    if index == 0 {
        index = 1
    }


    // Retrieve posts from database
    films := []*model.Film{}
    db := h.DB.Clone()
    if err = db.DB("webCourse").C("film").
		// Find(bson.M{"to": userID}).
		Find(nil).
        Skip((index - 1) * limit).
        Limit(limit).
        All(&films); err != nil {
        return
    }
    defer db.Close()

    c.Response().Header().Set("Access-Control-Allow-Origin", "*")
    return c.JSON(http.StatusOK, films)
}

