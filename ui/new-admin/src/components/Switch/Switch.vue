<template>
  <!--
    Switch
      defaultChecked: true/false
      inline: true/false
      size: sm, lg
      type: primary, secondary, success, info, warning, danger, dark, light
      theme: alt, outline
  -->
  <div :class="switchClass()">
    <input
      type="checkbox"
      class="custom-control-input"
      :name="name"
      :id="`switch-id-${name}`"
    />
    <label class="custom-control-label" :for="`switch-id-${name}`">{{ label }}</label>
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
    }
  },
  methods: {
    switchClass() {
      let cls = `custom-control custom-switch custom-control-${this.type} `
      if (this.size) {
        cls += `custom-control-${this.size} `
      }
      return cls
    }
  }
}
</script>