<template>
  <t-block title="所有分类" theme border rounded>
    <t-nestable v-model="nestableItems" class="push t-category-list">
      <template v-slot:default="props">
        <b-row>
          <b-col :cols="2"><img class="img-avatar img-avatar-thumb" :src="props.item.icon" alt=""></b-col>
          <b-col class="py-3">
            <p class="mb-0">
              <a class="link-fx font-w700 d-inline-block" href="javascript:void(0)">{{ props.item.text }}</a>
            </p>
            <p class="mb-0">
              <a class="font-size-sm font-w700 text-uppercase text-muted mr-3" href="javascript:void(0)">{{ props.item.article_count }} 篇文章</a>
            </p>
          </b-col>
          <b-col class="py-3">
            <p class="mb-0">
              <a class="font-size-sm font-w700 text-uppercase text-muted mr-3" href="javascript:void(0)">5680 Ratings</a>
            </p>
          </b-col>
          <b-col class="py-3">
            <p class="font-size-h3 text-black mb-0">
              <i class="fa fa-angle-up text-success mr-1"></i> 86.850
            </p>
          </b-col>
        </b-row>
      </template>
    </t-nestable>
  </t-block>
</template>
<script>
export default {
  data () {
    return {
      title: '分类和标签管理',
      description: '添加分类和查看标签，可以通过拖拽移动分类排序',
      nestableItems: [
        {
          id: 0,
          text: 'Andy',
          article_count: 10,
          icon: 'http://wx1.sinaimg.cn/mw600/0080ZTByly1gd5ep80flvj30dw0dt3zk.jpg'
        }, {
          id: 1,
          text: 'Harry',
          article_count: 10,
          icon: 'http://wx1.sinaimg.cn/mw600/0080ZTByly1gd5ep80flvj30dw0dt3zk.jpg',
          children: [{
            id: 2,
            text: 'David',
            article_count: 10,
            icon: 'http://wx1.sinaimg.cn/mw600/0080ZTByly1gd5ep80flvj30dw0dt3zk.jpg'
          }]
        }, {
          id: 3,
          text: 'Lisa',
          article_count: 10,
          icon: 'http://wx1.sinaimg.cn/mw600/0080ZTByly1gd5ep80flvj30dw0dt3zk.jpg'
        }
      ]
    }
  }
}
</script>