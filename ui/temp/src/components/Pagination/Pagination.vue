<template>
  <nav>
    <ul class="pagination justify-content-end">
      <li :class="'page-item ' + (pageData.pageNum == 1 ? 'disabled' : '')">
        <a class="page-link" href="javascript:void(0)" @click="pageOption(-1)">上一页</a>
      </li>
      <li :class="'page-item ' + (num == pageData.pageNum ? 'active' : '')" v-for="num of page" :key="num">
        <a class="page-link" href="javascript:void(0)" @click="pageSkip(num)">{{num}}</a>
      </li>
      <li :class="'page-item ' + (pageData.pageNum == pageData.totalPage ? 'disabled' : '')">
        <a class="page-link" href="javascript:void(0)" @click="pageOption(+1)">下一页</a>
      </li>
    </ul>
  </nav>
</template>
<script>
export default {
  model: {
    prop: 'data',
    event: 'change.data'
  },
  props: {
    data: {
      type: Object,
      default() {
        return {
          pageNum: 1,
          pageSize: 20,
          totalPage: 10,
          totalCount: 0,
        }
      }
    }
  },
  data() {
    return {
      page: [],
      showCount: 5,
      pageData: {}
    }
  },
  mounted() {
    this.initPage()
  },
  methods: {
    pageOption(opt) {
      const diff = this.pageData.pageNum + opt
      if (diff < 1 || diff > this.pageData.totalPage) {
        return
      }
      this.pageData.pageNum = diff
      this.$emit('change.data', this.pageData)
    },
    pageSkip(num) {
      this.pageData.pageNum = num
      this.$emit('change.data', this.pageData)
    },
    initPage() {
      this.pageData = JSON.parse(JSON.stringify(this.data))
      if (this.pageData.pageNum > this.pageData.totalPage) {
        this.pageData.pageNum = this.pageData.totalPage
      } else if (this.pageData < 1) {
        this.pageData = 1
      }
      if (this.showCount > this.pageData.totalPage) {
        this.showCount = this.pageData.totalPage
      }
      if (this.pageData.totalPage <= this.showCount) {
        this.setPage(1, this.pageData.totalPage)
        return
      }
      const mod = parseInt((this.showCount - 1) / 2)
      if (this.pageData.pageNum <= (mod + 1)) {
        this.setPage(1, this.showCount)
        return
      }
      let start = this.pageData.pageNum - mod
      let end = this.pageData.pageNum + mod
      if (end > this.pageData.totalPage) {
        const diff = end - this.pageData.totalPage
        start -= diff
        end = this.pageData.totalPage
      }
      this.setPage(start, end)
    },
    setPage(min, max){
      for (let i=min;i<=max;i++) {
        this.page.push(i)
      }
    }
  },
  watch: {
    data() {
      this.initPage()
    }
  }
}
</script>