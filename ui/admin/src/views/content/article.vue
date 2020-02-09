<template>
  <a-card :bordered="false">
    <a-row :gutter="16">
      <a-col :span="17">
        <a-input placeholder="文章标题" class="form-item" />
        <div ref="editorSection"></div>
      </a-col>
      <a-col :span="7">
        <a-input placeholder="选择频道" class="form-item" />
        <a-upload
          name="avatar"
          listType="picture-card"
          class="avatar-uploader"
          :showUploadList="false"
          action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
          :beforeUpload="beforeUpload"
          @change="handleChange"
        >
          <img v-if="imageUrl" :src="imageUrl" alt="avatar" />
          <div v-else>
            <a-icon :type="loading ? 'loading' : 'plus'" />
            <div class="ant-upload-text">封面图片</div>
          </div>
        </a-upload>
        <a-textarea
          class="form-item"
          placeholder="简介摘要"
          :autosize="{ minRows: 2, maxRows: 6 }"
        />
        <a-button type="primary" class="btn-right">发布</a-button>
        <a-button type="dashed">保存草稿</a-button>
      </a-col>
    </a-row>
  </a-card>
</template>

<script>
import 'codemirror/lib/codemirror.css'
import 'tui-editor/dist/tui-editor.css'
import 'tui-editor/dist/tui-editor-contents.css'
import 'highlight.js/styles/github.css'
import Editor from 'tui-editor'

export default {
  data () {
    return {
      markdown: null
    }
  },
  mounted () {
    /* eslint-disable no-new */
    this.markdown = new Editor({
      el: this.$refs.editorSection,
      initialEditType: 'markdown',
      previewStyle: 'vertical',
      height: '700px'
    })
  }
}
</script>
<style scoped>
.form-item {
  margin-bottom: 16px;
}
.btn-right {
  margin-right: 16px;
}
.avatar-uploader > .ant-upload {
  width: 100%;
  height: 130px;
}
.avatar-uploader /deep/ .ant-upload-select-picture-card {
  width: 100%;
  margin-bottom: 16px;
}
.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card .ant-upload-text {
  color: #666;
}
</style>
