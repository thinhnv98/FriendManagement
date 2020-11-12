package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/require"
)

func TestUserHandler_CreateUser(t *testing.T) {
	type mockIsUserExisted struct {
		input  string
		result bool
		err    error
	}

	type mockCreateUserService struct {
		input *model.UserServiceInput
		err   error
	}
	testCases := []struct {
		name                   string
		requestBody            interface{}
		expectedResponseBody   string
		expectedResponseStatus int
		mockIsUserExisted      mockIsUserExisted
		mockCreateUserService  mockCreateUserService
	}{
		{
			name: "Validate request body failed",
			requestBody: map[string]interface{}{
				"email": "",
			},
			expectedResponseBody:   "\"email\" is required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name: "Email's format is not valid",
			requestBody: map[string]interface{}{
				"Email": "abc",
			},
			expectedResponseBody:   "\"email\"'s format is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name: "User email existed",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody:   "this email address existed\n",
			expectedResponseStatus: http.StatusAlreadyReported,
			mockIsUserExisted: mockIsUserExisted{
				input:  "abc@xyz.com",
				result: true,
				err:    nil,
			},
		},
		{
			name: "Check existed user's email with error",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody:   "check existed user's email process error\n",
			expectedResponseStatus: http.StatusInternalServerError,
			mockIsUserExisted: mockIsUserExisted{
				input:  "abc@xyz.com",
				result: false,
				err:    errors.New("check existed user's email process error"),
			},
		},
		{
			name: "Call services return with error",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody:   "services error\n",
			expectedResponseStatus: http.StatusInternalServerError,
			mockIsUserExisted: mockIsUserExisted{
				input:  "abc@xyz.com",
				result: false,
				err:    nil,
			},
			mockCreateUserService: mockCreateUserService{
				input: &model.UserServiceInput{
					Email: "abc@xyz.com",
				},
				err: errors.New("services error"),
			},
		},
		{
			name: "Everything success",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody:   "{\"Success\":true}\n",
			expectedResponseStatus: http.StatusOK,
			mockIsUserExisted: mockIsUserExisted{
				input:  "abc@xyz.com",
				result: false,
				err:    nil,
			},
			mockCreateUserService: mockCreateUserService{
				input: &model.UserServiceInput{
					Email: "abc@xyz.com",
				},
				err: nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockService := new(mockUserService)

			mockService.On("IsExistedUser", testCase.mockIsUserExisted.input).
				Return(testCase.mockIsUserExisted.result, testCase.mockIsUserExisted.err)
			mockService.On("CreateUser", testCase.mockCreateUserService.input).
				Return(testCase.mockCreateUserService.err)

			handlers := UserHandler{
				IUserService: mockService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			//When
			req, err := http.NewRequest(http.MethodPost, "/user", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateUser)
			handler.ServeHTTP(responseRecorder, req)

			//Then
			require.Equal(t, testCase.expectedResponseStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
