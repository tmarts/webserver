package handlers

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
)

// HomePageHandler handles request/response to the home page of our lame, fake webserver
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Welcome</h1>Click <a href=\"/viewxml?xmlfile=data/ugly.xml\">here</a> to see some ugly unformatted XML.</body></html>")
}

// XMLHandler handles request/response to the xml page of our lame, fake webserver.
//  requires a parameter of xmlfile for the XML file to dump on the page.
//  Returns a 500 if this parameter isn't specified or if it can't load the file.
func XMLHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("xmlfile")
	if filename == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("No xmlfile parameter specified"))
		return
	}
	// try to read ugly.xml file.  Throw 500 if we can't
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read " + filename + ": " + err.Error()))
		return
	}
	xmlstring := html.EscapeString(string(dat))

	fmt.Fprintf(w, "<html><body><h1>Here is some nasty, unformatted XML</h1>"+xmlstring+"<p><a href=\"./\">Go Home</a></body></html>")
}
