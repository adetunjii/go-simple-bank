package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	mockdb "github.com/Adetunjii/simplebank/db/mock"
	db "github.com/Adetunjii/simplebank/db/models"
	db2 "github.com/Adetunjii/simplebank/db/repository"
	"github.com/Adetunjii/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)


// custom go matcher
type eqCreateUserParamMatcher struct {
	arg db2.CreateUserDto
	password string
}

func (e eqCreateUserParamMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db2.CreateUserDto)
	if !ok {
		return false
	}

	err := util.PasswordMatch(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamMatcher) String() string {
	return fmt.Sprintf("Matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserDto(arg db2.CreateUserDto, password string) gomock.Matcher {
	return eqCreateUserParamMatcher{arg, password}
}


func TestServer_CreateUser(t *testing.T) {
	user, password := randomUser(t)

	testCases := []struct{
		name 				string
		body				gin.H
		buildStubs 			func(store *mockdb.MockIStore)
		checkResponse		func(recorder *httptest.ResponseRecorder)
	} {
		{
			name: "OK",
			body: gin.H {
				"username": user.Username,
				"full_name": user.FullName,
				"password": password,
				"email": user.Email,
			},
			buildStubs: func(store *mockdb.MockIStore) {

				arg := db2.CreateUserDto{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}

				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserDto(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockIStore(ctrl)
			testCase.buildStubs(store)

			server := CreateNewServer(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			endpoint := "/users"
			request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			testCase.checkResponse(recorder)
		})
	}
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(8)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.User{
		Username:       util.RandomOwnerName(),
		Password:		 hashedPassword,
		FullName:       util.RandomOwnerName(),
		Email:          util.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	fmt.Println(gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.Password)
}