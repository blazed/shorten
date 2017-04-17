package server

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/blazed/shorten/storage"
	"github.com/pressly/chi"
)

const salt = "XoX^B#5utID2s36MYW!zl!fpd0hxY!#7"

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", "")
}

func (s *Server) handleURLSlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "urlSlug")
	// Check the database if urlSlug exists, if not throw error
	url, err := s.storage.GetUrl(slug)
	if err != nil {
		http.Error(w, "No url found for "+slug, http.StatusNotFound)
		return
	}
	// if slug exists, set 302 and of we go!
	http.Redirect(w, r, url.URL, http.StatusFound)
}

func (s *Server) handleCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formURL := r.FormValue("url")

	if !strings.HasPrefix(formURL, "http://") || !strings.HasPrefix(formURL, "https://") {
		fmt.Println(formURL)
		formURL = "http://" + formURL
	}

	slug := genetareSlug(formURL)
	if url, _ := s.storage.GetUrl(slug); len(url.Slug) != 0 {
		w.Write([]byte(fmt.Sprintf("%s/%s", r.Host, url.Slug)))
		return
	}

	short := storage.URL{Slug: slug, URL: formURL}
	if err := s.storage.CreateShortUrl(short); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(fmt.Sprintf("%s/%s", r.Host, slug)))
}

func genetareSlug(url string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	h.Write([]byte(url))
	sum := h.Sum(nil)
	b := make([]byte, base64.URLEncoding.EncodedLen(len(sum)))
	base64.URLEncoding.Encode(b, sum)
	slug := string(b)[:10]
	return slug
}
