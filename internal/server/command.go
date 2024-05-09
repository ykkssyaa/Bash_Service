package server

import (
	"encoding/json"
	"errors"
	"github.com/ykkssyaa/Bash_Service/internal/consts"
	"github.com/ykkssyaa/Bash_Service/internal/models"
	"net/http"
)
import sErr "github.com/ykkssyaa/Bash_Service/pkg/serverError"

func (s *HttpServer) CommandGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HttpServer) CommandIdGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *HttpServer) CommandPost(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("Content-Type") != "application/json" {
		sErr.ErrorResponse(w, sErr.ServerError{
			Message:    consts.ErrorNotJSON,
			StatusCode: http.StatusUnsupportedMediaType,
		})
		return
	}

	var cmd models.Command
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&cmd)

	if err != nil {
		if errors.As(err, &unmarshalErr) {
			sErr.ErrorResponse(w, sErr.ServerError{
				Message:    consts.ErrorBadRequestWrongField + unmarshalErr.Field,
				StatusCode: http.StatusBadRequest,
			})
		} else {
			sErr.ErrorResponse(w, sErr.ServerError{
				Message:    consts.ErrorBadRequest + err.Error(),
				StatusCode: http.StatusBadRequest,
			})
		}
		return
	}

	cmd, err = s.services.Command.CreateCommand(cmd.Script)
	if err != nil {
		sErr.ErrorResponse(w, err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(cmd)
}

func (s *HttpServer) StopIdPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}
