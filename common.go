package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
)

type Request struct {
	URL string `json:"url"`
}

func getRandID() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(100))
}

func getEpochTime() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func parseRequest(inputs string) (*Request, error) {
	var req Request
	err := json.NewDecoder(strings.NewReader(inputs)).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("cannot decode body: %v: %v", err, inputs)
	}
	return &req, nil
}

func checkAmazonURL(url string) error {
	isAmazpnURL, err := regexp.MatchString(`^http(s)?://(www.)?amazon.(co.)?jp[\w!\?/\+\-_~=;\.,\*&@#\$%\(\)'\[\]]*`, url)
	if err != nil {
		return err
	}
	if !isAmazpnURL {
		return fmt.Errorf("this is not amazon url: %v", url)
	}
	return nil
}

func errorResponse(err error, statusCode int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       err.Error(),
		StatusCode: statusCode,
	}, err
}
