package repositories

import (
	"database/sql"
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestFriendRepo_CreateFriend(t *testing.T) {
	testCases := []struct {
		name        string
		input       *model.FriendsRepoInput
		expectedErr error
		preparePath string
		mockDB      *sql.DB
	}{
		{
			name: "Create failed with error",
			input: &model.FriendsRepoInput{
				FirstID:  1,
				SecondID: 2,
			},
			expectedErr: errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath: "../testhelpers/preparedata/datafortest",
			mockDB:      testhelpers.ConnectDBFailed(),
		},
		{
			name: "Create friend connection success",
			input: &model.FriendsRepoInput{
				FirstID:  1,
				SecondID: 2,
			},
			expectedErr: nil,
			preparePath: "./testhelpers/preparedata/datafortest",
			mockDB:      testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDB, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDB,
			}

			// When
			err := friendRepo.CreateFriend(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestFriendRepo_IsExistedFriend(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Connect DB failed with error",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Friend connection existed",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
		{
			name:           "Friend connection is not exist",
			input:          []int{1, 5},
			expectedResult: false,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.IsExistedFriend(testCase.input[0], testCase.input[1])

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

func TestFriendRepo_IsBlockedByOtherEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Check is blocked failed with error",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "Is blocked by each other",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
		{
			name:           "is not blocked by the other one",
			input:          []int{3, 4},
			expectedResult: false,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.IsBlockedByOtherEmail(testCase.input[0], testCase.input[1])

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

func TestFriendRepo_GetFriendListByID(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult []int
		expectedError  error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get friend list failed with error",
			input:          1,
			expectedResult: nil,
			expectedError:  errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Get friends list success",
			input:          2,
			expectedResult: []int{1},
			expectedError:  nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.GetFriendListByID(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, testCase.expectedResult, result)
			}
		})
	}
}

func TestFriendRepo_GetBlockedListByID(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult []int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get blocked list failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Get blocked list success",
			input:          2,
			expectedResult: []int{1},
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.GetBlockedListByID(testCase.input)

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

func TestFriendRepo_GetBlockingListByID(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult []int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get blocking list failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Get blocking list success",
			input:          1,
			expectedResult: []int{2},
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.GetBlockingListByID(testCase.input)

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

func TestFriendRepo_GetSubscriberList(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult []int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Get success",
			input:          1,
			expectedResult: []int{2},
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.GetSubscriberList(testCase.input)

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

func TestFriendRepo_GetEmailsFriendOrSubscribedWithNoBlocked(t *testing.T) {
	testCases := []struct {
		name           string
		input          int
		expectedResult []int
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Get failed with error",
			input:          1,
			expectedResult: nil,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Get success with no block",
			input:          1,
			expectedResult: []int{2},
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
		{
			name:           "Get success with removed block",
			input:          2,
			expectedResult: []int{},
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDb, testCase.preparePath)

			friendRepo := FriendRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := friendRepo.GetEmailsFriendOrSubscribedWithNoBlocked(testCase.input)

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
