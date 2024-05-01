/*
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 17:35:10
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-12-06 22:38:31
 * @FilePath: \frontend\src\main.js
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

import { createApp } from 'vue';
import App from './App.vue';
import router from './router'; // 导入Vue Router实例
import store from './store'; // 导入Vuex实例
import axios from 'axios';
import Vant from 'vant';

import 'vant/lib/index.css'
// import 'bootstrap/dist/css/bootstrap.css';


const app = createApp(App);
// 设置全局的axios实例
app.config.globalProperties.$axios = axios.create({
  baseURL: '/proxy_url' // 设置全局的BASEURL
});

app.use(Vant).use(router).use(store).mount('#app');