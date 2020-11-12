package repositories

import (
	"database/sql"
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestBlockingRepo_CreateBlocking(t *testing.T) {
	testCases := []struct {
		name        string
		input       *model.BlockingRepoInput
		expectedErr error
		preparePath string
		mockDB      *sql.DB
	}{
		{
			name: "Create failed with error",
			input: &model.BlockingRepoInput{
				Requestor: 1,
				Target:    10,
			},
			expectedErr: errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath: "",
			mockDB:      testhelpers.ConnectDBFailed(),
		},
		{
			name: "Create success",
			input: &model.BlockingRepoInput{
				Requestor: 2,
				Target:    1,
			},
			expectedErr: nil,
			preparePath: "../testhelpers/preparedata/datafortest",
			mockDB:      testhelpers.ConnectDB(),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			testhelpers.PrepareDBForTest(testCase.mockDB, testCase.preparePath)

			blockingRepo := BlockingRepo{
				Db: testCase.mockDB,
			}

			// When
			err := blockingRepo.CreateBlocking(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBlockingRepo_IsExistedBlocking(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Check is existed blocking failed with error",
			input:          []int{1, 10},
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath:    "",
			mockDb:         testhelpers.ConnectDBFailed(),
		},
		{
			name:           "Blocking existed",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    nil,
			preparePath:    "../testhelpers/preparedata/datafortest",
			mockDb:         testhelpers.ConnectDB(),
		},
		{
			name:           "Blocking is not exist",
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

			blockingRepo := BlockingRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := blockingRepo.IsExistedBlocking(testCase.input[0], testCase.input[1])

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
