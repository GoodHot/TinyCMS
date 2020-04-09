<template>
  <div class="content t-article-view">
    <b-overlay :show="loading" rounded="sm">
      <div class="block block-fx-pop">
        <div class="block-content block-content-sm block-content-full bg-body-light">
          <div class="media py-3">
            <div class="mr-3 ml-2 overlay-container overlay-right">
              <img class="img-avatar img-avatar48" src="http://tva2.sinaimg.cn/mw600/007R6ip9gy1gd3wfle7qxj30j60onjul.jpg" alt />
            </div>
            <div class="media-body">
              <div class="row">
                <div class="col-sm-7">
                  <a class="font-w600 link-fx" href="be_pages_generic_profile.html">{{ article.title }}</a>
                  <div class="font-size-sm text-muted">{{ dateFmt(article.created_at) }}</div>
                </div>
                <div class="col-sm-5 d-sm-flex align-items-sm-center">
                  <div class="font-size-sm font-italic text-muted text-sm-right w-100 mt-2 mt-sm-0">
                    <b-button size="sm" variant="success" @click="$router.push('/article/edit/' + article.id)"><t-icon icon="edit" pack="fa" /> 编辑</b-button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="block-content" v-html="content.html">
        </div>
        <div class="block-content bg-body-light">
          <p class="font-size-sm text-muted pt-2">
            <t-icon icon="tags" class="mr-2" />
            <b-badge class="mr-1" variant="info" v-for="(tag, index) in tagSplit(article.tags)" :key="index">{{ tag }}</b-badge>
          </p>
        </div>
      </div>
    </b-overlay>
  </div>
</template>
<script>
import { getArticleById } from '@/api/article'
import moment from 'moment'
// import userInput from '@/commons/user/user-input'
// import categoryInput from '@/commons/category/category-input'

export default {
  components: {
    // userInput,
    // categoryInput
  },
  data () {
    return {
      loading: false,
      article: {
        tags: ''
      },
      content: {}
    }
  },
  methods: {
    loadArticle (id) {
      this.loading = true
      getArticleById(id).then(res => {
        this.article = res.article
        this.content = res.content
        this.loading = false
      })
    },
    dateFmt (dt) {
      return moment(dt).format('YYYY-MM-DD H:mm:ss')
    },
    tagSplit (tags) {
      return tags.split(',')
    }
  }
}
</script>