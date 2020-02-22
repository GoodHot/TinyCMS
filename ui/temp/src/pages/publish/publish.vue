<template>
  <div class="push">
    <div class="t-publish-content">
      <input class="t-publish-title" placeholder="请输入文章标题" @focus="showFocusOver = true" />
      <TMarkdown @click="showFocusOver = true"></TMarkdown>
      <ArticleOption @showAdvance="showAdvance = !showAdvance"></ArticleOption>
      <AdvancedOption v-show="showAdvance"></AdvancedOption>
    </div>
    <transition name="fade">
      <div v-if="showFocusOver" class="t-publish-focus_ovr" @click="showFocusOver = false"></div>
    </transition>
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
      showAdvance: false,
      showFocusOver: false
    }
  },
  methods: {
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to /* .fade-leave-active below version 2.1.8 */ {
  opacity: 0;
}
</style>