package config

type AdminConfig struct {
	Http HttpServer
}

type HttpServer struct {
}

func (receiver HttpServer) GetPort() string {
	return "8080"
}
