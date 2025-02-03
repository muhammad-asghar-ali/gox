package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/config"
	"github.com/muhammad-asghar-ali/go/sysdesigns/shortly/internal/utils"
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

func (u *UrlStoreRepository) FindByCode(code string) (string, error) {
	collection := config.GetUrlStoreCollection(config.GetDatabase())
	filter := bson.M{"short_url": code, "expiration_date": bson.M{"$gt": time.Now()}}

	e := &UrlStore{}
	if err := collection.FindOne(context.Background(), filter).Decode(e); err != nil {
		return "", err
	}

	if err := utils.CheckExpiration(&e.ExpirationDate); err != nil {
		return "", err
	}

	return e.LongUrl, nil
}
