package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/sorfian/go-contact-management-api/model/web"
	"github.com/sorfian/go-contact-management-api/model/web/address"
	"github.com/stretchr/testify/assert"
)

func TestCreateAddressSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress1", "password123", "Test Address User 1")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")

	// Create address
	requestBody := address.AddressCreateRequest{
		Street:     "Jl. Sudirman No. 123",
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Country:    "Indonesia",
		PostalCode: "12345",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/"+contactID+"/addresses/", bytes.NewReader(bodyJSON))
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

	// Verify address response
	addressResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.NotEmpty(t, addressResponse["id"])
	assert.Equal(t, "Jl. Sudirman No. 123", addressResponse["street"])
	assert.Equal(t, "Jakarta", addressResponse["city"])
	assert.Equal(t, "Indonesia", addressResponse["country"])

	cleanupTestData()
}

func TestCreateAddressValidationFailed(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress2", "password123", "Test Address User 2")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")

	// Create address with invalid data
	requestBody := address.AddressCreateRequest{
		Street:     "",
		City:       "",
		Province:   "",
		Country:    "",
		PostalCode: "",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/"+contactID+"/addresses/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode)

	cleanupTestData()
}

func TestCreateAddressContactNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress3", "password123", "Test Address User 3")

	requestBody := address.AddressCreateRequest{
		Street:     "Jl. Sudirman No. 123",
		City:       "Jakarta",
		Province:   "DKI Jakarta",
		Country:    "Indonesia",
		PostalCode: "12345",
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/99999/addresses/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

func TestGetAddressSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress4", "password123", "Test Address User 4")
	contactID := createTestContact(t, token, "Jane", "Doe", "jane@example.com", "08123456789")
	addressID := createTestAddress(t, token, contactID, "Jl. Thamrin", "Jakarta", "DKI Jakarta", "Indonesia", "12340")

	// Get address
	req := httptest.NewRequest("GET", "/api/contacts/"+contactID+"/addresses/"+addressID, nil)
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

	addressResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "Jl. Thamrin", addressResponse["street"])
	assert.Equal(t, "Jakarta", addressResponse["city"])

	cleanupTestData()
}

func TestGetAddressNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress5", "password123", "Test Address User 5")
	contactID := createTestContact(t, token, "Jane", "Doe", "jane@example.com", "08123456789")

	req := httptest.NewRequest("GET", "/api/contacts/"+contactID+"/addresses/99999", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

func TestGetAllAddressesSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress6", "password123", "Test Address User 6")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")

	// Create multiple addresses
	createTestAddress(t, token, contactID, "Jl. Sudirman", "Jakarta", "DKI Jakarta", "Indonesia", "12345")
	createTestAddress(t, token, contactID, "Jl. Thamrin", "Jakarta", "DKI Jakarta", "Indonesia", "12340")

	// Get all addresses
	req := httptest.NewRequest("GET", "/api/contacts/"+contactID+"/addresses/", nil)
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

	addresses, ok := response.Data.([]interface{})
	assert.True(t, ok)
	assert.GreaterOrEqual(t, len(addresses), 2)

	cleanupTestData()
}

func TestUpdateAddressSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress7", "password123", "Test Address User 7")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")
	addressID := createTestAddress(t, token, contactID, "Old Street", "Old City", "Old Province", "Indonesia", "11111")

	// Update address
	updateBody := address.AddressUpdateRequest{
		Street:     "New Street",
		City:       "New City",
		Province:   "New Province",
		Country:    "Indonesia",
		PostalCode: "99999",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/contacts/"+contactID+"/addresses/"+addressID, bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	addressResponse, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "New Street", addressResponse["street"])
	assert.Equal(t, "New City", addressResponse["city"])
	assert.Equal(t, "99999", addressResponse["postal_code"])

	cleanupTestData()
}

func TestUpdateAddressNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress8", "password123", "Test Address User 8")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")

	updateBody := address.AddressUpdateRequest{
		Street: "Updated Street",
	}

	updateJSON, _ := json.Marshal(updateBody)
	req := httptest.NewRequest("PATCH", "/api/contacts/"+contactID+"/addresses/99999", bytes.NewReader(updateJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

func TestDeleteAddressSuccess(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress9", "password123", "Test Address User 9")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")
	addressID := createTestAddress(t, token, contactID, "To Delete", "City", "Province", "Indonesia", "12345")

	// Delete address
	req := httptest.NewRequest("DELETE", "/api/contacts/"+contactID+"/addresses/"+addressID, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	// Verify address is deleted
	req2 := httptest.NewRequest("GET", "/api/contacts/"+contactID+"/addresses/"+addressID, nil)
	req2.Header.Set("Authorization", "Bearer "+token)

	resp2, err := testApp.Test(req2, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp2.StatusCode)

	cleanupTestData()
}

func TestDeleteAddressNotFound(t *testing.T) {
	cleanupTestData()

	token := registerAndLogin(t, "testaddress10", "password123", "Test Address User 10")
	contactID := createTestContact(t, token, "John", "Doe", "john@example.com", "08123456789")

	req := httptest.NewRequest("DELETE", "/api/contacts/"+contactID+"/addresses/99999", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := testApp.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode)

	cleanupTestData()
}

// Helper function to create a test address and return its ID
func createTestAddress(t *testing.T, token, contactID, street, city, province, country, postalCode string) string {
	requestBody := address.AddressCreateRequest{
		Street:     street,
		City:       city,
		Province:   province,
		Country:    country,
		PostalCode: postalCode,
	}

	bodyJSON, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/api/contacts/"+contactID+"/addresses/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, _ := testApp.Test(req, -1)
	body, _ := io.ReadAll(resp.Body)
	var response web.Response
	err := json.Unmarshal(body, &response)
	if err != nil {
		return ""
	}

	addressResponse := response.Data.(map[string]interface{})
	addressID := addressResponse["id"].(float64)

	return strconv.FormatInt(int64(addressID), 10)
}
