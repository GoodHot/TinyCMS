<template>
  <div class="content p-0">
    <collapse :show="active === 'write'" title="写作设置" @onshow="showHandler('write')">
      <div class="push">
        <b-form-checkbox>自动保存为草稿（30秒一次）</b-form-checkbox>
        <b-form-checkbox>自动获取封面图</b-form-checkbox>
        <b-form-checkbox>自动保存网络图片</b-form-checkbox>
        <b-form-checkbox>自动获取内容简介</b-form-checkbox>
      </div>
    </collapse>
    <collapse :show="active === 'visiable'" title="发布设置" @onshow="showHandler('visiable')">
      <b-form-group
        label-cols="4"
        label="可见性:"
      >
        <b-form-select>
          <b-form-select-option>所有人可见</b-form-select-option>
          <b-form-select-option>仅自己可见</b-form-select-option>
        </b-form-select>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="作者:"
      >
        <user-input clearVariant="dark"></user-input>
      </b-form-group>
      <!-- <b-form-group
        label-cols="4"
        label="定时发布:"
      >
        <b-form-checkbox-group
          switches
        ></b-form-checkbox-group>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="发布时间:"
      >
        <b-form-datepicker calendar-width="425px" :right="true" offset="210"></b-form-datepicker>
      </b-form-group> -->
    </collapse>
    <collapse :show="active === 'seo'" title="SEO设置" @onshow="showHandler('seo')">
      <b-form-group
        label-cols="4"
        label="SEO地址:"
        description="系统会根据标题自动生成SEO地址，也可手动设置"
      >
        <b-form-input placeholder="SEO地址"></b-form-input>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="SEO描述:"
        description="如不单独填写，系统会自动截取文章第一段内容作为描述"
      >
        <b-form-textarea placeholder="SEO描述"></b-form-textarea>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="SEO关键字:"
        description="如不单独填写，如不单独填写，系统会使用标签作为关键字"
      >
        <b-form-textarea placeholder="SEO关键字"></b-form-textarea>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="内容描述:"
        description="此内容描述可用作列表文章简介"
      >
        <b-form-textarea placeholder="内容描述"></b-form-textarea>
      </b-form-group>
    </collapse>
    <collapse :show="active === 'tags'" title="分类和标签" @onshow="showHandler('tags')">
      <b-form-group
        label-cols="4"
        label="分类:"
      >
        <category-input clearVariant="dark"></category-input>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="标签:"
      >
        <b-form-tags></b-form-tags>
      </b-form-group>
    </collapse>
    <collapse :show="active === 'cover'" title="封面设置" @onshow="showHandler('cover')" style="border-bottom: 0">
      <t-image-upload class="upload-img" ></t-image-upload>
    </collapse>
  </div>
</template>

<script>
import collapse from '@/commons/article/collapse'
import userInput from '@/commons/user/user-input'
import categoryInput from '@/commons/category/category-input'

export default {
  components: {
    collapse,
    userInput,
    categoryInput
  },
  data() {
    return {
      active: 'write'
    }
  },
  methods: {
    showHandler (type) {
      this.active = type
    }
  }
}
</script>
<style lang="scss" scoped>
.upload-img {
  width: 100%;
  height: 250px;
  /deep/ .fa-plus {
    line-height: 225px;
  }
}
</style>