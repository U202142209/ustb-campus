<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 19:04:11
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-12 14:11:06
 * @FilePath: \ustb-campus\frontend\src\views\UserCenter.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <p class="fontlarge">用户中心</p>
    <p> {{ this.$route.params.openid }}</p>
    <div class="row">
      <span class="col-4" :class="{ 'btn btn-primary': activeIndex === 0, 'btn btn-default': activeIndex !== 0 }"
        @click="activeIndex = 0; this.refresh();">TA的发布</span>
      <span class="col-4" :class="{ 'btn btn-primary': activeIndex === 1, 'btn btn-default': activeIndex !== 1 }"
        @click="activeIndex = 1; this.refresh();">TA的评论</span>
      <span class="col-4" :class="{ 'btn btn-primary': activeIndex === 2, 'btn btn-default': activeIndex !== 2 }"
        @click="activeIndex = 2; this.refresh();">TA的的回复</span>
    </div>
    <div v-if="loading == false">
      <div v-if="activeIndex == 0">
        <ul v-if="data.length">
          <li v-for="obj in data" v-bind:key="obj.id">
            <router-link class="colorblack" :to="'/blog/' + obj.id">
              <blog-item :blog="obj"></blog-item>
            </router-link>
          </li>
        </ul>
        <p v-else v-text="'此用户没有发布的帖子记录'"></p>
      </div>
      <div v-else-if="activeIndex == 1">
        <ul v-if="data.length">
          <UserComments :comments="data"></UserComments>
        </ul>
        <p v-else v-text="'此用户没有发布的评论记录'"></p>
      </div>
      <div v-else-if="activeIndex == 2">
        <ul v-if="data.length">
          <UserComments :comments="data"></UserComments>
        </ul>
        <p v-else v-text="'此用户没有回复通知'"></p>
      </div>
    </div>
    <h2 v-else>
      加载中
    </h2>
    <p v-if="data.length">
      -----到底了，没有更多了-----
    </p>
  </div>
</template>

<script>
import Blog from '../components/BlogItems.vue';
import UserComments from '../components/UserComments.vue'

export default {
  name: 'UserCenterPage',
  components: {
    "blog-item": Blog,
    "UserComments": UserComments
  },
  data() {
    return {
      activeIndex: 0,
      type: ["gettaskbyOpenIdQuanzi", "getCommentByOpenidQuanzi", "getCommentByApplytoQuanzi"],
      page: 0,
      data: [],
      loading: false,
    }
  },
  created() {
    this.$watch('$route', this.init);
    this.init();
  },
  // 加载更多
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
  },
  unmounted() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    handleScroll() {
      // 判断是否触底
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
      const windowHeight = document.documentElement.clientHeight || document.body.clientHeight;
      const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight;
      const distanceToBottom = scrollHeight - (scrollTop + windowHeight);
      const threshold = Math.min(300, windowHeight * 0.2); // 设置触底阈值为100px或窗口高度的20%
      if (distanceToBottom <= threshold && !this.isLoading) {
        this.getMoreData();
      }
    },
    init() {
      this.activeIndex = 0;
      this.refresh()
    },
    getMoreData() {
      // 模拟异步加载数据
      setTimeout(() => {
        this.page += 10
        this.getUserData(this.page);
      }, 100);
    },
    refresh() {
      // 模拟异步加载数据
      setTimeout(() => {
        this.page = 0;
        this.data = [];
        this.getUserData(this.page);
      }, 100);
    },

    getUserData(page = 0) {
      if (window.navigator.webdriver == true) {
        alert("Warnning ! 系统检测到爬虫，无法获取数据。")
      } else {
        this.loading = true;
        this.$axios({
          methods: 'get',
          url: "/api/user/" + this.type[this.activeIndex],
          params: {
            openid: this.$route.params.openid,
            length: page
          }
        }).then(res => {
          if (this.activeIndex == 0) {
            this.data = this.data.concat(res.data.data.taskList);
          } else {
            this.data = this.data.concat(res.data.data.commentList);
            console.log("获取评论数据")
            console.log(this.data)
          }
          console.log(this.data.length)
        }, err => {
          console.log(err)
          alert("加载数据失败")
        })
        this.loading = false;
      }
    }
  },
}
</script>