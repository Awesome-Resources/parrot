package render

import (
	"encoding/json"
	"net/http"

	"github.com/anthonynsimon/parrot/errors"
)

var jsonContentType = "application/json; charset=utf-8"

func JSONError(w http.ResponseWriter, e *errors.Error) {
	data := map[string]interface{}{
		"status": e.Code,
		"error":  e.Message,
	}
	JSON(w, e.Code, data)
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", jsonContentType)
	w.WriteHeader(status)

	encoded, err := json.Marshal(data)
	if err != nil {
		http.Error(w, errors.ErrInternal.Message, errors.ErrInternal.Code)
	}

	w.Write(encoded)
}