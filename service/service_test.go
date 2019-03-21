package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/ggerritsen/building-blocks-app/model"
)

type testRepo struct {
	db        map[int]*model.Document
	idCounter int
}

func (r *testRepo) Insert(d *model.Document) (int, error) {
	r.idCounter = r.idCounter + 1
	id := r.idCounter

	d.ID = id
	r.db[id] = d

	return id, nil
}

func (r *testRepo) QueryByID(id int) (*model.Document, error) {
	return r.db[id], nil
}

func TestStoreAndRead(t *testing.T) {
	now := time.Now()
	nowFunc = func() time.Time {
		return now
	}

	r := &testRepo{db: map[int]*model.Document{}}
	svc := NewDocService(r)

	// test store
	got, err := svc.Store("test")
	if err != nil {
		t.Error(err)
	}
	want := &model.Document{ID: r.idCounter, Name: "test", CreateDate: now}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}

	// test retrieve
	got, err = svc.Retrieve(got.ID)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
