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

func TestFriendHandler_CreateFriend(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockIsExistedFriend struct {
		input  []int
		result bool
		err    error
	}
	type mockIsBlockedEachOther struct {
		input  []int
		result bool
		err    error
	}
	type mockCreateFriendService struct {
		input *model.FriendsServiceInput
		err   error
	}
	testCases := []struct {
		name                    string
		requestBody             interface{}
		expectedResponseBody    string
		expectedStatus          int
		mockGetFirstUserID      mockGetUserIDByEmail
		mockGetSecondUserID     mockGetUserIDByEmail
		mockIsExistedFriend     mockIsExistedFriend
		mockIsBlocked           mockIsBlockedEachOther
		mockCreateFriendService mockCreateFriendService
	}{
		{
			name: "Decode failed",
			requestBody: map[string]interface{}{
				"friends": "abc",
			},
			expectedResponseBody: "json: cannot unmarshal string into Go struct field FriendConnectionRequest.friends of type []string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "No data request body",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: "\"friends\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "First email format invalid",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc",
					"xyz@abc.com",
				},
			},
			expectedResponseBody: "first \"email\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Second email format invalid",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"xyz",
				},
			},
			expectedResponseBody: "second \"email\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Not enough email addresses",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
				},
			},
			expectedResponseBody: "needs exactly two email addresses\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Two email address must be different",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"xyz@abc.com",
				},
			},
			expectedResponseBody: "two email addresses must be different\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Get UserID failed with error",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "get UserID failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 0,
				err:    errors.New("get UserID failed with error"),
			},
		},
		{
			name: "First email address's UserID is not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "the first email does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "Second email address's UserID is not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "the second email does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "Check existed friend connection failed with error",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "check existed connection friend failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 11,
				err:    nil,
			},
			mockIsExistedFriend: mockIsExistedFriend{
				input:  []int{10, 11},
				result: false,
				err:    errors.New("check existed connection friend failed with error"),
			},
		},
		{
			name: "Friend connection existed",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "friend connection existed\n",
			expectedStatus:       http.StatusAlreadyReported,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 11,
				err:    nil,
			},
			mockIsExistedFriend: mockIsExistedFriend{
				input:  []int{10, 11},
				result: true,
				err:    nil,
			},
		},
		{
			name: "Check is blocked each other failed with error",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "check failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 11,
				err:    nil,
			},
			mockIsExistedFriend: mockIsExistedFriend{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: false,
				err:    errors.New("check failed with error"),
			},
		},
		{
			name: "Email addresses blocked each other",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "emails blocked each other\n",
			expectedStatus:       http.StatusPreconditionFailed,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 11,
				err:    nil,
			},
			mockIsExistedFriend: mockIsExistedFriend{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: true,
				err:    nil,
			},
		},
		{
			name: "Create friend failed with error",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "create failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 11,
				err:    nil,
			},
			mockIsExistedFriend: mockIsExistedFriend{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockCreateFriendService: mockCreateFriendService{
				input: &model.FriendsServiceInput{
					FirstID:  10,
					SecondID: 11,
				},
				err: errors.New("create failed with error"),
			},
		},
		{
			name: "Create friend connection success",
			requestBody: map[string]interface{}{
				"friends": []string{
					"xyz@abc.com",
					"abc@xyz.com",
				},
			},
			expectedResponseBody: "{\"Success\":true}\n",
			expectedStatus:       http.StatusOK,
			mockGetFirstUserID: mockGetUserIDByEmail{
				input:  "xyz@abc.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 11,
				err:    nil,
			},
			mockIsExistedFriend: mockIsExistedFriend{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockIsBlocked: mockIsBlockedEachOther{
				input:  []int{10, 11},
				result: false,
				err:    nil,
			},
			mockCreateFriendService: mockCreateFriendService{
				input: &model.FriendsServiceInput{
					FirstID:  10,
					SecondID: 11,
				},
				err: nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockFriendService := new(mockFriendService)
			mockUserService := new(mockUserService)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetFirstUserID.input).
				Return(testCase.mockGetFirstUserID.result, testCase.mockGetFirstUserID.err)
			mockUserService.On("GetUserIDByEmail", testCase.mockGetSecondUserID.input).
				Return(testCase.mockGetSecondUserID.result, testCase.mockGetSecondUserID.err)

			mockFriendService.On("CreateFriend", testCase.mockCreateFriendService.input).
				Return(testCase.mockCreateFriendService.err)
			if testCase.mockIsExistedFriend.input != nil {
				mockFriendService.On("IsExistedFriend", testCase.mockIsExistedFriend.input[0], testCase.mockIsExistedFriend.input[1]).
					Return(testCase.mockIsExistedFriend.result, testCase.mockIsExistedFriend.err)
			}
			if testCase.mockIsBlocked.input != nil {
				mockFriendService.On("IsBlockedByOtherEmail", testCase.mockIsBlocked.input[0], testCase.mockIsBlocked.input[1]).
					Return(testCase.mockIsBlocked.result, testCase.mockIsBlocked.err)
			}

			handlers := FriendHandler{
				IUserService:    mockUserService,
				IFriendServices: mockFriendService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			//When
			req, err := http.NewRequest(http.MethodPost, "/friend", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.CreateFriend)
			handler.ServeHTTP(responseRecorder, req)

			//Then
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}

func TestFriendHandler_GetFriendListByEmail(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockGetFriendsList struct {
		input  int
		result []string
		err    error
	}
	testCases := []struct {
		name                 string
		requestBody          interface{}
		expectedResponseBody string
		expectedStatus       int
		mockGetUserIDByEmail mockGetUserIDByEmail
		mockGetFriendList    mockGetFriendsList
	}{
		{
			name: "Validate request body failed",
			requestBody: map[string]interface{}{
				"email": "",
			},
			expectedResponseBody: "\"Email\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Email request format is not valid",
			requestBody: map[string]interface{}{
				"email": "abc",
			},
			expectedResponseBody: "\"email\" format is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Get UserID failed with error",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody: "get UserID failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 0,
				err:    errors.New("get UserID failed with error"),
			},
		},
		{
			name: "Get UserID return not exist",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody: "email does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "Get friend list failed with error",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody: "get friend list failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 1,
				err:    nil,
			},
			mockGetFriendList: mockGetFriendsList{
				input:  1,
				result: nil,
				err:    errors.New("get friend list failed with error"),
			},
		},
		{
			name: "Success request",
			requestBody: map[string]interface{}{
				"email": "abc@xyz.com",
			},
			expectedResponseBody: "{\"success\":true,\"friends\":[\"xyz@gmail.com\"],\"count\":1}\n",
			expectedStatus:       http.StatusOK,
			mockGetUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 1,
				err:    nil,
			},
			mockGetFriendList: mockGetFriendsList{
				input: 1,
				result: []string{
					"xyz@gmail.com",
				},
				err: nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserService := new(mockUserService)
			mockFriendService := new(mockFriendService)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetUserIDByEmail.input).
				Return(testCase.mockGetUserIDByEmail.result, testCase.mockGetUserIDByEmail.err)

			mockFriendService.On("GetFriendListByID", testCase.mockGetFriendList.input).
				Return(testCase.mockGetFriendList.result, testCase.mockGetFriendList.err)

			handlers := FriendHandler{
				IUserService:    mockUserService,
				IFriendServices: mockFriendService,
			}
			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			//When
			req, err := http.NewRequest(http.MethodGet, "/friends", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetFriendListByEmail)
			handler.ServeHTTP(responseRecorder, req)

			//Then
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}

func TestFriendHandler_GetCommonFriendListByEmails(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockGetCommonFriendList struct {
		input  []int
		result []string
		err    error
	}
	testCases := []struct {
		name                       string
		requestBody                interface{}
		expectedResponseBody       string
		expectedStatus             int
		mockGetFirstUserIDByEmail  mockGetUserIDByEmail
		mockGetSecondUserIDByEmail mockGetUserIDByEmail
		mockGetCommonFriendList    mockGetCommonFriendList
	}{
		{
			name: "Validate request body failed",
			requestBody: map[string]interface{}{
				"friends": "././",
			},
			expectedResponseBody: "json: cannot unmarshal string into Go struct field FriendGetCommonFriendsRequest.friends of type []string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "No data request body",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: "\"friends\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Not enough email",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc",
				},
			},
			expectedResponseBody: "needs exactly two email addresses\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "First email format invalid",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc",
					"xyz",
				},
			},
			expectedResponseBody: "first \"email\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Second email format invalid",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"xyz",
				},
			},
			expectedResponseBody: "second \"email\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Get UserID by email failed with error",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"xyz@gmail.com",
				},
			},
			expectedResponseBody: "get userID failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetFirstUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 0,
				err:    errors.New("get userID failed with error"),
			},
		},
		{
			name: "First email is not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"xyz@gmail.com",
				},
			},
			expectedResponseBody: "first email does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetFirstUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 0,
				err:    nil,
			},
			mockGetCommonFriendList: mockGetCommonFriendList{},
		},
		{
			name: "Second email is not exist",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"xyz@gmail.com",
				},
			},
			expectedResponseBody: "second email does not exist\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetFirstUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserIDByEmail: mockGetUserIDByEmail{
				input:  "xyz@gmail.com",
				result: 0,
				err:    nil,
			},
		},
		{
			name: "Two email must different each other",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"abc@gmail.com",
				},
			},
			expectedResponseBody: "two email addresses must be different\n",
			expectedStatus:       http.StatusBadRequest,
			mockGetFirstUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 10,
				err:    nil,
			},
		},
		{
			name: "Get common friend list failed with error",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"xyz@gmail.com",
				},
			},
			expectedResponseBody: "get common friend list failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetFirstUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserIDByEmail: mockGetUserIDByEmail{
				input:  "xyz@gmail.com",
				result: 11,
				err:    nil,
			},
			mockGetCommonFriendList: mockGetCommonFriendList{
				input:  []int{10, 11},
				result: nil,
				err:    errors.New("get common friend list failed with error"),
			},
		},
		{
			name: "Get Success",
			requestBody: map[string]interface{}{
				"friends": []string{
					"abc@gmail.com",
					"xyz@gmail.com",
				},
			},
			expectedResponseBody: "{\"success\":true,\"friends\":[\"abc@xyz.com\",\"xyz@abc.com\"],\"count\":2}\n",
			expectedStatus:       http.StatusOK,
			mockGetFirstUserIDByEmail: mockGetUserIDByEmail{
				input:  "abc@gmail.com",
				result: 10,
				err:    nil,
			},
			mockGetSecondUserIDByEmail: mockGetUserIDByEmail{
				input:  "xyz@gmail.com",
				result: 11,
				err:    nil,
			},
			mockGetCommonFriendList: mockGetCommonFriendList{
				input: []int{10, 11},
				result: []string{
					"abc@xyz.com",
					"xyz@abc.com",
				},
				err: nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserService := new(mockUserService)
			mockFriendService := new(mockFriendService)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetFirstUserIDByEmail.input).
				Return(testCase.mockGetFirstUserIDByEmail.result, testCase.mockGetFirstUserIDByEmail.err)
			mockUserService.On("GetUserIDByEmail", testCase.mockGetSecondUserIDByEmail.input).
				Return(testCase.mockGetSecondUserIDByEmail.result, testCase.mockGetSecondUserIDByEmail.err)

			mockFriendService.On("GetCommonFriendListByID", testCase.mockGetCommonFriendList.input).
				Return(testCase.mockGetCommonFriendList.result, testCase.mockGetCommonFriendList.err)

			handlers := FriendHandler{
				IUserService:    mockUserService,
				IFriendServices: mockFriendService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}
			//When
			req, err := http.NewRequest(http.MethodGet, "/friend/common-friend", bytes.NewBuffer(requestBody))

			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetCommonFriendListByEmails)
			handler.ServeHTTP(responseRecorder, req)

			//Then
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())
		})
	}
}

