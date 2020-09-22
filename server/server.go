package server

type HTTPServer struct {
	Age  int `val:"${dddd}"`
	Name string
}

func (*HTTPServer) Startup() {
}
