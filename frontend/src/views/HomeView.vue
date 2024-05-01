<!--
 * @Author: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @Date: 2023-10-05 17:55:36
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2024-01-03 23:23:21
 * @FilePath: \vueLearning\src\views\Home.vue
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
<template>
  <div>
    <h4>贝壳校园墙网页版</h4>
    <!-- 搜索栏 -->
    <div class="row">
      <div class="col-2"></div>
      <div class="col-8 d-flex">
        <input v-model="keyword" class="form-control me-2" type="search" placeholder="请输入关键字搜索帖子">
        <button @click="searchData()" class="btn btn-outline-success">Search</button>
      </div>
    </div>
    <br>
    <!-- 下拉选择框 -->
    <div>
      <div class="row">
        <div class="col-1"></div>
        <div class="form-group col-5">
          <select @change="classifyoptionChanged" v-model="curclassify" name="categroy" class="form-control" id="categroy"
            required>
            <option v-for="(item, index) in classifyoptions" v-bind:key="index" v-bind:value="item.value"
              v-text="item.text"></option>
          </select>
        </div>
        <div class="form-group col-5">
          <select @change="classifyoptionChanged" v-model="order" name="categroy" class="form-control" id="order"
            required>
            <option v-for="(item, index) in orderoptions" v-bind:key="index" v-bind:value="item.value" v-text="item.text">
            </option>
          </select>
        </div>
      </div>
    </div>

    <!-- 数据 -->
    <div id="tiezidata">
      <ul id="xiaoyuanqiangdatacontent">
        <li v-for="obj in data" v-bind:key="obj.id">
          <router-link target="_blank" class="colorblack" :to="'/blog/' + obj.id">
            <blog-item :blog="obj"></blog-item>
          </router-link>
        </li>
      </ul>
    </div>
    <div v-if="loading == true" id="waiting" class="tac">
      <div>
        <img width="30%" src="https://www.intogif.com/resource/image/loading/dots.gif" alt="">
      </div>
      <h4>正在加载数据，请稍后</h4>
    </div>
    <div v-if="data.length == 0">
      <p>没有查询到数据</p>
    </div>
    <!-- 分页 -->
    <p class="tac">
      <button @click="classifyoptionChanged" class="btn btn-warning">刷新</button>
      <button @click="getMore()" class="btn btn-success">加载更多</button>
    </p>
  </div>
</template>


<script>
import Blog from '../components/BlogItems.vue';

export default {
  name: 'HomePage',
  components: {
    "blog-item": Blog
  },
  data() {
    return {
      // 新的数据
      curclassify: '%5B%22radio4%22%2C%22radio40%22%2C%22radio41%22%2C%22radio42%22%2C%22radio43%22%5D',
      classifyoptions: [
        {
          value: "%5B%22radio4%22%2C%22radio40%22%2C%22radio41%22%2C%22radio42%22%2C%22radio43%22%5D",
          text: "全部"
        },
        {
          value: "%5B%22radio4%22%2C%22radio40%22%5D",
          text: "吐槽"
        },
        {
          value: "%5B%22radio41%22%5D",
          text: "表白"
        },
        {
          value: "%5B%22radio42%22%5D",
          text: "心愿"
        },
        {
          value: "%5B%22radio43%22%5D",
          text: "知乎"
        }
      ],
      order: 0, // 帖子分类
      orderoptions: [
        {
          value: 0,
          text: "新发"
        },
        {
          value: 1,
          text: "新回"
        },
        {
          value: 2,
          text: "最热"
        },
        {
          value: 3,
          text: "精选"
        }
      ],
      length: 0,  // 分页
      search: "",
      loading: true,
      data: [],
      keyword: '',
    }
  },

  created() {
    // 在页面第一次被渲染时执行获取数据的函数
    this.initData();
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll);
  },
  unmounted() {
    window.removeEventListener('scroll', this.handleScroll);
  },
  methods: {
    // 页面初始化的函数
    initData() {
      // 发送axios请求获取数据
      this.getQuanziData()
    },
    // 获取数据
    getQuanziData(length = this.length, campus = 1) {
      this.loading = true;
      this.$axios({
        methods: 'get',
        url: "/api/post/gettaskbyTypeQuanzi/",
        params: {
          length: length,
          radioGroup: this.curclassify,
          type: this.order,
          search: this.keyword,
          campus: campus
        },
      }).then(res => {
        this.data = this.data.concat(res.data.data.taskList);
      }, err => {
        console.log(err)
        alert("加载数据失败")
      });
      this.loading = false;
    },
    getMore() {
      this.length += 10;
      // 模拟异步加载数据
      setTimeout(() => {
        this.getQuanziData();
      }, 1);
    },
    classifyoptionChanged() {
      this.length = 0;
      this.data = [];
      this.keyword = "";
      this.getQuanziData(this.length)
    },
    searchData() {
      this.length = 0;
      this.data = [];
      this.getQuanziData();
    },
    handleScroll() {
      // 判断是否触底
      const scrollTop = document.documentElement.scrollTop || document.body.scrollTop;
      const windowHeight = document.documentElement.clientHeight || document.body.clientHeight;
      const scrollHeight = document.documentElement.scrollHeight || document.body.scrollHeight;
      const distanceToBottom = scrollHeight - (scrollTop + windowHeight);
      const threshold = Math.min(300, windowHeight * 0.2); // 设置触底阈值为100px或窗口高度的20%
      if (distanceToBottom <= threshold && !this.isLoading) {
        this.getMore();
      }
    },

  }
}
</script>