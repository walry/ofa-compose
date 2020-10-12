package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Code 		int						`json:"code"`
	Msg 		string					`json:"msg"`
	Data 		interface{}				`json:"data"`
	HttpCode 	int						`json:"-"`
}


type ResExt interface {
	apply(*Res)
}
//
type funcResExt struct {
	f func(*Res)
}
//继承接口函数
func (fdo *funcResExt)apply(r *Res)  {
	fdo.f(r)
}



//创建中间操作函数
func newFuncResExt(f func(*Res)) *funcResExt{
	return &funcResExt{ f: f}
}

//单独设置响应码
func WithCode(c int) ResExt {
	return newFuncResExt(func(res *Res) {
		res.Code = c
	})
}
//单独设置文本消息
func WithMsg(m string) ResExt {
	return newFuncResExt(func(res *Res) {
		res.Msg = m
	})
}
//单独设置返回数据
func WithData(d interface{}) ResExt {
	return newFuncResExt(func(res *Res) {
		res.Data = d
	})
}
//单独设置HTTP响应码
func WithHttpCode(hc int) ResExt {
	return newFuncResExt(func(res *Res) {
		res.HttpCode = hc
	})
}




//统一响应格式
func Success(c *gin.Context,opts ...ResExt){
	//默认响应结果
	rData := &Res{
		Code: 0,
		Msg:  "Success",
		Data: make([]interface{},0),
		HttpCode: http.StatusOK,
	}
	//自定义返回私有字段
	for _,opt := range opts {
		opt.apply(rData)
	}
	c.JSON(rData.HttpCode,rData)
}

//统一返回失败时响应格式
func Fail(c *gin.Context,opts ...ResExt){
	res := &Res{
		Code:     -1,
		Msg:      "Error",
		Data:     nil,
		HttpCode: http.StatusOK,
	}
	//自定义返回私有字段
	for _,opt := range opts {
		opt.apply(res)
	}
	c.JSON(res.HttpCode,res)
}