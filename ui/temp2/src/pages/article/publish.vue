<template>
  <div class="row no-gutters flex-md-10-auto">
    <div class="col-md-4 col-lg-5 col-xl-3 order-md-1 bg-white">
      <article-setting ref="setting" style="position: sticky; top: 70px" @autosave="autosaveHandler"></article-setting>
    </div>
    <div class="col-md-8 col-lg-7 col-xl-9 order-md-0 bg-body-dark">
      <div class="content content-full">
        <div class="block block-fx-pop">
          <div class="block-content block-content-full d-flex justify-content-between border-bottom">
            <b-button size="sm" variant="primary"><t-icon icon="eye" /> 预览</b-button>
            <div>
              <b-button size="sm" variant="secondary" class="mr-1" @click="saveArticle('draft')"><t-icon icon="save" /> 保存草稿</b-button>
              <b-button size="sm" variant="info" @click="saveArticle('publish')"><t-icon icon="paper-plane" pack="far" /> 立即发布</b-button>
            </div>
          </div>
          <div class="block-content">
            <b-form>
              <b-form-group>
                <b-form-input size="lg" autocomplete="off" v-model="article.title" placeholder="请输入文章标题"></b-form-input>
              </b-form-group>
              <b-form-group>
                <t-markdown ref="markdown" v-model="article.markdown"></t-markdown>
              </b-form-group>
            </b-form>
          </div>
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
      contentSide: true,
      article: {
        title: 'Hello World',
        markdown: '# 你好，世界'
      }
    }
  },
  methods: {
    autosaveHandler () {
      this.$emit('onloading', true, '保存中……')
      console.log('save')
    },
    saveArticle (type) {
      const content = this.$refs.markdown.getValue()
      const setting = this.$refs.setting.getValue()
      const tags = []
      if (setting.meta.tags && setting.meta.tags.length > 0) {
        setting.meta.tags.map(t => {
          tags.push({
            name: t
          })
        })
      }
      const post = {
        article: {
          title: this.article.title,
          seo_title: setting.seoSetting.seo_title,
          seo_description: setting.seoSetting.seo_description,
          seo_keyword: setting.seoSetting.seo_keyword,
          category_id: setting.meta.category,
          description: setting.seoSetting.description,
          cover: setting.cover,
          status: type === 'draft' ? 2 : 1,
          author: setting.publishSetting.authorId,
          template: setting.publishSetting.template,
          visibility: setting.publishSetting.visibility
        },
        content: {
          markdown: content.markdown,
          html: content.html
        },
        get_cover: setting.writeSetting.getCover,
        get_net_image: setting.writeSetting.getNetImage,
        get_description: setting.writeSetting.getDescription,
        tags
      }
      console.log(post)
    }
  }
}
</script>