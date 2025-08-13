package variadic_func

import "fmt"

// 变长参数函数示例三：基于变长参数实现WithOption 功能选项模式

type Server struct {
	host        string
	port        string
	maxTimeout  int
	connTimeout int
}

// 传统的构造方法
func NewServer(host, port string, maxTimeout, connTimeout int) Server {
	// 每一个参数都要传
	// 新增参数需要修改NewServer API，不兼容
	// 参数很多情况下 很难维护
	return Server{
		host:        host,
		port:        port,
		maxTimeout:  maxTimeout,
		connTimeout: connTimeout,
	}
}

func NewServerV2(opts ...Option) Server {
	server := Server{}
	for _, opt := range opts {
		opt(&server)
	}
	return server
}

type Option func(*Server)

func WithHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithPort(port string) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithMaxTimeout(maxTimeout int) Option {
	return func(s *Server) {
		s.maxTimeout = maxTimeout
	}
}

func WithConnTimeout(connTimeout int) Option {
	return func(s *Server) {
		s.connTimeout = connTimeout
	}
}

func main() {
	// 使用变长参数函数实现构造方法
	server := NewServer("localhost", "8080", 30, 10)
	fmt.Println(server)

	// 使用变长参数函数实现构造方法
	server = NewServerV2(
		WithHost("127.0.0.1"),
		WithPort("8081"),
		WithMaxTimeout(60),
		WithConnTimeout(20))
	
}
