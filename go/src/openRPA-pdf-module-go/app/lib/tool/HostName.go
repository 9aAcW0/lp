package tool

import "os"

func GetHostURL() (string, error) {

	name, err := os.Hostname()
	if err != nil {
		return "", err
	}

	return name, err
}
