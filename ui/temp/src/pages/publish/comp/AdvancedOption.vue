<template>
  <div class="t-publish-advanced">
    <div class="row justify-content-center">
      <div class="col-md-8 col-lg-6">
        <b-form>
          <b-form-group
            label="SEO标题"
            description="当不填写时，会自动解析标题名称"
          >
            <b-form-input placeholder="SEO标题" v-model="data.seoTitle"></b-form-input>
          </b-form-group>
          <b-form-group
            label="文章简介"
            description="当不填写时，会自动从文章中获取解析简介信息"
          >
            <b-form-textarea
            v-model="data.description"
              rows="3"
              max-rows="6"
            ></b-form-textarea>
          </b-form-group>

          <b-form-group
            label="封面图"
            description="当不上传时，会自动获取文章中的第一张图片作为封面图"
          >
            <TUpload v-model="data.cover" @uploaded="uploadHandler"></TUpload>
          </b-form-group>
        </b-form>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  model: {
    prop: 'data',
    event: 'change.data'
  },
  props: {
    data: {
      type: Object,
      default() {
        return {}
      }
    }
  },
  methods: {
    uploadHandler(src) {
      this.data.cover = src
    },
    getValue() {
      return {
        cover: this.data.cover,
        description: this.data.description,
        seo_title: this.data.seoTitle
      }
    }
  }
}
</script>