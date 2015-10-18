package hello

import (
	"encoding/json"
	"net/http"

	"github.com/lytics/multibayes"
)

//
//  default response stuff
//

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/multibayes", multibayesHandler)

	initClassifier()
}

type servicesResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func writeResponse(w http.ResponseWriter, status int, data interface{}, msg string) {
	response := servicesResponse{
		Status: status,
		Data:   data,
	}

	if len(msg) > 0 {
		response.Msg = msg
	}

	jsonData, err := json.Marshal(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response.Status = http.StatusInternalServerError
		response.Msg = err.Error()
	}
	response.Data = jsonData

	w.WriteHeader(status)
	w.Write(jsonData)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusNotFound, nil, "nothing here to see...")
	return
}

//
//  multibayes specific stuff
//

var (
	classifier *multibayes.Classifier
)

type multibayesResponse struct {
	Probs map[string]float64 `json:"probs"`
}

func initClassifier() {
	// documents to train the classifier
	documents := []struct {
		Text    string
		Classes []string
	}{
		{
			Text:    "My dog has fleas.",
			Classes: []string{"vet"},
		},
		{
			Text:    "My cat has ebola.",
			Classes: []string{"vet", "cdc"},
		},
		{
			Text:    "My cat has ebola.",
			Classes: []string{"vet", "cdc"},
		},
		{
			Text:    "My cat has ebola.",
			Classes: []string{"vet", "cdc"},
		},
		{
			Text:    "My cat has ebola.",
			Classes: []string{"vet", "cdc"},
		},
		{
			Text:    "Aaron has ebola.",
			Classes: []string{"cdc"},
		},
	}

	classifier = multibayes.NewClassifier()

	// train the classifier
	for _, document := range documents {
		classifier.Add(document.Text, document.Classes)
	}
}

func multibayesHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	text := query.Get("text")
	if len(text) == 0 {
		writeResponse(w, http.StatusBadRequest, nil, "empty text parameter")
		return
	}

	probs := classifier.Posterior(text)
	writeResponse(w, http.StatusOK, probs, "")
	return
}
