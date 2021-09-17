package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//var name string -
	name, Academy := "Rezaul Karim", "DevOps"

	show := fmt.Sprintf("%s is a student of %s Master-Academy. \n", name, Academy)

	io.WriteString(os.Stdout, show)
}
