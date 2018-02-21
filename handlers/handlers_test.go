package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHomePageHandlerReturns200(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	HomePageHandler(w, r)
	if w.Code != 200 {
		t.Fatalf("Wrong code returned from home page: %d", w.Code)
	}
}

func TestHomePageHandlerReturnsExpected(t *testing.T) {
	expectedBody := "Click <a href=\"/viewxml?xmlfile=ugly.xml\">here</a> to see some ugly unformatted XML."
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	HomePageHandler(w, r)
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Home page body did not contain %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}
}

func TestXMLHandlerNoParamsReturns500(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if w.Code != 500 {
		t.Fatalf("Wrong code returned from xml page when no params are passed: %d", w.Code)
	}
}

func TestXMLHandlerNoParamsReturnsExpectedContent(t *testing.T) {
	expectedBody := "Invalid xmlfile parameter specified"
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Response body of xml page when called with no parms did not contain expected text %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}
}

func TestXMLHandlerBadFileReturns500(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	q := r.URL.Query()         // Get a copy of the query values.
	q.Add("xmlfile", "foobar") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if w.Code != 500 {
		t.Fatalf("Wrong code returned from xml page bad file param is passed: %d", w.Code)
	}
}

func TestXMLHandlerBadFileReturnsExpectedContent(t *testing.T) {
	expectedBody := "Failed to read foobar from project's data dir"
	r, _ := http.NewRequest("GET", "", nil)
	q := r.URL.Query()         // Get a copy of the query values.
	q.Add("xmlfile", "foobar") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Response body of xml page when called with bad file param did not contain expected text %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}
}

func TestXMLHandlerGoodFileReturns200(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	q := r.URL.Query()           // Get a copy of the query values.
	q.Add("xmlfile", "ugly.xml") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if w.Code != 200 {
		t.Fatalf("Wrong code returned from xml page good file param is passed: %d", w.Code)
	}
}

func TestXMLHandlerGoodFileReturnsExpectedContent(t *testing.T) {
	expectedBody := "Here is some nasty, unformatted XML"
	r, _ := http.NewRequest("GET", "", nil)
	q := r.URL.Query()           // Get a copy of the query values.
	q.Add("xmlfile", "ugly.xml") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Response body of xml page when called with bad file param did not contain expected text %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}
}

func TestXmlHanderCannotServeEtcPasswd(t *testing.T) {
	expectedBody := "Failed to read etcpasswd from project's data dir"
	r, _ := http.NewRequest("GET", "", nil)
	q := r.URL.Query()              // Get a copy of the query values.
	q.Add("xmlfile", "/etc/passwd") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if w.Code != 500 {
		t.Fatalf("Wrong code returned from xml page bad file param is passed: %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Response body of xml page when called with no parms did not contain expected text %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}
}

func TestXmlHanderCannotServeFilesOutsideOfDataDir(t *testing.T) {
	expectedBody := "Failed to read webserver.go from project's data dir"
	r, _ := http.NewRequest("GET", "", nil)
	q := r.URL.Query()                  // Get a copy of the query values.
	q.Add("xmlfile", "../webserver.go") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if w.Code != 500 {
		t.Fatalf("Wrong code returned from xml page bad file param is passed: %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Response body of xml page when called with no parms did not contain expected text %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}

}
