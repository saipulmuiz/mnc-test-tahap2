package services

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_UserService_RegisterUser(t *testing.T) {
	type testCase struct {
		name               string
		wantError          bool
		expectedResponse   *params.Response
		request            params.RegisterUser
		onCheckUserByPhone func(mock *mocks.MockUserRepo)
		onRegisterUser     func(mock *mocks.MockUserRepo)
	}

	var testTable []testCase
	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		request: params.RegisterUser{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "08123456789",
			Address:     "123 Street",
			PIN:         "123456",
		},
		onCheckUserByPhone: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByPhoneNumber("08123456789").Return(&models.User{
				UserID: "",
			}, gorm.ErrRecordNotFound)
		},
		onRegisterUser: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().RegisterUser(gomock.Any()).Return(&models.User{
				UserID:      "123e4567-e89b-12d3-a456-426614174000",
				FirstName:   "John",
				LastName:    "Doe",
				PhoneNumber: "08123456789",
				Address:     "123 Street",
				CreatedDate: time.Now(),
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusCreated,
			Payload: params.ResponseSuccess{
				Status: "SUCCESS",
				Data: params.RegisterUserResponse{
					UserID:      "123e4567-e89b-12d3-a456-426614174000",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "08123456789",
					Address:     "123 Street",
					CreatedDate: helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, time.Now()),
				},
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "phone number already registered",
		wantError: true,
		request: params.RegisterUser{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "08123456789",
			Address:     "123 Street",
			PIN:         "1234",
		},
		onCheckUserByPhone: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByPhoneNumber("08123456789").Return(&models.User{
				UserID: "existing-id",
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusBadRequest,
			Payload: params.ResponseErrorMessage{
				Message: "Phone Number already registered",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "error checking user by phone",
		wantError: true,
		request: params.RegisterUser{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "08123456789",
			Address:     "123 Street",
			PIN:         "1234",
		},
		onCheckUserByPhone: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByPhoneNumber("08123456789").Return(nil, errors.New("some error"))
		},
		expectedResponse: &params.Response{
			Status: http.StatusBadRequest,
			Payload: params.ResponseErrorMessage{
				Message: "some error",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "error registering user",
		wantError: true,
		request: params.RegisterUser{
			FirstName:   "John",
			LastName:    "Doe",
			PhoneNumber: "08123456789",
			Address:     "123 Street",
			PIN:         "1234",
		},
		onCheckUserByPhone: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByPhoneNumber("08123456789").Return(&models.User{
				UserID: "",
			}, gorm.ErrRecordNotFound)
		},
		onRegisterUser: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().RegisterUser(gomock.Any()).Return(nil, errors.New("some error"))
		},
		expectedResponse: &params.Response{
			Status: http.StatusBadRequest,
			Payload: params.ResponseErrorMessage{
				Message: "some error",
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepo(mockCtrl)

			if tc.onCheckUserByPhone != nil {
				tc.onCheckUserByPhone(userRepo)
			}

			if tc.onRegisterUser != nil {
				tc.onRegisterUser(userRepo)
			}

			service := &UserService{userRepo: userRepo}

			resp := service.RegisterUser(tc.request)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResponse.Status, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			} else {
				assert.Equal(t, http.StatusCreated, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			}
		})
	}
}

func Test_UserService_UpdateProfile(t *testing.T) {
	type testCase struct {
		name             string
		wantError        bool
		expectedResponse *params.Response
		request          params.UpdateProfile
		onCheckUserByID  func(mock *mocks.MockUserRepo)
		onUpdateUser     func(mock *mocks.MockUserRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		request: params.UpdateProfile{
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 New Address",
		},
		onCheckUserByID: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByID("123e4567-e89b-12d3-a456-426614174000", gomock.Any()).Return(&models.User{
				UserID: "123e4567-e89b-12d3-a456-426614174000",
			}, nil)
		},
		onUpdateUser: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().UpdateUser("123e4567-e89b-12d3-a456-426614174000", gomock.Any()).Return(&models.User{
				UserID:      "123e4567-e89b-12d3-a456-426614174000",
				FirstName:   "John",
				LastName:    "Doe",
				Address:     "123 New Address",
				UpdatedDate: time.Now(),
			}, nil)
		},
		expectedResponse: &params.Response{
			Status: http.StatusOK,
			Payload: params.ResponseSuccess{
				Status: "SUCCESS",
				Data: params.UpdateProfileResponse{
					UserID:      "123e4567-e89b-12d3-a456-426614174000",
					FirstName:   "John",
					LastName:    "Doe",
					Address:     "123 New Address",
					UpdatedDate: helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, time.Now()),
				},
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "user not found",
		wantError: true,
		request: params.UpdateProfile{
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 New Address",
		},
		onCheckUserByID: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByID("123e4567-e89b-12d3-a456-426614174000", gomock.Any()).Return(&models.User{
				UserID: "",
			}, gorm.ErrRecordNotFound)
		},
		expectedResponse: &params.Response{
			Status: http.StatusNotFound,
			Payload: params.ResponseErrorMessage{
				Message: "User not found",
			},
		},
	})

	testTable = append(testTable, testCase{
		name:      "error updating user",
		wantError: true,
		request: params.UpdateProfile{
			FirstName: "John",
			LastName:  "Doe",
			Address:   "123 New Address",
		},
		onCheckUserByID: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().CheckUserByID("123e4567-e89b-12d3-a456-426614174000", gomock.Any()).Return(&models.User{
				UserID: "123e4567-e89b-12d3-a456-426614174000",
			}, nil)
		},
		onUpdateUser: func(mock *mocks.MockUserRepo) {
			mock.EXPECT().UpdateUser("123e4567-e89b-12d3-a456-426614174000", gomock.Any()).Return(nil, errors.New("failed to update user"))
		},
		expectedResponse: &params.Response{
			Status: http.StatusInternalServerError,
			Payload: params.ResponseErrorMessage{
				Message: "An unexpected error occurred. Please try again later.",
			},
		},
	})

	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			userRepo := mocks.NewMockUserRepo(mockCtrl)

			if tc.onCheckUserByID != nil {
				tc.onCheckUserByID(userRepo)
			}

			if tc.onUpdateUser != nil {
				tc.onUpdateUser(userRepo)
			}

			service := &UserService{
				userRepo: userRepo,
			}

			resp := service.UpdateProfile("123e4567-e89b-12d3-a456-426614174000", tc.request)

			if tc.wantError {
				assert.NotNil(t, resp)
				assert.Equal(t, tc.expectedResponse.Status, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			} else {
				assert.Equal(t, http.StatusOK, resp.Status)
				assert.Equal(t, tc.expectedResponse.Payload, resp.Payload)
			}
		})
	}
}
