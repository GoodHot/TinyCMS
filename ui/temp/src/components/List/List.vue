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
    }
  },
  methods: {
    selectItem(row) {
      this.$emit('onclick', row)
    }
  }
}
</script>