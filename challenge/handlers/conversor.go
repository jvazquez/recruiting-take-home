package handlers

import (
	"github/jvazquez/recruiting-take-home/challenge/pkg"
	"log"
	"net/http"
	"strconv"
)

// Conversor translates a number
func ConversorHandler(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("userInput")
	lang := r.URL.Query().Get("lang")
	number, err := strconv.Atoi(userInput)

	if len(lang) == 0 {
		log.Printf("User did not provide a language to convert the input.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Printf("We could not convert the user provided number. User sent %s",
			r.URL.Query().Get("userInput"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	converted, err := pkg.Converter(number, rune(lang[0]))
	if err != nil {
		log.Printf("We couldn't convert the number. See library logs %s",
			r.URL.Query().Get("userInput"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
