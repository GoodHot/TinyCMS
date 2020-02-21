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
      :key="item[rowkey]"
      :img="item.img"
      :children="item.children"
      :icon="item.icon"
      :title="item.title"
      :remark="item.remark"
      :remarkType="item.remarkType"
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
    rowkey: {
      type: String
    }
  }
};
</script>