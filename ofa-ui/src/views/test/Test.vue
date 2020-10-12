<template>
    <div class="wrap">
        <div class="box">
            <el-button type="primary" @click="showAddIndustryDialog = true">新建</el-button>
            <el-button type="primary" @click="showDialog = true">录入</el-button>
        </div>
        <div class="box">
            <el-table :data="industryInfo" max-height="400" @row-click="onIndustryClick">
                <el-table-column prop="industry" label="行业大类"></el-table-column>
                <el-table-column prop="name" label="行业细分"></el-table-column>
                <el-table-column fixed="right" label="操作">
                    <template slot-scope="scope">
                        <el-button @click="addCompany(scope.row)" type="text" size="small">添加</el-button>
                        <el-button @click="delIndustry(scope.row)" type="text" size="small">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>

        <div class="box">
            <el-table :data="companyInfo" max-height="400">
                <el-table-column prop="code" label="代码"></el-table-column>
                <el-table-column prop="short_name" label="名称"></el-table-column>
                <el-table-column prop="industry" label="行业大类"></el-table-column>
                <el-table-column prop="name" label="行业"></el-table-column>
                <el-table-column fixed="right" label="操作">
                    <template slot-scope="scope">
                        <el-button @click="uploadStatementFile(scope.row)" type="text" size="small">上传</el-button>
                        <el-button @click="getFinanceInfo(scope.row)" type="text" size="small">查看</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </div>

        <el-dialog title="报表数据" :visible.sync="showDialog" :show-close="false">
            <el-form v-model="financeStatementForm" label-width="80px" label-position="left" class="dialog-form">
                <el-form-item label="公司">
                    <el-select filterable remote :remote-method="searchCompany" placeholder="搜索上市公司"
                        v-model="financeStatementForm.companyId" value-key="value">
                        <el-option v-for="(option,index) in companyOptions" :key="index" :label="option.label"
                            :value="option.value">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="报表年份">
                    <el-date-picker v-model="financeStatementForm.year" type="year" value-format="yyyy"
                        placeholder="报表年度"></el-date-picker>
                </el-form-item>
                <el-form-item label="年报类型">
                    <el-radio-group v-model="financeStatementForm.period">
                        <el-radio label="0">年报</el-radio>
                        <el-radio label="1">中报</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="公布日期">
                    <el-date-picker v-model="financeStatementForm.publishDate" type="date" value-format="yyyy-MM-dd"
                        placeholder="公布日期"></el-date-picker>
                </el-form-item>
                <el-form-item label="备注">
                    <el-input type="textarea" v-model="financeStatementForm.remark"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button v-if="showAccountView == 0 " type="primary" @click="showAccountView++">财务报表</el-button>
                    <account-form v-else-if="showAccountView == 1" @data-finish="renderAccountList"></account-form>
                    <div v-else>
                        <el-table :data="financeStatementForm.account" :span-method="tableRowSpan">
                            <el-table-column prop="table_name" label="报表"></el-table-column>
                            <el-table-column prop="sublist" label="子项目"></el-table-column>
                            <el-table-column prop="account_name" label="科目" width="120px"></el-table-column>
                            <el-table-column prop="end_value" label="期末余额" width="160px"></el-table-column>
                            <el-table-column prop="begin_value" label="期初余额" width="160px"></el-table-column>
                        </el-table>
                    </div>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="closeFinanceDialog">取 消</el-button>
                <el-button type="primary" @click="submitFinanceData">完成</el-button>
            </div>
        </el-dialog>
        <el-dialog title="添加行业或组" :visible.sync="showAddIndustryDialog">
            <el-form ref="form" :model="industry" size="small" class="dialog-form" label-width="120px">
                <el-form-item label="门类名称及代码">
                    <el-input v-model="industry.industry"></el-input>
                </el-form-item>
                <el-form-item label="行业代码">
                    <el-input v-model="industry.code"></el-input>
                </el-form-item>
                <el-form-item label="行业大类名称">
                    <el-input v-model="industry.name"></el-input>
                </el-form-item>
                <el-form-item label="分类">
                    <el-select v-model="industry.type">
                        <el-option label="证监会" value="0"></el-option>
                        <el-option label="自定义" value="1"></el-option>
                    </el-select>
                </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
                <el-button @click="showAddIndustryDialog = false">取 消</el-button>
                <el-button type="primary" @click="submit">提交</el-button>
            </div>
        </el-dialog>
        <el-dialog title="新增公司" :visible.sync="showAddCompanyDialog">
            <el-form :model="company" class="dialog-form" label-width="80px">
                <el-form-item label="行业ID">
                    <el-input v-model="company.industry_id"></el-input>
                </el-form-item>
                <el-form-item label="代码">
                    <el-input v-model="company.code"></el-input>
                </el-form-item>
                <el-form-item label="名称">
                    <el-input v-model="company.short_name"></el-input>
                </el-form-item>
            </el-form>
             <div slot="footer" class="dialog-footer">
                <el-button @click="showAddCompanyDialog = false">取 消</el-button>
                <el-button type="primary" @click="addNewCompany">提交</el-button>
            </div>
        </el-dialog>
    </div>
