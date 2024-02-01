package foohandler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGetFoo(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handleGetFoo))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}

	expected := "FOO"
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}

func TestHandleGetFooRR(t *testing.T) {

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}
	handleGetFoo(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Result().StatusCode)
	}

	expected := "FOO"
	b, err := io.ReadAll(rr.Result().Body)
	if err != nil {
		t.Error(err)
	}
	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}
}
