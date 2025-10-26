package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/user"
	"github.com/sorfian/go-contact-management-api/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	testDB             *gorm.DB
	testApp            *fiber.App
	testUserRepository repository.UserRepository
)

func setupTestApp() {
	// Initialize app with all dependencies using Wire
	deps := InitializeTestApp()
	testApp = deps.App
	testDB = deps.DB
	testUserRepository = deps.UserRepository
}

func cleanupTestData() {
	testDB.Exec("DELETE FROM users WHERE username LIKE 'test%'")
}

func TestMain(m *testing.M) {
	setupTestApp()
	m.Run()
}

func TestRegisterSuccess(t *testing.T) {
	cleanupTestData()

	requestBody := user.UserRegisterRequest{
		Username: "testuser1",
		Password: "password123",
		Name:     "Test User 1",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, "Created", response.Status)

	// Verify token response
	tokenResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, tokenResponse["token"])
	assert.NotEmpty(t, tokenResponse["token_exp"])

	cleanupTestData()
}

func TestRegisterValidationFailed(t *testing.T) {
	cleanupTestData()

	requestBody := user.UserRegisterRequest{
		Username: "te", // Too short
		Password: "pa", // Too short
		Name:     "T",  // Too short
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 400, response.Code)
	assert.Equal(t, "Bad Request", response.Status)
	assert.Contains(t, response.Data.(string), "Validation failed")

	cleanupTestData()
}

func TestRegisterDuplicateUsername(t *testing.T) {
	cleanupTestData()

	// Register first user
	requestBody := user.UserRegisterRequest{
		Username: "testuser2",
		Password: "password123",
		Name:     "Test User 2",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")

	_, err := testApp.Test(req, -1)
	if err != nil {
		return
	}

	// Try to register with the same username
	req2 := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(bodyJSON))
	req2.Header.Set("Content-Type", "application/json")

	resp, err := testApp.Test(req2, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusConflict, resp.StatusCode)

	cleanupTestData()
}

func TestLoginSuccess(t *testing.T) {
	cleanupTestData()

	// Register user first
	registerBody := user.UserRegisterRequest{
		Username: "testuser3",
		Password: "password123",
		Name:     "Test User 3",
	}

	registerJSON, _ := json.Marshal(registerBody)
	regReq := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(registerJSON))
	regReq.Header.Set("Content-Type", "application/json")
	_, err := testApp.Test(regReq, -1)
	if err != nil {
		return
	}

	// Login
	loginBody := user.UserLoginRequest{
		Username: "testuser3",
		Password: "password123",
	}

	loginJSON, _ := json.Marshal(loginBody)
	req := httptest.NewRequest("POST", "/api/users/login", bytes.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Status)

	tokenResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, tokenResponse["token"])
	assert.NotEmpty(t, tokenResponse["token_exp"])

	cleanupTestData()
}

func TestLoginWrongUsername(t *testing.T) {
	cleanupTestData()

	loginBody := user.UserLoginRequest{
		Username: "nonexistentuser",
		Password: "password123",
	}

	loginJSON, _ := json.Marshal(loginBody)
	req := httptest.NewRequest("POST", "/api/users/login", bytes.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 404, response.Code)
	assert.Equal(t, "Not Found", response.Status)

	cleanupTestData()
}

func TestLoginWrongPassword(t *testing.T) {
	cleanupTestData()

	// Register user first
	registerBody := user.UserRegisterRequest{
		Username: "testuser4",
		Password: "password123",
		Name:     "Test User 4",
	}

	registerJSON, _ := json.Marshal(registerBody)
	regReq := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(registerJSON))
	regReq.Header.Set("Content-Type", "application/json")
	_, err := testApp.Test(regReq, -1)
	if err != nil {
		return
	}

	// Login with the wrong password
	loginBody := user.UserLoginRequest{
		Username: "testuser4",
		Password: "wrongpassword",
	}

	loginJSON, _ := json.Marshal(loginBody)
	req := httptest.NewRequest("POST", "/api/users/login", bytes.NewReader(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)

	cleanupTestData()
}

