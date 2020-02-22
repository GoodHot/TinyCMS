<template>
  <!--
    ListItem
      title: string
      remark: string
      remarkType: primary, secondary(default), success, info, warning, danger, dark, light
  -->
  <li class="nav-item my-1">
    <span :class="'nav-link d-flex justify-content-between align-items-center t-list-item-tree ' + (active ? 'active' : '')">
      <span>
        <TIcon :icon="iconAngle" class="mr-1 t-list-item-tree" v-if="children" @click="showChildrenHandler"/>
        <img v-if="img && !children" class="rounded mr-1 t-list-item-img" :src="img">
        <TIcon v-if="icon && !img && !children" :icon="icon" :pack="iconPack" class="mr-1" /> <span @click="clickTitle">{{title}}</span>
      </span>
      <span>
        <span v-if="remark" :class="`badge badge-pill badge-${remarkType}`">{{remark}}</span>
      </span>
    </span>
    <TList v-if="children" :rowKey="rowKey" v-show="showChildren" :data="children" class="ml-3" @onclick="selectItem"></TList>
  </li>
</template>
<script>
export default {
  model: {
    prop: 'active',
    event: 'change.active'
  },
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
    },
    active: {
      type: Boolean,
      default: true
    },
    row: {
      type: [String, Number, Object]
    },
    rowKey: {
      type: String
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
    },
    clickTitle() {
      this.$emit('change.active', true)
      this.$emit('onclick', this.row)
    },
    selectItem(row) {
      this.$emit('onclick', row)
    }
  }
}
</script>