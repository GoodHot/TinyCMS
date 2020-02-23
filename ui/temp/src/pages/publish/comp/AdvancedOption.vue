<template>
  <div class="t-publish-advanced">
    <div class="row justify-content-center">
      <div class="col-md-8 col-lg-6">
        <b-form>
          <b-form-group
            label="文章简介"
            description="当不填写时，会自动从文章中获取解析简介信息"
          >
            <b-form-textarea
              rows="3"
              max-rows="6"
            ></b-form-textarea>
          </b-form-group>

          <b-form-group
            label="封面图"
            description="当不上传时，会自动获取文章中的第一张图片作为封面图"
          >
            <TUpload @uploaded="uploadHandler"></TUpload>
          </b-form-group>
        </b-form>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      cover: null
    }
  },
  methods: {
    uploadHandler(src) {
      this.cover = src
    }
  }
}
</script>