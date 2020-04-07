<template>
  <div class="row no-gutters flex-md-10-auto">
    <div class="col-md-4 col-lg-5 col-xl-3 order-md-1 bg-white">
      <article-setting></article-setting>
    </div>
    <div class="col-md-8 col-lg-7 col-xl-9 order-md-0 bg-body-dark">
      <div class="content content-full">
        <div class="block block-fx-pop">
          <div class="block-content block-content-full d-flex justify-content-between border-bottom">
              <div>
                <b-button size="sm" variant="primary"><t-icon icon="eye" /> 预览</b-button>
              </div>
              <div>
                <b-button size="sm" variant="secondary" class="mr-1"><t-icon icon="save" /> 保存草稿</b-button>
                <b-dropdown size="sm" split text="立即发布" variant="info">
                  <b-dropdown-item href="#">定时发布</b-dropdown-item>
                </b-dropdown>
              </div>
          </div>
          <div class="block-content">
            <b-form>
              <b-form-group>
                <b-form-input size="lg" autocomplete="off" placeholder="请输入文章标题"></b-form-input>
              </b-form-group>
              <b-form-group>
                <t-markdown></t-markdown>
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
      contentSide: true
    }
  }
}
</script>
<style lang="scss" scoped>
.article-title {
}
</style>