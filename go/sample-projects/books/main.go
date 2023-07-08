package main

import (
	"Go-Sample-Projects/books/model"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {

	book := &model.Book{
		Id:    1,
		Title: "Harry Potter",
		Author: []*model.Author{
			{Id: 1, Name: "JK. Rowling"},
		},
		Category: model.Category_Novel,
	}

	// encoding data and exporting to a file

	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("book.protobuf", data, 0600)

	// decoding data back from file
	data, err = ioutil.ReadFile("book.protobuf")

	if err != nil {
		log.Fatal(err)
	}

	bookFromFile := model.Book{}
	err = proto.Unmarshal(data, &bookFromFile)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("book from protobuf file %+v\n", bookFromFile)
	}

}
