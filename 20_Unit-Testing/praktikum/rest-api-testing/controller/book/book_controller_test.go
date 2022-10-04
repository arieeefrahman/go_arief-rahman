package book

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"prak-testing/database"
	"prak-testing/models"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	database.InitTestDB()
	e := echo.New()

	return e
}

func TestGetAllBooksController_Success(t *testing.T) {
	var testCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "success",
			path: "/books",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"books\":",
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetAllBooksController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetBookController_Success(t *testing.T) {
	var testCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "success",
			path: "/books",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"book\":",
		},
	}
	
	e := InitEcho()
	book := database.SeedBook()
	bookId := strconv.Itoa(int(book.ID))
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookId)

		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetBookController_Failed(t *testing.T) {
	var testErrorCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "failed",
			path: "/books",
			expectedStatus: 400,
			expectedBodyStartsWith: "record not found",
		},
	}
	e := InitEcho()
	bookId := "-1"

	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookId)

		if assert.Error(t, GetBookController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}

func TestCreateBookController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/books",
			expectedStatus: http.StatusCreated,
			expectedBodyStartsWith: "{\"book\":",
		},
	}

	e := InitEcho()

	bookInput := models.BookInput {
		Title: "test title",
		Isbn: "123456",
		Description: "test title description",
	}

	jsonbody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonbody)
	req := httptest.NewRequest(http.MethodPost, "/books", bodyReader)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestDeleteBookController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/books",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()
	book := database.SeedBook()
	bookId := strconv.Itoa(int(book.ID))

	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookId)

		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestDeleteBookController_Failed(t *testing.T) {
	var testErrorCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "failed",
			path: "/users",
			expectedStatus: 400,
			expectedBodyStartsWith: "record not found",
		},
	}
	e := InitEcho()
	bookId := "-1"

	req := httptest.NewRequest(http.MethodDelete, "/books", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookId)

		if assert.Error(t, DeleteBookController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}

func TestUpdateBookController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/books",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()
	book := database.SeedBook()
	bookId := strconv.Itoa(int(book.ID))

	bookInput := models.BookInput {
		Title: "test title",
		Isbn: "123456",
		Description: "test title description",
	}

	jsonbody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonbody)
	
	req := httptest.NewRequest(http.MethodPut, "/books", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookId)

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestUpdateBookController_Failed(t *testing.T) {
	var testErrorCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "failed",
			path: "/books",
			expectedStatus: 400,
			expectedBodyStartsWith: "record not found",
		},
	}
	e := InitEcho()
	bookId := "-1"

	bookInput := models.BookInput {
		Title: "test title",
		Isbn: "123456",
		Description: "test title description",
	}

	jsonbody, _ := json.Marshal(&bookInput)
	bodyReader := bytes.NewReader(jsonbody)

	req := httptest.NewRequest(http.MethodPut, "/books", bodyReader)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(bookId)

		if assert.Error(t, UpdateBookController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}