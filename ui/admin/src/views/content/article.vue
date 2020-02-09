<template>
  <a-card :bordered="false">
    <a-row :gutter="16">
      <a-col :span="17">
        <a-input placeholder="文章标题" class="form-item" />
        <div ref="editorSection"></div>
      </a-col>
      <a-col :span="7">
        <a-upload
          name="avatar"
          listType="picture-card"
          class="avatar-uploader"
          :showUploadList="false"
          action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
          :beforeUpload="beforeUpload"
          @change="handleChange"
        >
          <img v-if="coverUrl" :src="coverUrl" alt="avatar" />
          <div v-else>
            <a-icon :type="coverLoading ? 'loading' : 'plus'" />
            <div class="ant-upload-text">封面图片</div>
          </div>
        </a-upload>
        <a-tree-select
          class="form-item"
          style="width:100%"
          :dropdownStyle="{ maxHeight: '400px', overflow: 'auto' }"
          :treeData="channelData"
          placeholder="选择频道"
          treeDefaultExpandAll
        >
          <span style="color: #08c" slot="title" slot-scope="{key, value}" v-if="key='0-0-1'">
            Child Node1 {{ value }}
          </span>
        </a-tree-select>
        <a-textarea
          class="form-item"
          placeholder="简介摘要"
          :autosize="{ minRows: 4, maxRows: 6 }"
        />
        <div class="form-item">
          <a-date-picker
            showTime
            placeholder="定时发布"
            style="width: 100%"
            format="YYYY-MM-DD HH:mm:ss"
          />
        </div>
        <div class="form-item">
          <a-select mode="tags" :tokenSeparators="[',']" style="width:100%;" placeholder="文章标签"></a-select>
        </div>
        <a-button type="primary" class="btn-right" @click="publish">发布</a-button>
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
import { getChannelTree } from '@/api/channel'

export default {
  data () {
    return {
      markdown: null,
      coverUrl: null,
      coverLoading: false,
      channelData: []
    }
  },
  mounted () {
    /* eslint-disable no-new */
    this.markdown = new Editor({
      el: this.$refs.editorSection,
      initialEditType: 'markdown',
      previewStyle: 'vertical',
      height: '700px',
      exts: ['scrollSync']
    })
    this.loadChannelTree()
  },
  methods: {
    publish () {
      console.log(this.markdown.getMarkdown())
    },
    beforeUpload () {},
    handleChange () {},
    loadChannelTree () {
      getChannelTree().then(res => {
        this.channelData = res.tree
      })
    }
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
