package api

import (
	"net/http"
	"tdd/Interface"
	req "tdd/model/request"
	res "tdd/model/response"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.RouterGroup, db Interface.Datastore, jwt Interface.JwtTool) {

	r.POST("signUp", handlePostSignUp(db, jwt))
	r.POST("signIn", handlePostSignIn(db, jwt))

}

func handlePostSignUp(db Interface.DatastoreCustomer, jwt Interface.JwtTool) gin.HandlerFunc {

	return func(c *gin.Context) {
		var requestModel req.SignUpRequest

		if e := c.ShouldBindJSON(&requestModel); e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.APIMessage("input noCorrect."))
			return
		}

		if !db.CheckUserNameCanUse(requestModel.Username) {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.APIMessage("username is taken."))
			return
		}

		id, err := db.CreateCustomer(requestModel.Username, requestModel.Password)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.APIBadRequest())
			return
		}

		token, err := jwt.GenerateUserToken(id)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.APIBadRequest())
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "data": token})
	}

}

func handlePostSignIn(db Interface.DatastoreCustomer, jwt Interface.JwtTool) gin.HandlerFunc {

	return func(c *gin.Context) {

	}

}
