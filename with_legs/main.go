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
	fmt.Println("Hi, %s, this is my toy language!", user.Username)
	fmt.Println("You can now type commands")
	repl.Start(os.Stdin, os.Stdout)
}
