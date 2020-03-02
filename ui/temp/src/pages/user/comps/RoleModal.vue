<template>
  <b-form>
    <b-form-group label="角色编码" >
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
    <b-form-group label="拥有权限" >
      <PermissionList></PermissionList>
    </b-form-group>
  </b-form>
</template>
<script>
import PermissionList from './PermissionList'
export default {
  components: {
    PermissionList
  },
  data() {
    return {
      form: {
        code: {
          value: '',
          require: '请输入角色编码'
        },
        role_name: {
          value: '',
          require: '请输入角色名称'
        }
      }
    }
  }
}
</script>