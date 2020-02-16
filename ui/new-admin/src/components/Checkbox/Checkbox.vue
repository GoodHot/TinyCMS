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
      v-bind:checked="checked"
      @change="changeHandler"
    />
    <label class="custom-control-label" :for="`checkbox-id-${name}`">{{label}}</label>
  </div>
</template>
<script>
export default {
  model: {
    prop: 'checked',
    event: 'updateChecked'
  },
  props: {
    label: {
      type: String,
      default: ''
    },
    checked: {
      type: [Boolean, Object],
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
    },
    val: {
      type: [String, Number, Boolean, Array, Object, Function],
      default: null
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
      this.$emit('updateChecked', e.target.checked)
      this.$emit('change', {checked: e.target.checked, value: this.val})
    }
  }
}
</script>