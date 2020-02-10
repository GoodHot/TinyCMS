<template>
  <a-card :bordered="false">
    <a-row :gutter="16">
      <a-col :span="17">
        <a-input placeholder="文章标题" class="form-item" v-model="publishData.title" />
        <div ref="editorSection"></div>
      </a-col>
      <a-col :span="7">
        <a-upload
          name="file"
          accept=".jpg,.jpeg,.png,.gif"
          listType="picture-card"
          :multiple="false"
          class="avatar-uploader"
          :showUploadList="false"
          :action="actionURL('/upload')"
          @change="uploadHandleChange"
        >
          <img class="upload-cover" v-if="publishData.coverView" :src="publishData.coverView" alt="avatar" />
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
          v-model="publishData.channelId"
        >
        </a-tree-select>
        <a-textarea
          class="form-item"
          placeholder="简介摘要"
          :autosize="{ minRows: 4, maxRows: 6 }"
          v-model="publishData.description"
        />
        <div class="form-item">
          <a-date-picker
            showTime
            placeholder="定时发布"
            style="width: 100%"
            :format="dateFormat"
            @change="publishTimeChange"
          />
        </div>
        <div class="form-item">
          <a-select mode="tags" :tokenSeparators="[',']" style="width:100%;" placeholder="文章标签" v-model="publishData.tags"></a-select>
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
import { uploadBase64 } from '@/api/upload'

export default {
  data () {
    return {
      markdown: null,
      coverLoading: false,
      channelData: [],
      dateFormat: 'YYYY-MM-DD HH:mm:ss',
      publishData: {
        title: '',
        html: '',
        markdown: '',
        cover: '',
        coverView: '',
        channelId: '',
        description: '',
        publishTime: '',
        tags: [],
        status: null
      }
    }
  },
  mounted () {
    /* eslint-disable no-new */
    const _this = this
    this.markdown = new Editor({
      el: this.$refs.editorSection,
      initialEditType: 'markdown',
      previewStyle: 'vertical',
      height: '700px',
      exts: ['scrollSync'],
      hooks: {
        addImageBlobHook (blob, callback) {
          _this.imageUpload(blob, callback)
        }
      }
    })
    this.loadChannelTree()
  },
  methods: {
    imageUpload (blob, mdCall) {
      const _this = this
      const reader = new FileReader()
      reader.addEventListener('load', () => {
        console.log(reader.result)
        uploadBase64({
          base64: reader.result
        }).then(res => {
          mdCall(_this.assetsURL(res.path), 'image')
        })
      })
      reader.readAsDataURL(blob)
    },
    publish () {
      this.publishData.markdown = this.markdown.getMarkdown()
      this.publishData.html = this.markdown.getHtml()
      console.log(this.publishData)
    },
    uploadHandleChange (info) {
      const status = info.file.status
      if (status === 'done') {
        if (info.file.response.status !== 200) {
          this.$message.error(`${info.file.name} , ${info.file.response.msg}`)
          return
        }
        this.publishData.cover = info.file.response.data.path
        this.publishData.coverView = this.assetsURL(info.file.response.data.path)
        this.$message.success(`${info.file.name} file uploaded successfully.`)
      } else if (status === 'error') {
        this.$message.error(`${info.file.name} file upload failed.`)
      }
    },
    loadChannelTree () {
      getChannelTree().then(res => {
        this.channelData = res.tree
      })
    },
    publishTimeChange (value) {
      this.publishData.publishTime = value.format(this.dateFormat)
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
.upload-cover {
  max-width: 100%;
  max-height: 200px;
}
</style>
