package server

type IOCTester struct {
	Name string `val:"我是IOC测试"`
	Age  int    `val:"${props.age}"`
}

type HTTPServer struct {
	Age       int        `val:"${props.age}"`
	Name      string     `val:"我是name"`
	IOCTester *IOCTester `ioc:"auto"`
}

func (s *HTTPServer) Startup() {

}
