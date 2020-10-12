package logic

import (
	"ofa/lib/db"
	"ofa/models/finance"
	"ofa/services/sql"
)

type FinanceAccount struct {

}

type FinanceRequest struct {
	finance.CompanyStateInfo
	Account 							[]finance.StateAccountDetail 						`json:"account" form:"account"`
}

////保存会计科目
//func (fa *FinanceAccount)Save()(int64,error){
//	store := db.Use("default")
//	res,_ := store.Exec(sql.FinanceSql["saveAccount"],fa.TableName,fa.AccountName,fa.AccountValue)
//	return res.LastInsertId()
//}

func (fa *FinanceAccount)AccountList(name string)([]string,[]string) {
	DB := db.Use("default")
	data,_ := DB.Query(sql.FinanceSql["getAccountList"],name)
	sublist := make(map[string]string)
	var account []string
	for _,item := range data {
		account = append(account,string(item["account_name"]))
		sub := string(item["sublist"])
		sublist[sub] = sub
	}
	var sublistArr []string
	for _,e := range sublist {
		sublistArr = append(sublistArr,e)
	}
	return sublistArr,account
}

func (fa *FinanceAccount)SaveFinanceData(param FinanceRequest) error {
	DB := db.Use("default")
	transaction := DB.NewSession()
	defer transaction.Close()

	if err := transaction.Begin(); err != nil {
		return err
	}
	res,err := transaction.Exec(sql.FinanceSql["saveStateData"],param.CompanyId,param.StateYear,param.StatePeriod,param.StateDate,param.Annotation,param.Remark)
	if err != nil {
		return err
	}
	stateId,_ := res.LastInsertId()
	var detailList  []*finance.StateAccountDetail
	for _,item := range param.Account {
		detail := &finance.StateAccountDetail{
			StateId:     uint(stateId),
			TableName:   item.TableName,
			Sublist:     item.Sublist,
			AccountName: item.AccountName,
			BeginValue:  item.BeginValue,
			EndValue:    item.EndValue,
		}
		detailList = append(detailList,detail)
	}
	if _,err := DB.Insert(detailList); err != nil {
		return err
	}
	return transaction.Commit()
}

func (fa *FinanceAccount)GetAllStatement(cid string) interface{} {
	var data []*finance.CompanyStateInfo
	DB := db.Use("default")
	_ = DB.Table(&finance.CompanyStateInfo{}).Select("id,state_year,state_period, state_date,remark,file_path,annotation").Where("company_id = ?",cid).Find(&data)
	return data
}
