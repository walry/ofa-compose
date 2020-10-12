package sql

var IdySqlList = map[string]string{
	"insertIndustry" : "INSERT INTO ofa_industry_info(industry,code,name,type) VALUES(?,?,?,?)",
	"insertCompany"  : "INSERT INTO ofa_company_info(code,short_name,industry_id) VALUES(?,?,?)",
	"queryAllIndustry" : "SELECT id,industry,`name` FROM ofa_industry_info WHERE is_deleted = 0 and type = 0",
	"queryCompanyInfo" : "SELECT a.id,a.industry_id,a.`code`,a.short_name,b.industry,b.`name` FROM ofa_company_info a,ofa_industry_info b WHERE a.industry_id = b.id AND a.is_deleted = 0 AND b.is_deleted = 0 and b.type = 0 ",

}


