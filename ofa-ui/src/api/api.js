import request from './request'


export function saveIndustry(data){
    return request({
        url: '/ofa/v1/industry/save',
        method: 'post',
        data: data
    })
}

export function search(data){
    return request({
        url:'/ofa/v1/company/info/get',
        method: 'get',
        params: data
    })
}

//报表项目列表
export function getTableAccount(params){
    return request({
        url: '/ofa/v1/account/get',
        method: 'get',
        params
    })
}

//上市公司表单提交
export function submitFinanceData(data){
    return request({
        url: '/ofa/v1/finance/statement/save',
        method: 'post',
        data,
        headers: {'Content-Type':'application/json; charset=utf-8'}
    })
}

//获取行业列表数据
export function getIndustryList(){
    return request({
        url: '/ofa/v1/enterprise/list',
        method: 'get'
    })
}

//获取行业包含的公司
export function getIndustryInfo(params){
    return request({
        url: '/ofa/v1/enterprise/get',
        method: 'get',
        params
    })
}

//公司列表
export function getCompanyList(){
    return request({
        url: '/ofa/v1/company/list',
        method: 'get'
    })
}