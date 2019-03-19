package service

import (
	"fmt"
	"time"

	"github.com/ggerritsen/building-blocks-app/model"
)

type repo interface {
	QueryByID(id int) (*model.Document, error)
	Insert(d *model.Document) (id int, e error)
}

type docService struct {
	r repo
}

// NewDocService creates a new Document service
func NewDocService(r repo) *docService {
	return &docService{r}
}

func (svc *docService) Read(id int) (*model.Document, error) {
	return svc.r.QueryByID(id)
}

func (svc *docService) Store(name string) (*model.Document, error) {
	if name == "" {
		return nil, fmt.Errorf("No name provided")
	}

	d := &model.Document{Name: name, CreateDate: time.Now()}
	id, err := svc.r.Insert(d)
	if err != nil {
		return nil, err
	}

	d.ID = id
	return d, nil
}
