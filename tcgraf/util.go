package tcgraf

import (
	"log"
	"net/http"
)

func writeErr(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	log.Print(err.Error())
	http.Error(w, err.Error(), http.StatusBadRequest)
}
