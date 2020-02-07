<template>
  <a-modal :title="title" v-model="visible" @ok="save">
    <a-form :form="form">
      <a-form-item
        label="上级频道"
        :labelCol="labelCol"
        :wrapperCol="wrapperCol"
      >
        <a-tree-select
          showSearch
          :value="channelParentId"
          :dropdownStyle="{ maxHeight: '200px', overflow: 'auto' }"
          placeholder="Please select"
          allowClear
          treeDefaultExpandAll
          @change="channelChange"
        >
          <a-tree-select-node value="parent 1" title="parent 1" key="0-1">
            <a-tree-select-node value="parent 1-0" title="parent 1-0" key="0-1-1">
              <a-tree-select-node :selectable="false" value="leaf1" title="my leaf" key="random" />
              <a-tree-select-node value="leaf2" title="your leaf" key="random1" />
            </a-tree-select-node>
            <a-tree-select-node value="parent 1-1" title="parent 1-1" key="random2">
              <a-tree-select-node value="sss" key="random3">
                <b style="color: #08c" slot="title">sss</b>
              </a-tree-select-node>
            </a-tree-select-node>
          </a-tree-select-node>
        </a-tree-select>
      </a-form-item>
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
    </a-form>
  </a-modal>
</template>
<script>
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
      channelParentId: ''
    }
  },
  methods: {
    loadAdd () {
      this.formReset()
      this.visible = true
      this.title = '新建频道'
    },
    save () {
      const { form: { validateFields } } = this
      validateFields((errors, values) => {
        if (!errors) {
        }
      })
    },
    channelChange () {
    },
    formReset () {
      const { form: { resetFields } } = this
      resetFields()
    }
  }
}
</script>
