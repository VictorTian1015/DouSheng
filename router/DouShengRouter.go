package router

import (
	"dousheng/handlers/sociality"
	"dousheng/handlers/user"
	"dousheng/handlers/video"
	"dousheng/middleware"
	"dousheng/models"
	"github.com/gin-gonic/gin"
)

func InitDouyinRouter() *gin.Engine {
	models.InitDB()
	r := gin.Default()

	r.Static("static", "./static")

	baseGroup := r.Group("/douyin")
	//根据灵活性考虑是否加入JWT中间件来进行鉴权，还是在之后再做鉴权
	// basic apis
	baseGroup.GET("/feed/", video.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware.JWTMiddleWare(), user.UserInfoHandler)
	baseGroup.POST("/user/login/", middleware.SHAMiddleWare(), user.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware.SHAMiddleWare(), user.UserRegisterHandler)
	baseGroup.POST("/publish/action/", middleware.JWTMiddleWare(), video.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware.NoAuthToGetUserId(), video.QueryVideoListHandler)

	//extend 1
	baseGroup.POST("/favorite/action/", middleware.JWTMiddleWare(), video.PostFavorHandler)
	baseGroup.GET("/favorite/list/", middleware.NoAuthToGetUserId(), video.QueryFavorVideoListHandler)
	baseGroup.POST("/sociality/action/", middleware.JWTMiddleWare(), sociality.PostCommentHandler)
	baseGroup.GET("/sociality/list/", middleware.JWTMiddleWare(), sociality.QueryCommentListHandler)

	//extend 2
	baseGroup.POST("/relation/action/", middleware.JWTMiddleWare(), sociality.PostFollowActionHandler)
	baseGroup.GET("/relation/sociality/list/", middleware.NoAuthToGetUserId(), sociality.QueryFollowListHandler)
	baseGroup.GET("/relation/follower/list/", middleware.NoAuthToGetUserId(), sociality.QueryFollowerHandler)
	return r
}
