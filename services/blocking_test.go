package services

import (
	"errors"
	"testing"

	"S3_FriendManagement_ThinhNguyen/model"
	"github.com/stretchr/testify/require"
)

func TestBlockingService_CreateBlocking(t *testing.T) {
	testCases := []struct {
		name          string
		input         *model.BlockingServiceInput
		expectedErr   error
		mockRepoInput *model.BlockingRepoInput
		mockRepoError error
	}{
		{
			name: "Create blocking failed with error",
			input: &model.BlockingServiceInput{
				Requestor: 1,
				Target:    2,
			},
			expectedErr: errors.New("create blocking failed with error"),
			mockRepoInput: &model.BlockingRepoInput{
				Requestor: 1,
				Target:    2,
			},
			mockRepoError: errors.New("create blocking failed with error"),
		},
		{
			name: "Create blocking success",
			input: &model.BlockingServiceInput{
				Requestor: 3,
				Target:    4,
			},
			expectedErr: nil,
			mockRepoInput: &model.BlockingRepoInput{
				Requestor: 3,
				Target:    4,
			},
			mockRepoError: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// Given
			mockBlockingRepo := new(mockBlockingRepo)
			mockBlockingRepo.On("CreateBlocking", testCase.mockRepoInput).
				Return(testCase.mockRepoError)

			service := BlockingService{
				IBlockingRepo: mockBlockingRepo,
			}

			// Then
			err := service.CreateBlocking(testCase.input)

			// Then
			if testCase.expectedErr != nil {
				require.EqualError(t, err, testCase.expectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBlockingService_IsExistedBlocking(t *testing.T) {
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
			name:           "Check is existed blocking failed with error",
			input:          []int{1, 2},
			expectedResult: false,
			expectedErr:    errors.New("check is existed blocking failed with error"),
			mockRepoInput:  []int{1, 2},
			mockRepoResult: false,
			mockRepoError:  errors.New("check is existed blocking failed with error"),
		},
		{
			name:           "Check is existed blocking success",
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
			mockBlockingRepo := new(mockBlockingRepo)
			mockBlockingRepo.On("IsExistedBlocking", testCase.mockRepoInput[0], testCase.mockRepoInput[1]).
				Return(testCase.mockRepoResult, testCase.mockRepoError)

			service := BlockingService{
				IBlockingRepo: mockBlockingRepo,
			}

			// When
			result, err := service.IsExistedBlocking(testCase.input[0], testCase.input[1])

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
