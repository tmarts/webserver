package handlers

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
	"path/filepath"
  "runtime"
)

// HomePageHandler handles request/response to the home page of our lame, fake webserver
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Welcome</h1>Click <a href=\"/viewxml?xmlfile=ugly.xml\">here</a> to see some ugly unformatted XML.</body></html>")
}

// XMLHandler handles request/response to the xml page of our lame, fake webserver.
//  requires a parameter of xmlfile for the XML file from our data dir to dump on the page (which is super-dangerous for a non-B.S. demo).
//  Returns a 500 if this parameter isn't specified or if it can't load the file.
func XMLHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("xmlfile")
        filename = strings.Replace(filename, "/", "", -1) // trim out path separators so they can't traverse dirs
        filename = strings.Replace(filename, "..", "", -1) // trim out .. so they can't traverse dirs
	if filename == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Invalid xmlfile parameter specified"))
		return
	}

	// get project's data dir:
	 _, b, _, _ := runtime.Caller(0)
	 basepath   := filepath.Dir(b)
	 //base       := filepath.Base(b)
	 datadir := basepath + "/../data/"


	// try to read ugly.xml file from data dir.  Throw 500 if we can't
	dat, err := ioutil.ReadFile(datadir + filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read " + filename + " from project's data dir"))
		return
	}
	xmlstring := html.EscapeString(string(dat))

	fmt.Fprintf(w, "<html><body><h1>Here is some nasty, unformatted XML</h1>"+xmlstring+"<p><a href=\"./\">Go Home</a></body></html>")
}
