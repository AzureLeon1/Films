package model

import (

)


type Film struct {
	Id string `bson:"_id" json:"_id"`
	Directors []Person `bson:"directors" json:"directors"`
	Genres    []string   `bson:"genres" json:"genres"`
	SeasonCount string   `bson:"season_count" json:"season_count"`
	Rating    Rating   `bson:"rating" json:"rating"`
	Pubdate   []string   `bson:"pubdate" json:"pubdate"`
	Countries  []string   `bson:"countries" json:"countries"`
	LensId  int   `bson:"lens_id" json:"lens_id"`
	Casts   []Person   `bson:"casts" json:"casts"`
	Title   string   `bson:"title" json:"title"`
	Site    string   `bson:"site" json:"site"`
	Poster  string   `bson:"poster" json:"poster"`
	Summary  string   `bson:"summary" json:"summary"`
	Languages []string   `bson:"languages" json:"languages"`
	Episodes string   `bson:"episodes" json:"episodes"`
	Writers []Person   `bson:"writers" json:"writers"`
	Imdb  string   `bson:"imdb" json:"imdb"`
	Year string   `bson:"year" json:"year"`
	Duration string   `bson:"duration" json:"duration"`
	DoubanSite string    `bson:"douban_site" json:"douban_site"`
	Aka []string   `bson:"aka" json:"aka"`
}

type Person struct {
	Name string `bson:"name" json:"name"`
	Id   string `bson:"id"   json:"id"`
}

type Rating struct {
	Average string `bson:"average" json:"average"`
	RatingPeople   string `bson:"rating_people"   json:"rating_people"`
	Stars  []string   `bson:"stars" json:"stars"`
}