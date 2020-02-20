<template>
  <!--
    ListItem
      title: string
      remark: string
      remarkType: primary, secondary(default), success, info, warning, danger, dark, light
  -->
  <li class="nav-item my-1">
    <span class="nav-link d-flex justify-content-between align-items-center t-list-item-tree">
      <span>
        <TIcon :icon="iconAngle" class="mr-1 t-list-item-tree" v-if="children" @click="showChildrenHandler"/>
        <img v-if="img && !children" class="rounded mr-1 t-list-item-img" :src="img">
        <TIcon v-if="icon && !img && !children" :icon="icon" :pack="iconPack" class="mr-1" /> {{title}}
      </span>
      <span>
        <span v-if="remark" :class="`badge badge-pill badge-${remarkType}`">{{remark}}</span>
      </span>
    </span>
    <TList v-if="children" v-show="showChildren" :data="children" class="ml-3"></TList>
  </li>
</template>
<script>
export default {
  props: {
    title: {
      type: String,
      default: ''
    },
    remark: {
      type: String,
      default: null
    },
    remarkType: {
      type: String,
      default: 'secondary'
    },
    icon: {
      type: String,
      default: null
    },
    iconPack: {
      type: String,
      default: 'fa'
    },
    img: {
      type: String,
      default: null
    },
    children: {
      type: Array,
      default() {
        return null
      }
    }
  },
  data() {
    return {
      showChildren: false,
      iconAngle: 'angle-right'
    }
  },
  methods: {
    showChildrenHandler(){
      this.showChildren = !this.showChildren
      this.iconAngle = 'angle-' + (this.showChildren ? 'down' : 'right')
    }
  }
}
</script>