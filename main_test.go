package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/practical-coder/booknotes/models/book"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	e := SetEngine()
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")

	req, err := http.NewRequest("GET", "/notes", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	var notes []book.Note
	json.Unmarshal(w.Body.Bytes(), &notes)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Greater(t, len(notes), 0)
}

func TestCreateDelete(t *testing.T) {
	e := SetEngine()
	os.Setenv("MONGODB_URI", "mongodb://localhost:27017")

	note := book.Note{
		Title: "Testing in Go",
	}

	jsonNote, _ := json.Marshal(note)
	req, err := http.NewRequest("POST", "/notes", bytes.NewReader(jsonNote))
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	e.ServeHTTP(w, req)

	var respNote book.Note
	json.Unmarshal(w.Body.Bytes(), &respNote)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, note.Title, respNote.Title)

	endpoint := fmt.Sprintf("/notes/%s", respNote.UUID)
	req, err = http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		t.Error(err)
	}

	w = httptest.NewRecorder()
	e.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
