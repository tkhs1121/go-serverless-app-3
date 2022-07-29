package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestParseRequest(t *testing.T) {
	tests := []struct {
		name   string
		inputs string
		want   Request
	}{
		{
			name:   "Parse Request URL",
			inputs: `{"url": "https://www.amazon.co.jp"}`,
			want: Request{
				URL: "https://www.amazon.co.jp",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := parseRequest(tt.inputs); *got != tt.want {
				if err != nil {
					t.Errorf("%v", err)
				}
				t.Errorf("parseRequest() = %v, want %v", *got, tt.want)
			}
		})
	}

}

func TestCheckAmazonURL(t *testing.T) {
	t.Run("Correct URL", func(t *testing.T) {
		if err := checkAmazonURL("https://www.amazon.co.jp"); err != nil {
			t.Errorf("%v", err)
		}
	})
	t.Run("Invalid URL", func(t *testing.T) {
		if err := checkAmazonURL("http://konozama.co.jp"); err == nil {
			t.Errorf("%v", err)
		}
	})
}

func TestErrorResponse(t *testing.T) {
	body, sc := "ERROR", 500
	expected := events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: sc,
	}
	t.Run("Error Response", func(t *testing.T) {
		got, _ := errorResponse(fmt.Errorf(body), sc)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("errorResponse() = %v, want %v", got, expected)
		}
	})
}
