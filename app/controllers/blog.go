package controllers

import (
	"github.com/robfig/revel"
	"html"
	"myapp/app/models"
	"time"
	"tux21b.org/v1/gocql"
)

type Blog struct {
	CassandraController
}

func (c Blog) Listing() revel.Result {
	var posts []*models.Post
	var id uint
	var title, tags, content string
	var date time.Time
	iter := c.Sess.Query(`SELECT id,title,tags,content,date FROM posts`).Iter()
	for iter.Scan(&id, &title, &tags, &content, &date) {
		newPost := models.Post{Id: id, Title: html.EscapeString(title), Tags: tags, Content: content, Date: date}
		posts = append(posts, &newPost)
	}
	if err := iter.Close(); err != nil {
		panic(err)
	}
	return c.Render(posts)
}

func (c Blog) Single(title string) revel.Result {
	c.Validation.Required(title).Message("Title not found.")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Blog.Listing)
	}
	// do stuff
	var id uint
	var titleOut, tags, content string
	var date time.Time
	if err := c.Sess.Query(`SELECT id,title,tags,content,date FROM posts WHERE title = ? LIMIT 1`,
		c.Params.Get("title")).Consistency(gocql.One).Scan(&id, &titleOut, &tags, &content, &date); err != nil {
		panic(err)
	}
	titleOut = html.EscapeString(titleOut)
	return c.Render(title, id, titleOut, tags, content, date)
}
