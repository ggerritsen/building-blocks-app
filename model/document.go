package model

import "time"

// Document is the object managed by this service
type Document struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreateDate time.Time `json:"createDate"`
}
