package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/models"
	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/utils"
)

type (
	UrlStoreHanlder struct {
		repo *models.UrlStoreRepository
	}

	ShortenRequest struct {
		LongUrl        string     `json:"long_url"`
		ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	}
)

func NewUrlStoreHandler() *UrlStoreHanlder {
	return &UrlStoreHanlder{}
}

func (ush *UrlStoreHanlder) Shorten(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &ShortenRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	expiration := time.Now().Add(30 * 24 * time.Hour) // Default to 30 days
	if req.ExpirationDate != nil {
		expiration = *req.ExpirationDate
	}

	u := &models.UrlStore{
		ShortUrl:       utils.GenerateShortURL(req.LongUrl),
		LongUrl:        req.LongUrl,
		ExpirationDate: expiration,
		UserID:         "",
	}

	if err := ush.repo.Insert(u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}

func (ush *UrlStoreHanlder) Redirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
