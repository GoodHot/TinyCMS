<template>
  <div class="row no-gutters flex-md-10-auto">
    <div class="col-md-4 col-lg-5 col-xl-3">
      <article-list></article-list>
    </div>
    <div class="col-md-8 col-lg-7 col-xl-9 bg-body-dark">
      <article-view></article-view>
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
  }
};
</script>