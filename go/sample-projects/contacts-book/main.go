package main

import (
	"Go-Sample-Projects/contacts-book/model"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

type Contact struct {
	FirstName string
	LastName  string
	Email     string
}

type PageData struct {
	Contacts *model.Contacts
}

var tmplHome *template.Template
var tmplAdd *template.Template

var contacts []Contact

func FetchContacts() {
	data, err := ioutil.ReadFile("contacts.protobuf")
	if err != nil {
		fmt.Println("No Contacts")
	}
	contacts := model.Contacts{}
	err = proto.Unmarshal(data, &contacts)
	if err != nil {
		log.Fatal(err)
	}
	contacts = append(contacts,contact);
	return &contacts

}

func UpdateContacts() {

}

func ContactsBookPage(w http.ResponseWriter, r *http.Request) {
	// FetchContacts()
	contacts := FetchContacts()
	fmt.Printf("Contacts: %+v\n", contacts)
	data := PageData{
		Contacts: contacts,
	}
	tmplHome.Execute(w, data)
}

func AddContactPage(w http.ResponseWriter, r *http.Request) {
	FetchContacts()
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s = %s\n", key, value)
	}
	// data := PageData{
	// 	Contacts: contacts,
	// }
	// tmplAdd.Execute(w, data)
}

func main() {

	tmplHome = template.Must(template.ParseFiles("templates/home.gohtml"))
	tmplAdd = template.Must(template.ParseFiles("templates/add.gohtml"))
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/home", ContactsBookPage)
	mux.HandleFunc("/add", AddContactPage)

	log.Fatal(http.ListenAndServe(":9000", mux))

}
