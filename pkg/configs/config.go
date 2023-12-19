package configs

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host    string `yaml:"host"`
		Port    string `yaml:"port"`
		Timeout struct {
			Server time.Duration `yaml:"server"`
			Read   time.Duration `yaml:"read"`
			Write  time.Duration `yaml:"write"`
			Idle   time.Duration `yaml:"idle"`
		}
	} `yaml:"server"`

	Postgres struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		DbName string `yaml:"dbname"`
		DbUser string `yaml:"dbuser"`
		DbPass string `yaml:"dbpass"`
	} `yaml:"postgres"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ParseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}

	if s.IsDir() {
		return fmt.Errorf("invalid file format")
	}

	return nil
}

func (c *Config) Run(r *mux.Router) {
	log.Println("Start to run server")
	var runChan = make(chan os.Signal)

	ctx, cancel := context.WithTimeout(context.Background(), c.Server.Timeout.Server)
	defer cancel()

	server := &http.Server{
		Addr:         c.Server.Host + ":" + c.Server.Port,
		Handler:      r,
		ReadTimeout:  c.Server.Timeout.Read,
		WriteTimeout: c.Server.Timeout.Write,
		IdleTimeout:  c.Server.Timeout.Idle,
	}

	signal.Notify(runChan, os.Interrupt, syscall.SIGTSTP)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				return
			}
		}
	}()

	interrupt := <-runChan
	log.Printf("Server is shutting down due to %+v\n", interrupt)
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server was unable to gracefull shutdown due to err : %+v\n", err)
	}
}
