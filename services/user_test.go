package services

import (
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/require"
)

func TestUserService_CreateUser(t *testing.T) {
	testCases := []struct {
		name          string
		input         *model.UserServiceInput
		expectedErr   error
		mockRepoInput *model.UserRepoInput
		mockRepoErr   error
	}{
		{
			name: "Create user failed with error",
			input: &model.UserServiceInput{
				Email: "xyz@gmail.com",
			},
			expectedErr: errors.New("create user failed with error"),
			mockRepoInput: &model.UserRepoInput{
				Email: "xyz@gmail.com",
			},
			mockRepoErr: errors.New("create user failed with error"),
		},
		{
			name: "Create user success",
			input: &model.UserServiceInput{
				Email: "xyz@gmail.com",
			},
			expectedErr: nil,
			mockRepoInput: &model.UserRepoInput{
				Email: "xyz@gmail.com",
			},
			mockRepoErr: nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserRepo := new(mockUserRepo)
			mockUserRepo.On("CreateUser", testCase.mockRepoInput).
				Return(testCase.mockRepoErr)

			service := UserService{
				IUserRepo: mockUserRepo,
			}

			//When
			err := service.CreateUser(testCase.input)

			//Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserService_IsExistedUser(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedErr    error
		expectedResult bool
		mockRepoInput  string
		mockRepoResult bool
		mockRepoErr    error
	}{
		{
			name:           "Check existed user failed with error",
			input:          "abc@email.com",
			expectedErr:    errors.New("check existed failed with error"),
			expectedResult: false,
			mockRepoInput:  "abc@email.com",
			mockRepoResult: false,
			mockRepoErr:    errors.New("check existed failed with error"),
		},
		{
			name:           "Check existed success",
			input:          "abc@email.com",
			expectedErr:    nil,
			expectedResult: true,
			mockRepoInput:  "abc@email.com",
			mockRepoResult: true,
			mockRepoErr:    nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserRepo := new(mockUserRepo)
			mockUserRepo.On("IsExistedUser", testCase.mockRepoInput).
				Return(testCase.mockRepoResult, testCase.mockRepoErr)

			service := UserService{
				IUserRepo: mockUserRepo,
			}

			//When
			existed, err := service.IsExistedUser(testCase.input)

			//Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, existed)
			}
		})
	}
}

func TestUserService_GetUserIDByEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedErr    error
		expectedResult int
		mockRepoInput  string
		mockRepoResult int
		mockRepoErr    error
	}{
		{
			name:           "Get failed with error",
			input:          "abc@email.com",
			expectedErr:    errors.New("get failed with error"),
			expectedResult: 0,
			mockRepoInput:  "abc@email.com",
			mockRepoResult: 0,
			mockRepoErr:    errors.New("get failed with error"),
		},
		{
			name:           "Check existed success",
			input:          "abc@email.com",
			expectedErr:    nil,
			expectedResult: 10,
			mockRepoInput:  "abc@email.com",
			mockRepoResult: 10,
			mockRepoErr:    nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserRepo := new(mockUserRepo)
			mockUserRepo.On("GetUserIDByEmail", testCase.mockRepoInput).
				Return(testCase.mockRepoResult, testCase.mockRepoErr)

			service := UserService{
				IUserRepo: mockUserRepo,
			}

			//When
			existed, err := service.GetUserIDByEmail(testCase.input)

			//Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, existed)
			}
		})
	}
}

func TestUserService_CheckInvalidEmails(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		expectedErr    error
		expectedResult []string
		mockRepoInput  []string
		mockRepoResult []string
		mockRepoErr    error
	}{
		{
			name:           "Get failed with error",
			input:          []string{"abc@email.com"},
			expectedErr:    errors.New("get failed with error"),
			expectedResult: []string{"abc@email.com"},
			mockRepoInput:  []string{"abc@email.com"},
			mockRepoResult: []string{"abc@email.com"},
			mockRepoErr:    errors.New("get failed with error"),
		},
		{
			name:           "Get failed with error",
			input:          []string{"abc@email.com", "xyz@email.com"},
			expectedErr:    nil,
			expectedResult: []string{"xyz@email.com"},
			mockRepoInput:  []string{"abc@email.com", "xyz@email.com"},
			mockRepoResult: []string{"xyz@email.com"},
			mockRepoErr:    nil,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			//Given
			mockUserRepo := new(mockUserRepo)
			mockUserRepo.On("CheckInvalidEmails", testCase.mockRepoInput).
				Return(testCase.mockRepoResult, testCase.mockRepoErr)

			service := UserService{
				IUserRepo: mockUserRepo,
			}

			//When
			existed, err := service.CheckInvalidEmails(testCase.input)

			//Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, existed)
			}
		})
	}
}
