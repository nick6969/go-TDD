package middleware

import (
	"tdd/Tools/helper"
	"testing"

	"github.com/gin-gonic/gin"
)

type mockDB struct {
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

func Test_AuthMiddleware(t *testing.T) {
	tests := []struct {
		name string
		args args
		want helper.Response
	}{
		{
			name: "no input",
			args: args{
				req: helper.Request{Method: "POST", Path: "/", Body: nil},
			},
			want: helper.Response{
				Code: 400,
				Body: []byte(`{"message":"No Token Provider."}`),
				Err:  nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := helper.TestApiCall(tt.args.req, []gin.HandlerFunc{}, AuthCustomer(tt.args.db, tt.args.jwt))
			helper.JudgementApicallResponse(t, res, tt.want)
		})
	}
}
