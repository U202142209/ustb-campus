<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-10 10:41:10
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-05-01 12:57:37
 * @FilePath: \ustb-campus\frontend\src\views\NewPost.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <van-form @submit="submitForm">
      <van-field v-model="nickname" label="昵称" placeholder="请输入昵称（必填）"></van-field>
      <van-field v-model="content" label="发帖内容" type="textarea" rows="5" placeholder="请输入发帖内容（必填）"></van-field>
      <van-field v-model="secretKey" label="发帖密钥" placeholder="请输入发帖密钥（必填）"></van-field>
      <van-field v-model="imageUrl" label="图片链接" placeholder="请输入图片链接（选填）"></van-field>
      <van-field v-model="openid" label="openid" placeholder="请输入openid（选填）"></van-field>
      <van-button type="default" block round >发布</van-button>
    </van-form>
  </div>
</template>

<script>
// import { ErrorTypes } from 'vue-router';
import { encroyptedContent } from '@/utils';

export default {
  data() {
    return {
      nickname: '',
      content: '',
      openid: '',
      secretKey: '',
      imageUrl: '',
    };
  },
  methods: {
    submitForm() {
      // 测试加密内容
      var encrypted = encroyptedContent(this.content, "123")
      if (this.content.length == 0) {
        alert("发帖内容不能为空"); return
      } else if (this.nickname.length == 0) {
        alert("昵称不能为空"); return
      } else if (this.secretKey.length == 0) {
        alert("发帖密钥不能为空"); return
      } else {
        const postData = {
          nickname: this.nickname,
          content: this.content,
          openid: this.openid,
          secretKey: this.secretKey,
          imageUrl: this.imageUrl,
          encrypted: encrypted,
          ctime: "123",
        };

        this.$axios.post('/api/post/', postData, {
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
        })
          .then(res => {
            console.log(res.data)
            alert(JSON.stringify(res.data.data))
            // 刷新当前页面
            location.reload();
          })
          .catch(error => {
            console.log(error)
            alert("服务错误.")
          });
      }
    },
  },
};
</script>

<style>
/* 添加Bootstrap样式 */
/* @import '~bootstrap/dist/css/bootstrap.css'; */
</style>