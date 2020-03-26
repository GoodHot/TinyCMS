<template>
  <b-modal v-model="visible" @shown="shownHandler" title="分类排序" ok-title="确定" cancel-title="取消" @ok="okHandler">
    <b-alert show><t-icon icon="hand-rock" pack="far" /> 点击分类拖拽排序</b-alert>
    <t-nestable :data="data" :loading="loading" @change="changeHandler">
      <template v-slot:default="props">
        <img :src="assetsURL(props.item.icon)" width="20" height="20" class="mr-2" /> {{ props.item.id }} {{ props.item.name }}
        <span class="float-right"><t-icon icon="hand-rock" pack="far" /></span>
      </template>
    </t-nestable>
  </b-modal>
</template>
<script>
// import {PropTypes} from '@/utils/types'
import { categoryTree, categorySort } from '@/api/category'

export default {
  data () {
    return {
      data: [],
      visible: false,
      loading: false
    }
  },
  methods: {
    show () {
      this.visible = true
    },
    loadTree () {
      this.loading = true
      categoryTree().then(res => {
        this.data = res.tree
        this.loading = false
      })
    },
    shownHandler () {
      this.loadTree()
    },
    okHandler (env) {
      env.preventDefault()
      console.log(this.data)
    },
    changeHandler (val, opt) {
      const pathTo = opt.pathTo
      if (!pathTo) {
        return
      }
      let sort = pathTo[pathTo.length -1]
      let tmpArr = null
      for (let i = 0; i < pathTo.length - 1; i++) {
        if (i == 0) {
          tmpArr = opt.items[pathTo[i]]
        } else {
          tmpArr = tmpArr.children[pathTo[i]]
        }
      }
      const param = {
        id: val.id,
        parent_id: tmpArr ? tmpArr.id : 0,
        sort: sort
      }
      console.log(param)
      this.loading = true
      categorySort(param).then(() => {
        this.loading = false
      })
    }
  }
}
</script>