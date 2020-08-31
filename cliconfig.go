package cliconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"gopkg.in/yaml.v2"
)

const (
	dirModePerm  = 755
	fileModePerm = 644
)

// NewFunc is the function that creates the configuration struct with default values
type NewFunc func() interface{}

// Client is the configuration client for a cli
type Client struct {
	path string
	new  NewFunc
}

// New creates a new cliconfig client
func New(path string, newFunc NewFunc) *Client {
	return &Client{
		path: path,
		new:  newFunc,
	}
}

// FromFile creates a new config from a file
func (client *Client) FromFile() (interface{}, error) {
	config := client.new()

	content, err := ioutil.ReadFile(client.Path())

	if err != nil {
		return nil, errors.Wrap(err, "Error reading config file")
	}

	err = yaml.Unmarshal(content, &config)

	if err != nil {
		return nil, errors.Wrap(err, "Error decoding config file content")
	}

	return config, nil
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
func (client *Client) Init() error {
	path := client.Path()

	if _, err := os.Stat(path); err == nil {
		return nil
	}

	err := os.MkdirAll(filepath.Dir(path), dirModePerm)

	if err != nil {
		return errors.Wrap(err, "Error creating config folder")
	}

	config := client.new()
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

// Reset updates the configuration file with default values
func (client *Client) Reset() error {
	path := client.Path()

	err := os.MkdirAll(filepath.Dir(path), dirModePerm)

	if err != nil {
		return errors.Wrap(err, "Error creating config folder")
	}

	config := client.new()
	content, err := yaml.Marshal(config)

	if err != nil {
		return errors.Wrap(err, "Error encoding yaml config file")
	}

	err = ioutil.WriteFile(path, content, fileModePerm)

	if err != nil {
		return errors.Wrap(err, "Error writing config file")
	}

	return nil
}

// Open opens the configuration file inside the default file editor
func (client *Client) Open() {
	open.Run(client.Path())
}

// Show prints the content of the config file inside the console
func (client *Client) Show() error {
	config, err := client.FromFile()

	if err != nil {
		return errors.Wrap(err, "Error reading config file")
	}

	content, err := yaml.Marshal(config)

	if err != nil {
		return errors.Wrap(err, "Error decoding content of config file")
	}

	fmt.Println(string(content))

	return nil
}
