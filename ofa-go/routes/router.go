package routes

import (
	"github.com/gin-gonic/gin"
	"ofa/services"
)

func Load(g *gin.Engine) *gin.Engine {
	// 防止 Panic 把进程干死
	g.Use(gin.Recovery())
	// 自定义的中间件
	mw := globalMw()
	g.Use(mw...)
	// 默认404
	g.NoRoute(func(context *gin.Context) {
		context.JSON(404, gin.H{
			"code":    404,
			"msg": "请求地址有误，请核实",
			"data":    "",
		})
	})
	//健康检查
	g.GET("/health", func(context *gin.Context) {
		context.JSON(200,gin.H{ "result": "ok" })
	})
	//路由列表
	routeList(g)
	return g
}

func globalMw() []gin.HandlerFunc {
	return []gin.HandlerFunc{

	}
}

func routeList(g *gin.Engine)  {
	//v1版本
	v1 := g.Group("/ofa/v1")
	//保存行业数据
	v1.POST("/industry/save",services.SaveIndustry)
	//保存上市公司数据
	v1.POST("/company/save",services.SaveCompany)
	//查询上市公司信息
	v1.GET("/company/info/get",services.QueryCompany)
	//公司列表
	v1.GET("/company/list",services.CompanyList)
	//行业列表
	v1.GET("/enterprise/list",services.IndustryList)
	//查询行业中的公司
	v1.GET("/enterprise/get",services.IndustryGet)

	//录入会计科目
	v1.POST("/finance/account/save",services.InsertFinanceAccount)
	//获取会计科目
	v1.GET("/account/get",services.GetAccountList)

	//保存会计报表
	v1.POST("/finance/statement/save",services.SaveFinanceStatement)
	//上传文件
	//v1.POST("/uploads",services.SaveFinanceStatement)

	//获取公司报表数据
	v1.GET("/finance/statement/data",services.GetCompanyStatement)

	v1.GET("/rpc",services.TestRpc)
}

