package tcgraf

import (
	"net/http"
	"testing"
)

func TestWriteErr(t *testing.T) {
	t.Run("it handles nil errors", func(t *testing.T) {
		var err error
		w := new(http.ResponseWriter)
		writeErr(*w, err)
	})
}
