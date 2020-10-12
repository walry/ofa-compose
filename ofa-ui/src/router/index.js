import Vue from 'vue'
import VueRouter from 'vue-router'


Vue.use(VueRouter)

export const routerMap = [
    {
        path: '/',
        component:() => import('@/views/test/Test.vue')
    },
    {
        path: '/component',
        component: () => import('@/views/test/component/AccountForm')
    },
    {
        path: '/test',
        component: () => import('@/views/test.vue')
    },
    {
        path: '/state/:id',
        component: () => import('@/views/statement/Statement'),
        name: 'statement'
    }
] 
export default new VueRouter({
    routes: routerMap
})