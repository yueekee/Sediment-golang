package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC")
	executeTemplate("Dot is: {{.}}!\n",  123.4)
	/*
	Dot is: ABC!
	Dot is: 123.4!
	*/

	executeTemplate("Start {{if .}}Dot is true!{{end}} finish\n",  true)
	executeTemplate("Start {{if .}}Dot is true!{{end}} finish\n",  false)
	/*
	Start Dot is true! finish
	Start  finish
	*/

	//--------------------range

	// 在{{range}}action与其对应的{{end}}标记之间的模版部分，将对数组、切片、映射、channel中的每个值进行重复
	templateText := "Before loop: {{.}}\n{{range .}}In loop: {{.}}\n{{end}}After loop: {{.}}\n"
	executeTemplate(templateText, []string{"do", "re", "mi"})
	/*
	Before loop: [do re mi]
	In loop: do
	In loop: re
	In loop: mi
	After loop: [do re mi]
	*/

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	executeTemplate(templateText, []float64{1.25, 2.34, 10.2})
	/*
	Prices:
	$1.25
	$2.34
	$10.2
	*/
	executeTemplate(templateText, nil)	// 如果提供给{{range}}action的值为空或nil，则循环不会运行。
	/*Prices:*/

	//---------------------struct
	type Part struct {
		Name string
		Count int
	}
	templateText = "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "eric", Count: 5})
	executeTemplate(templateText, Part{Name: "Lily", Count: 2})
	/*
	Name: eric
	Count: 5
	Name: Lily
	Count: 2
	*/

	type Subscriber struct {
		Name string
		Rate float64
		Active bool
	}
	templateText = "Name: {{.Name}}\n{{if .Active}}Rate: ${{.Rate}}\n{{end}}"
	subscriber := Subscriber{Name: "pan", Rate: 4.99, Active: true}
	executeTemplate(templateText, subscriber)
	subscriber = Subscriber{Name: "apple", Rate: 5.99, Active: false}
	executeTemplate(templateText, subscriber)
	/*
	Name: pan
	Rate: $4.99
	Name: apple
	*/
	// 这里对于Active字段不为true的情况，没有了Rate字段
}

