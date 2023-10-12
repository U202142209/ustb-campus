<template>
  <div class="blog">
    <div class="row">
      <div @click="goToUserCenter(blog.openid)" class="fl col-6 goto">
        <img :src="blog.avatar" alt="avatar" class="avatar">
        <div class="user-name">
          <div><strong v-text="blog.userName"></strong></div>
          <div class="ip-address">
            <a target="_blank" :href="'https://ip.cn/ip/' + blog.ip + '.html'" :text="blog.ip"></a>
          </div>
        </div>
      </div>
      <div class="tar col-6">
        <div v-text="blog.c_time"> </div>
        <div v-text="'帖子id:' + blog.id"></div>
      </div>
    </div>
    <div class="row">
      <div :class="'fontlarge tal ' + (!isBlogDetail && blog.cover ? 'col-8' : 'col-12')" v-text="blog.content"></div>
      <div v-if="blog.cover != ''" :class="!isBlogDetail ? 'col-4' : 'col-12'">
        <img :src="blog.cover" alt="图片加载失败" class="post-image">
      </div>
    </div>
    <div class="stats">
      <div class="stat-item">
        <van-icon name="eye-o" class="stat-icon"></van-icon>
        <div class="stat-value">{{ blog.watchNum }}</div>
      </div>
      <div class="stat-item">
        <van-icon name="like-o" class="stat-icon"></van-icon>
        <div class="stat-value">{{ blog.likeNum }}</div>
      </div>
      <div class="stat-item">
        <van-icon name="comment-o" class="stat-icon"></van-icon>
        <div class="stat-value">{{ blog.commentNum }}</div>
      </div>
      <div class="comment-time">{{ blog.comment_time }}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'BlogItems',
  props: {
    blog: {
      type: Object,
      required: true
    },
    isBlogDetail: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    goToUserCenter(openid) {
      this.$router.push('/user/' + openid);
    }
  }
}
</script>

<style scoped>
.blog {
  background-color: #f5f5f5;
  padding: 10px;
  margin-bottom: 20px;
}


.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  margin-right: 8px;
}


.nickname {
  font-size: 14px;
  color: #999;
}

.post-image {
  width: 100%;
  height: 100%;
}

.stats {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
}

.stat-item {
  display: flex;
  align-items: center;
}

.stat-icon {
  font-size: 16px;
  margin-right: 4px;
}

.stat-value {
  font-size: 14px;
}

.comment-time {
  font-size: 12px;
  color: #999;
}

.post-id {
  font-size: 12px;
  color: #999;
}
</style>
