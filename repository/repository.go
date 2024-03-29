package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // justifying comment for golint

	"github.com/ggerritsen/building-blocks-app/model"
)

type repository struct {
	db *sql.DB
}

// NewRepositoryWithDb connects to database and uses that connection to create a repository
func NewRepositoryWithDb(host, user, password, dbname string, port int) (*repository, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return NewRepository(db), nil
}

// NewRepository creates a repository based on a db connection
func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

// Close should be called to properly close the repository's db connection
func (r *repository) Close() error {
	return r.db.Close()
}

// CreateTable will create the table in the db that is backing this repository
func (r *repository) CreateTable() error {
	q := "SELECT 1 FROM documents;"
	_, err := r.db.Exec(q)
	if err == nil {
		// table already exists, nothing to do here
		return nil
	}

	q = "CREATE TABLE documents (id SERIAL PRIMARY KEY, name text, description text, createDateTime timestamptz);"
	_, err = r.db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}

// QueryByID will return the Document with the provided id
func (r *repository) QueryByID(id int) (*model.Document, error) {
	q := "SELECT id, name, description, createDateTime FROM documents WHERE ID = $1;"
	row := r.db.QueryRow(q, id)

	d := &model.Document{}
	err := row.Scan(&d.ID, &d.Name, &d.Description, &d.CreateDate)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return d, nil
}

// Insert inserts a document and returns the id of the newly inserted document
func (r *repository) Insert(d *model.Document) (int, error) {
	q := "INSERT INTO documents (name, description, createDateTime) VALUES ($1, $2, $3) RETURNING id;"

	var id int
	if err := r.db.QueryRow(q, d.Name, d.Description, d.CreateDate).Scan(&id); err != nil {
		return -1, err
	}

	return id, nil
}

// Update updates the name of document with the specified id
func (r *repository) Update(id int, updatedName string) error {
	q := "UPDATE documents SET name = $1 WHERE id = $2;"
	_, err := r.db.Exec(q, updatedName, id)
	if err != nil {
		return err
	}
	return nil
}

// Delete removes the document with the specified id
func (r *repository) Delete(id int) error {
	q := "DELETE FROM documents WHERE id = $1;"
	_, err := r.db.Exec(q, id)
	if err != nil {
		return err
	}
	return nil
}
