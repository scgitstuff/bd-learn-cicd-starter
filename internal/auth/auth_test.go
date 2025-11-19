package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	if !errorCheck(t) {
		t.Fatal("errorCheck failed")
	}

	if !validCheck(t) {
		t.Fatal("validCheck failed")
	}
}

func validCheck(t *testing.T) bool {
	headers := http.Header{}

	headers.Add("Authorization", "ApiKey 123")
	_, err := GetAPIKey(headers)
	// t.Logf("validCheck: %v", err)

	return err == nil
}

func errorCheck(t *testing.T) bool {

	headers := http.Header{}

	_, err := GetAPIKey(headers)
	// t.Logf("Authorization: %v", err)
	if err == nil {
		return false
	}

	headers.Add("Authorization", "")
	_, err = GetAPIKey(headers)
	// t.Logf("blank: %v", err)
	if err == nil {
		return false
	}

	headers.Add("Authorization", "ApiKey")
	_, err = GetAPIKey(headers)
	// t.Logf("ApiKey: %v", err)
	if err == nil {
		return false
	}

	return true
}
