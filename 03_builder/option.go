package main

/**
Option 模式（或称为可选参数模式）是一种创建对象或配置对象的方式，它可以避免构造函数带过多的参数。
Option 模式常用于构建具有多个可选参数的对象，通过函数式编程来提供灵活的参数配置。
*/

/**
使用场景：假设我们有一个复杂的对象 Server，它有一些必需的参数和一些可选参数。为了避免构造函数过于庞大，可以使用 Option 模式来提供可选参数的设置方法。
*/

// Server 代表要创建多个对象，有多个可选参数
type Server struct {
	Host string
	Port int
	SSL  bool
}

// Option 是一个函数类型，用于修改 Server 对象
type Option func(*Server)

// NewServer 创建一个带有默认值的 Server，并应用传入的 Option 函数
func NewServer(host string, options ...Option) *Server {
	srv := &Server{
		Host: host,
		Port: 80,
		SSL:  false,
	}

	for _, option := range options {
		option(srv)
	}
	return srv
}

// WithPort 设置服务器端口

func WithPort(port int) Option {
	return func(srv *Server) {
		srv.Port = port
	}
}

func WithSSL(ssl bool) Option {
	return func(srv *Server) {
		srv.SSL = ssl
	}
}
