package sample

import (
	"log"
	"os"
	"text/template"
)

type Member struct {
	Name    string
	Comment string
}

func Sample4() {
	list := []Member{
		{
			Name:    "Ken",
			Comment: "hello! :)",
		},
		{
			Name:    "Nancy",
			Comment: "",
		},
	}

	t := template.Must(template.ParseFiles(`template/message_2.tmpl`))
	if err := t.Execute(os.Stdout, list); err != nil {
		log.Fatal(err)
	}
}
