package models

import (
	"ofa/lib/db"
	"ofa/lib/log"
	"ofa/models/enterprise"
	"ofa/models/finance"
)

func Init(){
	instance := db.Use("default")
	//创建CompanyStateDetail表
	if has,_ := instance.IsTableExist(&finance.StateAccountDetail{}); !has {
		err := instance.CreateTables(&finance.StateAccountDetail{})
		_ = instance.CreateIndexes(&finance.StateAccountDetail{})
		log.Info("创建ofa_company_state_detail表：",err)
	}
	//创建CompanyInfo表
	if has,_ := instance.IsTableExist(&enterprise.CompanyInfo{}); !has {
		err := instance.CreateTables(&enterprise.CompanyInfo{})
		_ = instance.CreateIndexes(&enterprise.CompanyInfo{})
		log.Info("创建ofa_company_info表:",err)
	}
	//创建IndustryInfo表
	if has,_ := instance.IsTableExist(&enterprise.IndustryInfo{}); !has {
		err := instance.CreateTables(&enterprise.IndustryInfo{})
		log.Info("创建ofa_industry_info表:",err)
	}
	//创建CustomGroupInfo表
	if has,_ := instance.IsTableExist(&enterprise.CustomGroupInfo{}); !has {
		err := instance.CreateTables(&enterprise.CustomGroupInfo{})
		_ = instance.CreateIndexes(&enterprise.CustomGroupInfo{})
		log.Info("创建ofa_custom_group_info表：",err)
	}
	//创建CompanyStateInfo表
	if has,_ := instance.IsTableExist(&finance.CompanyStateInfo{}); !has {
		err := instance.CreateTables(&finance.CompanyStateInfo{})
		_ = instance.CreateIndexes(&finance.CompanyStateInfo{})
		log.Info("创建ofa_company_state_info表：",err)
	}
}
