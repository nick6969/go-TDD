package api

import (
	"errors"
	"strings"
	"tdd/Database/mysql"
	"tdd/Tools/helper"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockDB struct {
	nameCanUse  string
	createID    uint
	createErr   error
	nameCanFind string
	findUserErr error
}

func (mock mockDB) CheckUserNameCanUse(name string) bool {
	return name == mock.nameCanUse
}

func (mock mockDB) CreateCustomer(username, password string) (id uint, err error) {
	if mock.createErr != nil {
		err = mock.createErr
	} else {
		id = mock.createID
	}
	return
}

func (mock mockDB) FindUserWithUserName(name string) (user mysql.Customer, err error) {
	if name == mock.nameCanFind {
		user = mysql.Customer{ID: 1, Username: name, Password: "086420Pp"}
	} else {
		err = mock.findUserErr
	}
	return
}

type mockJWT struct {
	token string
	id    uint
	err   error
}

func (j mockJWT) GenerateUserToken(id uint) (token string, err error) {
	if j.err != nil {
		err = j.err
		return
	}
	token = j.token
	return
}

func (j mockJWT) VerifyUserToken(token string) (id uint, err error) {
	return j.id, nil
}

type args struct {
	req helper.Request
	db  mockDB
	jwt mockJWT
}

func Test_handlePostSignUp(t *testing.T) {

	tests := []struct {
		name string
		args args
		want helper.Response
	}{
		{
			name: "no input",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signUp", Body: nil},
				db:  mockDB{nameCanUse: "Nick"},
				jwt: mockJWT{},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"input noCorrect."}`),
				Err:  nil,
			},
		},
		{
			name: "username taken",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signUp", Body: strings.NewReader(`{"username":"Nick1"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1},
				jwt: mockJWT{},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"username is taken."}`),
				Err:  nil,
			},
		},
		{
			name: "database failure",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signUp", Body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1, createErr: errors.New("error")},
				jwt: mockJWT{},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"Bad Request."}`),
				Err:  nil,
			},
		},
		{
			name: "Jwt generate fail",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signUp", Body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1},
				jwt: mockJWT{err: errors.New("error")},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"Bad Request."}`),
				Err:  nil,
			},
		},
		{
			name: "success",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signUp", Body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1},
				jwt: mockJWT{token: "123456"},
			},
			want: helper.Response{
				Code: 200,
				Body: []byte(`{"data":"123456","success":true}`),
				Err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := helper.TestApiCall(tt.args.req, []gin.HandlerFunc{}, handlePostSignUp(tt.args.db, tt.args.jwt))
			helper.JudgementApicallResponse(t, res, tt.want)
		})
	}
}

func Test_handlePostSignIn(t *testing.T) {
	tests := []struct {
		name string
		args args
		want helper.Response
	}{
		{
			name: "no input",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signIn", Body: nil},
				db:  mockDB{},
				jwt: mockJWT{},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"input noCorrect."}`),
				Err:  nil,
			},
		},
		{
			name: "no correct input.",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signIn", Body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{findUserErr: errors.New("user not find.")},
				jwt: mockJWT{},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"user not find."}`),
				Err:  nil,
			},
		},
		{
			name: "no correct password",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signIn", Body: strings.NewReader(`{"username":"Nick","password":"086420PpXx"}`)},
				db:  mockDB{nameCanFind: "Nick"},
				jwt: mockJWT{},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"password not correct."}`),
				Err:  nil,
			},
		},
		{
			name: "Jwt generate fail",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signIn", Body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanFind: "Nick"},
				jwt: mockJWT{err: errors.New("error")},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"Bad Request."}`),
				Err:  nil,
			},
		},
		{
			name: "success",
			args: args{
				req: helper.Request{Method: "POST", Path: "/api/signIn", Body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanFind: "Nick"},
				jwt: mockJWT{token: "123456"},
			},
			want: helper.Response{
				Code: 200,
				Body: []byte(`{"data":"123456","success":true}`),
				Err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := helper.TestApiCall(tt.args.req, []gin.HandlerFunc{}, handlePostSignIn(tt.args.db, tt.args.jwt))
			helper.JudgementApicallResponse(t, res, tt.want)
		})
	}
}
