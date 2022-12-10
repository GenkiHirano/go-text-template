package sample

import (
	"log"
	"os"
	"text/template"
)

func Sample1() {
	tmpl := "Hello {{.}}!\n"
	// template.New(<テンプレート名>).Parse(<文字列>)
	t, err := template.New("sample").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	// Execute(io.Writer(出力先), <データ>)
	if err = t.Execute(os.Stdout, "World"); err != nil {
		log.Fatal(err)
	}
}
