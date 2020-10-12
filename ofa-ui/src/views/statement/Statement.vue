<template>
    <div class="wrap">
        <div class="box">
            <el-table :data="conpamyFinance">
                <el-table-column prop="id" label="ID"></el-table-column>
                <el-table-column prop="state_year" label="年份"></el-table-column>
                <el-table-column prop="period" label="报告期"></el-table-column>
                <el-table-column prop="state_date" label="公布日期"></el-table-column>
                <el-table-column prop="remark" label="备注"></el-table-column>
            </el-table>
        </div>

        <div class="box">
            <el-tabs type="border-card">
                <el-tab-pane v-for="(item,i) in financeData" :key="i" :label="item.table">
                    <el-table :data="item.data">
                        <el-table-column v-for="(k,j) in item.key" :key="j" :prop="k.key" :label="k.label"></el-table-column>
                    </el-table>
                </el-tab-pane>
            </el-tabs>
        </div>
    </div>
</template>
<script>
    export default {
        data() {
            return {
                conpamyFinance: [],
                financeData: []
            }
        }
    }
</script>

<style lang="scss" scoped>
    @import '@/styles/common.scss'
</style>