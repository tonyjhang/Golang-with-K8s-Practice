package main

import "fmt"

func main() {
	server := RestServer{ port:":8888" }
	server.run()
	fmt.Println("Server has setup.")
}