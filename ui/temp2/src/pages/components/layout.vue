<template>
  <t-block title="Layout API" theme border rounded>

    <h2 class="content-heading">Sidebar Visibility</h2>
    <b-button variant="dark" class="mr-1" @click="sidebarVisibility(true)">Open Sidebar</b-button>
    <b-button variant="dark" @click="sidebarVisibility(false)">Close Sidebar</b-button>

    <h2 class="content-heading">Sidebar Position</h2>
    <b-button variant="dark" class="mr-1" @click="sidebarPosition('right')">Right Sidebar Position</b-button>
    <b-button variant="dark" @click="sidebarPosition('left')">Left Sidebar Position</b-button>

    <h2 class="content-heading">Sidebar Mini</h2>
    <b-button variant="dark" class="mr-1" @click="sidebarMin(true)">Enable Mini Mode</b-button>
    <b-button variant="dark" @click="sidebarMin(false)">Disable Mini Mode</b-button>

    <h2 class="content-heading">Sidebar Styles</h2>
    <b-button variant="dark" class="mr-1" @click="sidebarTheme('light')">Light Themed Sidebar</b-button>
    <b-button variant="dark" @click="sidebarTheme('dark')">Dark Themed Sidebar</b-button>

    <h2 class="content-heading">Side Overlay Visibility</h2>
    <b-button variant="dark" class="mr-1" @click="sideVisiable(true)">Open Side Overlay</b-button>
    <b-button variant="dark" @click="sideVisiable(false)">Close Side Overlay</b-button>

    <h2 class="content-heading">Side Overlay Hover Mode</h2>
    <b-button variant="dark" class="mr-1" @click="sideHoverMode(true)">Enable Hover Mode</b-button>
    <b-button variant="dark" @click="sideHoverMode(false)">Disable Hover Mode</b-button>

    <h2 class="content-heading">Header Mode</h2>
    <b-button variant="dark" class="mr-1" @click="headerMode('static')">Static Header</b-button>
    <b-button variant="dark" @click="headerMode('fixed')">Fixed Header</b-button>

    <h2 class="content-heading">Header Glass</h2>
    <b-button variant="dark" class="mr-1" @click="headerGlass(true)">Enable Header Glass</b-button>
    <b-button variant="dark" @click="headerGlass(false)">Disable Header Glass</b-button>

    <h2 class="content-heading">Header Styles</h2>
    <b-button variant="dark" class="mr-1" @click="headerTheme('light')">Light Themed Header</b-button>
    <b-button variant="dark" @click="headerTheme('dark')">Dark Themed Header</b-button>

    <h2 class="content-heading">Footer Mode</h2>
    <b-button variant="dark" class="mr-1" @click="footerMode('static')">Static Footer</b-button>
    <b-button variant="dark" @click="footerMode('fixed')">Fixed Footer</b-button>

    <h2 class="content-heading">Main Content</h2>
    <b-button variant="dark" class="push mr-1" @click="contentMode('boxed')">Boxed Content Layout</b-button>
    <b-button variant="dark" class="push mr-1" @click="contentMode('narrow')">Narrow Content Layout</b-button>
    <b-button variant="dark" class="push" @click="contentMode('full')">Full Width Content Layout</b-button>

  </t-block>
</template>
<script>
export default {
  methods: {
    sidebarVisibility (type) {
      this.$emit('sidebarVisiable', type)
    },
    sidebarPosition (pos) {
      this.$emit('sidebarPosition', pos)
    },
    sidebarMin (type) {
      this.$emit('sidebarMin', type)
    },
    sidebarTheme (theme) {
      this.$emit('sidebarTheme', theme)
    },
    sideVisiable (type) {
      this.$emit('sideVisiable', type)
    },
    sideHoverMode (type) {
      this.$emit('sideHoverMode', type)
    },
    headerMode (mode) {
      this.$emit('headerMode', mode)
    },
    headerGlass (type) {
      this.$emit('headerGlass', type)
    },
    headerTheme (theme) {
      this.$emit('headerTheme', theme)
    },
    footerMode (mode) {
      this.$emit('footerMode', mode)
    },
    contentMode (mode) {
      this.$emit('contentMode', mode)
    }
  }
}
</script>
