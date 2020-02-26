<template>
  <div class="push">
    <div class="t-publish-content">
      <input class="t-publish-title" placeholder="请输入文章标题" v-model="article.title" @focus="showFocusOver = true" />
      <TMarkdown ref="markdown" @click="showFocusOver = true" v-model="article.markdown"></TMarkdown>
      <ArticleOption @showAdvance="showAdvance = !showAdvance" ref="options" v-model="optionsData"  @onpublish="publishHandler" @onsave="draftHandler"></ArticleOption>
      <AdvancedOption v-show="showAdvance" ref="advance" v-model="advancedData"></AdvancedOption>
    </div>
    <transition name="fade">
      <div v-if="showFocusOver" class="t-publish-focus_ovr" @click="showFocusOver = false"></div>
    </transition>
  </div>
</template>
<script>
import ArticleOption from './comp/ArticleOption'
import AdvancedOption from './comp/AdvancedOption'
import {saveArticle, getArticleById} from '@/api/article'
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
        seo_title: '',
        markdown: '# 你好，世界',
        html: null,
        cover: null,
        category_id: null,
        description: null,
        tags: [],
        type: null
      },
      advancedData: {
        cover: '',
        seoTitle: '',
        description: ''
      },
      optionsData: {
        tags: [],
        categoryId: 0
      }
    }
  },
  mounted() {
    this.loadArticle()
  },
  methods: {
    loadArticle() {
      const id = this.$route.query.id
      if (!id) {
        return
      }
      getArticleById(id).then(res => {
        this.article = res.article
        this.advancedData = {
          cover: this.article.cover,
          seoTitle: this.article.seo_title,
          description: this.article.description
        }
        this.optionsData = {
          tags: this.article.tags,
          categoryId: this.article.category_id
        }
        console.log(this.article)
      })
    },
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
          toaster: 'b-toaster-top-center',
          solid: true
        })
        return
      }
      const advance = this.$refs.advance.getValue()
      const editor = this.$refs.markdown.getValue()
      const options = this.$refs.options.getValue()
      this.article.seo_title = advance.seo_title
      this.article.markdown = editor.markdown
      this.article.html = editor.html
      this.article.cover = advance.cover
      this.article.category_id = options.category_id
      this.article.description = advance.description
      this.article.tags = options.tags
      this.article.type = type
      saveArticle(this.article).then(() => {
        this.$bvToast.toast('保存成功', {
          title: '提示',
          toaster: 'b-toaster-top-center',
          solid: true
        })
        window.setTimeout(() => {
          this.$router.push('/content')
        }, 1500)
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