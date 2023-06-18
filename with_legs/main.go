package main

import (
	"fmt"
	"os"
	"os/user"
	"with_legs/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi, %s, this is my toy language!\n", user.Username)
	fmt.Printf("You can now type commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
