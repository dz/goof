package main

import (
	"fmt"
	"io/ioutil"
)

type Post struct {
	Title string
	Body []byte
}

func (p *Post) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func load(title string) (*Post, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Post{Title: title, Body: body}, nil
}

func main() {
	// create and save
	(&Post{Title: "Hello", Body: []byte("Hello World")}).save()

	// load
	p, _ := load("Hello")

	fmt.Println(string(p.Title))
	fmt.Println(string(p.Body))
}
