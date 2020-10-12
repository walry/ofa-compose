package services

import (
	"github.com/gin-gonic/gin"
	"ofa/services/logic"
)

//插入会计科目
func InsertFinanceAccount(c *gin.Context){
	var fa logic.FinanceAccount
	if err := c.ShouldBind(&fa); err != nil {
		opts := []ResExt{
			WithMsg(err.Error()),
		}
		Fail(c,opts...)
		return
	}
	//if _,err := fa.Save(); err != nil{
	//	//	log.Error("保存失败：",err)
	//	//}
	Success(c)
}

//获取某个表的会计科目
func GetAccountList(c *gin.Context){
	table := c.Query("name")
	fa := &logic.FinanceAccount{}
	sublist,accounts := fa.AccountList(table)
	data := map[string][]string{"sublist": sublist, "account": accounts}
	opts := []ResExt{
		WithData(data),
	}
	Success(c,opts...)
}

//处理会计报表
func SaveFinanceStatement(c *gin.Context){
	var param logic.FinanceRequest
	if err := c.ShouldBindJSON(&param); err != nil {
		opts := []ResExt{
			WithMsg(err.Error()),
		}
		Fail(c,opts...)
		return
	}

	var fa logic.FinanceAccount
	if err := fa.SaveFinanceData(param); err != nil {
		opts := []ResExt{
			WithMsg(err.Error()),
		}
		Fail(c,opts...)
		return
	}
	Success(c)
}

func GetCompanyStatement(c *gin.Context){
	cid := c.Query("cid")
	var l logic.FinanceAccount
	data :=  l.GetAllStatement(cid)
	res := []ResExt{
		WithData(data),
	}
	Success(c,res...)
}
