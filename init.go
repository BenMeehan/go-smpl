package smpl

import (
	"io/ioutil"
	"net/http"
)

// InitializeFromFile reads a file from a local path and returns its contents
func (c *Configuration) InitializeFromFile(filePath string) error {
	c.Get = map[string]string{}
	c.Geta = make(map[string][]string)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = c.ProcessData(string(data))
	if err != nil {
		return err
	}

	return nil
}

// InitializeFromURL reads a file from a URL and returns its contents
func (c *Configuration) InitializeFromURL(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = c.ProcessData(string(data))
	if err != nil {
		return err
	}

	return nil
}
