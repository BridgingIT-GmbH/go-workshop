package repository

import (
	"github.com/go-workshop/moviedb/model"
)

func testData() map[string]model.Movie {
	guardiansOfTheGalaxy := model.Movie{
		Id:                "1",
		Title:             "Guardians of the Galaxy",
		Description:       "A group of intergalactic criminals must pull together to stop a fanatical warrior with plans to purge the universe.",
		Actors:            []string{"Chris Pratt", "Zoe Saldaña", "Dave Bautista"},
		DurationInMinutes: 122,
	}
	starWars := model.Movie{
		Id:                "2",
		Title:             "Star Wars: A New Hope",
		Description:       "Luke Skywalker joins forces with a Jedi Knight, a cocky pilot, a Wookiee and two droids to save the galaxy from the Empire's world-destroying battle station.",
		Actors:            []string{"Mark Hamill", "Harrison Ford", "Carrie Fisher"},
		DurationInMinutes: 121,
	}
	galaxyQuest := model.Movie{
		Id:                "3",
		Title:             "Galaxy Quest",
		Description:       "The alumni cast of a space opera television series have to play their roles as the real thing when an alien race needs their help.",
		Actors:            []string{"Tim Allen", "Sigourney Weaver", "Alan Rickman"},
		DurationInMinutes: 102,
	}
	return map[string]model.Movie{
		guardiansOfTheGalaxy.Id: guardiansOfTheGalaxy,
		starWars.Id:             starWars,
		galaxyQuest.Id:          galaxyQuest,
	}
}
