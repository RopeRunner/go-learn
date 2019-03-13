package main

import (
	"fmt"
	"io/ioutil"
)

// Page is page structer type
type Page struct {
	Title string
	Body  []byte
}

// Save is method to write daya to file system from Page
func (p *Page) Save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Function to load file from the system
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "SamplePage", Body: []byte("This is a simple page")}
	p1.Save()

	p2, _ := loadPage("SamplePage")

	fmt.Println(string(p2.Body))
}
