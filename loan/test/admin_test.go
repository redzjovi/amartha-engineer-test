package test

import (
	"encoding/json"
	"fmt"
	"loan/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestAdminLoanApproveUnauthorized(t *testing.T) {
	ClearAll()

	request := httptest.NewRequest(http.MethodPost, "/api/admin/loan/1/approve", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestAdminLoanApproveBadRequest(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)

	request := httptest.NewRequest(http.MethodPost, "/api/admin/loan/1/approve", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestAdminLoanApproveNotFound(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)

	requestBody := model.LoanApproveRequest{
		PictureProof: faker.URL(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/admin/loan/1/approve", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestAdminLoanApproveConflict(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)
	loan := CreateLoanApproved(t)

	requestBody := model.LoanApproveRequest{
		PictureProof: faker.URL(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/admin/loan/%v/approve", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
}

func TestAdminLoanApprove(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)
	loan := CreateLoanProposed(t)

	requestBody := model.LoanApproveRequest{
		PictureProof: faker.URL(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/admin/loan/%v/approve", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}

func TestAdminLoanDisburseUnauthorized(t *testing.T) {
	ClearAll()

	request := httptest.NewRequest(http.MethodPost, "/api/admin/loan/1/disburse", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestAdminLoanDisburseBadRequest(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)

	request := httptest.NewRequest(http.MethodPost, "/api/admin/loan/1/disburse", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestAdminLoanDisburseNotFound(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)

	requestBody := model.LoanDisburseRequest{
		AgreementLetter: faker.URL(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/admin/loan/1/disburse", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestAdminLoanDisburseConflict(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)
	loan := CreateLoanApproved(t)

	requestBody := model.LoanDisburseRequest{
		AgreementLetter: faker.URL(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/admin/loan/%v/disburse", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
}

func TestAdminLoanDisburse(t *testing.T) {
	ClearAll()
	user := CreateUserAdmin(t)
	loan := CreateLoanInvested(t)

	requestBody := model.LoanDisburseRequest{
		AgreementLetter: faker.URL(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/admin/loan/%v/disburse", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
