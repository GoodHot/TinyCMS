<template>
  <div class="t-tree">
    <TIcon :icon="'angle-' + (showChild ? 'down' : 'right')" class="arrow" @click="showChild = !showChild" v-if="data.child && data.child.length > 0" />
    <b-form-checkbox :class="checkboxClass" v-model="data.checked" @change="checkHandler">
      {{data.name}}
    </b-form-checkbox>
    <div class="t-tree-child" v-show="showChild">
      <TTreeChild ref="childs" v-for="item of data.child" :key="item.name" :data="item" :expand="expand" @select="selectHandler"></TTreeChild>
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
      if (this.data.child) {
        this.selectAll(checked)
      }
    },
    selectHandler(data) {
      this.$emit('select', data)
    },
    selectAll(status) {
      this.data.checked = status
      console.log(this.data)
      // if (this.data.child) {
      //   for (let i in this.data.child) {
      //     const c = this.data.child[i]
      //     console.log('c', c)
      //     c.checked = status
      //   }
      // }
      if (this.data.child) {
        for (let i in this.$refs.childs) {
          const child = this.$refs.childs[i]
          child.selectAll(status)
        }
      }
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