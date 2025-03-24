package models

import "time"

// Note struct
type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

// In-memory store
var NoteStore = make(map[string]Note)
var IDCounter int = 0
