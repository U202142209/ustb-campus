<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-26 13:28:53
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-26 13:42:20
 * @FilePath: \frontend\src\views\NotFound.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="not-found">
    <van-image src="https://ts1.cn.mm.bing.net/th?id=OIP-C.Az_SQ6Juv4XTVIU8ukzN9QHaEK&w=333&h=187&c=8&rs=1&qlt=90&o=6&dpr=2.2&pid=3.1&rm=2" fit="contain" alt="404" />
    <h1>页面未找到</h1>
    <p>抱歉，您访问的页面不存在。</p>
    <router-link to="/home">返回首页</router-link>
  </div>
</template>

<script>
export default {
  name: 'NotFound',
};
</script>

<style scoped>
.not-found {
  text-align: center;
  padding: 20px;
}

h1 {
  font-size: 24px;
  margin-bottom: 10px;
}

p {
  font-size: 16px;
  color: #999;
  margin-bottom: 20px;
}

router-link {
  color: #007aff;
}
</style>