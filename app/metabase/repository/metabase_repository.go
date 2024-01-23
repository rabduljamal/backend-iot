package repository

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rabduljamal/backend-iot/config"
)

type MetabaseParam struct {
	Question int         `json:"question"`
	Params   interface{} `json:"params"`
}

func GetMetabaseData(dataInput *MetabaseParam) (interface{}, error) {
	METABASE_SECRET_KEY := config.Config("METABASE_SECRET_KEY")
	METABASE_SITE_URL := config.Config("METABASE_SITE_URL")

	// Create the payload
	expiration := time.Now().Add(10 * time.Minute).Unix()
	payload := jwt.MapClaims{
		"resource": map[string]interface{}{
			"question": dataInput.Question,
		},
		"params": dataInput.Params,
		"exp":    expiration,
	}

	// Generate the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(METABASE_SECRET_KEY))
	if err != nil {
		return nil, err
	}

	// Build the URL
	url := METABASE_SITE_URL + "/api/embed/card/" + tokenString + "/query/json"

	// Perform the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	// Parse the response body
	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		return nil, err
	}

	return data, nil
}
