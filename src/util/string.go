package util

import (
	"bytes"
	"encoding/json"
)

func PrettyString(response []byte) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, response, "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
