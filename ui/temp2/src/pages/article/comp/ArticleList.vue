<template>
  <div class="content" style="height: calc(100vh - 70px); overflow-y: scroll">
    <div class="d-md-none push">
      <b-button block variant="hero-primary">
        Article List
      </b-button>
    </div>
    <div class="d-md-block push d-none t-article-list">
      <b-button block variant="hero-success" class="mb-3"><t-icon icon="plus" class="mr-1" /> 发布新文章</b-button>

      <div class="mb-3">
        <b-input-group>
          <b-form-input v-model="search.keyword" class="border-0" placeholder="文章搜索" :disabled="advancedSearch"></b-form-input>
          <b-input-group-append>
            <b-button class="bg-white px-3" variant="light" size="sm" text="Button" :disabled="advancedSearch"><t-icon size="fw" icon="search" /></b-button>
            <b-button class="bg-white px-3" variant="light" size="sm" text="Button" @click="advancedSearch = !advancedSearch">高级搜索</b-button>
          </b-input-group-append>
        </b-input-group>
        <b-form class="p-3 bg-white" v-show="advancedSearch">
          <b-form-group>
            <b-form-input v-model="search.keyword" class="form-control-alt" placeholder="文章搜索" size="sm" ></b-form-input>
          </b-form-group>
          <b-form-group>
            <user-input size="sm" inputClass="form-control-alt"></user-input>
          </b-form-group>
          <b-form-group>
            <category-input size="sm" inputClass="form-control-alt"></category-input>
          </b-form-group>
          <b-form-group>
            <b-button size="sm" block variant="primary">搜索</b-button>
          </b-form-group>
        </b-form>
      </div>
      <b-list-group class="font-size-sm">
        <b-list-group-item href="#" class="ribbon ribbon-light" v-for="article in articles" :key="article.id" @mouseover="showTime(article, 'full')" @mouseout="showTime(article, 'ago')">
          <div class="ribbon-box" v-if="article.status === 2"><t-icon icon="save" pack="far" v-if="article.status === 2" /></div>
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
      advancedSearch: false,
      search: {
        keyword: ''
      }
    }
  },
  mounted () {
    this.loadArticlePage()
  },
  methods: {
    loadArticlePage () {
      articlePage({
        page: 1
      }).then(res => {
        const tmp = res.page.list
        tmp.map(art => {
          art.created_at_fmt = this.fmtDateAgo(art.created_at)
        })
        this.articles = tmp
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
    }
  }
}
</script>