func TestGetCurrentUserSuccess(t *testing.T) {
	cleanupTestData()

	// Register and login to get a token
	token := registerAndLogin(t, "testuser5", "password123", "Test User 5")

	// Get current user
	req := httptest.NewRequest("GET", "/api/users/current", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Status)

	userData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "testuser5", userData["username"])
	assert.Equal(t, "Test User 5", userData["name"])

	cleanupTestData()
}

func TestGetCurrentUserUnauthorized(t *testing.T) {
	cleanupTestData()

	req := httptest.NewRequest("GET", "/api/users/current", nil)
	// No Authorization header

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 401, response.Code)
	assert.Equal(t, "Unauthorized", response.Status)

	cleanupTestData()
}

func TestGetCurrentUserInvalidToken(t *testing.T) {
	cleanupTestData()

	req := httptest.NewRequest("GET", "/api/users/current", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken123")

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	cleanupTestData()
}

func TestUpdateCurrentUserSuccess(t *testing.T) {
	cleanupTestData()

	// Register and login to get a token
	token := registerAndLogin(t, "testuser6", "password123", "Test User 6")

	// Update user
	updateBody := user.UserUpdateRequest{
		Name:     "Updated Name",
		Password: "newpassword123",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/users/current", bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Status)

	userData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "testuser6", userData["username"])
	assert.Equal(t, "Updated Name", userData["name"])

	cleanupTestData()
}

func TestUpdateCurrentUserOnlyName(t *testing.T) {
	cleanupTestData()

	// Register and login to get a token
	token := registerAndLogin(t, "testuser7", "password123", "Test User 7")

	// Update only name
	updateBody := user.UserUpdateRequest{
		Name: "Only Name Updated",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/users/current", bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	userData, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Only Name Updated", userData["name"])

	cleanupTestData()
}

func TestUpdateCurrentUserUnauthorized(t *testing.T) {
	cleanupTestData()

	updateBody := user.UserUpdateRequest{
		Name: "Should Fail",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/users/current", bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	// No Authorization header

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	cleanupTestData()
}

func TestLogoutSuccess(t *testing.T) {
	cleanupTestData()

	// Register and login to get a token
	token := registerAndLogin(t, "testuser8", "password123", "Test User 8")

	// Logout
	req := httptest.NewRequest("DELETE", "/api/users/current", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Status)

	// Verify token is invalidated
	req2 := httptest.NewRequest("GET", "/api/users/current", nil)
	req2.Header.Set("Authorization", "Bearer "+token)

	resp2, err := testApp.Test(req2, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp2.StatusCode)

	cleanupTestData()
}

func TestLogoutUnauthorized(t *testing.T) {
	cleanupTestData()

	req := httptest.NewRequest("DELETE", "/api/users/current", nil)
	// No Authorization header

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)

	cleanupTestData()
}

// Helper function to register and login a user, returns token
func registerAndLogin(t *testing.T, username, password, name string) string {
	// Register
	registerBody := user.UserRegisterRequest{
		Username: username,
		Password: password,
		Name:     name,
	}

	registerJSON, _ := json.Marshal(registerBody)
	regReq := httptest.NewRequest("POST", "/api/users/register", bytes.NewReader(registerJSON))
	regReq.Header.Set("Content-Type", "application/json")
	regResp, _ := testApp.Test(regReq, -1)

	regBody, _ := io.ReadAll(regResp.Body)
	var regResponse web.Response
	err := json.Unmarshal(regBody, &regResponse)
	if err != nil {
		return ""
	}

	tokenResponse, ok := regResponse.Data.(map[string]interface{})
	if !ok {
		t.Fatal("Failed to get token from registration")
	}

	token, ok := tokenResponse["token"].(string)
	if !ok {
		t.Fatal("Failed to extract token string")
	}

	return token
}

// Helper function to verify a user in a database
//func getUserFromDB(username string) (*domain.User, error) {
//	var user domain.User
//	err := testDB.Where("username = ?", username).First(&user).Error
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
