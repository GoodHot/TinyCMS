<template>
  <b-modal v-model="visible" title="分类" ok-title="确定" cancel-title="取消" size="lg" no-close-on-backdrop @ok="okHandler" :ok-disabled="loading">
    <b-overlay :show="loading" rounded="sm">
      <t-form :form="form" ref="categoryForm">
      </t-form>
    </b-overlay>
  </b-modal>
</template>
<script>
import { vType } from '@/utils/valid'
import { saveCategory } from '@/api/category'

export default {
  data () {
    return {
      visible: false,
      loading: false,
      form: [
        {
          name: 'name',
          value: '',
          type: 'input',
          label: '分类名称',
          placeholder: '这个分类叫什么名字？',
          valid: [vType.require('请输入分类名称'), vType.length(1, 255, '分类名称长度为1~255个字符')]
        },
        {
          name: 'path',
          value: '',
          type: 'input',
          label: '路径',
          placeholder: '好的路径有利于SEO哦',
          valid: [vType.require('请输入分类路径'), vType.regex('^[a-zA-Z0-9\\-]+$', '路径只能为字母和数字或横杠(-)'), vType.length(1, 255, '路径长度为1~255个字符')]
        },
        {
          name: 'seo_title',
          value: '',
          type: 'input',
          label: 'SEO标题',
          placeholder: '用于SEO的标题，可以配合分类名称一起使用',
          valid: [vType.length(1, 255, 'SEO标题长度为1~255个字符')]
        },
        {
          name: 'keyword',
          value: '',
          type: 'input',
          label: '关键字',
          placeholder: '好的关键字可以提升SEO效果',
          valid: [vType.length(1, 255, '关键字长度为1~255个字符')]
        },
        {
          name: 'desccription',
          value: '',
          type: 'textarea',
          label: '描述',
          placeholder: '输入分类的描述',
          valid: [vType.length(1, 255, '描述长度为1~255个字符')]
        },
        {
          name: 'remark',
          value: '',
          type: 'textarea',
          label: '备注',
          placeholder: '',
          valid: [vType.length(1, 255, '备注长度为1~255个字符')]
        },
        {
          name: 'icon',
          value: '',
          type: 'imageUpload',
          label: '图标',
          placeholder: '上传分类图标',
          valid: [vType.require('请上传分类图标')]
        }
      ]
    }
  },
  methods: {
    show () {
      this.visible = true
    },
    okHandler (env) {
      env.preventDefault()
      if (this.$refs.categoryForm.validate()) {
        this.loading = true
        this.$refs.categoryForm.submit(data => {
          saveCategory(data).then(() => {
            this.loading = false
            this.$emit('onsave', 'insert')
            this.visible = false
          })
        })
      }
    }
  }
}
</script>