package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	storyFile, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}
	story := make(map[string]StoryArc)
	if err = json.Unmarshal(storyFile, &story); err != nil {
		panic(err)
	}

	tmp, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arc, ok := story[r.URL.Path[1:]]
		if !ok {
			arc = story["intro"]
		}
		tmp.Execute(w, arc)
	})

	if err = http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
