package middleware

import (
	"net/http"
	"tdd/Interface"
	res "tdd/model/response"

	"github.com/gin-gonic/gin"
)

func AuthCustomer(db Interface.DatastoreAuthCustomer, jwt Interface.JwtTool) gin.HandlerFunc {

	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.NoTokenProvider())
			return
		}

		id, err := jwt.VerifyUserToken(token)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.APIBadRequest())
			return
		}

		if ok := db.ConfirmCustomerHas(id); !ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, res.APIBadRequest())
			return
		}

	}

}
