package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the an example of interpreter Based on the book 'Writing An Interpreter In Go' by Thorsten Ball!\n",
		user.Username)
}
