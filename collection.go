package adalo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Collection provides a CRUD interface to an Adalo collection.
type Collection struct {
	// ID of collection in Adalo
	ID string
}

// NewCollection initializes a Collection.
func NewCollection(collectionID string) *Collection {
	return &Collection{ID: collectionID}
}

// collectionAPIBaseURL returns the base url for api calls.
func (c *Collection) collectionAPIBaseURL() string {
	return fmt.Sprintf("http://api.adalo.com/apps/%s/collections/%s", AppID, c.ID)
}

// All gets all items in collection and binds result to the passed result variable.
func (c *Collection) All(result interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.collectionAPIBaseURL(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// check if an explicit error message was returned
	var errorResponse apiErrorResponse
	_ = json.Unmarshal(body, &errorResponse)
	if errorResponse != (apiErrorResponse{}) {
		return errors.New(strings.ToLower(errorResponse.Error))
	}

	return json.Unmarshal(body, &result)
}

// Get fetches a record from the collection by its id and binds it to passed result variable.
func (c *Collection) Get(id int, result interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%d", c.collectionAPIBaseURL(), id), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// check if an explicit error message was returned
	var errorResponse apiErrorResponse
	_ = json.Unmarshal(body, &errorResponse)
	if errorResponse != (apiErrorResponse{}) {
		return errors.New(strings.ToLower(errorResponse.Error))
	}

	return json.Unmarshal(body, &result)
}

// Insert will insert a new record to the collection and bind created item to passed result variable.
func (c *Collection) Insert(input interface{}, result interface{}) error {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(inputBytes)

	client := &http.Client{}
	req, err := http.NewRequest("POST", c.collectionAPIBaseURL(), payload)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// check if an explicit error message was returned
	var errorResponse apiErrorResponse
	_ = json.Unmarshal(body, &errorResponse)
	if errorResponse != (apiErrorResponse{}) {
		return errors.New(strings.ToLower(errorResponse.Error))
	}

	return json.Unmarshal(body, &result)
}

// Update will update the record with given id in the Adalo collection and bind updated item to passed result variable.
func (c *Collection) Update(id int, input interface{}, result interface{}) error {
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	payload := bytes.NewReader(inputBytes)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%d", c.collectionAPIBaseURL(), id), payload)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// check if an explicit error message was returned
	var errorResponse apiErrorResponse
	_ = json.Unmarshal(body, &errorResponse)
	if errorResponse != (apiErrorResponse{}) {
		return errors.New(strings.ToLower(errorResponse.Error))
	}

	return json.Unmarshal(body, &result)
}

// Delete removes a record from the Adalo collection.
func (c *Collection) Delete(id int) error {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%d", c.collectionAPIBaseURL(), id), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// check if an explicit error message was returned
	var errorResponse apiErrorResponse
	_ = json.Unmarshal(body, &errorResponse)
	if errorResponse != (apiErrorResponse{}) {
		return errors.New(strings.ToLower(errorResponse.Error))
	}

	switch res.StatusCode {
	case 204:
		return ErrorResourceNotFound // BUG: Adalo will return with 204 even with successful requests
	case 201:
		return nil
	default:
		return nil
	}
}
