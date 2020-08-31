package cliconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	dirModePerm  = 755
	fileModePerm = 644
)

// Client is the configuration client for a cli
type Client struct {
	path string
}

// New creates a new cliconfig client
func New(path string) *Client {
	return &Client{
		path: path,
	}
}

// NewFromFile creates a new config from a file
func (client *Client) NewFromFile(config interface{}) error {
	content, err := ioutil.ReadFile(client.Path())

	if err != nil {
		return errors.Wrap(err, "Error reading config file")
	}

	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return errors.Wrap(err, "Error decoding config file content")
	}

	return nil
}

func homeDir() string {
	h := os.Getenv("HOME")

	if h != "" {
		return h
	}

	return os.Getenv("USERPROFILE") // windows
}

// Dir returns the default directory to the configuration directory
func (client *Client) Dir() string {
	return filepath.Join(homeDir(), client.path)
}

// Path returns the default path to the configuration file
func (client *Client) Path() string {
	return filepath.Join(homeDir(), client.path, "config.yml")
}

// Init initialize the configuration file with default values
func (client *Client) Init(config interface{}) error {
	path := client.Path()

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	err := os.MkdirAll(filepath.Dir(path), dirModePerm)

	if err != nil {
		return errors.Wrap(err, "Error creating config folder")
	}

	content, err := yaml.Marshal(config)

	if err != nil {
		return errors.Wrap(err, "Error encoding config content")
	}

	err = ioutil.WriteFile(path, content, fileModePerm)

	if err != nil {
		return errors.Wrap(err, "Error creating config file")
	}

	return nil
}
