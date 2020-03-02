<template>
  <div class="t-tree">
    <TIcon :icon="'angle-' + (showChild ? 'down' : 'right')" class="arrow" @click="showChild = !showChild" v-if="data.child && data.child.length > 0" />
    <b-form-checkbox :class="checkboxClass" v-model="data.checked" @change="checkHandler">
      {{data.name}}
    </b-form-checkbox>
    <div class="t-tree-child" v-show="showChild">
      <TTree :data="data.child" isChild :expand="expand"></TTree>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    data: {
      type: Object,
      default() {
        return {}
      }
    },
    expand: {
      type: Boolean,
      default: false
    }
  },
  data() {
    return {
      showChild: false
    }
  },
  mounted() {
    if (this.expand) {
      this.showChild = true
    }
  },
  computed: {
    checkboxClass() {
      let cls = 'node '
      if (!this.data.child || this.data.child.length == 0) {
        cls += 'node-align'
      }
      return cls
    }
  },
  methods: {
    checkHandler(checked) {
      this.$emit('select', {checked, data: this.data})
    }
  }
}
</script>
<style scoped>
.t-tree {
  display: inline-block;
  cursor: pointer;
  width: 100%;
}
.t-tree .arrow {
  float: left;
  line-height: 25px;
}
.t-tree .node {
  float: left;
  margin-left: 5px;
}
.t-tree .node-align {
  margin-left: 15px;
}
.t-tree-child {
  margin-left: 1rem;
}
.t-tree .fa-angle-right {
  margin-right: 2px;
}
</style>