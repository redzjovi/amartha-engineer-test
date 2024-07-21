package test

import (
	"encoding/json"
	"fmt"
	"io"
	"loan/internal/entity"
	"loan/internal/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserAuthLogoutUnauthorized(t *testing.T) {
	ClearAll()
	TestGuestAuthLogin(t)

	request := httptest.NewRequest(http.MethodDelete, "/api/user/auth", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "wrong")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)
}

func TestUserAuthLogout(t *testing.T) {
	ClearAll()
	TestGuestAuthLogin(t)

	user := new(entity.User)
	err := db.Where("email = ?", "email@email.com").Take(user).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/user/auth", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}

func TestUserLoanProposeUnauthorized(t *testing.T) {
	ClearAll()

	request := httptest.NewRequest(http.MethodPost, "/api/user/loan/propose", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestUserLoanProposeBadRequest(t *testing.T) {
	ClearAll()
	user := CreateUser(t)

	request := httptest.NewRequest(http.MethodPost, "/api/user/loan/propose", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestUserLoanPropose(t *testing.T) {
	ClearAll()
	user := CreateUser(t)

	requestBody := model.LoanProposeRequest{
		PrincipalAmount: 100000000,
		Rate:            0.1,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/user/loan/propose", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}

func TestUserLoanInvestUnauthorized(t *testing.T) {
	ClearAll()

	request := httptest.NewRequest(http.MethodPost, "/api/user/loan/1/invest", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
}

func TestUserLoanInvestBadRequest(t *testing.T) {
	ClearAll()
	user := CreateUser(t)

	request := httptest.NewRequest(http.MethodPost, "/api/user/loan/1/invest", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
}

func TestUserLoanInvestNotFound(t *testing.T) {
	ClearAll()
	user := CreateUser(t)

	requestBody := model.LoanInvestRequest{
		Amount: 1000000,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/user/loan/1/invest", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUserLoanInvestConflict(t *testing.T) {
	ClearAll()
	user := CreateUser(t)
	loan := CreateLoanProposed(t)

	requestBody := model.LoanInvestRequest{
		Amount: 1000000,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/user/loan/%v/invest", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, response.StatusCode)
}

func TestUserLoanInvestNotAcceptableFullfilled(t *testing.T) {
	ClearAll()
	user := CreateUser(t)
	loan := CreateLoanApprovedAndFullfilled(t)

	requestBody := model.LoanInvestRequest{
		Amount: 1000000,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/user/loan/%v/invest", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotAcceptable, response.StatusCode)
}

func TestUserLoanInvestNotAcceptableInvestAmountInvalid(t *testing.T) {
	ClearAll()
	user := CreateUser(t)
	loan := CreateLoanApproved(t)

	requestBody := model.LoanInvestRequest{
		Amount: 200000000,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/user/loan/%v/invest", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotAcceptable, response.StatusCode)
}

func TestUserLoanInvest(t *testing.T) {
	ClearAll()
	user := CreateUser(t)
	loan := CreateLoanApproved(t)

	requestBody := model.LoanInvestRequest{
		Amount: 1000000,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api/user/loan/%v/invest", loan.ID), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
