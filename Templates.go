package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Blog type
type Blog struct {
	Title  string
	Author string
	Header string
}

// Post type
type Post struct {
	Title       string
	Author      string
	Body        string
	PublishDate string
}

// BlogViewModel type
type BlogViewModel struct {
	Blog  Blog
	Posts []Post
}

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func loadPosts() []Post {
	bytes, _ := ioutil.ReadFile("posts.json")
	var posts []Post
	json.Unmarshal(bytes, &posts)
	return posts
}

func handler(w http.ResponseWriter, r *http.Request) {
	blog := Blog{Title: "My Blog", Author: "Zoerab Tchahkiev", Header: "Welcome to my blog"}
	posts := loadPosts()
	viewModel := BlogViewModel{Blog: blog, Posts: posts}
	t, _ := template.ParseFiles("blog.html")
	t.Execute(w, viewModel)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Running server on: http://127.0.0.1:9000 👽")
	http.ListenAndServe(":9000", nil)
}
