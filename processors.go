package smpl

import (
	"errors"
	"strings"
)

func (c *Configuration) ProcessData(data string) error {
	c.Get = make(map[string]string)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Split(line, " : ")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if value[0] == '[' {
				c.ProcessArray(key, value)
			} else {
				c.Get[key] = value
			}
		} else {
			return errors.New("more than 1 ':' per line")
		}
	}
	return nil
}

func (c *Configuration) ProcessArray(key string, value string) {
	arr := strings.Split(value[1:len(value)-1], ",")
	c.Geta[key] = append(c.Geta[key], arr...)
}
