package user

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

func TestGetAllUsersController_Success(t *testing.T) {
	var testCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "success",
			path: "/users",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetAllUsersController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetUserController_Success(t *testing.T) {
	var testCases = []struct {
		name string
		path string
		expectedStatus int
		expectedBodyStartsWith string
	}{
		{
			name: "success",
			path: "/users",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}
	
	e := InitEcho()
	user := database.SeedUser()
	userId := strconv.Itoa(int(user.ID))
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestGetUserController_Failed(t *testing.T) {
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
	userId := "-1"

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.Error(t, GetUserController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}

func TestCreateUserController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/users",
			expectedStatus: http.StatusCreated,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()

	userInput := models.UserInput {
		Name: "test",
		Email: "test@gmail.com",
		Password: "test123",
	}

	jsonbody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonbody)
	
	req := httptest.NewRequest(http.MethodPost, "/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}

}

func TestDeleteUserController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/users",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()
	user := database.SeedUser()
	userId := strconv.Itoa(int(user.ID))

	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestDeleteUserController_Failed(t *testing.T) {
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
	userId := "-1"

	req := httptest.NewRequest(http.MethodDelete, "/users", nil)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.Error(t, DeleteUserController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}

func TestUpdateUserController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/users",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()
	user := database.SeedUser()
	userId := strconv.Itoa(int(user.ID))

	userInput := models.UserInput {
		Name: "test updt1",
		Email: "test@gmail.com",
		Password: "test123",
	}

	jsonbody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonbody)
	
	req := httptest.NewRequest(http.MethodPut, "/users", bodyReader)
	rec := httptest.NewRecorder()

	req.Header.Add("Content-Type", "application/json")

	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestUpdateUserController_Failed(t *testing.T) {
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
	userId := "-1"

	userInput := models.UserInput {
		Name: "test updt1",
		Email: "test@gmail.com",
		Password: "test123",
	}

	jsonbody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonbody)

	req := httptest.NewRequest(http.MethodPut, "/users", bodyReader)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.Error(t, UpdateUserController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}

func TestLoginUserController_Success(t *testing.T) {
	var testCases = []struct {
		name 					string
		path 					string
		expectedStatus 			int
		expectedBodyStartsWith 	string
	}{
		{
			name: "success",
			path: "/users",
			expectedStatus: http.StatusOK,
			expectedBodyStartsWith: "{\"message\":",
		},
	}

	e := InitEcho()

	userInput := models.UserInput {
		Email: "test@gmail.com",
		Password: "test123",
	}

	jsonbody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonbody)
	
	req := httptest.NewRequest(http.MethodPost, "/users", bodyReader)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			assert.True(t, strings.HasPrefix(body, testCase.expectedBodyStartsWith))
		}
	}
}

func TestLoginUserController_Failed(t *testing.T) {
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
	userId := "-1"

	userInput := models.UserInput {
		Email: "testfail@gmail.com",
		Password: "testfail123",
	}

	jsonbody, _ := json.Marshal(&userInput)
	bodyReader := bytes.NewReader(jsonbody)

	req := httptest.NewRequest(http.MethodPost, "/users", bodyReader)
	rec := httptest.NewRecorder()
	req.Header.Add("Content-Type", "application/json")
	c := e.NewContext(req, rec)
	
	for _, testErrorCase := range testErrorCases {
		c.SetPath(testErrorCase.path)
		c.SetParamNames("id")
		c.SetParamValues(userId)

		if assert.Error(t, GetUserController(c)) {
			if testErrorCase.expectedStatus != 400 {
				t.Error("error code")
			}
		}
	}
}