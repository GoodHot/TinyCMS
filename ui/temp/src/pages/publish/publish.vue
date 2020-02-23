<template>
  <div class="push">
    <div class="t-publish-content">
      <input class="t-publish-title" placeholder="请输入文章标题" v-model="article.title" @focus="showFocusOver = true" />
      <TMarkdown ref="markdown" @click="showFocusOver = true" v-model="article.markdown"></TMarkdown>
      <ArticleOption @showAdvance="showAdvance = !showAdvance" ref="options"  @onpublish="publishHandler" @onsave="draftHandler"></ArticleOption>
      <AdvancedOption v-show="showAdvance" ref="advance"></AdvancedOption>
    </div>
    <transition name="fade">
      <div v-if="showFocusOver" class="t-publish-focus_ovr" @click="showFocusOver = false"></div>
    </transition>
  </div>
</template>
<script>
import ArticleOption from './comp/ArticleOption'
import AdvancedOption from './comp/AdvancedOption'
import {saveArticle} from '@/api/article'
export default {
  components:{
    ArticleOption,
    AdvancedOption
  },
  data() {
    return {
      showAdvance: false,
      showFocusOver: false,
      article: {
        id: null,
        title: 'Hello World',
        markdown: '# 你好，世界',
        html: null,
        cover: null,
        category_id: null,
        description: null,
        tags: [],
        type: null
      }
    }
  },
  methods: {
    publishHandler() {
      this.save('publish')
    },
    draftHandler() {
      this.save('draft')
    },
    save(type) {
      this.article.title = this.article.title.trim()
      if (this.article.title === '') {
        this.$bvToast.toast(`请输入文章标题`, {
          title: '提示',
          variant: 'danger',
          solid: true
        })
        return
      }
      const advance = this.$refs.advance.getValue()
      const editor = this.$refs.markdown.getValue()
      const options = this.$refs.options.getValue()
      this.article.markdown = editor.markdown
      this.article.html = editor.html
      this.article.cover = advance.cover
      this.article.category_id = options.category_id
      this.article.description = advance.description
      this.article.tags = options.tags
      this.article.type = type
      saveArticle(this.article).then(res => {
        console.log('res', res)
      })
    }
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