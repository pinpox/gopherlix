package main

func main() {
	sv := NewGopherServer(
		"8000",
		"localhost",
		"localhost",
		"testdata",
	)
	sv.Run()
}
