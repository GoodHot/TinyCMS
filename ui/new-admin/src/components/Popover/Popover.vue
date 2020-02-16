<template>
  <!--
    popover:
      title: string
      content: string
      type: text(default), html
      placement: top(default), right, left, bottom
      trigger: click, hover, focus
  -->
  <span ref="popover">
    <slot name="html"></slot>
    <slot></slot>
  </span>
</template>
<script>
export default {
  props: {
    title: {
      type: String,
      default: ''
    },
    content: {
      type: String,
      default: ''
    },
    placement: {
      type: String,
      default: 'top'
    },
    trigger: {
      type: String,
      default: 'click'
    }
  },
  mounted() {
    this.showPopover()
  },
  methods: {
    showPopover() {
      console.log(123)
      const opt = {
        animation: true,
        title: this.title,
        content: this.content,
        placement: this.placement,
        trigger: this.trigger
      }
      if (this.$slots.html) {
        const html = this.$slots.html[0]
        if (html.elm) {
          opt.content = html.elm
          opt.template = '<div class="popover" role="tooltip"><div class="arrow"></div><div class="popover-body"></div></div>'
          opt.html = true
        }
      }
      this.$(this.$refs.popover).popover(opt)
    }
  }
}
</script>