package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"simpleserver.com/m/models"
)

// Create a new note
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	note.CreatedOn = time.Now()
	models.IDCounter++
	key := strconv.Itoa(models.IDCounter)
	models.NoteStore[key] = note

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

// Get all notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	for _, v := range models.NoteStore {
		notes = append(notes, v)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

// Update a note
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var updatedNote models.Note
	err := json.NewDecoder(r.Body).Decode(&updatedNote)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if existingNote, ok := models.NoteStore[key]; ok {
		updatedNote.CreatedOn = existingNote.CreatedOn // Preserve the original creation date
		models.NoteStore[key] = updatedNote
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Note not found", http.StatusNotFound)
	}
}

// Delete a note
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	if _, ok := models.NoteStore[key]; ok {
		delete(models.NoteStore, key)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Note not found", http.StatusNotFound)
	}
}
