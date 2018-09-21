package main

import (
        "html/template"
        "log"
        "os"
        "fmt"
)

type Person struct {
        Name string
        Age  int
}

func OddOrEven(s string) string {

        if len(s)%2 == 0 {
                return "even"
        } else {
                return "odd"
        }

}

func getFormatString() string {
        placeHolderFormat := "{{range .}}%s\n{{end}}"
        defaultFormatString := "{{.Name}} {{.Age}} {{ OddOrEven .Name}}"
        if len(os.Args) == 2 {
                return fmt.Sprintf(placeHolderFormat, os.Args[1])
        } else {
                return fmt.Sprintf(placeHolderFormat, defaultFormatString)
        }

}

func main() {

        var names = []Person{
                Person{Name: "Tabby", Age: 21},
                Person{Name: "Jill", Age: 19},
        }

        funcMap := template.FuncMap{
                "OddOrEven": OddOrEven,
        }

        formatString := getFormatString()

        tmpl := template.New("test").Funcs(funcMap)
        tmpl, err := tmpl.Parse(formatString)
        if err != nil {
                log.Fatal("Error Parsing template: ", err)
                return
        }
        err1 := tmpl.Execute(os.Stdout, names)
        if err1 != nil {
                log.Fatal("Error executing template: ", err1)

     