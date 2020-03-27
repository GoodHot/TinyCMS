<template>
  <b-modal v-model="visible" title="分类" ok-title="确定" cancel-title="取消" size="lg" no-close-on-backdrop @ok="okHandler">
    <b-form>
      <b-form-group label="分类名称:">
        <b-form-input placeholder="这个分类叫什么名字？" :state="true"></b-form-input>
      </b-form-group>
      <b-form-group
        label="路径:"
        description="分类路径为URL展示路径，不可重复，只允许输入字母和数字，尽量使用简单易懂的单词或拼音"
      >
        <b-form-input placeholder="好的路径有利于SEO哦"></b-form-input>
      </b-form-group>
      <b-form-group label="SEO标题:">
        <b-form-input placeholder="针对标题的SEO优化, SEO标题可以和分类名称组合使用"></b-form-input>
      </b-form-group>
      <b-form-group label="关键字:">
        <b-form-input placeholder="合理的关键字对SEO帮助很大"></b-form-input>
      </b-form-group>
      <b-form-group label="描述:">
        <b-form-textarea rows="3" max-rows="6"  placeholder="SEO中的Description"></b-form-textarea>
      </b-form-group>
      <b-form-group label="备注:">
        <b-form-textarea rows="3" max-rows="6"></b-form-textarea>
      </b-form-group>
    </b-form>
  </b-modal>
</template>
<script>
export default {
  data () {
    return {
      visible: false,
    }
  },
  methods: {
    show () {
      this.visible = true
    },
    okHandler (env) {
      env.preventDefault()
    }
  }
}
</script>