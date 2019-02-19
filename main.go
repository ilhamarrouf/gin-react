package main

import (
	"net/http"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Joke struct {
	ID    int `json:"id" binding:"required"`
	Likes int `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
	Joke{8, 0, "Dad, did you get a haircut? No I got them all cut."},
	Joke{9, 0, "What do you call a Mexican who has lost his car? Carlos."},
	Joke{10, 0, "Dad, can you put my shoes on? No, I don't think they'll fit me."},
	Joke{11, 0, "Why did the scarecrow win an award? Because he was outstanding in his field."},
	Joke{12, 0, "Why don't skeletons ever go trick or treating? Because they have no body to go with."},
}

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
		api.GET("/jokes", JokeHandler)
		api.POST("/jokes/like/:jokeID", LikeJoke)
	}

	router.Run(":3000")
}

func JokeHandler(c *gin.Context) {
	c.Header("Content-type", "application/json")
	c.JSON(http.StatusOK, gin.H {
		"message": "JokeHandler not implemented yet",
	})
}

func LikeJoke(c *gin.Context) {
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes = jokes[i].Likes + 1
			}
		}
		c.JSON(http.StatusOK, &jokes)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}