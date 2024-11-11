package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Config struct with generics
type Config[T any] struct {
	Filename string
	ModTime  time.Time
	Data     T
}

func New[T any](filename string) Config[T] {
	return Config[T]{
		Filename: filepath.Clean(filename),
		ModTime:  time.Time{},
		Data:     *new(T),
	}
}

// loadYAMLConfig function for YAML files
func LoadYAMLConfig[T any](filename string, target T) (T, error) {
	var err error

	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return target, err
	}

	// Unmarshal YAML content into target struct
	err = yaml.Unmarshal(fileContent, &target)
	if err != nil {
		return target, err
	}

	return target, nil
}

// Load config file. The file is only read once unless it is modified.
func (c *Config[T]) Load() (err error) {

	info, err := os.Stat(c.Filename)
	if err != nil {
		return err
	}

	if info.ModTime().Equal(c.ModTime) {
		return
	}

	c.Data, err = LoadYAMLConfig(c.Filename, c.Data)
	if err != nil {
		return err
	}

	c.ModTime = info.ModTime()
	fmt.Printf("config file %v loaded\n", c.Filename)
	return
}

// Same as load, but print error on console and doesn't return error
func (c *Config[T]) Reload() {
	err := c.Load()
	if err != nil {
		fmt.Printf("error loading %v config file: %v\n", c.Filename, err)
	}
}

func JoinPath(folderPath string, filePath string) string {
	if filepath.IsAbs(filePath) {
		return filePath
	}

	return filepath.Join(folderPath, filePath)
}

// Get the path to the config folder
func ConfigPath(args []string) (path string, err error) {

	// path given as command-line argument
	if len(args) > 1 {
		path = filepath.Clean(args[1])
	} else {
		// path is a /conf folder in the executable folder
		var executablePath string
		if executablePath, err = os.Executable(); err != nil {
			return
		}
		path = filepath.Join(filepath.Dir(executablePath), "conf")
	}

	_, err = os.Stat(path)

	return
}
