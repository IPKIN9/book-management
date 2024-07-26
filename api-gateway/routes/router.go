package routes

import (
	"api-gateway/usecases"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	bookUseCase usecases.BookUsecase
}

func NewRouter(bookUsecase usecases.BookUsecase) *mux.Router {
	r := mux.NewRouter()
	router := &Router{bookUseCase: bookUsecase}

	r.HandleFunc("/books/{id}", router.getBookHandler).Methods("GET")

	return r
}

func (router *Router) getBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	ctx := r.Context()
	book, err := router.bookUseCase.GetBook(ctx, id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Could not fetch book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}
