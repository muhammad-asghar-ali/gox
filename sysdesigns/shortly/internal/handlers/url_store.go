package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type (
	UrlStoreHanlder struct{}
)

func NewUrlStoreHandler() *UrlStoreHanlder {
	return &UrlStoreHanlder{}
}

func (ush *UrlStoreHanlder) Shorten(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (ush *UrlStoreHanlder) Redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
