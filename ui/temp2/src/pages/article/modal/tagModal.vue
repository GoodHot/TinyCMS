<template>
  <b-modal v-model="visible" title="标签" ok-title="确定" cancel-title="取消" size="lg" no-close-on-backdrop @ok="okHandler" :ok-disabled="loading">
    <b-overlay :show="loading" rounded="sm">
      <t-form :form="form" ref="tagForm">
      </t-form>
    </b-overlay>
  </b-modal>
</template>
<script>
import { vType } from '@/utils/valid'
import { saveTag, getTag } from '@/api/tag'

export default {
  data () {
    return {
      visible: false,
      loading: false,
      form: [
        {
          name: 'id',
          value: 0,
          type: 'hidden'
        },
        {
          name: 'name',
          value: '',
          type: 'input',
          label: '标签名称',
          placeholder: '这个标签叫什么名字？',
          description: '标签名唯一，不可重复',
          valid: [vType.require('请输入标签名称'), vType.length(1, 255, '标签名称长度为1~255个字符')]
        },
        {
          name: 'path',
          value: '',
          type: 'input',
          label: '路径',
          placeholder: '好的路径有利于SEO哦，路径不可重复',
          description: '路径名称唯一，不可重复',
          valid: [vType.require('请输入标签路径'), vType.regex('^[a-zA-Z0-9\\-]+$', '路径只能为字母和数字或横杠(-)'), vType.length(1, 255, '路径长度为1~255个字符')]
        },
        {
          name: 'description',
          value: '',
          type: 'textarea',
          label: '标签描述',
          placeholder: '',
          valid: [vType.length(1, 255, '描述长度为1~255个字符')]
        },
        {
          name: 'meta_title',
          value: '',
          type: 'input',
          label: 'Meta Title',
          placeholder: 'SEO标题',
          description: '建议控制在70个字符以内，最多可输入255个字符',
          valid: [vType.length(1, 255, 'Meta Title 长度为1~255个字符')]
        },
        {
          name: 'meta_description',
          value: '',
          type: 'textarea',
          label: 'Meta Description',
          placeholder: 'SEO描述',
          description: '建议控制在150个字符以内，最多可输入255个字符',
          valid: [vType.length(1, 255, '描述长度为1~255个字符')]
        },
        {
          name: 'icon',
          value: '',
          type: 'imageUpload',
          label: '图标',
          placeholder: '上传标签图标',
          valid: []
        }
      ]
    }
  },
  methods: {
    show () {
      this.cleanForm()
      this.visible = true
    },
    cleanForm () {
      const temp = []
      this.form.map(item => {
        item.value = null
        temp.push(item)
      })
      this.form = temp
    },
    edit (id) {
      this.visible = true
      this.loading = true
      getTag(id).then(res => {
        const tag = res.tag
        const temp = []
        this.form.map(item => {
          item.value = tag[item.name]
          temp.push(item)
        })
        this.form = temp
        this.loading = false
      })
    },
    okHandler (env) {
      env.preventDefault()
      if (this.$refs.tagForm.validate()) {
        this.loading = true
        this.$refs.tagForm.submit(data => {
          saveTag(data).then(() => {
            this.loading = false
            this.$emit('onsave')
            this.visible = false
          }).catch(() => {
            this.loading = false
          })
        })
      }
    }
  }
}
</script>