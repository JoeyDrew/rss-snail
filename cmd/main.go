package main

import (

	"fmt"
	"log"
	"os"

  	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/mmcdole/gofeed"
    //"github.com/mmcdole/gofeed/rss"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"

)

type rssFeed struct {
	feed *gofeed.Feed
	To string
	From string

}

func main() {

	database, _ := sql.Open("sqlite3", "./testdb.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
    statement.Exec()

	var feed = NewFeed()
	feed.To = os.Getenv("SENDGRID_TO")
	feed.From = os.Getenv("SENDGRID_FROM")

	//sendMail(feed)

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