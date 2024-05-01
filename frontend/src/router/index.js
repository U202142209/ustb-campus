/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 19:03:38
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-26 13:37:41
 * @FilePath: \ustb-campus\frontend\src\router\index.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AEim
 */

import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: "/",
    name: "Index",
    component: () => import("@/views/IndexView.vue")
  }, {
    path: "/home",
    name: "Home",
    component: () => import("@/views/HomeView.vue")
  },
  {
    path: "/center",// 用户个人中心
    name: "UserCenter",
    component: () => import("@/views/MySelf.vue")
  },
  {
    path: "/blog/:blog_id",
    name: "BlogDetail",
    component: () => import("@/views/BlogDetail.vue")
  },
  {
    path: "/user/:openid",
    name: "User",
    component: () => import("@/views/UserCenter.vue")
  },
  {
    path:"/new",
    name:"NewPost",
    component: () => import("@/views/NewPost.vue")
  },
  {
    path: '/404',
    name: 'NotFound',
    component:() => import('@/views/NotFound.vue') 
  },
  // {
  //   path: '*',
  //   redirect: '/404',
  // },
]

const router = createRouter({
  history: createWebHistory(),
  routes: routes
});

export default router










