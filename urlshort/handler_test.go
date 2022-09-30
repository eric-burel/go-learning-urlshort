package urlshort

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMapHandlerFallback(t *testing.T) {
	fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	mapHandler := MapHandler(map[string]string{"/foo": "https://www.foo.bar"}, fallback)

	w := httptest.NewRecorder()
	rFallback := httptest.NewRequest(http.MethodGet, "/", nil)
	mapHandler(w, rFallback)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(data) != "Hello world" {
		t.Errorf("Expected 'Hello world', got %s", data)
	}
}

func TestMapHandlerRedirect(t *testing.T) {
	fallback := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	mapHandler := MapHandler(map[string]string{"/foo": "https://www.foo.bar"}, fallback)

	w := httptest.NewRecorder()
	rFoo := httptest.NewRequest(http.MethodGet, "/foo", nil)
	mapHandler(w, rFoo)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if string(data) != "<a href=\"https://www.foo.bar\">Found</a>" {
		t.Errorf("Expected data to be a link, got %s", data)
	}
}
