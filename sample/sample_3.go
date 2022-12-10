package sample

import (
	"log"
	"os"
	"text/template"
)

func Sample3() {
	data := map[string]string{
		"Name":     "hiro",
		"Language": "Go",
	}

	t := template.Must(template.ParseFiles(`template/message_1.tmpl`))
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
