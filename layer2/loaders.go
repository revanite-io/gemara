package layer2

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

// loadYamlFromURL is a sub-function of loadYaml for HTTP only
// sourcePath is the URL. data is a pointer to the recieving object.
func loadYamlFromURL(sourcePath string, data *Layer2) error {
	resp, err := http.Get(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch URL; response status: %v", resp.Status)
	}

	err = decode(resp.Body, data)
	if err != nil {
		return fmt.Errorf("failed to decode YAML from URL: %v", err)
	}
	return nil
}

// loadYaml opens a provided path to unmarshal its data as YAML.
// sourcePath is a URL or local path to a file.
// data is a pointer to the recieving object.
func loadYaml(sourcePath string, data *Layer2) error {
	if strings.HasPrefix(sourcePath, "http") {
		return loadYamlFromURL(sourcePath, data)
	}

	file, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	err = decode(file, data)
	if err != nil {
		return fmt.Errorf("error decoding YAML: %w (%s)", err, sourcePath)
	}
	return nil
}

// loadYaml opens a provided path to unmarshal its data as JSON.
// sourcePath is a URL or local path to a file.
// data is a pointer to the recieving object.
func loadJson(sourcePath string, data *Layer2) error {
	return fmt.Errorf("loadJson not implemented [%s, %v]", sourcePath, data)
}

// LoadControlFamiliesFile loads data from any number of YAML
// files at the provided paths. JSON support is pending development.
// If run multiple times, this method will append new data to previous data.
func (c *Layer2) LoadFiles(sourcePaths []string) error {
	for _, sourcePath := range sourcePaths {
		catalog := &Layer2{}
		err := c.LoadFile(sourcePath)
		if err != nil {
			return err
		}
		c.ControlFamilies = append(c.ControlFamilies, catalog.ControlFamilies...)
		c.Capabilities = append(c.Capabilities, catalog.Capabilities...)
		c.Threats = append(c.Threats, catalog.Threats...)
	}
	return nil
}

// LoadControlFamiliesFile loads data from a single YAML
// file at the provided path. JSON support is pending development.
// If run multiple times for the same data type, this method will override previous data.
func (c *Layer2) LoadFile(sourcePath string) error {
	if strings.Contains(sourcePath, ".yaml") || strings.Contains(sourcePath, ".yml") {
		err := loadYaml(sourcePath, c)
		if err != nil {
			return err
		}
	} else if strings.Contains(sourcePath, ".json") {
		err := loadJson(sourcePath, c)
		if err != nil {
			return fmt.Errorf("error loading json: %w", err)
		}
	} else {
		return fmt.Errorf("unsupported file type")
	}
	return nil
}

func decode(reader io.Reader, data *Layer2) error {
	decoder := yaml.NewDecoder(reader, yaml.DisallowUnknownField())
	err := decoder.Decode(data)
	if err != nil {
		return fmt.Errorf("error decoding YAML: %w", err)
	}
	return nil
}
