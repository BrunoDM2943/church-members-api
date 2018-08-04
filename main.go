package main

import (
	"github.com/BrunoDM2943/church-members-api/handler"
	"github.com/BrunoDM2943/church-members-api/handler/filters"
	"github.com/BrunoDM2943/church-members-api/infra"
	"github.com/BrunoDM2943/church-members-api/member"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	auth := filters.NewAuthFilter()
	r.Use(auth.Validate())
	mongo := infra.NewMongoConnection()
	repo := member.NewMemberRepository(mongo.Connect())
	service := member.NewMemberService(repo)

	memberHandler := handler.NewMemberHandler(service)
	memberHandler.SetUpRoutes(r)
	r.Run()
}
