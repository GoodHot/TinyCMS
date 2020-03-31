<template>
  <div class="content">
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
                    <b-button size="sm" variant="success"><t-icon icon="edit" pack="fa" /> 编辑内容</b-button>
                    <b-button size="sm" class="ml-1"><t-icon icon="cog" pack="fa" /> 选项设置</b-button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="block-content" v-html="content.html">
        </div>
      </div>
    </b-overlay>
  </div>
</template>
<script>
import { getArticleById } from '@/api/article'
import moment from 'moment'

export default {
  data () {
    return {
      loading: false,
      article: {},
      content: {}
    }
  },
  methods: {
    loadArticle (id) {
      this.loading = true
      getArticleById(id).then(res => {
        console.log(res)
        this.article = res.article
        this.content = res.content
        this.loading = false
      })
    },
    dateFmt (dt) {
      return moment(dt).format('YYYY-MM-DD H:mm:ss')
    }
  }
}
</script>