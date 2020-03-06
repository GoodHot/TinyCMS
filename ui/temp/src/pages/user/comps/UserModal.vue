<template>
  <b-form>
    <b-form-group label="用户昵称">
      <b-form-input placeholder="用户昵称" :state="form.nickname.valid" v-model="form.nickname.value"></b-form-input>
      <b-form-invalid-feedback :state="form.nickname.valid">
        {{form.nickname.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="密码" >
      <b-form-input placeholder="密码" type="password" :state="form.password.valid" v-model="form.password.value"></b-form-input>
      <b-form-invalid-feedback :state="form.password.valid">
        {{form.password.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="用户名" >
      <b-form-input placeholder="用户名" :state="form.username.valid" v-model="form.username.value"></b-form-input>
      <b-form-invalid-feedback :state="form.username.valid">
        {{form.username.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="选择角色" >
      <b-form-select :state="form.username.valid" v-model="form.role.value" :options="roleOptions"></b-form-select>
      <b-form-invalid-feedback :state="form.role.valid">
        {{form.role.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="简介" >
      <b-form-textarea placeholder="简介" rows="3" max-rows="6" v-model="form.resume.value"></b-form-textarea>
      <b-form-invalid-feedback :state="form.resume.valid">
        {{form.resume.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
    <b-form-group label="上传头像">
      <TUpload @uploaded="avatarUploaded"></TUpload>
      <b-form-invalid-feedback :state="form.avatar.valid">
        {{form.avatar.msg}}
      </b-form-invalid-feedback>
    </b-form-group>
  </b-form>
</template>
<script>
import valid from '@/utils/valid'
import {getAllRole} from '@/api/role'
import {saveAdmin} from '@/api/admin'

export default {
  data() {
    return {
      form: {
        nickname: {
          value: '',
          require: '请输入用户昵称'
        },
        password: {
          value: '',
          require: '请输入密码'
        },
        username: {
          value: '',
          require: '请输入用户名'
        },
        resume: {
          value: ''
        },
        role: {
          value: null,
          require: '请选择角色'
        },
        avatar: {
          value: null
        }
      },
      roleOptions: []
    }
  },
  mounted() {
    this.loadRoles()
  },
  methods: {
    submit() {
      const rst = valid.check(this.form)
      this.form = rst.values
      if (!rst.valid) {
        return false
      }
      const param = {
        nickname: this.form.nickname.value,
        password: this.form.password.value,
        user_name: this.form.username.value,
        role_id: this.form.role.value,
        resume: this.form.resume.value,
        avatar: this.form.avatar.value
      }
      saveAdmin(param).then(() => {
        this.$emit('savedUser')
      })
    },
    loadRoles() {
      getAllRole().then(res => {
        const roles = [{
          value: null,
          text: '请选择角色'
        }]
        for (let i in res.roles) {
          const role = res.roles[i]
          roles.push({
            value: role.id,
            text: role.role_name
          })
        }
        this.roleOptions = roles
      })
    },
    avatarUploaded(url) {
      this.form.avatar.value = url
    }
  }
}
</script>