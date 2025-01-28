package models

import (
	"context"
	"time"

	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/config"
)

type (
	UrlStore struct {
		ShortUrl       string    `bson:"short_url" json:"short_url"`
		LongUrl        string    `bson:"long_url" json:"long_url"`
		ExpirationDate time.Time `bson:"expiration_date" json:"expiration_date"`
		UserID         string    `bson:"user_id" json:"user_id"`
	}

	UrlStoreRepository struct{}
)

func NewUrlStoreRepository() *UrlStoreRepository {
	return &UrlStoreRepository{}
}

func (u *UrlStoreRepository) Insert(us *UrlStore) error {
	collection := config.GetUrlStoreCollection(config.GetDatabase())
	_, err := collection.InsertOne(context.Background(), us)
	return err
}
