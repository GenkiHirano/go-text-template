package sample

import (
	"log"
	"os"
	"text/template"
)

func Sample5() {
	funcMap := template.FuncMap{
		"getHogeString": getHogeString,
		"sum":           func(x, y int) int { return x + y },
	}

	t, err := template.New("message_3.tmpl").Funcs(funcMap).ParseFiles("template/message_3.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}

func getHogeString() string {
	return "Hoge!"
}
