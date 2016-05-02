package config

import (
	"testing"
)

func TestParseMasterUrl_EmptyString(t *testing.T) {
	uri, err := stringToMasterUrl("")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_TooSimple(t *testing.T) {
	uri, err := stringToMasterUrl("apa")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_NoPort(t *testing.T) {
	uri, err := stringToMasterUrl("http://apa:")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_NoHost(t *testing.T) {
	uri, err := stringToMasterUrl("http://:apa")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_NoProtocol(t *testing.T) {
	uri, err := stringToMasterUrl("apa:10")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_TooManyParts(t *testing.T) {
	uri, err := stringToMasterUrl("http://a:b:c")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_TooManyProtocols(t *testing.T) {
	uri, err := stringToMasterUrl("http://ftp://a:b:c")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_InvalidPort_Zero(t *testing.T) {
	uri, err := stringToMasterUrl("http://apa:0")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_InvalidPort_Negative(t *testing.T) {
	uri, err := stringToMasterUrl("http://apa:-1")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_InvalidPort_TooLarge(t *testing.T) {
	uri, err := stringToMasterUrl("http://apa:65536")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseMasterUrl_Valid(t *testing.T) {
	uri, err := stringToMasterUrl("http://example.com:1337")
	if err != nil {
		t.Error("Expected no error but got", err)
	}

	if uri.protocol != "http" {
		t.Error("Expected http, got", uri.protocol)
	}
	if uri.port != 1337 {
		t.Error("Expected 1337, got", uri.port)
	}
	if uri.host != "example.com" {
		t.Error("Expected example.com, got", uri.host)
	}
}