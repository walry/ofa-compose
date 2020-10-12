create table `ofa_company_finance` (
id int(10) not null comment '自增id',
company_id int(10) not null comment  '公司id',
year char(4) not null comment '年份'
period_flag tinyint(4) not null comment '年报，中报'
table_type tinyint(4) not null comment '资产负债/现金流量/利润表/附注'
table_name varchar(30) not null comment '表名称',
key_id smallint(6) not null comment '科目id',
key_name varchar (50) not null comment '科目名称',
account_value float (10) not null comment '科目数值'
is_delete tinyint '记录是否已删除'
created_at '记录创建时间'
updated_at '记录更新时间'
)

insert into ofa_company_finance(company_id,year,period_flag,table_type,table_name,account_id,account_name,account_value) values(?,?,?,?,?,?,?,?)