</template>

<script>
    import {
        saveIndustry,
        search,
        submitFinanceData,
        getIndustryList,
        getIndustryInfo,
        getCompanyList
    } from '@/api/api'
    import AccountForm from '@/views/test/component/AccountForm'
    export default {
        components: {
            AccountForm
        },
        data() {
            return {
                industry: {
                    industry: '',
                    code: '',
                    name: '',
                    type: ''
                },
                companyOptions: [],
                showAccountView: 0,
                financeStatementForm: {
                    companyId: '',
                    year: '',
                    period: '',
                    publishDate: '',
                    account: [],
                    annotation: '',
                    remark: ''
                },
                showDialog: false,
                showAddIndustryDialog: false,
                industryInfo: [],
                companyInfo: [],
                showAddCompanyDialog: false,
                company: {
                    industry_id: '',
                    code: '',
                    short_name: ''
                }
            }
        },
        methods: {
            getFinanceInfo(row){
                this.$router.push({ path: `/state/${row.id}` })
            },
            uploadStatementFile(row){
                console.log(row)
            },
            delIndustry(row){
                console.log(row)
            },
            addCompany(row){
                this.showAddCompanyDialog = true
                this.company = {
                    industry_id: row.id,
                    code: '',
                    short_name: ''
                }
            },
            addNewCompany(){
                this.showAddCompanyDialog = false

            },
            submit() {
                saveIndustry(this.industry).then(res => {
                    this.showAddIndustryDialog = false
                    this.$message({
                        message: res.msg,
                        type: 'success'
                    })
                })
            },
            searchCompany(query) {
                if (query) {
                    let self = this
                    self.companyOptions = []
                    search({
                        value: query
                    }).then(res => {
                        if (res.data) {
                            res.data.forEach(item => {
                                self.companyOptions.push({
                                    label: item.code + item.short_name,
                                    value: item.id
                                })
                            })
                        }

                    })
                }
            },
            renderAccountList(data) {
                this.financeStatementForm.account = data.table
                this.financeStatementForm.annotation = data.annotation
                this.showAccountView++
            },
            submitFinanceData() {
                this.closeFinanceDialog()
                this.financeStatementForm.period = parseInt(this.financeStatementForm.period)
                this.financeStatementForm.account.map(item => {
                    item.begin_value = parseFloat(item.begin_value)
                    item.end_value = parseFloat(item.end_value)
                    return item
                })
                submitFinanceData(this.financeStatementForm).then(res => {
                    this.$message({
                        type: 'success',
                        message: res.msg
                    })
                })
            },
            closeFinanceDialog() {
                this.showAccountView = 0
                this.showDialog = false
            },
            tableRowSpan({
                row,
                rowIndex,
                columnIndex
            }) {
                let accountInfo = this.financeStatementForm.account
                if (columnIndex === 0) {
                    if (rowIndex == 0 || row.table_name != accountInfo[rowIndex - 1].table_name) {
                        let same = accountInfo.filter(item => {
                            return item.table_name == row.table_name
                        })
                        return {
                            rowspan: same.length,
                            colspan: 1
                        }
                    }
                    return {
                        rowspan: 0,
                        colspan: 0
                    }
                }
                if (columnIndex === 1) {
                    if (rowIndex == 0 || row.sublist != accountInfo[rowIndex - 1].sublist) {
                        let same = accountInfo.filter(item => {
                            return item.sublist == row.sublist
                        })
                        return {
                            rowspan: same.length,
                            colspan: 1
                        }
                    }
                    return {
                        rowspan: 0,
                        colspan: 0
                    }
                }
            },
            getIndustryList(){
                getIndustryList().then(res => { this.industryInfo = res.data })
            },
            getCompanyList(){
                getCompanyList().then( res => { this.companyInfo = res.data })
            },
            onIndustryClick(row){
                getIndustryInfo({ pid: row.id }).then(res => { this.companyInfo = res.data })
            }
        },
        created() {
            this.getIndustryList()
            this.getCompanyList()
        },
    }
</script>

<style>
    .box {
        border: solid 1px #cacaca;
        margin: 25px 10px;
        padding: 30px 50px;
        border-radius: 8px;
        box-shadow: 6px 6px 6px #cacaca;
        display: flex;
        justify-content: left;
        flex-wrap: wrap;
    }

    .dialog-form {
        width: 80%;
        text-align: left;
    }
</style>