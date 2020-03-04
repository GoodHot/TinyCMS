<template>
  <b-form>
    <b-form-group label="角色编码" description="只允许输入字母和数字">
      <b-form-input placeholder="角色编码" :state="form.code.valid" v-model="form.code.value"></b-form-input>
      <b-form-invalid-feedback :state="form.code.valid">
        {{form.code.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="角色名称" >
      <b-form-input placeholder="角色名称" :state="form.role_name.valid" v-model="form.role_name.value"></b-form-input>
      <b-form-invalid-feedback :state="form.role_name.valid">
        {{form.role_name.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-alert variant="info" show v-if="isSuper">超级管理员拥有所有权限</b-alert>
    <b-form-group label="拥有权限" v-if="!isSuper" >
      <PermissionList ref="permissionList" :checkedValues="checkedValues" :isSuper="isSuper"></PermissionList>
      <b-form-invalid-feedback :state="form.permission.valid">
        {{form.permission.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
  </b-form>
</template>
<script>
import valid from '@/utils/valid'
import PermissionList from './PermissionList'
import {saveRole, getRole} from '@/api/role'

export default {
  components: {
    PermissionList
  },
  props: {
    editId: {
      type: Number,
      default: 0
    }
  },
  data() {
    return {
      form: {
        id: {
          value: 0
        },
        code: {
          value: '',
          require: '请输入角色编码'
        },
        role_name: {
          value: '',
          require: '请输入角色名称'
        },
        permission: {
          value: []
        }
      },
      checkedValues: [],
      isSuper: false
    }
  },
  mounted() {
    this.loadEditRole()
  },
  methods: {
    submit() {
      const rst = valid.check(this.form)
      this.form = rst.values
      const permission = []
      if (!this.isSuper) {
        this.$refs.permissionList.getValue()
      }
      if (!permission || permission == 0) {
        this.form.permission.valid = false
        this.form.permission.msg = "请选择角色拥有权限"
        rst.valid = false
      } else {
        this.form.permission.value = permission
        this.form.permission.valid = true
      }
      if (!rst.valid) {
        return false
      }
      const param = {
        id: this.form.id.value,
        code: this.form.code.value,
        name: this.form.role_name.value,
        permission: this.form.permission.value
      }
      saveRole(param).then(() => {
        this.$emit('savedRole')
      })
    },
    loadEditRole() {
      if (this.editId === 0) {
        return
      }
      getRole(this.editId).then(res => {
        this.form.code.value = res.role.code
        this.form.role_name.value = res.role.role_name
        this.form.id.value = res.role.id
        this.isSuper = res.role.is_super
        for (let i in res.role.permissions) {
          for (let j in res.role.permissions[i]) {
            this.checkedValues.push(res.role.permissions[i][j].id)
          }
        }
      })
    }
  }
}
</script>