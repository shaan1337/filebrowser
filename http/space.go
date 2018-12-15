package http

import (
	"net/http"

	fb "github.com/shaan1337/filebrowser"
)

func spaceHandler(c *fb.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	r.URL.Path = sanitizeURL(r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		return spaceGetHandler(c, w, r)
	}

	return http.StatusNotImplemented, nil
}

func spaceGetHandler(c *fb.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	f := c.File
	listing, err := f.GetFullListing(c.FileBrowser, c.User, r)
	if err != nil {
		return ErrorToHTTP(err, true), err
	}

	return renderJSON(w, listing)
}
