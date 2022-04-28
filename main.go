package main

import (
	"fmt"
	"gitlab-variables/src/app"
)

func init() {
	app.ReadConfig("gitlab_one")
}

func main() {
	fmt.Println("Hello")
}
