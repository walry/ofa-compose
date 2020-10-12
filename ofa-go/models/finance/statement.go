package finance

import (
	"time"
)


type StateAccountDetail struct {
	Id 						uint 					`xorm:"int(10) pk autoincr comment('自增主键')"`
	StateId 				uint 					`xorm:"int(10) index not null default 0 comment('报表外键')"`
	TableName 				string 					`json:"table_name" form:"table_name" xorm:"varchar(30) not null default '' comment('报表名称')"`
	Sublist 				string 					`json:"sublist" form:"sublist" xorm:"varchar(100) not null default '' comment('报表子项目')"`
	AccountName 			string 					`json:"account_name" form:"account_name" xorm:"varchar(100) not null default '' comment('科目名称')"`
	BeginValue				float64 				`json:"begin_value" form:"begin_value" xorm:"decimal(16,2) not null default 0 comment('期初金额，单位:元')"`
	EndValue				float64 				`json:"end_value" form:"end_value" xorm:"decimal(16,2) not null default 0 comment('期末金额，单位:元')"`
	IsDeleted 				uint8					`xorm:"tinyint(4) not null default 0 comment('记录是否已删除')"`
	CreatedAt 				time.Time				`xorm:"created default CURRENT_TIMESTAMP comment('记录创建时间')"`
}


type CompanyStateInfo struct {
	Id 						uint 					`json:"id" xorm:"int(10) pk autoincr comment('主键')"`
	CompanyId 				uint 					`json:"companyId" form:"companyId" xorm:"int(10) not null index(idx_query) default 0 comment('公司表外键')"`
	StateYear				string 					`json:"year" form:"year" xorm:"char(4) not null index(idx_query) default '' comment('报表年份')"`
	StatePeriod				uint8 					`json:"period" form:"period" xorm:"tinyint(4) not null index(idx_query) default 0 comment('0：年报，1：中报')"`
	StateDate				string					`json:"publishDate" form:"publishDate" xorm:"varchar(20) not null default '' comment('报告公布时间')"`
	Remark 					string 					`json:"remark" form:"remark" xorm:"varchar(500) not null default '' comment('备注')"`
	FilePath 				string 					`json:"file_path" xorm:"varchar(200) not null default '' comment('报表存放目录')"`
	Annotation				string 					`json:"annotation" form:"annotation" xorm:"text not null default '' comment('报表附注')"`
	CreatedAt 				time.Time 				`json:"created_at" xorm:"created default CURRENT_TIMESTAMP"`
}


