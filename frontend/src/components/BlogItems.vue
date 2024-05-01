<template>
  <div class="blog">
    <div class="row fontsmall">
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
      <div :class="'tal ' + (!isBlogDetail && blog.cover ? 'col-8' : 'col-12')" v-text="blog.content"></div>
      <div v-if="getImages.length" :class="!isBlogDetail ? 'col-4' : 'col-12'">
        <div v-if="isBlogDetail == false">
          <img :src="getImages[0]" alt="图片加载失败" class="post-image fontsmall">
        </div>
        <div v-else>
          <div v-for=" src in getImages " v-bind:key="src">
            <img :src="src" alt="图片加载失败" class="post-image fontsmall">
          </div>
        </div>
      </div>
    </div>
    <div class="stats fontsmall">
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
    <div class="tar" v-if="isBlogDetail">
      <span @click="deleteBlog(blog.id)" class="delete">删除</span>
    </div>
  </div>
</template>

<script>
import { jiama, splitImageUrl } from '@/utils';
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
  computed: {
    // 计算属性
    getImages: function () {
      var res = splitImageUrl(this.blog.cover)
      return res ? res : []
    }
  },
  methods: {
    goToUserCenter(openid) {
      this.$router.push('/user/' + openid);
    },
    deleteBlog(blogId) {
      var pwd = prompt("请输入操作密钥")
      if (pwd) {
        this.$axios.delete('/api/post/delete/' + blogId + "/" + pwd + "/" + jiama(blogId), {})
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
  width: 30px;
  height: 30px;
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
  margin-right: 4px;
}

.comment-time {
  color: #999;
}

.post-id {
  color: #999;
}
</style>
