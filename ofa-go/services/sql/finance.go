package sql

var FinanceSql = map[string]string{
	"saveAccount" : "insert into ofa_company_state_detail(company_id,state_year,state_period,table_index,table_name,account_id,account_name,account_value,state_date) values(?,?,?,?,?,?,?,?,?)",
	"saveStateData" : "INSERT INTO ofa_company_state_info(company_id,state_year,state_period,state_date,annotation,remark) VALUES (?,?,?,?,?,?)",
	"getAccountList" : "SELECT DISTINCT sublist,account_name FROM ofa_state_account_detail where table_name = ?",
}
