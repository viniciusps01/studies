package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	Server ServerConfig
	DBpath string
}

func Load() *Environment {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to initialize environment:", err)
	}

	serverConfig, err := loadServerConfig()

	if err != nil {
		log.Fatal("failed to load server config:", err)
	}

	dbPath := os.Getenv(dbPathKey)

	env := &Environment{
		Server: *serverConfig,
		DBpath: dbPath,
	}

	return env
}

func loadServerConfig() (*ServerConfig, error) {
	port, err := strconv.ParseUint(os.Getenv(portKey), 10, 16)

	if err != nil {
		return nil, err
	}

	host := os.Getenv(hostKey)

	config := &ServerConfig{
		Port: uint16(port),
		Host: host,
	}

	return config, nil
}
