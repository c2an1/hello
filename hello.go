package main

import(
	"fmt"
	"github.com/longyingsuifeng/greetings"
)

func main(){
	message := greetings.Hello("lysf")
	fmt.Println(message)
}