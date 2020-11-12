package services

import (
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/require"
)

func TestSubscriptionService_CreateSubscription(t *testing.T) {
	testCases := []struct {
		name          string
		input         *model.SubscriptionServiceInput
		expectedError error
		mockRepoInput *model.SubscriptionRepoInput
		mockRepoError error
	}{
		{
			name: "create subscription failed",
			input: &model.SubscriptionServiceInput{
				Requestor: 1,
				Target:    2,
			},
			expectedError: errors.New("create failed with error"),
			mockRepoInput: &model.SubscriptionRepoInput{
				Requestor: 1,
				Target:    2,
			},
			mockRepoError: errors.New("create failed with error"),
		},
		{
			name: "create subscription successfully",
			input: &model.SubscriptionServiceInput{
				Requestor: 3,
				Target:    4,
			},
			expectedError: nil,
			mockRepoInput: &model.SubscriptionRepoInput{
				Requestor: 3,
				Target:    4,
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockSubscriptionRepo := new(mockSubscriptionRepo)
			mockSubscriptionRepo.On("CreateSubscription", testCase.mockRepoInput).
				Return(testCase.mockRepoError)

			service := SubscriptionService{
				ISubscriptionRepo: mockSubscriptionRepo,
			}

			// When
			err := service.CreateSubscription(testCase.input)

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestSubscriptionService_IsExistedSubscription(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedErr    error
		mockRepoInput  []int
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "Check is existed subscription failed with error",
			input:          []int{1, 2},
			expectedResult: false,
			expectedErr:    errors.New("check is existed subscription failed with error"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("check is existed subscription failed with error"),
		},
		{
			name:           "Subscription is not exist",
			input:          []int{1, 2},
			expectedResult: false,
			expectedErr:    nil,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  nil,
		},
		{
			name:           "Check is existed subscription success",
			input:          []int{1, 2},
			expectedErr:    nil,
			expectedResult: true,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockSubscriptionRepo := new(mockSubscriptionRepo)
			mockSubscriptionRepo.On("IsExistedSubscription", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := SubscriptionService{
				ISubscriptionRepo: mockSubscriptionRepo,
			}

			// When
			result, err := service.IsExistedSubscription(testCase.input[0], testCase.input[1])

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

func TestSubscriptionService_IsBlockedByOtherEmail(t *testing.T) {
	testCases := []struct {
		name           string
		input          []int
		expectedResult bool
		expectedError  error
		mockRepoInput  []int
		mockRepoResult bool
		mockRepoError  error
	}{
		{
			name:           "Check is blocked failed with error",
			input:          []int{1, 2},
			expectedResult: false,
			expectedError:  errors.New("check is blocked failed with error"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("check is blocked failed with error"),
		},
		{
			name:           "Check is blocked successfully",
			input:          []int{1, 2},
			expectedError:  nil,
			expectedResult: true,
			mockRepoInput:  []int{1, 2},
			mockRepoResult: true,
			mockRepoError:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockSubscriptionRepo := new(mockSubscriptionRepo)
			mockSubscriptionRepo.On("IsBlockedByOtherEmail", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := SubscriptionService{
				ISubscriptionRepo: mockSubscriptionRepo,
			}

			// When
			result, err := service.IsBlockedByOtherEmail(testCase.input[0], testCase.input[1])

			// Then
			if testCase.expectedError != nil {
				require.EqualError(t, err, testCase.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, result, testCase.expectedResult)
			}
		})
	}
}
