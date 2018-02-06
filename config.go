package nobreak

import (
	"errors"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Config type.
type Config struct {
	// Host target host.
	Host string `yaml:"host"`

	// Port target port.
	Port int `yaml:"port"`

	// Auto flag for automatic cache return.
	Auto bool `yaml:"auto"`

	// AdminPort port for admin configuration dashboard. WIP.
	AdminPort int `yaml:"admin_port"`

	// Timeout http client timeout in milliseconds.
	Timeout int `yaml:"timeout"`

	// Database connection uri for sqlite3. If empty in memory will be used.
	Database string `yaml:"database"`

	// TLSEnabled enable tls server.
	TLSEnabled bool `yaml:"tls_enabled"`

	// TLSCert tls cert pem file.
	TLSCert string `yaml:"tls_cert"`

	// TLSKey tls key pem file.
	TLSKey string `yaml:"tls_key"`
}

// LoadConfig load and parses yaml file.
func LoadConfig(file string) (*Config, error) {
	if file == "" {
		return nil, errors.New("empty filename")
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	if err := f.Close(); err != nil {
		return nil, err
	}
	var c *Config
	if err := yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return c, nil
}
