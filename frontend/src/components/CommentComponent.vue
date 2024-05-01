<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-11 13:45:21
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-05-01 12:48:43
 * @FilePath: \ustb-campus\frontend\src\components\CommentComponent.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="comment-section">
    <ul class="comment-list">
      <li v-for="comment in comments" :key="comment.id">
        <div class="comment-item">
          <img @click="goToUserCenter(comment.openid)" :src="comment.avatar" alt="Avatar" class="goto avatar">
          <div @click="replySecondComment(comment)" class="comment-content">
            <div class="tal comment-header">
              <span class="user-name">{{ comment.userName }}</span>
            </div>
            <div class="tal comment-text">{{ comment.comment }}</div>
            <div v-if="comment.img">
              <a :href="comment.img" target="_blank">
                <img :src="comment.img" alt="" class="img">
              </a>
            </div>
            <div class="tal fontsmall">
              <span class="col-10 comment-time">{{ comment.c_time }}</span>
              <span @click="deleteComment(comment.id)" class="comment-time delete">删除</span>
              <span @click="replySecondComment(comment)" class="comment-time delete">回复</span>
              <span class="fr">
                <van-icon name="like-o" class="stat-icon"></van-icon>
                <!-- <span>{{ comment.like_num }}</span> -->
              </span>
            </div>
          </div>
        </div>
        <!-- 展示二级评论 -->
        <ul class="reply-list">
          <li v-for="reply in comment.commentList" :key="reply.id">
            <div class="reply-item">
              <img @click="goToUserCenter(reply.openid)" :src="reply.avatar" alt="Avatar" class="goto avatar">
              <!-- 全部增加回复评论的事件 -->
              <div @click="replySecondComment(reply)" class="reply-content">
                <div class="tal reply-header">
                  <span class="user-name">{{ reply.userName }}</span>
                </div>
                <div class="tal comment-text">{{ reply.comment }}</div>
                <div v-if="reply.img">
                  <a :href="reply.img" target="_blank">
                    <img :src="reply.img" alt="" class="img">
                  </a>
                </div>
                <div class="tal fontsmall">
                  <span class="col-10 comment-time">{{ reply.c_time }}</span>
                  <span @click="deleteComment(reply.id)" class="col-10 comment-time delete">删除</span>
                  <span @click="replySecondComment(reply)" class="comment-time delete">回复</span>
                  <span class="fr">
                    <van-icon name="like-o" class="stat-icon"></van-icon>
                    <!-- <span>{{ reply.like_num }}</span> -->
                  </span>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </li>
    </ul>
    <!--        弹窗-->
    <van-popup v-model:show="showCommentReplyModel" position="bottom" :style="{ height: 'auto' }" closeable
      close-icon="close">
      <div class="modal-content">
        <div class="modal-body">
          <p class="tac">回复:{{ commentReplyOnject.userName }}</p>
          <van-form>
            <!-- <van-field v-model="commentReplyOnject.userName" label="回复用户" readonly></van-field> -->
            <van-field v-model="commentReplyOnject.comment" label="回复评论" readonly type="textarea" rows="3"
              class="auto-height"></van-field>
            <van-field v-model="reply_nickname" label="回帖昵称" placeholder="请输入回帖昵称（必填）"></van-field>
            <van-field v-model="reply_content" label="回帖内容" type="textarea" rows="5"
              placeholder="请输入回复内容（必填）"></van-field>
            <van-field v-model="reply_secretKey" label="回帖密钥" placeholder="请输入回复密钥（必填）"></van-field>
            <van-field v-model="relpy_imageUrl" label="图片链接" placeholder="请输入图片链接（选填）"></van-field>
            <van-field v-model="reply_openid" label="openid" placeholder="请输入openid（选填）"></van-field>
            <van-button block @click="sendSecondComment" type="primary" native-type="submit">回复:{{ commentReplyOnject.userName }}</van-button>
          </van-form>
        </div>
      </div>
    </van-popup>
  </div>
</template>

<script>
import { Dialog } from 'vant';


