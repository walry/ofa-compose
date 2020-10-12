<template>
    <div class="wrap">
        <div class="steps">
            <el-steps :active="active" finish-status="success" process-status="finish" align-center>
                <el-step v-for="(name,index) in tableInfo" :title="name.label" :key="index"></el-step>
            </el-steps>
        </div>
        <div class="content">
            <div v-if="active < 3">
                <div class="item" v-for="(group,index) in currentAccountArr" :key="index">
                    <el-select size="mini" v-model="group.subTable" placeholder="报表子项目" filterable allow-create default-first-option>
                        <el-option v-for="(sub,index_sub) in sublist" :key="index_sub" :label="sub" :value="sub"></el-option>
                    </el-select>
                    <div class="item-row" v-for="(item,index_c) in group.children" :key="index_c">
                    <el-select size="small" v-model="item.account" filterable allow-create default-first-option placeholder="会计科目">
                        <el-option v-for="(a,i) in accountList" :key="i" :label="a" :value="a"></el-option>
                    </el-select>
                    <el-input v-model="item.beginValue" type="number" placeholder="期末金额（单位：元）" size="small"></el-input>
                    -
                    <el-input v-model="item.endValue" type="number" placeholder="期初金额（单位：元）" size="small"></el-input>
                    <span v-if="showIcon(index,index_c)" @click="addItem(index)" style="color:#67C23A;"><i class="el-icon-circle-plus"></i></span>
                    <span v-else @click="delItem(index,index_c)" style="color:#F56C6C;"><i class="el-icon-remove"></i></span>
                </div>
                </div>
                <div style="text-align:center;"><el-button type="success" icon="el-icon-plus" circle @click="addGroup"></el-button></div>
            </div>
            <div v-if="active == 3">
                <el-input type="textarea" v-model="tableInfo[active].data"></el-input>
            </div>
        </div>
        <div class="bottom-tool">
            <el-button type="primary" plain @click="preClick">上一步</el-button>
            <el-button type="primary" @click="submitAccountInfo">录入</el-button>
        </div>
    </div>
</template>

<script>
    import {
        getTableAccount
    } from '@/api/api'

    export default {
        data() {
            return {
                active: 0,
                tableInfo: [
                    { label: '资产负债表', data: [] },
                    { label: '现金流量表', data: [] },
                    { label: '利润表', data: [] },
                    { label: '报表附注', data: '' }
                ],
                accountList: [],
                sublist: []
            }
        },
        computed: {
            currentAccountArr() {
                return this.tableInfo[this.active].data
            },
            tableName() {
                return this.tableInfo[this.active].label
            }

        },
        methods: {
            preClick() {
                if (this.active-- < 0) {
                    this.active = 0
                }
            },
            nextClick() {
                if (this.active++ > 3) {
                    this.active = 3
                }
                if (this.currentAccountArr instanceof Array && this.currentAccountArr.length == 0) {
                    this.addGroup()
                }
            },
            getDropdownData(table) {
                getTableAccount({
                    name: table
                }).then(res => {
                    this.sublist = res.data.sublist
                    this.accountList = res.data.account
                })
            },
            addItem(index) {
                this.currentAccountArr[index].children.push({ account: '', beginValue: '', endValue: ''})
            },
            delItem(index,subIndex) {
                this.currentAccountArr[index].children.splice(subIndex, 1)
            },
            submitAccountInfo() {
                if (this.active != 3) {
                    this.nextClick()
                    return
                }
                let table = []
                let annotation
                this.tableInfo.forEach(item => {
                    if (item.data instanceof Array){
                        item.data.forEach(child => {
                            child.children.forEach(v => {
                                let tmp = {}
                                tmp.table_name = item.label
                                tmp.sublist = child.subTable
                                tmp.account_name = v.account
                                tmp.begin_value = v.beginValue
                                tmp.end_value = v.endValue
                                table.push(tmp)
                            })
                        })
                    }else{
                        if(item.label == '报表附注'){
                            annotation = item.data
                        }
                    }
                })

                this.$emit('data-finish', {
                    table,
                    annotation
                })
            },
            showIcon(i,index) {
                return this.currentAccountArr[i].children.length == index + 1
            },
            addGroup(){
                this.currentAccountArr.push({
                    subTable: '',
                    children:[{account: '',beginValue: '',endValue: ''}]
                })
            }
        },
        created() {
            this.getDropdownData(this.tableName)
            this.addGroup()
        }
    }
</script>

<style scoped>
    .wrap {
        margin: 5px 10px;
        border-radius: 5px;
        background-color: #F3F3F3;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
        padding: 20px 10px;
    }

    .steps {
        margin: 25px 10px;
        background-color: #fff;
        padding: 10px;
        align-self: stretch;
    }

    .content {
        flex: auto;
    }
    .item{
        margin-top: 80px;
    }

    .item-row {
        width: 100%;
        margin: 20px 0;
        display: flex;
        
        justify-content: space-between;
    }

    .bottom-tool {
        margin: 60px 0;
    }
</style>