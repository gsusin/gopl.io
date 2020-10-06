// Exercício 8.1

package main

import (
	"fmt"
	"net"
	"os"
	"text/template"
	"time"
)

const tmpl = `Times from all servers:
{{range $}}
{{getTime .}}
{{end}}`

var clockwall = template.Must(template.New("clockwall").
	Funcs(template.FuncMap{"getTime": getTime}).
	Parse(tmpl))

func getTime(url string) string {
	conn, err := net.Dial("tcp", url)
	var s string
	if err != nil {
		return fmt.Sprintln(err)
	}
	defer conn.Close()
	_, err = fmt.Fscanf(conn, "%s", &s)
	if err != nil {
		return fmt.Sprintln(err)
	}
	return fmt.Sprintf("%s", s)
}

func main() {
	for {
		if err := clockwall.Execute(os.Stdout, os.Args[1:]); err != nil {
			println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
