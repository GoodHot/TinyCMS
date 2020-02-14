<template>
  <i :class="classes"></i>
</template>
<script>
export default {
  props: {
    pack: {
      type: String,
      default: "fa"
    },
    icon: {
      type: String
    },
    size: {
      type: String,
      default: ''
    }
  },
  computed: {
    classes() {
      let icon = ''
      switch(this.pack){
        case 'fa':
        case 'fab':
        case 'far':
          icon = `fa-${this.icon}`
          break;
        case 'si':
          icon = `si-${this.icon}`
          break;
      }
      return `${this.pack} ${icon} ${this.size}`
    }
  }
}
</script>