package enterprise

import "time"

type CompanyInfo struct {
	Id 					uint 						`xorm:"int(10) pk autoincr comment('自增id')" json:"id"`
	Code 				string						`xorm:"char(6) not null index default '' comment('公司代码')" json:"code"`
	ShortName 			string						`xorm:"varchar(30) not null index default '' comment('名称')" json:"short_name"`
	IndustryId			uint 						`xorm:"int(10) not null INDEX default 0 comment('行业表外键')" json:"industry_id"`
	IsDeleted 			uint8						`xorm:"tinyint(4) not null default 0 comment('记录是否删除')" json:"-"`
	CreatedAt 			time.Time 					`xorm:"created default CURRENT_TIMESTAMP comment('记录创建时间')" json:"created_at"`
}

type IndustryInfo struct {
	Id 					uint 						`json:"id" xorm:"int(10) pk autoincr comment('自增主键')"`
	Industry 			string 						`json:"industry" form:"industry" xorm:"varchar(100) not null default '' comment('行业总称')"`
	Code 				string 						`json:"code" form:"code" xorm:"varchar(10) not null default '' comment('行业代码')"`
	Name 				string						`json:"name" form:"name" xorm:"varchar(150) not null default '' comment('行业名称')"`
	Type 				uint8 						`form:"type" xorm:"tinyint(4) not null default 0 comment('行业分类：0，证监会分类；1，自定义分类')"`
	IsDeleted			uint8 						`xorm:"tinyint(4) not null default 0 comment('记录是否已删除')"`
	CreatedAt 			time.Time 					`xorm:"created default CURRENT_TIMESTAMP comment('记录创建时间')"`
}

type CustomGroupInfo struct {
	Id 					uint 						`xorm:"int(10) pk autoincr comment('自增id')"`
	GroupId 			uint 						`xorm:"int(10) not null index(search) default 0 comment('自定义组id，外键')"`
	CompanyId 			uint 						`xorm:"int(10) not null index(search) default 0 comment('公司id')"`
	IsDeleted 			uint8 						`xorm:"tinyint(4) not null default 0 comment('记录是否已删除')"`
	CreatedAt 			time.Time 					`xorm:"created default CURRENT_TIMESTAMP comment('记录创建时间')"`
	UpdatedAt 			time.Time 					`xorm:"updated default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment('记录更新时间')"`
}
