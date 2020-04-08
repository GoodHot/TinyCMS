<template>
  <t-input-drawer :clearVariant="clearVariant" ref="categoryInput" placeholder="选择分类" :dataValue="choose.name" :size="size" :inputClass="inputClass" @onshow="onshowHandler" :readonly="true" @clear="clearHandler">
    <b-overlay :show="loading" rounded="sm">
      <div style="max-height: 200px; overflow-y: scroll">
        <ul class="t-category-drawer">
          <li v-for="cate in categorys" :style="calcPaddingLeft(cate.level)" :key="cate.id" @click="chooseItem(cate)">
            <img :src="assetsURL(cate.icon)" width="20" height="20" class="mr-2" />
            {{ cate.name }}
          </li>
        </ul>
      </div>
    </b-overlay>
  </t-input-drawer>
</template>
<script>
import { PropTypes } from '@/utils/types'
import { categoryTree } from '@/api/category'

export default {
  model: {
    prop: 'value',
    event: 'value.change'
  },
  data () {
    return {
      loading: false,
      hasData: false,
      categorys: [],
      choose: {
        name: ''
      }
    }
  },
  props: {
    size: PropTypes.String,
    inputClass: PropTypes.String,
    clearVariant: PropTypes.String.def('light'),
    value: PropTypes.Number
  },
  mounted () {
    if (this.value !== 0) {
      this.loadCategoryTree()
    }
  },
  methods: {
    chooseItem (item) {
      this.choose.name = item.name
      this.$refs.categoryInput.hideDrawer()
      this.$emit('change', item.id)
    },
    calcPaddingLeft (level) {
      if (level === 0) {
        return ''
      }
      return `padding-left: ${1 + level}rem`
    },
    onshowHandler () {
      if (!this.hasData) {
        this.loadCategoryTree()
      }
    },
    loadCategoryTree () {
      if (this.categorys && this.categorys.length > 0) {
        return
      }
      this.loading = true
      categoryTree().then(res => {
        this.loading = false
        this.hasData = true
        this.categorys = this.setData(res.tree, 0)
        this.setValue()
      })
    },
    setValue () {
      this.categorys.map(c => {
        if (c.id === this.value) {
          this.choose.name = c.name
        }
      })
    },
    setData (tree, level) {
      const data = []
      for (let i in tree) {
        const tmp = tree[i]
        data.push({
          id: tmp.id,
          name: tmp.name,
          icon: tmp.category.icon,
          level: level
        })
        if (tmp.children && tmp.children.length > 0) {
          data.push(...this.setData(tmp.children, level + 1))
        }
      }
      return data
    },
    clearHandler () {
      this.choose.name = ''
      this.$emit('change', 0)
    }
  },
  watch: {
    value () {
      this.loadCategoryTree()
    }
  }
}
</script>