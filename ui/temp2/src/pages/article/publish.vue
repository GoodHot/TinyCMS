<template>
  <div class="row no-gutters flex-md-10-auto">
    <div class="col-md-4 col-lg-5 col-xl-3 order-md-1 bg-white">
      <article-setting></article-setting>
    </div>
    <div class="col-md-8 col-lg-7 col-xl-9 order-md-0 bg-body-dark">
      <div class="content content-full">
        <div class="block block-fx-pop">
          sdfsdf
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import ArticleSetting from './comp/ArticleSetting'
export default {
  components: {
    ArticleSetting
  },
  data() {
    return {
      contentSide: true
    }
  }
}
</script>