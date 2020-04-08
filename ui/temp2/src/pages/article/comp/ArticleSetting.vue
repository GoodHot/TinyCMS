<template>
  <div class="content p-0">
    <collapse :show="active === 'write'" title="写作设置" @onshow="showHandler('write')">
      <div class="push">
        <b-form-checkbox v-model="writeSetting.autoSave" @change="saveWriteSetting($event, 'autoSave')">自动保存内容为草稿（30秒一次）</b-form-checkbox>
        <b-form-checkbox v-model="writeSetting.getCover" @change="saveWriteSetting($event, 'getCover')">自动获取封面图</b-form-checkbox>
        <b-form-checkbox v-model="writeSetting.getNetImage" @change="saveWriteSetting($event, 'getNetImage')">自动保存网络图片</b-form-checkbox>
        <b-form-checkbox v-model="writeSetting.getDescription" @change="saveWriteSetting($event, 'getDescription')">自动获取内容简介</b-form-checkbox>
      </div>
    </collapse>
    <collapse :show="active === 'publish'" title="发布设置" @onshow="showHandler('publish')">
      <b-form-group
        label-cols="4"
        label="可见性:"
      >
        <b-form-select v-model="publishSetting.visibility">
          <b-form-select-option :value="1">所有人可见</b-form-select-option>
          <b-form-select-option :value="2">仅自己可见</b-form-select-option>
        </b-form-select>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="页面模板"
      >
        <b-form-select v-model="publishSetting.template">
          <b-form-select-option value="post.html">post.html</b-form-select-option>
          <b-form-select-option value="post_line.html">post_line.html</b-form-select-option>
        </b-form-select>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="作者:"
      >
        <user-input clearVariant="dark" @change="userChangeHandler"></user-input>
      </b-form-group>
    </collapse>
    <collapse :show="active === 'seo'" title="SEO设置" @onshow="showHandler('seo')">
      <b-form-group
        label-cols="4"
        label="文章地址:"
        description="系统会根据标题自动生成SEO地址，也可手动设置"
      >
        <b-form-input v-model="seoSetting.seo_title" placeholder="SEO地址"></b-form-input>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="SEO描述:"
        description="如不单独填写，系统会自动截取文章第一段内容作为描述"
      >
        <b-form-textarea v-model="seoSetting.seo_description" placeholder="SEO描述"></b-form-textarea>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="SEO关键字:"
        description="如不单独填写，如不单独填写，系统会使用标签作为关键字"
      >
        <b-form-textarea v-model="seoSetting.seo_keyword" placeholder="SEO关键字"></b-form-textarea>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="内容描述:"
        description="此内容描述可用作列表文章简介"
      >
        <b-form-textarea v-model="seoSetting.description" placeholder="内容描述"></b-form-textarea>
      </b-form-group>
    </collapse>
    <collapse :show="active === 'tags'" title="分类和标签" @onshow="showHandler('tags')">
      <b-form-group
        label-cols="4"
        label="分类:"
      >
        <category-input clearVariant="dark" @change="categoryChangeHandler"></category-input>
      </b-form-group>
      <b-form-group
        label-cols="4"
        label="标签:"
      >
        <b-form-tags v-model="meta.tags"></b-form-tags>
      </b-form-group>
    </collapse>
    <collapse :show="active === 'cover'" title="封面设置" @onshow="showHandler('cover')" style="border-bottom: 0">
      <t-image-upload class="upload-img" v-model="cover" ></t-image-upload>
    </collapse>
  </div>
</template>

<script>
import collapse from '@/commons/article/collapse'
import userInput from '@/commons/user/user-input'
import categoryInput from '@/commons/category/category-input'
import Vue from 'vue'

const ls_key_write_setting = 'w_setting'

export default {
  components: {
    collapse,
    userInput,
    categoryInput
  },
  data() {
    return {
      active: 'write',
      writeSetting: {
        autoSave: true,
        getCover: true,
        getNetImage: true,
        getDescription: true
      },
      publishSetting: {
        visibility: 1,
        template: "post.html",
        authorId: 0
      },
      seoSetting: {
        seo_title: '',
        seo_description: '',
        seo_keyword: '',
        description: ''
      },
      meta: {
        category: 0,
        tags: []
      },
      cover: '',
      autoSaveTimer: null
    }
  },
  mounted () {
    this.loadWriteSetting()
  },
  methods: {
    showHandler (type) {
      this.active = type
    },
    userChangeHandler (id) {
      this.publishSetting.authorId = id
    },
    categoryChangeHandler (id) {
      this.meta.category = id
    },
    loadWriteSetting () {
      const st = Vue.ls.get(ls_key_write_setting)
      this.writeSetting.autoSave = st.autoSave
      this.writeSetting.getCover = st.getCover
      this.writeSetting.getNetImage = st.getNetImage
      this.writeSetting.getDescription = st.getDescription
      this.enableAutoSave(st.autoSave)
    },
    saveWriteSetting (e, key) {
      this.writeSetting[key] = e
      Vue.ls.set(ls_key_write_setting, this.writeSetting)
      if (key === 'autoSave') {
        this.enableAutoSave(e)
      }
    },
    enableAutoSave (status) {
      if (status) {
        const _this = this
        this.autoSaveTimer = setInterval(() => {
          _this.$emit('autosave')
        }, 3 * 1000)
      } else {
        clearInterval(this.autoSaveTimer)
      }
    },
    getValue () {
      return {
        writeSetting: this.writeSetting,
        publishSetting: this.publishSetting,
        seoSetting: this.seoSetting,
        meta: this.meta,
        cover: this.cover
      }
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