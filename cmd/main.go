package main

import (

	"fmt"
	"log"
	"os"

  	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/mmcdole/gofeed"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/JoeyDrew/rss-snail/models"

)



// refactor into models.go
type rssFeed struct {
	feed *gofeed.Feed
	To string
	From string

}

func main() {

	models.DB, err = sql.Open("sqlite3", "./testdb.db")
	if err != nil {
        log.Fatal(err)
    }

	// initialize our 3 tables, I don't think this will work but it's a nice sentiment
	statement, _ := models.DB.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT)")

	/* CREATE TABLE IF NOT EXISTS userfeeds (userid INTEGER PRIMARY KEY, feedid INTEGER PRIMARY KEY) ;
	CREATE TABLE IF NOT EXISTS feeds (id INTEGER PRIMARY KEY, Url TEXT)") */

	statement.Exec()

	// this will get refactored
	/* var feed = NewFeed()
	feed.To = os.Getenv("SENDGRID_TO")
	feed.From = os.Getenv("SENDGRID_FROM")

	sendMail(feed) */

}

func NewFeed() *rssFeed {

	fp := gofeed.NewParser()
	newFeed, _ := fp.ParseURL("https://sanantonioreport.org/feed/")

	return &rssFeed{
		feed: newFeed,
	}

}

func sendMail(rf *rssFeed) {

	from := mail.NewEmail("RSS Snail", rf.From)
	subject := "RSS Digest"
	to := mail.NewEmail("", rf.To)
	plainTextContent := rf.feed.Title
	htmlContent := fmt.Sprintf("<strong>%s</strong>", rf.feed.Title)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	
}