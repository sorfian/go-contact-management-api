package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/sorfian/go-todo-list/model/web"
	"github.com/stretchr/testify/assert"
)

func TestCreateContactSuccess(t *testing.T) {
	cleanupTestData()

	// Register and login to get a token
	token := registerAndLogin(t, "testcontact1", "password123", "Test Contact User 1")

	// Create contact
	requestBody := web.ContactCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "08123456789",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 201, response.Code)
	assert.Equal(t, "Created", response.Status)

	// Verify contact response
	contactResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, contactResponse["id"])
	assert.Equal(t, "John", contactResponse["first_name"])
	assert.Equal(t, "Doe", contactResponse["last_name"])
	assert.Equal(t, "john.doe@example.com", contactResponse["email"])
	assert.Equal(t, "08123456789", contactResponse["phone"])

	cleanupTestData()
}

func TestCreateContactValidationFailed(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact2", "password123", "Test Contact User 2")

	// Create contact with invalid data
	requestBody := web.ContactCreateRequest{
		FirstName: "",
		LastName:  "",
		Email:     "invalid-email",
		Phone:     "",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	cleanupTestData()
}

func TestCreateContactUnauthorized(t *testing.T) {
	cleanupTestData()

	requestBody := web.ContactCreateRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Phone:     "08123456789",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	// No Authorization header

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode)

	cleanupTestData()
}

func TestGetContactSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact3", "password123", "Test Contact User 3")

	// Create contact first
	contactID := createTestContact(t, token, "Jane", "Doe", "jane.doe@example.com", "08123456789")

	// Get contact
	req := httptest.NewRequest("GET", "/api/contacts/"+contactID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Status)

	contactResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Jane", contactResponse["first_name"])
	assert.Equal(t, "Doe", contactResponse["last_name"])

	cleanupTestData()
}

func TestGetContactNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact4", "password123", "Test Contact User 4")

	req := httptest.NewRequest("GET", "/api/contacts/99999", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

func TestGetAllContactsSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact5", "password123", "Test Contact User 5")

	// Create multiple contacts
	createTestContact(t, token, "John", "Doe", "john@example.com", "08111111111")
	createTestContact(t, token, "Jane", "Smith", "jane@example.com", "08222222222")

	// Get all contacts
	req := httptest.NewRequest("GET", "/api/contacts/", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, 200, response.Code)
	assert.Equal(t, "OK", response.Status)

	contacts, ok := response.Data.([]interface{})
	assert.True(t, ok)
	assert.GreaterOrEqual(t, len(contacts), 2)

	cleanupTestData()
}

func TestUpdateContactSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact6", "password123", "Test Contact User 6")

	// Create contact first
	contactID := createTestContact(t, token, "Original", "Name", "original@example.com", "08111111111")

	// Update contact
	updateBody := web.ContactUpdateRequest{
		FirstName: "Updated",
		LastName:  "Name",
		Email:     "updated@example.com",
		Phone:     "08999999999",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/contacts/"+contactID, bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	contactResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Updated", contactResponse["first_name"])
	assert.Equal(t, "updated@example.com", contactResponse["email"])

	cleanupTestData()
}

func TestUpdateContactNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact7", "password123", "Test Contact User 7")

	updateBody := web.ContactUpdateRequest{
		FirstName: "Updated",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/contacts/99999", bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

func TestDeleteContactSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact8", "password123", "Test Contact User 8")

	// Create contact first
	contactID := createTestContact(t, token, "ToDelete", "User", "delete@example.com", "08111111111")

	// Delete contact
	req := httptest.NewRequest("DELETE", "/api/contacts/"+contactID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify contact is deleted
	req2 := httptest.NewRequest("GET", "/api/contacts/"+contactID, nil)
	req2.Header.Set("Authorization", "Bearer "+token)

	resp2, err := testApp.Test(req2, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp2.StatusCode)

	cleanupTestData()
}

func TestDeleteContactNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testcontact9", "password123", "Test Contact User 9")

	req := httptest.NewRequest("DELETE", "/api/contacts/99999", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

// Helper function to create a test contact and return its ID
func createTestContact(t *testing.T, token, firstName, lastName, email, phone string) string {
	requestBody := web.ContactCreateRequest{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := testApp.Test(req, -1)
	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	json.Unmarshal(body, &response)

	contactResponse := response.Data.(map[string]interface{})
	contactID := contactResponse["id"].(float64)

	return formatContactID(int64(contactID))
}

func formatContactID(id int64) string {
	return strconv.FormatInt(id, 10)
}
