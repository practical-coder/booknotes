package main

import (
	"encoding/json"
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
	json.Unmarshal([]byte(w.Body.String()), &notes)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Greater(t, len(notes), 0)
}
