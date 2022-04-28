package main

import (
	"fmt"
	"gitlab-variables/app"
)

func init() {
	app.ReadConfig("gitlab_one")
}

func main() {
	fmt.Println("Hello")
}
