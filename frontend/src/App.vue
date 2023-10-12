<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 17:35:10
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-10 13:09:40
 * @FilePath: \frontend\src\App.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE底部
-->
<template>
  <div>
    <!-- 顶部导航栏 -->
    <div>
      <top-nav></top-nav>
    </div>
    <div class="row">
      <div class="col=1 col-md-3"></div>
      <div class="col=10 col-md-6">
        <!-- 路由匹配的组件会被渲染到这里 -->
        <router-view></router-view>
      </div>
    </div>
    <!-- 底部tarbar -->
    <div>
      <tarbar-view></tarbar-view>
    </div>
  </div>
</template>


<script>
import Tarbar from './components/TarBar.vue'
import TopNav from './components/TopNav.vue'
export default {
  components: {
    "tarbar-view": Tarbar,
    "top-nav": TopNav,
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
