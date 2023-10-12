<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-11 13:45:21
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-11 23:18:39
 * @FilePath: \ustb-campus\frontend\src\components\CommentComponent.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div class="comment-section">
    <ul class="comment-list">
      <li v-for="comment in comments" :key="comment.id">
        <div class="comment-item">
          <img @click="goToUserCenter(comment.openid)" :src="comment.avatar" alt="Avatar" class="goto avatar">
          <div class="comment-content">
            <div class="tal comment-header">
              <span class="user-name">{{ comment.userName }}</span>
            </div>
            <div class="tal comment-text">{{ comment.comment }}</div>
            <div class="tal">
              <span class="col-10 comment-time">{{ comment.c_time }}</span>
              <span class="fr">
                <van-icon name="like-o" class="stat-icon"></van-icon>
                <span>{{ comment.like_num }}</span>
              </span>
            </div>
          </div>
        </div>
        <!-- 展示二级评论 -->
        <ul class="reply-list">
          <li v-for="reply in comment.commentList" :key="reply.id">
            <div class="reply-item">
              <img @click="goToUserCenter(reply.openid)" :src="reply.avatar" alt="Avatar" class="goto avatar">
              <div class="reply-content">
                <div class="tal reply-header">
                  <span class="user-name">{{ reply.userName }}</span>
                </div>
                <div class="tal comment-text">{{ reply.comment }}</div>
                <div class="tal">
                  <span class="col-10 comment-time">{{ reply.c_time }}</span>
                  <span class="fr">
                    <van-icon name="like-o" class="stat-icon"></van-icon>
                    <span>{{ reply.like_num }}</span>
                  </span>
                </div>
              </div>
            </div>
          </li>
        </ul>
      </li>
    </ul>
  </div>
</template>

<script>
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
  methods: {
    goToUserCenter(openid) {
      this.$router.push('/user/' + openid);
    }
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
  width: 40px;
  height: 40px;
  border-radius: 50%;
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
  margin-bottom: 10px;
  font-size: large;
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
</style>
