package middleware

import (
	"tdd/Interface"

	"github.com/gin-gonic/gin"
)

func AuthCustomer(db Interface.DatastoreAuthCustomer, jwt Interface.JwtTool) gin.HandlerFunc {

	return func(c *gin.Context) {

	}

}
