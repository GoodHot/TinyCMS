<template>
  <TBlock title="角色管理">
    <template slot="options">
      <b-button-group size="sm">
        <b-button class="btn-block-option" variant="light" @click="roleVisible = true">
          <TIcon icon="plus" />
        </b-button>
      </b-button-group>
    </template>
    <div class="pull-x">
      <TTable :pagination="{}" :column="column" :data="data" ref="dataTable" class="table table-hover table-vcenter font-size-sm">
        <template v-slot:roleName="{item}" class="d-none d-sm-table-cell font-w600">
          {{item.role_name}}
        </template>
        <template v-slot:permission="{item}">
          <div v-if="item.is_super">
            <b-badge :id="`superTip_${item.id}`" variant="dark">超级管理员</b-badge>
            <b-tooltip :target="`superTip_${item.id}`">
              超级管理员拥有所有权限
            </b-tooltip>
          </div>
          <div v-for="(val, key) of item.permissions" :key="key" class="mb-2">
            <b-badge variant="primary">{{key}}</b-badge>
            <b-badge class="ml-2" variant="info" v-for="per of val" :key="per.id">{{per.permission_name}}</b-badge>
          </div>
        </template>
        <template v-slot:action="{item}">
          <b-button-group size="sm">
            <b-button variant="light" @click="editRole(item.id)"><TIcon icon="edit" pack="far" /> 编辑</b-button>
            <b-button variant="light"><TIcon icon="trash-alt" pack="far" /> 删除</b-button>
          </b-button-group>
        </template>
      </TTable>
    </div>
    <b-modal
      v-model="roleVisible"
      size="lg"
      title="创建角色"
      no-close-on-backdrop
      @ok="submitRole"
      @hide="roleModalHide"
      ok-title="保存"
      cancel-title="关闭"
    >
      <RoleModal ref="roleModal" :editId="editId" @savedRole="savedRoleHandler"></RoleModal>
    </b-modal>
  </TBlock>
</template>
<script>
import RoleModal from './comps/RoleModal'
import {pageRole} from '@/api/role'
export default {
  components: {
    RoleModal
  },
  mounted() {
    this.loadPage()
  },
  data() {
    return {
      editId: 0,
      roleVisible: false,
      column: [
        {
          title: 'ID',
          index: 'id',
          align: 'center',
          width: '100px',
          class: 'd-none d-sm-table-cell font-w600'
        },
        {
          title: '编码',
          index: 'code',
          width: '150px',
          class: 'd-none d-sm-table-cell'
        },
        {
          title: '角色',
          index: 'role_name',
          slot: 'roleName',
          width: '250px',
          class: 'font-w600'
        },
        {
          title: '权限',
          slot: 'permission'
        },
        {
          title: '用户数',
          index: 'admin_count',
          width: '120px',
          class: 'd-none d-sm-table-cell font-w600'
        },
        {
          title: '操作',
          slot: 'action',
          width: '200px',
          class: 'd-none d-xl-table-cell text-muted'
        }
      ],
      data: []
    }
  },
  methods: {
    loadPage() {
      pageRole(1).then(res => {
        this.data = res.page.list
      })
    },
    submitRole(bvModalEvt) {
      bvModalEvt.preventDefault()
      this.$refs.roleModal.submit()
      return false;
    },
    savedRoleHandler() {
      this.$bvToast.toast('角色设置成功', {
        title: '提示',
        variant: 'info',
        solid: true
      })
      this.roleVisible = false
      this.loadPage()
    },
    editRole(id) {
      this.roleVisible = true
      this.editId = id
    },
    roleModalHide() {
      this.editId = 0
    }
  }
}
</script>