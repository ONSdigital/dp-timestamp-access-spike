package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ONSdigital/log.go/log"
	"net/http"
	"sort"
	"time"
)

type HelloResponse struct {
	Content     string    `json:"content,omitempty"`
	PublishTime time.Time `json:"publish_time,omitempty"`
}

type MetadataSlice []Metadata

func (p MetadataSlice) Len() int {
	fmt.Println("calling len")
	return len(p)
}

func (p MetadataSlice) Less(i, j int) bool {
	fmt.Println("calling less")
	return p[i].PublishTime.After(p[j].PublishTime)
}

func (p MetadataSlice) Swap(i, j int) {
	fmt.Println("calling swap")
	p[i], p[j] = p[j], p[i]
}

type Item struct {
	Metadata MetadataSlice
}

type Metadata struct {
	PublishTime time.Time `json:"publish_time,omitempty"`
	Content     string    `json:"content,omitempty"`
}

var m map[string]Item

func initialiseVars() {

	m = make(map[string]Item)

	// current page is published
	m["/publishedpage"] = Item{
		[]Metadata{
			// A page that has already been published
			{
				PublishTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Content:     "published page content",
			},
		},
	}

	// current page is published - and a page to be published
	m["/bulletin1"] = Item{
		[]Metadata{
			// A page that has already been published
			{
				PublishTime: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				Content:     "the old bulletin1 version content",
			},
			{
				PublishTime: time.Date(2020, 2, 1, 0, 0, 0, 0, time.UTC),
				Content:     "the latest bulletin1 published content ",
			},
			// A page that will never get published XD
			{
				PublishTime: time.Now().UTC().Add(time.Second * 15),
				Content:     "the bulletin1 content to be published",
			},
		},
	}

	// a page to be published
	m["/tobepublished"] = Item{
		[]Metadata{
			// A page that will never get published XD
			{
				PublishTime: time.Now().UTC().Add(time.Second * 15),
				Content:     "the bulletin1 content to be published",
			},
		},
	}
}

func selectLatest(item Item) Metadata {

	sort.Sort(item.Metadata)
	for _, metadata := range item.Metadata {
		if metadata.PublishTime.Before(time.Now()) {
			return metadata
		}
	}

	return Metadata{
		Content: "content not found",
	}
}

// HelloHandler returns function containing a simple hello world example of an api handler
func HelloHandler(ctx context.Context) http.HandlerFunc {
	log.Event(ctx, "api contains example endpoint, remove hello.go as soon as possible", log.INFO)
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		initialiseVars()

		var response Metadata

		item, ok := m[req.URL.String()]
		if !ok {
			response = Metadata{
				Content: "content not found - URL was not recognised",
			}
		} else {
			response = selectLatest(item)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			log.Event(ctx, "marshalling response failed", log.Error(err), log.ERROR)
			http.Error(w, "Failed to marshall json response", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(jsonResponse)
		if err != nil {
			log.Event(ctx, "writing response failed", log.Error(err), log.ERROR)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}
