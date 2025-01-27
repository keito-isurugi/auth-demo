package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func GetJsonWithBody(body io.ReadCloser) ([]byte, error) {
	reqBody, err := io.ReadAll(body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to read request body: %+v", err))
	}
	defer body.Close()

	return reqBody, nil
}

func GetPearseJson(body []byte, value any) error {
	if err := json.Unmarshal(body, &value); err != nil {
		return errors.New(fmt.Sprintf("Invalid JSON format: %+v", err))
	}
	return nil
}