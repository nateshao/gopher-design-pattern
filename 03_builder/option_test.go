package main

import "testing"

func TestServerWithOptions(t *testing.T) {
	// 默认配置
	server1 := NewServer("localhost")
	if server1.Port != 80 || server1.SSL != false {
		t.Errorf("Expected default port 80 and SSL false, got port %d and SSL %t", server1.Port, server1.SSL)
	}

	// 自定义端口和启用 SSL
	server2 := NewServer("example.com", WithPort(80), WithSSL(true))
	if server2.Port != 80 || server2.SSL != true {
		t.Errorf("Expected port 8080 and SSL true, got port %d and SSL %t", server2.Port, server2.SSL)
	}
}
