<template>
    <el-table :data="data">
        <el-table-column v-for="(item,index) in data"
            :key="index"
            :prop="item.key"
            :label="item.label"
        >
        </el-table-column>
    </el-table>
</template>

<script>
export default {
    name:'sample-table',
    props:['data','field']
}
</script>