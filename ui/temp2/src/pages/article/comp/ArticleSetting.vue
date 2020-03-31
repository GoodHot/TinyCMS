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
    <collapse :show="active === 'visiable'" title="可见性设置" @onshow="showHandler('visiable')"></collapse>
    <collapse :show="active === 'seo'" title="SEO设置" @onshow="showHandler('seo')"></collapse>
    <collapse :show="active === 'tags'" title="分类和标签" @onshow="showHandler('tags')"></collapse>
    <collapse :show="active === 'cover'" title="封面设置" @onshow="showHandler('cover')" style="border-bottom: 0"></collapse>
  </div>
</template>

<script>
import collapse from '@/commons/article/collapse'
export default {
  components: {
    collapse
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