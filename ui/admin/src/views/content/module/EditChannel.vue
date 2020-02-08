<template>
  <a-modal :title="title" v-model="visible" @ok="save">
    <a-form :form="form">
      <a-form-item
        label="频道名称"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-input v-decorator="['name', {rules:[{required: true, message: '请输入名称'}]}]" />
      </a-form-item>
      <a-form-item
        label="英文名"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-input v-decorator="['title', {rules:[{required: true, message: '请输入频道英文名'}]}]" />
      </a-form-item>
      <a-form-item
        label="SEO标题"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-input v-decorator="['seo_title', {rules:[{required: false, message: '请输入SEO标题'}]}]" />
      </a-form-item>
      <a-form-item
        label="SEO关键字"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-input v-decorator="['keyword', {rules:[{required: false, message: '请输入SEO关键字'}]}]" />
      </a-form-item>
      <a-form-item
        label="SEO描述"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-textarea
          :autosize="{ minRows: 2, maxRows: 6 }"
          v-decorator="['desccription', {rules:[{required: false, message: '请输入SEO描述'}]}]"
        />
      </a-form-item>
      <a-form-item
        label="备注"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-textarea
          :autosize="{ minRows: 2, maxRows: 6 }"
          v-decorator="['remark', {rules:[{required: false, message: '请输入备注'}]}]"
        />
      </a-form-item>
      <a-form-item
        label="图标"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-textarea
          :autosize="{ minRows: 2, maxRows: 6 }"
          v-decorator="['icon', {rules:[{required: false, message: '请输入svg图标'}]}]"
        />
      </a-form-item>
    </a-form>
  </a-modal>
</template>
<script>
import { saveChannel, getChannel } from '@/api/channel'

export default {
  data () {
    return {
      title: '',
      visible: false,
      labelCol: {
        xs: { span: 24 },
        sm: { span: 7 }
      },
      wrapperCol: {
        xs: { span: 24 },
        sm: { span: 13 }
      },
      form: this.$form.createForm(this),
      channelParentId: '',
      reqData: {
        id: 0,
        parentId: 0
      }
    }
  },
  methods: {
    loadAdd (parentId) {
      this.formReset()
      this.visible = true
      this.title = '新建频道'
      this.reqData.id = 0
      this.reqData.parentId = parentId
    },
    loadEdit (id) {
      this.formReset()
      this.visible = true
      this.title = '编辑频道'
      const { form: { setFieldsValue } } = this
      getChannel(id).then(res => {
        const channel = res.channel
        setFieldsValue(channel)
        this.reqData.id = channel.id
      })
    },
    save () {
      const { form: { validateFields } } = this
      validateFields((errors, values) => {
        if (!errors) {
          if (this.reqData.id !== 0) {
            values.id = this.reqData.id
          }
          if (this.reqData.parentId !== 0) {
            values.parent_id = this.reqData.parentId
          }
          saveChannel(values).then(res => {
            this.$message.success(`操作成功`)
            this.visible = false
            this.$emit('commitHandler')
            this.formReset()
          })
        }
      })
    },
    formReset () {
      const { form: { resetFields } } = this
      resetFields()
    }
  }
}
</script>
