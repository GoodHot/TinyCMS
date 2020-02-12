<template>
  <a-card :body-style="{padding: '24px 32px'}" :bordered="false">
    <a-form :form="form">
      <a-form-item
        :label="item.name"
        :labelCol="{lg: {span: 7}, sm: {span: 7}}"
        :wrapperCol="{lg: {span: 10}, sm: {span: 17} }"
        :help="`${item.description}, key: ${item.key}`"
        v-for="item in dicts"
        :key="item.id"
      >
        <a-input
          v-if="item.type == 'text'"
          :placeholder="`请输入${item.description}`"
          v-model="item.value"
        />
        <a-textarea
          v-else-if="item.type == 'textarea'"
          rows="4"
          v-model="item.value"
          :placeholder="`请输入${item.description}`"
        />
        <a-select v-model="item.value" v-if="item.type == 'select'" >
          <a-select-option v-for="opt in splitOpts(item.options)" :key="opt" :value="opt">{{ opt }}</a-select-option>
        </a-select>
      </a-form-item>
      <a-form-item
        :wrapperCol="{ span: 24 }"
        style="text-align: center"
      >
        <a-button @click="handleSubmit" type="primary">提交</a-button>
      </a-form-item>
    </a-form>
  </a-card>
</template>
<script>
import { getAllDict, editDict } from '@/api/dict'

export default {
  data () {
    return {
      form: this.$form.createForm(this),
      dicts: []
    }
  },
  mounted () {
    this.loadDict()
  },
  methods: {
    loadDict () {
      getAllDict().then(res => {
        this.dicts = res.dicts
      })
    },
    handleSubmit () {
      editDict({
        dicts: this.dicts
      }).then(res => {
        this.$message.success(`操作成功`)
        this.loadDict()
      })
    },
    splitOpts (options) {
      return options.split(',')
    }
  }
}
</script>
