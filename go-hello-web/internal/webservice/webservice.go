package webservice

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// SixSidedDie represents any object that can roll a six sided dice for us.
type SixSidedDie interface {
	Roll6() (int, error)
}

// TwentySidedDie represents any object that can roll a twenty sided dice for us.
type TwentySidedDie interface {
	Roll20() (int, error)
}

// New returns a new initialised Webserver.
func New(die SixSidedDie, twenty TwentySidedDie) *WebService {
	return &WebService{
		sixSidedDie:    die,
		twentySidedDie: twenty,
	}
}

// WebService provides a set of web endpoints that handle HTTP requests.
type WebService struct {
	sixSidedDie    SixSidedDie
	twentySidedDie TwentySidedDie
}

// Ping handles an HTTP ping request, returning pong.
func (wb *WebService) Ping(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Pong!\n")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Roll handles an HTTP dice roll request, returning a random number between 1-6 for us.
func (wb *WebService) Roll(w http.ResponseWriter, r *http.Request) {
	roll, err := wb.sixSidedDie.Roll6()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.WriteString(w, fmt.Sprintf("%d\n", roll))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Roll20 handles an HTTP dice roll request, returning a random number between 1-20 for us.
func (wb *WebService) Roll20(w http.ResponseWriter, r *http.Request) {
	roll, err := wb.twentySidedDie.Roll20()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.WriteString(w, fmt.Sprintf("%d\n", roll))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
