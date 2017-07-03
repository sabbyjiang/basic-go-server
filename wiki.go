package main

import (
	"fmt"
	"io/ioutil"
)

/*
	Describes how pages will be stores in MEMORY
	Not how they will be in storages
*/
type Page struct {
	Title string
	Body  []byte // io expects byte not strings
}

/*
	Method named save that takes a page and saves it
	Will return an error if anything goes wrong but nil if everything goes right
*/
func (p *Page) save() error {
	filename := p.Title + ".txt"
	/*
		The 0600 octal integer lets the program know that the file
		written should only be given read-write permissions for the
		current user only
	*/
	return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
	Loads the page file from the title param
	Returns a pointer to the page literal constructed from the file
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
