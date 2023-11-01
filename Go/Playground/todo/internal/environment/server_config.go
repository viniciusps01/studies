package environment

import "fmt"

type ServerConfig struct {
	Port uint16
	Host string
}

func (s ServerConfig) Address() string {
	address := fmt.Sprintf("%s:%d", s.Host, s.Port)
	return address
}
