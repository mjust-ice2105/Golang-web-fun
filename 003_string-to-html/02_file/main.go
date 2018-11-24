package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Justice"

	tpl := fmt.Sprint(`
<! DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Hey, Whats Up</title>
</head>
<body>
<h2>` + name + `</h2>
</body>
</html>
`)

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(tpl))

}
