package repositories

import (
	"database/sql"
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"S3_FriendManagement_ThinhNguyen/testhelpers"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionRepo_CreateSubscription(t *testing.T) {
	testCases := []struct {
		name        string
		input       *model.SubscriptionRepoInput
		expectedErr error
		preparePath string
		mockDB      *sql.DB
	}{
		{
			name: "Create subscription failed with error",
			input: &model.SubscriptionRepoInput{
				Requestor: 1,
				Target:    10,
			},
			expectedErr: errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			preparePath: "../testhelpers/preparedata/datafortest",
			mockDB:      testhelpers.ConnectDBFailed(),
		},
		{
			name: "Create subscription success",
			input: &model.SubscriptionRepoInput{
				Requestor: 1,
				Target:    2,
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

			subscriptionRepo := SubscriptionRepo{
				Db: testCase.mockDB,
			}

			// When
			err := subscriptionRepo.CreateSubscription(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSubscriptionRepo_IsExistedSubscription(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Check is existed subscription failed with error",
			input:          []int{1, 10},
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "subscription exists",
			input:          []int{2, 1},
			expectedResult: true,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
		{
			name:           "subscription does not exist",
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

			subscriptionRepo := SubscriptionRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := subscriptionRepo.IsExistedSubscription(testCase.input[0], testCase.input[1])

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

func TestSubscriptionRepo_IsBlockedByOtherEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		preparePath    string
		mockDb         *sql.DB
	}{
		{
			name:           "Check is blocked by other email failed with error",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    errors.New("pq: password authentication failed for user \"postgrespassword=000000\""),
			mockDb:         testhelpers.ConnectDBFailed(),
			preparePath:    "",
		},
		{
			name:           "Is blocked",
			input:          []int{1, 2},
			expectedResult: true,
			expectedErr:    nil,
			mockDb:         testhelpers.ConnectDB(),
			preparePath:    "../testhelpers/preparedata/datafortest",
		},
		{
			name:           "Not blocked",
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

			subscriptionRepo := SubscriptionRepo{
				Db: testCase.mockDb,
			}

			// When
			result, err := subscriptionRepo.IsBlockedByOtherEmail(testCase.input[0], testCase.input[1])

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
