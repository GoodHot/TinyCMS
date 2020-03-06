<template>
  <div class="block text-center">
    <div :class="'block-content block-content-full ribbon ribbon-left ribbon-modern ' + (admin.role.is_super ? 'ribbon-primary':'ribbon-info')">
        <div class="ribbon-box">
            <i class="fab fa-black-tie"></i> {{admin.role.role_name}}
        </div>
        <div class="px-2">
          <img
            class="img-avatar img-avatar96 img-avatar-thumb"
            :src="admin.avatar ? assetsURL(admin.avatar): defaultAvatar"
            alt
          />
        </div>
        <div class="block-content-full pt-2">
            <div class="font-size-h4 font-w600 mb-0">{{ admin.nickname }}({{ admin.username }})</div>
            <div class="font-size-h5 text-muted">{{admin.resume}}</div>
        </div>
    </div>
    <div class="block-content block-content-full bg-body-light">
      <b-button variant="outline-secondary" size="sm" v-b-tooltip.hover title="查看拥有权限" @click="permissionVisible = true"><TIcon icon="address-card" pack="far" /></b-button>
      <b-button variant="outline-secondary" size="sm" class="mx-1" v-b-tooltip.hover title="设置账号信息" @click="$emit('onedit', admin.id)"><TIcon icon="settings" pack="si" /></b-button>
      <b-button variant="outline-secondary" size="sm" v-b-tooltip.hover title="删除用户" @click="deleteUser(admin.id)"><TIcon icon="trash" pack="si" /></b-button>
    </div>
    <b-modal
      v-model="permissionVisible"
      title="查看权限"
      hide-footer
    >
      <PermissionModal :role="admin.role"></PermissionModal>
    </b-modal>
  </div>
</template>
<script>
import PermissionModal from './PermissionModal'
import defaultAvatar from '@/assets/img/default.jpg'
import {deleteAdmin} from '@/api/admin'

export default {
  props: {
    admin: {
      type: Object
    }
  },
  components: {
    PermissionModal
  },
  data() {
    return {
      permissionVisible: false,
      defaultAvatar
    }
  },
  methods: {
    confirmDelete() {
      this.$dialog.confirm()
    },
    deleteUser(id){
      this.$bvModal.msgBoxConfirm('确定要删除该用户吗？删除后不可恢复', {
        title: '请确认',
        size: 'sm',
        buttonSize: 'sm',
        okVariant: 'danger',
        okTitle: '确定删除',
        cancelTitle: '再想想',
        footerClass: 'p-2',
        hideHeaderClose: false,
        centered: true
      })
      .then(value => {
        if (value) {
          deleteAdmin(id).then(() => {
            this.$emit('ondeleted')
          })
        }
      })
    }
  }
}
</script>