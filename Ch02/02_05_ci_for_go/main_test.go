package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testGUID = "05024756-765e-41a9-89d7-1407436d9a58"

var testData = []Data{
	{
		GUID:     testGUID,
		School:   "Iowa State University",
		Mascot:   "Cy the Cardinal",
		Nickname: "Cyclones",
		Location: "Ames, IA, USA",
		LatLong:  "42.026111,-93.648333",
	},
}

func TestGetAllData(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetDataByID(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/"+testGUID, nil)
	router.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != 200 {
		b, _ := io.ReadAll(res.Body)
		t.Error(res.StatusCode, string(b))
	}

	assert.Equal(t, http.StatusOK, w.Code)

	var result Data

	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, testData[0], result)
}

func TestGetDataByIDNotFound(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/this-is-a-bad-guid", nil)
	router.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != 404 {
		b, _ := io.ReadAll(res.Body)
		t.Error(res.StatusCode, string(b))
	}

	assert.Equal(t, http.StatusNotFound, w.Code)

	var result map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, "Data not found", result["error"])
}