func TestFriendHandler_GetEmailsReceiveUpdate(t *testing.T) {
	type mockGetUserIDByEmail struct {
		input  string
		result int
		err    error
	}
	type mockGetEmailsReceiveUpdate struct {
		sender int
		text   string
		result []string
		err    error
	}
	testCases := []struct {
		name                       string
		requestBody                interface{}
		expectedResponseBody       string
		expectedStatus             int
		mockGetSenderUserID        mockGetUserIDByEmail
		mockGetEmailsReceiveUpdate mockGetEmailsReceiveUpdate
	}{
		{
			name: "decode request body failed",
			requestBody: map[string]interface{}{
				"sender": 1,
			},
			expectedResponseBody: "json: cannot unmarshal number into Go struct field EmailReceiveUpdateRequest.sender of type string\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Body no data",
			requestBody: map[string]interface{}{
				"": "",
			},
			expectedResponseBody: "\"sender\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "No text",
			requestBody: map[string]interface{}{
				"sender": "abc@xyz.com",
			},
			expectedResponseBody: "\"text\" is required\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "sender email is invalid",
			requestBody: map[string]interface{}{
				"sender": "abc",
				"text":   "abc",
			},
			expectedResponseBody: "\"sender\" is not valid. (ex: \"andy@abc.xyz\")\n",
			expectedStatus:       http.StatusBadRequest,
		},
		{
			name: "Get email list receive updates failed with error",
			requestBody: map[string]interface{}{
				"sender": "abc@xyz.com",
				"text":   "hello abc@xyz.com lmk@xyz.com",
			},
			expectedResponseBody: "failed with error\n",
			expectedStatus:       http.StatusInternalServerError,
			mockGetSenderUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetEmailsReceiveUpdate: mockGetEmailsReceiveUpdate{
				sender: 10,
				text:   "hello abc@xyz.com lmk@xyz.com",
				result: nil,
				err:    errors.New("failed with error"),
			},
		},
		{
			name: "Get success",
			requestBody: map[string]interface{}{
				"sender": "abc@xyz.com",
				"text":   "hello another@gmail.com",
			},
			expectedResponseBody: "{\"success\":true,\"recipients\":[\"lmk@xyz.com\",\"abc@gmail.com\"]}\n",
			expectedStatus:       http.StatusOK,
			mockGetSenderUserID: mockGetUserIDByEmail{
				input:  "abc@xyz.com",
				result: 10,
				err:    nil,
			},
			mockGetEmailsReceiveUpdate: mockGetEmailsReceiveUpdate{
				sender: 10,
				text:   "hello another@gmail.com",
				result: []string{"lmk@xyz.com", "abc@gmail.com"},
				err:    nil,
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockFriendService := new(mockFriendService)
			mockUserService := new(mockUserService)

			mockUserService.On("GetUserIDByEmail", testCase.mockGetSenderUserID.input).
				Return(testCase.mockGetSenderUserID.result, testCase.mockGetSenderUserID.err)

			mockFriendService.On("GetEmailsReceiveUpdate",
				testCase.mockGetEmailsReceiveUpdate.sender, testCase.mockGetEmailsReceiveUpdate.text).
				Return(testCase.mockGetEmailsReceiveUpdate.result, testCase.mockGetEmailsReceiveUpdate.err)

			handlers := FriendHandler{
				IUserService:    mockUserService,
				IFriendServices: mockFriendService,
			}

			requestBody, err := json.Marshal(testCase.requestBody)
			if err != nil {
				t.Error(err)
			}

			// When
			req, err := http.NewRequest(http.MethodGet, "/friends/emails-receiving-update", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Error(err)
			}
			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(handlers.GetEmailsReceiveUpdate)
			handler.ServeHTTP(responseRecorder, req)

			// Then
			require.Equal(t, testCase.expectedStatus, responseRecorder.Code)
			require.Equal(t, testCase.expectedResponseBody, responseRecorder.Body.String())

		})

	}
}
