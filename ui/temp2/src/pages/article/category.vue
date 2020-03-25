<template>
  <t-block title="所有分类" theme border rounded>
    <t-nestable v-model="nestableItems" class="push t-category-list">
      <template v-slot:default="props">
        <b-row>
          <b-col :cols="2"><img class="img-avatar img-avatar-thumb" :src="assetsURL(props.item.channel.icon)" alt=""></b-col>
          <b-col class="py-3">
            <p class="mb-0">
              <a class="link-fx font-w700 d-inline-block" href="javascript:void(0)">{{ props.item.name }}</a>
            </p>
            <p class="mb-0">
              <a class="font-size-sm font-w700 text-uppercase text-muted mr-3" href="javascript:void(0)">{{ props.item.channel.article_count }} 篇文章</a>
            </p>
          </b-col>
          <b-col class="py-3">
            <p class="mb-0">
              <a class="font-size-sm font-w700 text-uppercase text-muted mr-3" href="javascript:void(0)">{{ props.item.channel.path }}</a>
            </p>
          </b-col>
          <b-col class="py-3">
            <p class="mb-0">
              <a class="font-size-sm font-w700 text-uppercase text-muted mr-3" href="javascript:void(0)">{{ props.item.channel.keyword }}</a>
            </p>
          </b-col>
          <b-col class="py-3">
            <p class="mb-0">
              <a class="font-size-sm font-w700 text-uppercase text-muted mr-3" href="javascript:void(0)">{{ props.item.channel.remark }}</a>
            </p>
          </b-col>
        </b-row>
      </template>
    </t-nestable>
  </t-block>
</template>
<script>
import { categoryTree } from '@/api/category'

export default {
  data () {
    return {
      title: '分类和标签管理',
      description: '添加分类和查看标签，可以通过拖拽移动分类排序',
      nestableItems: []
    }
  },
  created () {
    this.loadCategoryTree()
  },
  methods: {
    loadCategoryTree () {
      categoryTree().then(res => {
        this.nestableItems = res.tree
      })
    }
  }
}
</script>