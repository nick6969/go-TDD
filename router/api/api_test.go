package api

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockDB struct {
	nameCanUse string
	createID   uint
	createErr  error
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
	req request
	db  mockDB
	jwt mockJWT
}

func Test_handlePostSignUp(t *testing.T) {

	tests := []struct {
		name string
		args args
		want response
	}{
		{
			name: "no input",
			args: args{
				req: request{method: "POST", path: "/api/signUp", body: nil},
				db:  mockDB{nameCanUse: "Nick"},
				jwt: mockJWT{},
			},
			want: response{
				body: []byte(`{"message":"input noCorrect."}`),
				code: 400,
				err:  nil,
			},
		},
		{
			name: "username taken",
			args: args{
				req: request{method: "POST", path: "/api/signUp", body: strings.NewReader(`{"username":"Nick1"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1},
				jwt: mockJWT{},
			},
			want: response{
				body: []byte(`{"message":"username is taken."}`),
				code: 400,
				err:  nil,
			},
		},
		{
			name: "database failure",
			args: args{
				req: request{method: "POST", path: "/api/signUp", body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1, createErr: errors.New("error")},
				jwt: mockJWT{},
			},
			want: response{
				body: []byte(`{"message":"Bad Request."}`),
				code: 400,
				err:  nil,
			},
		},
		{
			name: "Jwt generate fail",
			args: args{
				req: request{method: "POST", path: "/api/signUp", body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1},
				jwt: mockJWT{err: errors.New("error")},
			},
			want: response{
				body: []byte(`{"message":"Bad Request."}`),
				code: 400,
				err:  nil,
			},
		},
		{
			name: "success",
			args: args{
				req: request{method: "POST", path: "/api/signUp", body: strings.NewReader(`{"username":"Nick","password":"086420Pp"}`)},
				db:  mockDB{nameCanUse: "Nick", createID: 1},
				jwt: mockJWT{token: "123456"},
			},
			want: response{
				body: []byte(`{"data":"123456","success":true}`),
				code: 200,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := testApiCall(tt.args.req, handlePostSignUp(tt.args.db, tt.args.jwt))
			judgementApicallResponse(t, res, tt.want)
		})
	}
}

type request struct {
	method string
	path   string
	body   io.Reader
}

type response struct {
	body []byte
	code int
	err  error
}

func testApiCall(req request, handle gin.HandlerFunc) (res response) {

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(resp)
	router.POST(req.path, handle)
	request, err := http.NewRequest(req.method, req.path, req.body)

	if err != nil {
		res.err = err
		return
	}

	router.ServeHTTP(resp, request)

	res.code = resp.Result().StatusCode
	res.body, res.err = ioutil.ReadAll(resp.Result().Body)

	return
}

func judgementApicallResponse(t *testing.T, res response, want response) {
	if res.code != want.code {
		t.Errorf("status code should got %d, but got %d", want.code, res.code)
	}

	if res.err != want.err {
		t.Errorf("err should got %d, but got %d", res.err, want.err)
	}

	if string(res.body) != string(want.body) {
		t.Errorf("response body should got %s, but got %s", string(res.body), string(want.body))
	}
}
