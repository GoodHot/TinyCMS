<template>
  <!--
    list
      data: array
      slot
  -->
  <ul class="nav nav-pills flex-column font-size-sm">
    <slot v-if="!data"></slot>
    <TListItem
      v-for="item of data"
      :key="item[rowKey]"
      :img="item.img"
      :children="item.children"
      :icon="item.icon"
      :title="item.title"
      :remark="item.remark"
      :remarkType="item.remarkType"
      :row="item"
      :rowKey="rowKey"
      :selected="selected"
      v-model="item.active"
      @onclick="selectItem"
    ></TListItem>
  </ul>
</template>
<script>
export default {
  props: {
    data: {
      type: Array,
      default() {
        return null;
      }
    },
    rowKey: {
      type: String
    },
    selected: {
      type: [String, Number],
      default: null
    }
  },
  mounted() {
    this.checkActive()
  },
  methods: {
    selectItem(row) {
      this.$emit('onclick', row)
    },
    checkActive() {
      if (!this.selected) {
        return
      }
      for (let i=0;i<this.data.length;i++) {
        const item = this.data[i]
        item.active = item[this.rowKey] == this.selected
        if (item.active) {
          this.$emit('onclick', item)
        }
      }
    }
  },
  watch: {
    selected() {
      this.checkActive()
    }
  }
}
</script>