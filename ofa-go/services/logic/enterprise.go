package logic

import (
	"fmt"
	"ofa/lib/db"
	"ofa/models/enterprise"
	"ofa/services/sql"
	"regexp"
)


//行业
type Industry struct {
}

func (idy *Industry) TableName() string {
	return "ofa_industry_info"
}

func (idy *Industry) Save(data enterprise.IndustryInfo) (int64,error) {
	ofa := db.Use("default")
	res,_ := ofa.Exec(sql.IdySqlList["insertIndustry"],data.Industry,data.Code,data.Name,data.Type)
	return res.LastInsertId()
}

//是否已录入某个行业
func (idy *Industry) Has(code string,name string) bool {
	ofa := db.Use("default")
	has,_ := ofa.Where("code=? and name=?",code,name).Get(idy)
	return has
}

func (idy *Industry)List() interface{}{
	type ResData struct {
		Id 					uint 						`json:"id"`
		Industry 			string 						`json:"industry"`
		Name 				string						`json:"name"`
	}
	var list []ResData
	ofa := db.Use("default")
	_ = ofa.SQL(sql.IdySqlList["queryAllIndustry"]).Find(&list)
	return list
}

//查询行业里包含的公司
func (idy *Industry)Company(pid string) interface{} {
	var data []Company
	DB := db.Use("default")
	sqlStr := sql.IdySqlList["queryCompanyInfo"] + " and a.industry_id = ?"
	_ = DB.SQL(sqlStr,pid).Find(&data)
	return data
}


//公司
type Company struct {
	Id 					uint 		`json:"id"`
	Code 				string		`json:"code"`
	ShortName 			string		`json:"short_name"`
	IndustryId			uint 		` json:"industry_id"`
	Industry			string 		`json:"industry"`
	Name 				string		`json:"name"`
}

func (c *Company)Save() (int64,error){
	ofa := db.Use("default")
	res,_ := ofa.Exec(sql.IdySqlList["insertCompany"],c.Code,c.ShortName,c.IndustryId)
	return res.LastInsertId()
}

func (c *Company)Info(value string) (interface{},string) {
	clause := fmt.Sprintf(" and a.short_name like '%%%s%%'",value)
	matched,_ := regexp.MatchString("^[0-9]+$",value)
	if matched {
		clause = fmt.Sprintf(" and a.code like '%%%s%%'",value)
	}

	var list []Company
	ofa := db.Use("default")
	sqlStr := sql.IdySqlList["queryCompanyInfo"] + clause
	fmt.Println(sqlStr)
	_ = ofa.SQL(sqlStr).Find(&list)
	return list,""
}

//公司列表
func (c *Company)List()interface{}{
	var resList []Company
	ofa := db.Use("default")
	_ = ofa.SQL(sql.IdySqlList["queryCompanyInfo"]).Find(&resList)
	return resList
}

//批量导入数据
func BatchSave(item []string)  {
	data := enterprise.IndustryInfo{
			Industry:  item[0],
			Code:      item[1],
			Name:      item[2],
		}
	var idy Industry
	var industryId uint
	if !idy.Has(item[1],item[2]) {
		id,_ := idy.Save(data)
		industryId = uint(id)
	}else {
		industryId = data.Id
	}
	company := &Company{
		Code:       item[3],
		ShortName:  item[4],
		IndustryId: industryId,
	}
	company.Save()
}