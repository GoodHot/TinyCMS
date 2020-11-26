package server

import "fmt"

type HTTPServer struct {
	Age  int    `val:"${props.age}"`
	Name string `val:"我是name"`
}

func (s *HTTPServer) Startup() {
	fmt.Println(s.Age)
	fmt.Println(s.Name)
}
