package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
func checkHandlerResponse(verb string, handler http.Handler, query string, expectedBody string, expectedRetcode int, t *testing.T) {
  // We first create the http.Handler we wish to test
	n := handler

	// We create an http.Request object to test with. The http.Request is
	// totally customizable in every way that a real-life http request is, so
	// even the most intricate behavior can be tested
	r, _ := http.NewRequest("GET", "/one", nil)

	// httptest.Recorder implements the http.ResponseWriter interface, and as
	// such can be passed into ServeHTTP to receive the response. It will act as
	// if all data being given to it is being sent to a real client, when in
	// reality it's being buffered for later observation
	w := httptest.NewRecorder()

	// Pass in our httptest.Recorder and http.Request to our numberDumper. At
	// this point the numberDumper will act just as if it was responding to a
	// real request
	n.ServeHTTP(w, r)

	// httptest.Recorder gives a number of fields and methods which can be used
	// to observe the response made to our request. Here we check the response
	// code
	if w.Code != 200 {
		t.Fatalf("wrong code returned: %d", w.Code)
	}

	// We can also get the full body out of the httptest.Recorder, and check
	// that its contents are what we expect
	body := w.Body.String()

}*/

func TestHomePageHandlerReturns200(t *testing.T) {
	r, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	HomePageHandler(w, r)
	if w.Code != 200 {
		t.Fatalf("Wrong code returned from home page: %d", w.Code)
	}
}

func TestHomePageHandlerReturnsExpected(t *testing.T) {
	expectedBody := "Click <a href=\"/viewxml?xmlfile=data/ugly.xml\">here</a> to see some ugly unformatted XML."
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
	expectedBody := "No xmlfile parameter specified"
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
	expectedBody := "Failed to read foobar: "
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
	q := r.URL.Query()                   // Get a copy of the query values.
	q.Add("xmlfile", "../data/ugly.xml") // Add a new value to the set.
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
	q := r.URL.Query()                   // Get a copy of the query values.
	q.Add("xmlfile", "../data/ugly.xml") // Add a new value to the set.
	r.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	XMLHandler(w, r)
	if !strings.Contains(w.Body.String(), expectedBody) {
		t.Fatalf("Response body of xml page when called with bad file param did not contain expected text %s.  Actual Body: %s ", expectedBody, w.Body.String())
	}
}
