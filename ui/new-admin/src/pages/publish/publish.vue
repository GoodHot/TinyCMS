<template>
  <div class="push">
    <div class="t-publish-content">
      <input class="t-publish-title" placeholder="请输入文章标题" @focus="showOver(true)" />
      <TMarkdown @click="showOver(true)"></TMarkdown>
      <AdvancedOption v-show="showAdvance"></AdvancedOption>
      <ArticleOption @showAdvance="showAdvance = !showAdvance"></ArticleOption>
    </div>
    <div ref="focusOver" class="t-publish-focus_ovr" @click="showOver(false)"></div>
  </div>
</template>
<script>
import ArticleOption from './comp/ArticleOption'
import AdvancedOption from './comp/AdvancedOption'

export default {
  components:{
    ArticleOption,
    AdvancedOption
  },
  data() {
    return {
      showAdvance: false
    }
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