package main

func main() {

	const (
		CONN_HOST   = "localhost"
		CONN_DOMAIN = "localhost"
		CONN_PORT   = "8000"
		SERVER_ROOT = "testdata"
	)
	sv := NewGopherServer(CONN_PORT, CONN_DOMAIN, CONN_HOST, SERVER_ROOT)
	sv.Run()
}
