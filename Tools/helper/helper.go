package helper

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    io.Reader
}

type Response struct {
	Code int
	Body []byte
	Err  error
}

func TestApiCall(req Request, use []gin.HandlerFunc, handle gin.HandlerFunc) (res Response) {

	resp := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	_, router := gin.CreateTestContext(resp)

	for _, element := range use {
		router.Use(element)
	}

	router.POST(req.Path, handle)
	request, err := http.NewRequest(req.Method, req.Path, req.Body)

	if err != nil {
		res.Err = err
		return
	}

	for key, value := range req.Headers {
		request.Header.Set(key, value)
	}

	router.ServeHTTP(resp, request)

	res.Code = resp.Result().StatusCode
	res.Body, res.Err = ioutil.ReadAll(resp.Result().Body)

	return
}

func JudgementApicallResponse(t *testing.T, res Response, want Response) {
	if res.Code != want.Code {
		t.Errorf("status code should got %d, but got %d", want.Code, res.Code)
	}

	if res.Err != want.Err {
		t.Errorf("err should got %d, but got %d", res.Err, want.Err)
	}

	if string(res.Body) != string(want.Body) {
		t.Errorf("response body should got %s, but got %s", string(want.Body), string(res.Body))
	}
}
