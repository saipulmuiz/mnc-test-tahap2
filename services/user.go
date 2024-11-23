package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/saipulmuiz/mnc-test-tahap2/helpers"
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repositories.UserRepo
}

func NewUserService(repo repositories.UserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (u *UserService) RegisterUser(request params.RegisterUser) *params.Response {
	user := models.User{
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		PIN:         request.PIN,
	}

	userCheck, err := u.userRepo.CheckUserByPhoneNumber(request.PhoneNumber)
	if err != nil && err != gorm.ErrRecordNotFound {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	if userCheck.UserID != uuid.Nil {
		return helpers.HandleErrorService(http.StatusBadRequest, "Phone Number already registered")
	}

	userData, err := u.userRepo.RegisterUser(&user)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	result := params.ResponseSuccess{
		Status: "SUCCESS",
		Data:   userData,
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: result,
	}
}

func (u *UserService) Login(request params.UserLogin) *params.Response {
	user, err := u.userRepo.CheckUserByPhoneNumber(request.PhoneNumber)
	if err != nil {
		return helpers.HandleErrorService(http.StatusBadRequest, err.Error())
	}

	dataIsOK := helpers.CompareCredential([]byte(user.PIN), []byte(request.PIN))

	if !dataIsOK {
		return helpers.HandleErrorService(http.StatusBadRequest, "Phone Number and PIN doesn't match")
	}

	accessToken, err := helpers.GenerateAccessToken(user.UserID.String(), user.PhoneNumber)
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, "Failed to generate access token")
	}

	refreshToken, err := helpers.GenerateRefreshToken(user.UserID.String())
	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, "failed to generate refresh token")
	}

	return &params.Response{
		Status: http.StatusOK,
		Payload: params.ResponseSuccessLogin{
			Status:       "SUCCESS",
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}
}

func (u *UserService) UpdateProfile(id int, request params.UpdateProfile) *params.Response {
	checkData, _ := u.userRepo.CheckUserByID(id, &models.User{})

	if checkData.UserID == uuid.Nil {
		return helpers.HandleErrorService(http.StatusNotFound, "User not found")
	}

	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Address:   request.Address,
	}

	profileUpdated, err := u.userRepo.UpdateUser(checkData.UserID.String(), &user)

	if err != nil {
		return helpers.HandleErrorService(http.StatusInternalServerError, err.Error())
	}

	result := params.ResponseSuccess{
		Status: "SUCCESS",
		Data: params.UpdateProfileResponse{
			UserID:      profileUpdated.UserID.String(),
			FirstName:   profileUpdated.FirstName,
			LastName:    profileUpdated.LastName,
			Address:     profileUpdated.Address,
			UpdatedDate: helpers.ParseDateTime(helpers.DATE_FORMAT_YYYY_MM_DD_TIME, profileUpdated.UpdatedDate),
		},
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: result,
	}
}

func (u *UserService) Logout() *params.Response {
	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message":       "Logout successfull",
			"is_logged_out": true,
		},
	}
}