export default {
  name: 'commentComponent',
  props: {
    comments: {
      type: Array,
      required: true
    },
    ctext: {
      type: String,
      default: "(0)"
    }
  },
  data() {
    return {
      showCommentReplyModel: false,
      commentReplyOnject: {},
      reply_nickname: "",
      reply_content: "",
      reply_openid: "",
      reply_secretKey: "",
      relpy_imageUrl: "",
    }
  },
  methods: {
    goToUserCenter(openid) {
      this.$router.push('/user/' + openid);
    },
    deleteComment(commentId) {
      var pwd = prompt("请输入操作密钥")
      if (pwd) {
        this.$axios.delete('/api/comment/deleteCommentQuanzi/' + commentId + "/" + pwd, {
          data: {
            pwd: pwd,
            id: commentId
          }
        })
          .then(res => {
            console.log(res.data)
            alert(JSON.stringify(res.data.data))
          })
          .catch(error => {
            console.log(error)
            alert("服务错误")
          });
      }
    },

    replySecondComment(comment) {
      this.commentReplyOnject = comment;
      this.showCommentReplyModel = true;
    },
    // 发送二级评论
    sendSecondComment() {
      if (this.reply_nickname.length == 0) {
        Dialog({ message: '回复昵称不能为空' });
        return;
      } if (this.reply_content.length == 0) {
        Dialog({ message: '回复内容不能为空' });
        return;
      }
      if (this.reply_secretKey.length == 0) {
        Dialog({ message: '回复密钥不能为空' });
        return;
      }
      const postData = {
        reply_nickname: this.reply_nickname,
        reply_content: this.reply_content,
        reply_secretKey: this.reply_secretKey,
        reply_openid: this.reply_openid,
        relpy_imageUrl: this.relpy_imageUrl,
        reply_title: localStorage.getItem('blog_content'),
        reply_pk: localStorage.getItem('blog_id'),
        Receiver_UserName: this.commentReplyOnject.userName,
        Pid: this.commentReplyOnject.id,
        receiverOpenid: this.commentReplyOnject.openid,
        level: 2,
      }
      this.$axios.post('/api/comment/add/', postData, {
        headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
      })
        .then(res => {
          // console.log(res.data)
          alert(JSON.stringify(res.data.data))
          location.reload();
        })
        .catch(error => {
          console.log(error)
          alert("服务错误.")
        });
    }
    // hideModel() {
    //   var modal = document.getElementById('myModal');
    //   modal.style.display = 'none';
    //   modal.classList.remove('show');
    // },

    // submitForm() {
    //   // 测试加密内容
    //   if (this.reply_content.length == 0) {
    //     alert("回复内容不能为空"); return
    //   } else if (this.reply_nickname.length == 0) {
    //     alert("昵称不能为空"); return
    //   } else if (this.reply_secretKey.length == 0) {
    //     alert("密钥不能为空"); return
    //   } else {
    //     const postData = {
    //       reply_nickname: this.reply_nickname,
    //       reply_content: this.reply_content,
    //       reply_openid: this.reply_openid.length > 15 ? this.reply_openid : "",
    //       reply_secretKey: this.reply_secretKey,
    //       relpy_imageUrl: this.relpy_imageUrl,
    //       reply_title: this.commentReplyOnject.comment,
    //       receiverOpenid: this.commentReplyOnject.openid,
    //       reply_pk: localStorage.getItem('blog_id'),
    //       level: localStorage.getItem('level'),
    //     };
    //     this.$axios.post('/api/comment/add/', postData, {
    //       headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
    //     })
    //       .then(res => {
    //         console.log(res.data)
    //         alert(JSON.stringify(res.data.data))
    //         // 刷新当前页面
    //         location.reload();
    //       })
    //       .catch(error => {
    //         console.log(error)
    //         alert("服务错误.")
    //       });
    //   }
    // },
  }
};
</script>

<style scoped>
.comment-list {
  list-style-type: none;
  padding-left: 5px;
}

.comment-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 20px;
}

.avatar {
  width: 25px;
  height: 25px;
  border-radius: 50%;
}

.img {
  width: 100%;
  height: 100%;
}

.comment-content {
  margin-left: 10px;
}

.comment-header {
  align-items: center;
  margin-bottom: 5px;
}

.user-name {
  font-weight: bold;
  font-size: small;
}

.comment-time {
  color: gray;
  margin-right: 20px;
}

.comment-text {
  margin-bottom: 6px;
}

.reply-list {
  margin-top: 10px;
  margin-bottom: 20px;
  padding-left: 40px;
}

.reply-item {
  display: flex;
  align-items: flex-start;
  margin-bottom: 15px;
}

.reply-content {
  margin-left: 10px;
}




.reply-button {
  background-color: lightgray;
  border: none;
  border-radius: 4px;
  padding: 4px 8px;
  color: white;
}

.auto-height {
  height: auto;
  min-height: 50px;
  /* 最小高度 */
  max-height: 200px;
  /* 最大高度 */
}
</style>
