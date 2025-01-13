package servers

// import (
// 	"errors"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/entities"
// 	"github.com/suryasaputra2016/phase-3-graded-challenge-2/server/repos"
// )

// // user repo and user service mock
// var userRepoMock = &repos.UserRepoMock{Mock: mock.Mock{}}
// var userServiceMock = userService{ur: userRepoMock}

// // test CheckRegistrationData: email already in used
// func TestCheckRegistrationDataEmailUsed(t *testing.T) {
// 	userRes := entities.User{
// 		UserName:     "adam",
// 		PasswordHash: "hashed password",
// 	}
// 	userRepoMock.Mock.On("FindUserByID", "123").Return(userRes)

// 	err := (*userService).CheckRegistrationData(&userServiceMock, "adam@mail", "123abcABC!")
// 	assert.NotNil(t, err)
// 	assert.Equal(t, errors.New("validating registration data: email already in use"), err, "Test Error email in use")
// }

// // test CheckRegistrationData: email and password validation test
// func TestCheckRegistrationDataEmailAndPasswordValidation(t *testing.T) {
// 	userRepoMock.Mock.On("FindUserByEmail", "sample@mail.com").Return(nil)

// 	tests := []struct {
// 		Expected string
// 		Actual   string
// 		ErrMsg   string
// 	}{
// 		{
// 			Expected: "validating registration data: email is not well formatted",
// 			Actual:   fmt.Sprint((*userService).CheckRegistrationData(&userServiceMock, "sample", "123abcABC!")),
// 			ErrMsg:   "Test Error invalid email format",
// 		},
// 		{
// 			Expected: "validating registration data: password must contain upper case",
// 			Actual:   fmt.Sprint((*userService).CheckRegistrationData(&userServiceMock, "sample@mail.com", "123abcabc!")),
// 			ErrMsg:   "Test Error password not having capital",
// 		},
// 		{
// 			Expected: "validating registration data: password must contain punctuation or special symbol",
// 			Actual:   fmt.Sprint((*userService).CheckRegistrationData(&userServiceMock, "sample@mail.com", "123abcABCA")),
// 			ErrMsg:   "Test Error password not havng special symbol",
// 		},
// 		{
// 			Expected: "validating registration data: password must contain eight or more characters",
// 			Actual:   fmt.Sprint((*userService).CheckRegistrationData(&userServiceMock, "sample@mail.com", "12abAB!")),
// 			ErrMsg:   "Test Error password not long enough",
// 		},
// 		{
// 			Expected: "<nil>",
// 			Actual:   fmt.Sprint((*userService).CheckRegistrationData(&userServiceMock, "sample@mail.com", "123abcABC!")),
// 			ErrMsg:   "Test Success",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Test Registration %s : ", test.Actual), func(t *testing.T) {
// 			assert.Equal(t, test.Expected, test.Actual, test.ErrMsg)
// 		})
// 	}
// }

// func TestCreateNewUser(t *testing.T) {
// 	registrationData := []entity.RegisterRequest{
// 		{
// 			FirstName: "ben",
// 			LastName:  "benalu",
// 			Email:     "ben@mail.com",
// 			Password:  "test",
// 		},
// 		{
// 			FirstName: "chen",
// 			LastName:  "cincinnati",
// 			Email:     "chen@mail.com",
// 			Password:  "test",
// 		},
// 	}

// 	newUser := []entity.User{
// 		{
// 			FirstName:     "ben",
// 			LastName:      "benalu",
// 			Email:         "ben@mail.com",
// 			PasswordHash:  "hashed password",
// 			DepositAmount: 0.0,
// 		},
// 		{
// 			FirstName:     "chen",
// 			LastName:      "cincinnati",
// 			Email:         "chen@mail.com",
// 			PasswordHash:  "hashed password",
// 			DepositAmount: 0.0,
// 		},
// 	}

// 	userRepoMock.Mock.On("AddUser", &newUser[0]).Return(nil)
// 	userRepoMock.Mock.On("AddUser", &newUser[1]).Return(&newUser[1])

// 	tests := []struct {
// 		Expected string
// 		Actual   string
// 		ErrMsg   string
// 	}{
// 		{
// 			Expected: fmt.Sprintf("%v creating new user: adding user failed", nil),
// 			Actual:   fmt.Sprint((*userService).CreateNewUser(&userServiceMock, &registrationData[0])),
// 			ErrMsg:   "Test Error invalid email format",
// 		},
// 		{
// 			Expected: fmt.Sprintf("%v %v", &newUser[1], nil),
// 			Actual:   fmt.Sprint((*userService).CreateNewUser(&userServiceMock, &registrationData[1])),
// 			ErrMsg:   "Test Success",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Test Registration %s : ", test.Actual), func(t *testing.T) {
// 			assert.Equal(t, test.Expected, test.Actual, test.ErrMsg)
// 		})
// 	}
// }
