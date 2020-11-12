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

func TestBlockHandler_CreateBlocking(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockIsBlockedEachOther struct {
		input  []int
		result bool
		err    error
	}
	type mockCreateBlockingService struct {
		input *model.BlockingServiceInput
		err   error
	}
	testCases := []struct {
		name                      string
		requestBody               interface{}
		expectedResponseBody      string
		expectedStatus            int
		mockGetRequestorUserID    mockGetUserIDByEmail
		mockGetTargetUserID       mockGetUserIDByEmail
		mockIsBlocked             mockIsBlockedEachOther
		mockCreateBlockingService mockCreateBlockingService
	}{
		{
			name: "Body no data",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: "\"requestor\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "No target email",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
			},
			expectedResponseBody: "\"target\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Requestor email is invalid",
			requestBody: map[string]interface{}{
				"requestor": "abc",
				"target":    "abc@xyz.com",
			},
			expectedResponseBody: "\"requestor\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Target email is invalid",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "abc",
			},
			expectedResponseBody: "\"target\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Two email addresses must be different",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "abc@xyz.com",
			},
			expectedResponseBody: "two email addresses must be different\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Get requestor user ID failed with error",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "get requestor userID failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 0,
				err:    errors.New("get requestor userID failed with error"),
			},
		},
		{
			name: "Get target user ID failed with error",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "get target userID failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetTargetUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 0,
				err:    errors.New("get target userID failed with error"),
			},
		},
		{
			name: "Requestor userID is not exist",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "the requestor does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "Target userID is not exist",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "the target does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetTargetUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "Check exist blocking failed with error",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetTargetUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 11,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: false,
				err:    errors.New("failed with error"),
			},
		},
		{
			name: "Existed blocking",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "target's email have already been blocked by requestor's email\n",
			expectedStatus:       http.StatusPreconditionFailed,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetTargetUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 11,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: true,
				err:    nil,
			},
		},
		{
			name: "Create failed with error",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "create blocking failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetTargetUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 11,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockCreateBlockingService: mockCreateBlockingService{
				input: &model.BlockingServiceInput{
					Requestor: 10,
					Target:    11,
				},
				err: errors.New("create blocking failed with error"),
			},
		},
		{
			name: "Create success",
			requestBody: map[string]interface{}{
				"requestor": "abc@xyz.com",
				"target":    "xyz@abc.com",
			},
			expectedResponseBody: "{\"Success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockGetRequestorUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetTargetUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 11,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockCreateBlockingService: mockCreateBlockingService{
				input: &model.BlockingServiceInput{
					Requestor: 10,
					Target:    11,
				},
				err: nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserService := new(mockUserService)
			mockBlockingService := new(mockBlockingService)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetRequestorUserID.input).
				Return(testCase.mockGetRequestorUserID.result, testCase.mockGetRequestorUserID.err)
			mockUserService.On("GetUserIDByEmail", testCase.mockGetTargetUserID.input).
				Return(testCase.mockGetTargetUserID.result, testCase.mockGetTargetUserID.err)

			if testCase.mockIsBlocked.input != nil {
				mockBlockingService.On("IsExistedBlocking", testCase.mockIsBlocked.input[0], testCase.mockIsBlocked.input[1]).
					Return(testCase.mockIsBlocked.result, testCase.mockIsBlocked.err)
			}
			mockBlockingService.On("CreateBlocking", testCase.mockCreateBlockingService.input).
				Return(testCase.mockCreateBlockingService.err)

			handlers := BlockHandler{
				IUserService:     mockUserService,
				IBlockingService: mockBlockingService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			//When
			req, err := http.NewRequest(http.MethodPost, "/block", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateBlocking)
			handler.ServeHTTP(responseRecorder, req)

			//Then
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}
