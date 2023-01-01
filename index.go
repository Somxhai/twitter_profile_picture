package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)


func getProfilePicture(username string) (*twitter.User, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("CONSUMER_KEY"),
		ClientSecret: os.Getenv("CONSUMER_SECRET"),
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	httpClient := config.Client(oauth2.NoContext)
	// Twitter client
	client := twitter.NewClient(httpClient)
	userShowParams := &twitter.UserShowParams{ScreenName: username}
	user, _, err := client.Users.Show(userShowParams)
	return user, err
}

func getProfilePictueByUserName(context *gin.Context) {

	username := context.Param("username")
	result, err := getProfilePicture(username)
	if err != nil {
		return
	}
	context.IndentedJSON(http.StatusOK, result)
}

func main() {
	router := gin.Default()
	router.GET("/:username", getProfilePictueByUserName)
	router.Run(":8080")
}