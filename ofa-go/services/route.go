package services

import (
	"github.com/gin-gonic/gin"
	"ofa/lib/log"
	"ofa/models/enterprise"
	"ofa/services/logic"
	"ofa/services/rpc"
)

//保存行业数据
func SaveIndustry(c *gin.Context){
	var data enterprise.IndustryInfo
	if err := c.ShouldBind(&data); err != nil  {
		log.Error("parse enterprise param err:",err.Error())
		Fail(c)
		return
	}
	var idy logic.Industry
	if _,err := idy.Save(data); err != nil {
		log.Error("db save enterprise err:",err)
		Fail(c)
		return
	}
	opts := []ResExt{
		WithMsg("保存成功"),
	}
	Success(c,opts...)
}

//保存公司数据
func SaveCompany(c *gin.Context)  {
	var company logic.Company
	if err := c.ShouldBind(&company); err != nil {
		log.Error("parse company param err:",err.Error())
		Fail(c)
		return
	}
	if _,err := company.Save(); err != nil {
		log.Error("db save company err:",err)
		Fail(c)
		return
	}
	Success(c)
}

//查询上市公司信息
func QueryCompany(c *gin.Context)  {
	search := c.Query("value")
	var h logic.Company
	data,msg := h.Info(search)
	opts := []ResExt{
		WithMsg(msg),
		WithData(data),
	}
	Success(c,opts...)
}

//公司列表
func CompanyList(c *gin.Context){
	var company logic.Company
	data := company.List()
	opts := []ResExt{
		WithData(data),
	}
	Success(c,opts...)
}

//行业列表
func IndustryList(c *gin.Context){
	var idy logic.Industry
	data := idy.List()
	opts := []ResExt{
		WithData(data),
	}
	Success(c,opts...)
}

//行业中的公司
func IndustryGet(c *gin.Context){
	pid := c.Query("pid")
	var idy logic.Industry
	data := idy.Company(pid)
	res := []ResExt{
		WithData(data),
	}
	Success(c,res...)
}

func TestRpc(c *gin.Context){
	rpc.Client()
}