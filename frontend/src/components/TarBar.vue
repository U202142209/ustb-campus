<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-10 10:46:49
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-12 17:58:49
 * @FilePath: \ustb-campus\frontend\src\components\TarBar.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <van-tabbar v-model="active" route active-color="#F62301" inactive-color="#8B8B8B">
      <van-tabbar-item icon="guide-o" to="/">首页</van-tabbar-item>
      <van-tabbar-item icon="apps-o" to="/home">广场</van-tabbar-item>
      <van-tabbar-item icon="shield-o" to="/new">发布帖子</van-tabbar-item>
      <van-tabbar-item icon="envelop-o" to="/">消息</van-tabbar-item>
      <van-tabbar-item icon="user-o" to="/center">我的</van-tabbar-item>
    </van-tabbar>
  </div>
</template>

<script>


export default {
  data() {
    return {
      active: 0,
      cartList: []
    }
  }
}
</script>