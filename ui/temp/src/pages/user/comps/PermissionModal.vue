<template>
  <div>
    <b-alert variant="info" show v-if="role.is_super">超级管理员拥有所有权限</b-alert>
    <div v-for="(val, key) of role.permissions" :key="key">
      <b-badge variant="primary">{{key}}</b-badge>
      <b-badge class="ml-2" variant="info" v-for="per of val" :key="per.id">{{per.permission_name}}</b-badge>
    </div>
  </div>
</template>
<script>
export default {
  props: {
    role: {
      type: Object,
      default() {
        return {}
      }
    }
  }
}
</script>