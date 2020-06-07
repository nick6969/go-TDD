package router

import (
	"tdd/Interface"
	"tdd/router/api"

	"github.com/gin-gonic/gin"
)

func Start(db Interface.Datastore, jwt Interface.JwtTool) {
	r := gin.Default()
	setupRouter(r, db, jwt)
	r.Run()
}

func setupRouter(r *gin.Engine, db Interface.Datastore, jwt Interface.JwtTool) {
	addSubRouter(r, db, jwt)
}

func addSubRouter(r *gin.Engine, db Interface.Datastore, jwt Interface.JwtTool) {
	api.SetupRouter(r.Group("api"), db, jwt)
}
