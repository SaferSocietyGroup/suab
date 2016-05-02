package config

import (
	"testing"
)

func TestParseSwarmUri_EmptyString(t *testing.T) {
	uri, err := stringToSwarmUri("")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_NoColon(t *testing.T) {
	uri, err := stringToSwarmUri("apa")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_NoPort(t *testing.T) {
	uri, err := stringToSwarmUri("apa:")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_NoHost(t *testing.T) {
	uri, err := stringToSwarmUri(":apa")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_TooManyParts(t *testing.T) {
	uri, err := stringToSwarmUri("a:b:c")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_InvalidPort_Zero(t *testing.T) {
	uri, err := stringToSwarmUri("apa:0")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_InvalidPort_Negative(t *testing.T) {
	uri, err := stringToSwarmUri("apa:-1")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_InvalidPort_TooLarge(t *testing.T) {
	uri, err := stringToSwarmUri("apa:65536")
	if err == nil {
		t.Error("Expected error got", uri)
	}
}

func TestParseSwarmUri_Valid(t *testing.T) {
	uri, err := stringToSwarmUri("example.com:1337")
	if err != nil {
		t.Error("Expected no error but got", err)
	}

	if uri.port != 1337 {
		t.Error("Expected 1337, got", uri.port)
	}
	if uri.host != "example.com" {
		t.Error("Expected example.com, got", uri.host)
	}
}