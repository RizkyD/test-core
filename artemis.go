package artemis

import (
	"fmt"
	"os"
)

func serve() {
	argument := "dev"
	if len(os.Args) != 1 {
		argument = os.Args[1]
	}
	start(argument)
	SetServerAttribute(os.Getenv("db_link"), os.Getenv("db_path"))
	fmt.Println("success")
}
