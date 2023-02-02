package tcgraf

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httputil"

	graf "github.com/grafana/grafana-api-golang-client"
)

var client *graf.Client

func initClient(address string, apiKey string) (err error) {
	if client == nil {
		client, err = graf.New(address, graf.Config{APIKey: apiKey})
		if err != nil {
			return
		}
	}

	return
}

func Start(address string, apiKey string) (err error) {
	err = initClient(address, apiKey)
	if err != nil {
		return
	}

	http.HandleFunc("/", handler)

	return http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, req *http.Request) {
	reqDump, err := httputil.DumpRequest(req, true)
	if err == nil {
		log.Print(string(reqDump))
	}

	var ev Event
	err = json.NewDecoder(req.Body).Decode(&ev)
	if err != nil {
		writeErr(w, err)
		return
	}
	if ev.Type == "" {
		err = errors.New("an event type is required")
		writeErr(w, err)
		return
	}
	if ev.InstanceId == "" {
		err = errors.New("an instance id is required")
		writeErr(w, err)
		return
	}
	if ev.NasId == "" {
		err = errors.New("a nas id is required")
		writeErr(w, err)
		return
	}

	switch ev.Type {
	case CreateDash:
		createDash(ev)
		if err != nil {
			log.Print("dash creation failed", err.Error())
			writeErr(w, err)
			return
		}

		_ = json.NewEncoder(w).Encode(ResponseSuccess{
			Message: "nice",
			Success: true,
		})

	default:
		http.Error(w, "no matching event type found for event '"+ev.Type.String()+"'", http.StatusBadRequest)
	}
}
