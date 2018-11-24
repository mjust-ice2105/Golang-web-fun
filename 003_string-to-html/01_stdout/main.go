package main

import "fmt"

func main() {
	name := "Justice"

	tpl := `
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
`

	fmt.Println(tpl)
}
