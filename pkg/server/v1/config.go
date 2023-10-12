package server

import (
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// Config to define config
type Config struct {
	Server struct {
		Host         string        `toml:"host"`
		Port         string        `toml:"port"`
		PromPort     string        `toml:"prom_port"`
		ReadTimeout  time.Duration `toml:"read_timeout"`
		WriteTimeout time.Duration `toml:"write_timeout"`
		CertFile     string        `toml:"cert_file"`
		KeyFile      string        `toml:"key_file"`
	} `toml:"server"`
	Account struct {
		Connection string `toml:"conn"`
	} `toml:"account"`
	SQL struct {
		DownloaderDriver     string `toml:"downloader_driver"`
		DownloaderConnection string `toml:"downloader_conn"`

		Driver     string `toml:"driver"`
		Connection string `toml:"conn"`
	} `toml:"sql"`
	CACHE struct {
		Driver     string `toml:"driver"`
		Connection string `toml:"conn"`
	} `toml:"cache"`
	SFTP struct {
		Remote string `toml:"remote"`
		User   string `toml:"user"`
		Pass   string `toml:"pass"`
	}
	Log struct {
		Level string `toml:"level"`
	} `toml:"log"`
}

// LoadConfig to load config
func LoadConfig(file string) *Config {
	c := Config{}
	if _, err := toml.DecodeFile(file, &c); err != nil {
		log.Fatal(errors.Wrap(err, "fail to load config file"))
	}

	return &c
}
