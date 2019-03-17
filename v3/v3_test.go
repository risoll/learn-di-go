package v3

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type (
	MockHttpClient struct {
		resp *http.Response
		err error
	}

	errReader int
)

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test err")
}

func (m *MockHttpClient) Get(url string) (*http.Response, error) {
	return m.resp, m.err
}

func TestWithValidResponse(t *testing.T) {
	httpClient := &MockHttpClient{
		resp: &http.Response{
			Body: ioutil.NopCloser(bytes.NewBuffer([]byte("should works"))),
		},
	}
	err := send(httpClient, "should works")
	if err != nil {
		t.Errorf("shouldn't received error: got %s", err)
	}
}

func TestWithErr(t *testing.T) {
	httpClient := &MockHttpClient{err: errors.New("error")}
	err := send(httpClient, "should not works - got error")
	if err == nil {
		t.Errorf("should received error: got %s", err)
	}
}

func TestWithNilResp(t *testing.T) {
	httpClient := &MockHttpClient{
		resp: nil,
	}
	err := send(httpClient, "should not works - nil response")
	if err == nil {
		t.Errorf("should received error: got %s", err)
	}
}

func TestWithInvalidResp(t *testing.T) {
	req, _ := http.NewRequest("GET", "/hehe", errReader(0))
	httpClient := &MockHttpClient{
		resp: &http.Response{
			Body: req.Body,
		},
	}
	err := send(httpClient, "should not works - invalid response")
	if err == nil {
		t.Errorf("should received error: got %s", err)
	}
}