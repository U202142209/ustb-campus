<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-08 19:03:57
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-05-01 12:46:41
 * @FilePath: \ustb-campus\frontend\src\views\BlogDetail.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <div v-if="blogData">
      <Blog :blog="blogData" :isBlogDetail="true" />
      <span @click="replyBlog(blogData, 1)" class="cp">发表评论</span>
    </div>
    <h3 v-else>
      帖子不存在或已被删除
    </h3>
    <div>
      <!-- 评论区 -->
      <CommentComponent :comments="commentList" :ctext="'ctext test'"></CommentComponent>
    </div>
    <!--        弹窗-->
    <div id="myModal" class="modal" tabindex="-1" role="dialog">

    </div>

    <van-popup v-model:show="showReplyBlogPopup" position="bottom" :style="{ padding: '4px' }" closeable
      close-icon="close">
      <div class="modal-dialog modal-dialog-centered" role="document">
        <div class="modal-content">
          <div class="modal-header">
            <p class="modal-title">回复:{{ commentReplyOnject.userName }}</p>
            <p class="fontsmall">{{ commentReplyOnject.id }}</p>
            <button @click="hideModel()" type="button" class="close" data-dismiss="modal" aria-label="Close">
              <span aria-hidden="true">&times;</span>
            </button>
          </div>
          <div class="modal-body">
            <van-form @submit="submitFormBlog">
              <van-field v-model="reply_nickname" label="回帖昵称" placeholder="请输入回帖昵称（必填）"></van-field>
              <van-field v-model="reply_content" label="回帖内容" type="textarea" rows="5"
                placeholder="请输入回复内容（必填）"></van-field>
              <van-field v-model="reply_secretKey" label="回帖密钥" placeholder="请输入回复密钥（必填）"></van-field>
              <van-field v-model="relpy_imageUrl" label="图片链接" placeholder="请输入图片链接（选填）"></van-field>
              <van-field v-model="reply_openid" label="openid" placeholder="请输入openid（选填）"></van-field>
              <van-button type="primary" block native-type="submit">发表评论</van-button>
            </van-form>
          </div>
        </div>
      </div>
    </van-popup>

  </div>
  <br><br><br><br>
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
      showReplyBlogPopup: false,
      blogData: {},
      commentList: [],
      commentlength: 0,
      haveBlog: true,
      currentCommentsLength: 0,
      // 回复评论
      commentReplyOnject: {},
      reply_nickname: "",
      reply_content: "",
      reply_openid: "",
      reply_secretKey: "",
      relpy_imageUrl: "",
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
          this.blogData = res.data.data.taskList[0];
          // 缓存
          localStorage.setItem('blog_id', this.blogData.id);
          localStorage.setItem('blog_content', this.blogData.content);
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
    },

    // 回复评论
    replyBlog(comment, level) {
      // var modal = document.getElementById('myModal');
      // modal.style.display = 'block';
      // modal.classList.add('show');
      this.showReplyBlogPopup = true;
      console.log(11111111111111111111111);
      console.log(this.showReplyBlogPopup);
      localStorage.setItem('level', level);
      this.commentReplyOnject = comment;
      this.commentReplyOnject.comment = comment.content
    },
    hideModel() {
      var modal = document.getElementById('myModal');
      modal.style.display = 'none';
      modal.classList.remove('show');
    },
    submitFormBlog() {
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

        this.$axios.post('/api/comment/add/', postData, {
          headers: { 'Content-Type': 'application/x-www-form-urlencoded' }
        })
          .then(res => {
            console.log(res.data)
            alert(JSON.stringify(res.data.data))
            // 刷新当前页面
            if (JSON.stringify(res.data.data).length < 4) {
              location.reload();
            }
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