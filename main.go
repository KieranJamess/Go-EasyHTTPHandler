package main

func main() {
	server := NewHTTPServer(":4000")
	server.Run()
}
