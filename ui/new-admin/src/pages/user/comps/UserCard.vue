<template>
  <div class="block">
    <div class="block-header">
      <h3 class="block-title">超级管理员</h3>
      <div class="block-options">
        <a class="btn-block-option">
          <TIcon icon="settings" pack="si" />
        </a>
        <a class="btn-block-option js-tooltip-enabled">
          <TIcon icon="trash" pack="si" @click="confirmDelete" />
        </a>
      </div>
    </div>
    <div
      class="block-content block-content-full text-center bg-image"
      style="background-image: url('http://wx1.sinaimg.cn/mw600/798cfa5bly1gbybogzyiuj20b406qt8u.jpg');"
    >
      <img
        class="img-avatar img-avatar96 img-avatar-thumb"
        src="http://ww1.sinaimg.cn/mw1024/00745YaMgy1gbyi3gl2mhj33402c0hdw.jpg"
        alt
      />
    </div>
    <div class="block-content font-size-sm" style="height: 120px">
      <p>
        <slot></slot>
      </p>
    </div>
  </div>
</template>
<script>
export default {
  methods: {
    confirmDelete() {
      this.$dialog.confirm()
    }
  }
}
</script>