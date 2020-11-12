package repositories

import (
	"database/sql"
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestUserRepo_CreateUser(t *testing.T) {
	testCases := []struct {
		name        string
		input       *model.UserRepoInput
		expectedErr error
		mockDB      *sql.DB
	}{
		{
			name: "Create new user failed with error",
			input: &model.UserRepoInput{
				Email: "abc@gmail.com",
			},
			expectedErr: errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDB:      testhelpers.ConnectDBFailed(),
		},
		{
			name: "Create user success",
			input: &model.UserRepoInput{
				Email: "xyz@abc.com",
			},
			expectedErr: nil,
			mockDB:      testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			dbMock := testCase.mockDB

			UserRepo := UserRepo{
				Db: dbMock,
			}

			// When
			err := UserRepo.CreateUser(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserRepo_IsExistedUser(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Check existed failed with error",
			input:          "abc@xyz.com",
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "User existed",
			input:          "abc@xyz.com",
			expectedResult: true,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
		{
			name:           "User not exist",
			input:          "abcd@xyz.com",
			expectedResult: false,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			if err := testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath); err != nil {
				//t.Error(err)
			}

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.IsExistedUser(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}

func TestUserRepo_GetUserIDByEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get UserID failed with error",
			input:          "abc@xyz.com",
			expectedResult: 0,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "The user does not exist",
			input:          "mlk@xyz.com",
			expectedResult: 0,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
		{
			name:           "Get UserID by email success",
			input:          "abc@xyz.com",
			expectedResult: 1,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.GetUserIDByEmail(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestUserRepo_GetEmailListByIDs(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult []string
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "No data emailsInput",
			input:          []int{},
			expectedResult: []string{},
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "",
		},
		{
			name:           "Failed with error",
			input:          []int{1},
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "Get email list from UserID list success",
			input:          []int{1},
			expectedResult: []string{"abc@xyz.com"},
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.GetEmailListByIDs(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}

func TestUserRepo_GetUserIDsByEmails(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		expectedResult []int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "No data emailsInput",
			input:          []string{},
			expectedResult: []int{},
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "",
		},
		{
			name:           "Failed with error",
			input:          []string{"abc@xyz.com"},
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "Get email list from UserID list success",
			input:          []string{"abc@xyz.com", "xyz@abc.com"},
			expectedResult: []int{1, 2},
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			userRepo := UserRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := userRepo.GetUserIDsByEmails(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}
