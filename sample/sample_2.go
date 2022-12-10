package sample

import (
	"log"
	"os"
	"text/template"
)

type Language struct {
	Name string
	URL  string
}

func Sample2() {
	data := Language{
		Name: "Go",
		URL:  "https://golang.org/",
	}

	tmpl := "今回紹介する言語は、{{.Name}}です！\n詳しくは、{{.URL}} を見てください。\n"

	t, err := template.New("sample").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	if err = t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
