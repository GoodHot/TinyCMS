<template>
  <div class="push">
    <div class="t-publish-content">
      <input class="t-publish-title" placeholder="请输入文章标题" @focus="showOver(true)" />
      <TMarkdown @click="showOver(true)"></TMarkdown>
      <ArticleOption></ArticleOption>
    </div>
    <div ref="focusOver" class="t-publish-focus_ovr" @click="showOver(false)"></div>
  </div>
</template>
<script>
import ArticleOption from './comp/ArticleOption'

export default {
  components:{
    ArticleOption
  },
  methods: {
    showOver(state) {
      const over = this.$refs.focusOver
      if (state) {
        this.$(over).fadeIn()
      } else {
        this.$(over).fadeOut()
      }
    }
  }
}
</script>