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
	keys, _ := client.Keys("*").Result()
	for _, key := range keys {
		cachedUrl := rGet(key)
		if url == cachedUrl {
			t := map[string]string{
				"token": key,
			}
			payload, _ := json.Marshal(t)
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
			return
		}
	}
	if urlValidator.MatchString(url) {
		t := map[string]string{
			"token": uid.New(8),
		}
		if rSet(t["token"], url) {
			payload, _ := json.Marshal(t)
			w.Header().Set("Content-Type", "application/json")
			w.Write(payload)
			return
		}
	}
}

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	url := rGet(strings.TrimLeft(r.URL.Path, "/"))
	if url != ""{
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") { //Some of dirty hacks
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "http://" + url, http.StatusSeeOther)
		}
	}
}
