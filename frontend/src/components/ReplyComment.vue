<template>
  <div>
    <div id="myModal" class="modal" tabindex="-1" role="dialog">
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <p class="modal-title">回复:{{ commentReplyOnject.userName }}</p>
            <button @click="hideModel()" type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <van-form @submit="submitForm">
              <van-field v-model="reply_nickname" label="回帖昵称" placeholder="请输入回帖昵称（必填）"></van-field>
              <van-field v-model="reply_content" label="回帖内容" type="textarea" rows="5"
                placeholder="请输入回复内容（必填）"></van-field>
              <van-field v-model="reply_secretKey" label="回帖密钥" placeholder="请输入回复密钥（必填）"></van-field>
              <van-field v-model="relpy_imageUrl" label="图片链接" placeholder="请输入图片链接（选填）"></van-field>
              <van-field v-model="reply_openid" label="openid" placeholder="请输入openid（选填）"></van-field>
              <van-button type="primary" native-type="submit">回复</van-button>
            </van-form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default{
  data() {
    return {
      commentReplyOnject: {},
      reply_nickname: "",
      reply_content: "",
      reply_openid: "",
      reply_secretKey: "",
      relpy_imageUrl: "",
    }
  },
  methods: {
    submitForm() {
      // 测试加密内容
      if (this.reply_content.length == 0) {
        alert("回复内容不能为空"); return
      } else if (this.reply_nickname.length == 0) {
        alert("昵称不能为空"); return
      } else if (this.reply_secretKey.length == 0) {
        alert("密钥不能为空"); return
      } else {
        const postData = {
          reply_nickname: this.reply_nickname,
          reply_content: this.reply_content,
          reply_openid: this.reply_openid.length > 15 ? this.reply_openid : "",
          reply_secretKey: this.reply_secretKey,
          relpy_imageUrl: this.relpy_imageUrl,
          reply_title: this.commentReplyOnject.comment,
          receiverOpenid: this.commentReplyOnject.openid,
          reply_pk: localStorage.getItem('blog_id'),
          level: localStorage.getItem('level'),
        };

        alert(postData)
        console.log(postData)
        this.$axios.post('/api/comment/add/', postData, {
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
  }
}
</script>