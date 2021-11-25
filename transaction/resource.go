package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/matheusmv/financial-planning-assistant/util"
)

func sendDataAsJSON(data interface{}, rw http.ResponseWriter) {
	_ = json.NewEncoder(rw).Encode(data)
}

func FindAllTransactions(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().Set("content-type", "application/json")

	transactions, _ := GetDB().FindAll()

	rw.WriteHeader(http.StatusOK)
	sendDataAsJSON(transactions, rw)
}

func FindTransactionById(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().Set("content-type", "application/json")

	pathVar := strings.Split(r.URL.String(), "/")[2]

	if pathVar == "" {
		rw.WriteHeader(http.StatusBadRequest)
		sendDataAsJSON(util.NewErrorResponse("invalid parameter"), rw)
		return
	}

	id, err := strconv.ParseInt(pathVar, 10, 64)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		sendDataAsJSON(util.NewErrorResponse("invalid parameter"), rw)
		return
	}

	transaction, err := GetDB().FindById(id)

	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		sendDataAsJSON(util.NewErrorResponse(err.Error()), rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	sendDataAsJSON(transaction, rw)
}

func CreateTransaction(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	rw.Header().Set("content-type", "application/json")

	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		sendDataAsJSON(util.NewErrorResponse(err.Error()), rw)
		return
	}

	ct := r.Header.Get("content-type")

	if ct != "application/json" {
		rw.WriteHeader(http.StatusUnsupportedMediaType)
		msg := fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)
		sendDataAsJSON(util.NewErrorResponse(msg), rw)
		return
	}

	var transaction Transaction
	err = json.Unmarshal(requestBody, &transaction)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		sendDataAsJSON(util.NewErrorResponse(err.Error()), rw)
		return
	}

	transaction, _ = GetDB().Create(transaction)

	rw.WriteHeader(http.StatusCreated)
	sendDataAsJSON(transaction, rw)
}
