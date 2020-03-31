<template>
  <div class="content" style="height: calc(100vh - 70px); overflow-y: scroll">
    <div class="d-md-none push">
      <b-button block variant="hero-primary">
        Article List
      </b-button>
    </div>
    <div class="d-md-block push d-none t-article-list">
      <b-button block variant="hero-success" class="mb-3" @click="$router.push('/article/publish')"><t-icon icon="plus" class="mr-1" /> 发布新文章</b-button>

      <div class="mb-3">
        <b-input-group>
          <b-form-input v-model="search.keyword" class="border-0" placeholder="文章搜索" :disabled="advancedSearch"></b-form-input>
          <b-input-group-append>
            <b-button class="bg-white px-3" variant="light" size="sm" text="Button" :disabled="advancedSearch" @click="searchSubmit"><t-icon size="fw" icon="search" /></b-button>
            <b-button class="bg-white px-3" variant="light" size="sm" text="Button" @click="advancedSearch = !advancedSearch">高级搜索</b-button>
          </b-input-group-append>
        </b-input-group>
        <b-form class="p-3 bg-white" v-show="advancedSearch">
          <b-form-group>
            <b-form-input v-model="search.keyword" class="form-control-alt" placeholder="文章搜索" size="sm" ></b-form-input>
          </b-form-group>
          <b-form-group>
            <user-input size="sm" inputClass="form-control-alt" @change="userChangeHandler"></user-input>
          </b-form-group>
          <b-form-group>
            <category-input size="sm" inputClass="form-control-alt" @change="categoryChangeHandler"></category-input>
          </b-form-group>
          <b-form-group>
            <b-button size="sm" block variant="primary" @click="searchSubmit">搜索</b-button>
          </b-form-group>
        </b-form>
      </div>
      <b-overlay :show="loading" rounded="sm">
        <b-alert show variant="info" v-if="!articles || articles.length === 0">
          <h4 class="alert-heading">没有文章</h4>
          <p>
            抱歉，没有找到任何文章，你可以立即发布一篇新文章
          </p>
        </b-alert>
        <b-list-group class="font-size-sm">
          <b-list-group-item href="#" @click="onperview(article.id)" class="ribbon ribbon-light ribbon-modern" v-for="article in articles" :key="article.id" @mouseover="showTime(article, 'full')" @mouseout="showTime(article, 'ago')">
            <div class="ribbon-box" v-if="article.status === 2 || article.id === activeId">
              <t-icon icon="save" pack="far" v-if="article.status === 2" />
              <span class="mr-2" v-if="article.id === activeId && article.status === 2" ></span>
              <t-icon icon="check" class="text-success" v-if="article.id === activeId" />
            </div>
            <p class="font-size-h6 font-w700 mb-0">{{ article.title }}</p>
            <p class="text-muted mb-2">
              {{ article.description }}
            </p>
            <div class="d-flex justify-content-between mb-2">
              <a href="#" class="t-article-list-category">{{ article.category_name }}</a>
              <div>
                <b-badge variant="secondary" href="#" pill v-for="tag in splitTag(article.tags)" :key="tag" class="mr-1">{{ tag }}</b-badge>
              </div>
            </div>
            <small class="d-flex justify-content-between">
              <time datetime="article.created_at">{{ article.created_at_fmt }}</time>
              <span>Create by {{ article.author_name }}</span>
            </small>
          </b-list-group-item>
        </b-list-group>
      </b-overlay>
    </div>
  </div>
</template>
<script>
import { articlePage } from '@/api/article'
import moment from 'moment'
import categoryInput from '@/commons/category/category-input'
import userInput from '@/commons/user/user-input'

export default {
  components: {
    categoryInput,
    userInput
  },
  data () {
    return {
      articles: [],
      activeId: 0,
      advancedSearch: false,
      loading: false,
      search: {
        keyword: '',
        user: 0,
        category: 0,
        page: 1
      }
    }
  },
  mounted () {
    this.loadArticlePage('firstPerview')
  },
  methods: {
    loadArticlePage (perview) {
      this.loading = true
      const param = {}
      if (this.advancedSearch) {
        param.page = this.search.page
        param.keyword = this.search.keyword
        param.user = this.search.user
        param.category = this.search.category
      } else {
        param.page = this.search.page
        param.keyword = this.search.keyword
      }
      articlePage(param).then(res => {
        const tmp = res.page.list
        tmp.map(art => {
          art.created_at_fmt = this.fmtDateAgo(art.created_at)
        })
        this.articles = tmp
        this.loading = false
        if (perview && this.articles && this.articles.length > 0) {
          this.onperview(this.articles[0].id)
        }
      })
    },
    fmtDateAgo (dt) {
      return moment(dt).fromNow()
    },
    showTime (article, type) {
      if (type === 'ago') {
        article.created_at_fmt = this.fmtDateAgo(article.created_at)
      } else {
        article.created_at_fmt = moment(article.created_at).format('YYYY-MM-DD H:mm:ss')
      }
    },
    splitTag (tags) {
      return tags.split(',')
    },
    userChangeHandler (id) {
      this.search.user = id
    },
    categoryChangeHandler (id) {
      this.search.category = id
    },
    searchSubmit () {
      this.search.page = 1
      this.loadArticlePage()
    },
    onperview (id) {
      this.activeId = id
      this.$emit('perview', id)
    }
  }
}
</script>