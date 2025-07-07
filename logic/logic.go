package logic

import "errors"

// concatenates 2 strings
func stringConcat(str1, str2 string) (string, error) {
	if str1 == "" {
		return "", errors.New("error: str1 cannot be empty")
	}

	str := str1 + " " + str2
	return str, nil
}
