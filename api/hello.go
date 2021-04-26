package api

import (
	"context"
	"encoding/json"
	"github.com/ONSdigital/log.go/log"
	"net/http"
	"time"
)

// TODO: remove hello world handler

type HelloResponse struct {
	Message string `json:"message,omitempty"`
}

type Item struct {
	Metadata []Metadata
}

type Metadata struct {
	publishTime time.Time
	content     string
}

var m map[string]Metadata

func initialiseVars() Item {
	exampleItem := Item{
		[]Metadata{
			// A page that has already been published
			{
				publishTime: time.Date(2020,1,1, 0, 0, 0, 0, time.UTC),
				content:     "a page",
			},
			// A page that will never get published XD
			{
				publishTime: time.Now().UTC().Add(time.Second * 15),
				content:     "a second page",
			},
		},
	}

	return exampleItem
}

// HelloHandler returns function containing a simple hello world example of an api handler
func HelloHandler(ctx context.Context) http.HandlerFunc {
	log.Event(ctx, "api contains example endpoint, remove hello.go as soon as possible", log.INFO)
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		item := initialiseVars()

		response := HelloResponse{
			Message:  req.URL.String(),
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
