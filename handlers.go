package main

import (
	"net/http"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/zemirco/uid"
	"github.com/bitly/go-simplejson"
)

var urlValidator, _ = regexp.Compile(`^(?:http(s)?://)?[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=]+$`)

func GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	j, _ := simplejson.NewFromReader(r.Body)
	url, _ := j.Get("url").String()
	if urlValidator.MatchString(url) {
		t := map[string]string{
			"token": uid.New(8),
		}
		if rSet(t["token"], url) {
			payload, _ := json.Marshal(t)
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
		}
	}
}

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	url := rGet(strings.TrimLeft(r.URL.Path, "/"))
	if url != ""{
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
}
