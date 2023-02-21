package sociality

import (
	"dousheng/models"
	"dousheng/service/sociality"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FollowListResponse struct {
	models.CommonResponse
	*sociality.FollowList
}

func QueryFollowListHandler(c *gin.Context) {
	NewProxyQueryFollowList(c).Do()
}

type ProxyQueryFollowList struct {
	*gin.Context

	userId int64

	*sociality.FollowList
}

func NewProxyQueryFollowList(context *gin.Context) *ProxyQueryFollowList {
	return &ProxyQueryFollowList{Context: context}
}

func (p *ProxyQueryFollowList) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		p.SendError(err.Error())
		return
	}
	p.SendOk("请求成功")
}

func (p *ProxyQueryFollowList) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowList) prepareData() error {
	list, err := sociality.QueryFollowList(p.userId)
	if err != nil {
		return err
	}
	p.FollowList = list
	return nil
}

func (p *ProxyQueryFollowList) SendError(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg},
	})
}

func (p *ProxyQueryFollowList) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0, StatusMsg: msg},
		FollowList:     p.FollowList,
	})
}

type FollowerListResponse struct {
	models.CommonResponse
	*sociality.FollowerList
}

func QueryFollowerHandler(c *gin.Context) {
	NewProxyQueryFollowerHandler(c).Do()
}

type ProxyQueryFollowerHandler struct {
	*gin.Context

	userId int64

	*sociality.FollowerList
}

func NewProxyQueryFollowerHandler(context *gin.Context) *ProxyQueryFollowerHandler {
	return &ProxyQueryFollowerHandler{Context: context}
}

func (p *ProxyQueryFollowerHandler) Do() {
	var err error
	if err = p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	if err = p.prepareData(); err != nil {
		if errors.Is(err, sociality.ErrUserNotExist) {
			p.SendError(err.Error())
		} else {
			p.SendError("准备数据出错")
		}
		return
	}
	p.SendOk("成功")
}

func (p *ProxyQueryFollowerHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyQueryFollowerHandler) prepareData() error {
	list, err := sociality.QueryFollowerList(p.userId)
	if err != nil {
		return err
	}
	p.FollowerList = list
	return nil
}

func (p *ProxyQueryFollowerHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func (p *ProxyQueryFollowerHandler) SendOk(msg string) {
	p.JSON(http.StatusOK, FollowerListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
		FollowerList: p.FollowerList,
	})
}
