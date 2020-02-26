<template>
  <TList rowKey="id" v-if="category && category.length > 0" :selected="selected" :data="category" @onclick="selectItem"></TList>
</template>
<script>
import { categoryTree } from "@/api/category";

export default {
  props: {
    defaultCategory: {
      type: Array,
      default() {
        return null
      }
    },
    selected: {
      type: [String, Number],
      default: null
    }
  },
  data() {
    return {
      category: [],
      selectRow: null
    }
  },
  beforeMount() {
    this.getCategoryTree()
  },
  methods: {
    reload() {
      this.category = []
      this.getCategoryTree()
    },
    selectItem(row) {
      if (this.selectRow) {
        this.selectRow.active = false
      }
      this.selectRow = row
      this.$emit('selectItem', row)
    },
    getCategoryTree() {
      if (this.defaultCategory) {
        for (let i in this.defaultCategory) {
          this.category.push(this.defaultCategory[i])
        }
      }
      categoryTree().then(res => {
        if (res.tree) {
          const rst = this.setCategoryTree(res.tree);
          for (let k in rst) {
            this.category.push(rst[k])
          }
        }
        this.$emit('inited', this.category)
      });
    },
    setCategoryTree(tree) {
      const result = [];
      for (let i in tree) {
        const data = tree[i];
        let remark = null;
        const cat = {
          id: data.id,
          title: data.name,
          icon: "book",
          remark: remark,
          img: data.icon,
          active: false
        };
        if (data.children && data.children.length > 0) {
          cat.children = this.setCategoryTree(data.children);
        } else {
          cat.remark = `${data.channel.article_count ? data.channel.article_count : 0}篇`
        }
        result.push(cat);
      }
      return result;
    }
  }
}
</script>