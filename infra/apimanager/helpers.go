package apimanager

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/lawmatsuyama/pismo-transactions/domain"
	"github.com/sirupsen/logrus"
)

func Decode[T any](r io.Reader) (T, error) {
	var obj T
	err := json.NewDecoder(r).Decode(&obj)
	return obj, err
}

func HandleResponse[T any](w http.ResponseWriter, r *http.Request, in T, err error) {
	var errStr string
	statusCode := http.StatusOK
	if err != nil {
		errTr := domain.ErrorToErrorTransaction(err)
		errStr = errTr.Error()
		statusCode = errTr.Status()
	}

	genRes := GenericResponse[T]{
		Error:  errStr,
		Result: in,
	}

	res, err := json.Marshal(genRes)
	if err != nil {
		logrus.WithError(err).Error("couldnt marshal the response to json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		logrus.WithError(err).Error("couldnt send response to writer")
		http.Error(w, domain.ErrUnknown.Error(), http.StatusBadRequest)
	}
}
