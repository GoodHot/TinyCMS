<template>
  <div class="row no-gutters flex-md-10-auto">
    <div class="col-md-4 col-lg-5 col-xl-3">
      <article-list @perview="perviewHandler"></article-list>
    </div>
    <div class="col-md-8 col-lg-7 col-xl-9 bg-body-dark">
      <article-view ref="articleView"></article-view>
    </div>
  </div>
</template>
<script>
import ArticleList from './comp/ArticleList'
import ArticleView from './comp/ArticleView'

export default {
  components: {
    ArticleList,
    ArticleView
  },
  data() {
    return {
      contentSide: true
    };
  },
  methods: {
    perviewHandler (id) {
      this.$refs.articleView.loadArticle(id)
    }
  }
};
</script>