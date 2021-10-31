package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func getName(r *http.Request) (string, error) {
	var req struct{ Name string }
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(data, &req)
	return req.Name, err
}
