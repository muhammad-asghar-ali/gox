package services

import (
	"fmt"
	"matrix/internal/pb"
	"sync"

	"github.com/jinzhu/copier"
)

type (
	LaptopStore interface {
		Save(laptop *pb.Laptop) error
		Find(id string) (*pb.Laptop, error)
	}
)

type (
	InMemoryLaptopStore struct {
		mutex sync.RWMutex
		data  map[string]*pb.Laptop
	}
)

func NewInMemoryLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(lp *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[lp.Id] != nil {
		return ErrAlreadyExists
	}

	o := &pb.Laptop{}
	err := copier.Copy(o, lp)
	if err != nil {
		return err
	}

	store.data[o.Id] = o

	return nil
}

func (store *InMemoryLaptopStore) Find(id string) (*pb.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, nil
	}

	return deep_copy(laptop)
}

func deep_copy(laptop *pb.Laptop) (*pb.Laptop, error) {
	other := &pb.Laptop{}

	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}

	return other, nil
}
