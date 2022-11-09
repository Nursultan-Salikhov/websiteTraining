package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint8
	Money                 int16
	AvgGrades, HappyLevel float64
	Hobbies               []string
}

func init() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/contacts", contactsPage)
	http.ListenAndServe(":8080", nil)
}

func main() {

}

func homePage(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 26, 200, 0.9, 0.25, []string{"Football", "Reading", "TV"}}
	tmpl, _ := template.ParseFiles("templates/homePage.html")
	tmpl.Execute(w, bob)
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	bob := User{}
	bob.Name = "Alex"
	fmt.Fprintf(w, bob.getAllInfo())
	//tmpl, _ := template.ParseFiles("templates/contactsPage.html")
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is %s, \nAge - %d, Money -%d, AvdGraduate - %f, Happy level -%f", u.Name, u.Age, u.Money, u.AvgGrades, u.HappyLevel)
}
