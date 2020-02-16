<template>
  <!--
    checkbox
      checked: true/false
      inline: true/false
      size: sm, lg
      type: primary, secondary, success, info, warning, danger, dark, light
  -->
  <div :class="classes">
    <input
      type="checkbox"
      class="custom-control-input"
      :name="name"
      :id="`checkbox-id-${name}`"
      :checked="selected"
      @change="changeHandler"
    />
    <label class="custom-control-label" :for="`checkbox-id-${name}`">{{label}}</label>
  </div>
</template>
<script>
export default {
  props: {
    label: {
      type: String,
      default: ''
    },
    checked: {
      type: Boolean,
      default: false
    },
    inline: {
      type: Boolean,
      default: false
    },
    size: {
      type: String,
      default: null
    },
    type: {
      type: String,
      default: 'primary'
    },
    name: {
      type: String
    },
    shape: {
      type: String,
      default: 'circle'
    }
  },
  data() {
    return {
      selected: this.checked
    }
  },
  computed: {
    classes() {
      let cls = `custom-control custom-checkbox custom-control-${this.type} `
      if (this.size) {
        cls += `custom-control-${this.size} `
      }
      if (this.inline) {
        cls += 'custom-control-inline '
      }
      if (this.shape !== 'circle') {
        cls += 'custom-checkbox-square '
      }
      return cls
    }
  },
  methods: {
    changeHandler(e) {
      this.selected = !this.selected
      this.$emit("change", {
        checked: this.selected,
        event: e
      })
    },
    check(val) {
      this.selected = val
    },
    state() {
      return this.selected
    }
  }
}
</script>