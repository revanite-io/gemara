package layer2

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// loadYamlFromURL is a sub-function of loadYaml for HTTP only
// sourcePath is the URL. data is a pointer to the recieving object.
func loadYamlFromURL(sourcePath string, data interface{}) error {
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

	decoder := yaml.NewDecoder(resp.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return fmt.Errorf("failed to decode YAML from URL: %v", err)
	}
	return nil
}

// loadYaml opens a provided path to unmarshal its data as YAML.
// sourcePath is a URL or local path to a file.
// data is a pointer to the recieving object.
func loadYaml(sourcePath string, data interface{}) error {
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
	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)

	err = decoder.Decode(data)
	if err != nil {
		return fmt.Errorf("error decoding YAML: %w", err)
	}
	return nil
}

// loadYaml opens a provided path to unmarshal its data as JSON.
// sourcePath is a URL or local path to a file.
// data is a pointer to the recieving object.
func loadJson(sourcePath string, data interface{}) error {
	return fmt.Errorf("loadJson not implemented [%s, %s]", sourcePath, data)
}

// LoadControlFamiliesFile loads data from any number of YAML
// files at the provided paths. JSON support is pending development.
// If run multiple times, this method will append new data to previous data.
func (c *Catalog) LoadFiles(sourcePaths []string) error {
	for _, sourcePath := range sourcePaths {
		catalog := &Catalog{}
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
func (c *Catalog) LoadFile(sourcePath string) error {
	if strings.Contains(sourcePath, ".yaml") {
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
