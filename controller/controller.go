package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPResult struct {
	Error    error
	Code     int
	JSONData []byte
}

// Controller Commment - TODO
type Controller interface {
	post(*http.Request)
	get(*http.Request)
	put(*http.Request)
	delete(*http.Request)
}

func SetupControllers() *map[Controller]string {
	kevin := make(map[Controller]string)
	return &kevin
}

func HandleError(err error) {
	if err != nil {
		fmt.Printf("Fucked %+v", err)
	}
}

func ParseJSON(rc io.Reader) (map[string]interface{}, error) {
	// Read the request body
	body, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}

	// Parse data into a generic interface. This will get us our JSON!
	var JSON interface{}
	err = json.Unmarshal(body, &JSON)
	if err != nil {
		return nil, err
	}

	// Create a map from the unmarshalled interface
	JSONMap := JSON.(map[string]interface{}) // This is a type assertion
	return JSONMap, nil
}
