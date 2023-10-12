<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 19:03:57
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2023-10-12 13:59:47
 * @FilePath: \ustb-campus\frontend\src\views\BlogDetail.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <Blog v-if="blogData" :blog="blogData" :isBlogDetail="true" />
    <h3 v-else>
      帖子不存在或已被删除
    </h3>
    <div>
      <!-- 评论区 -->
      <CommentComponent :comments="commentList" :ctext="'ctext test'"></CommentComponent>
    </div>
  </div>
</template>
<script>
import Blog from '../components/BlogItems.vue';
import CommentComponent from '../components/CommentComponent.vue'
export default {
  name: 'App',
  components: {
    "Blog": Blog,
    "CommentComponent": CommentComponent

  },
  data() {
    return {
      blogData: {},
      commentList: [],
      commentlength: 0,
      haveBlog: true,
      currentCommentsLength: 0
    };
  },
  created() {
    this.getBlogDetail()
    this.getCommentByBlogId();
  },
  methods: {
    getBlogDetail() {
      this.$axios({
        methods: 'get',
        url: "/api/post/detail/",
        params: {
          pk: this.$route.params.blog_id,
        },
      }).then(res => {
        if (res.data.data.taskList) {
          this.blogData = res.data.data.taskList[0]
        } else {
          this.haveBlog = false
        }
      }, err => {
        console.log(err)
        alert("加载数据失败")
      })
    },

    async getCommentDataByLength(pk = this.$route.params.blog_id, Ttype = 0) {
      return new Promise((resolve, reject) => {
        this.$axios({
          methods: 'get',
          url: "/api/comment/getCommentByTypeQuanzi/",
          params: {
            pk: pk,
            length: this.commentlength,
            Ttype: Ttype
          },
        }).then(res => {
          this.currentCommentsLength = res.data.data.commentList.length
          // console.log("在函数中：" + this.currentCommentsLength)
          this.commentList = this.commentList.concat(res.data.data.commentList)
          resolve(res);
        }).catch(error => {
          // 处理错误...
          reject(error);
        });
      });
    },

    async getCommentByBlogId() {
      do {
        await new Promise((resolve, reject) => {
          this.getCommentDataByLength().then(res => {
            this.commentlength += 10;
            // console.log(res);
            console.log("加载了一次评论数据...");
            resolve(res);
          }).catch(error => {
            reject(error);
          });
        });
      } while (this.currentCommentsLength > 0);
    }
  }
}
</script>