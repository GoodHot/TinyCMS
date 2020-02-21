<template>
  <div class="row justify-content-center">
    <div class="col-md-8 col-lg-6 col-xl-4">
      <TBlock title="登录" subtitle="TinyCMS" class="block-themed block-fx-shadow mb-0">
        <div class="p-sm-3 px-lg-4 py-lg-5">
          <h1 class="mb-2">TinyCMS</h1>
          <p>欢迎，请登录</p>
          <b-form @submit="submitHandler">
            <b-form-group>
              <b-form-input
                size="lg"
                class="form-control-alt"
                type="text"
                placeholder="用户名"
                v-model="form.username.value"
              ></b-form-input>
              <b-form-invalid-feedback :state="form.username.valid">
                {{form.username.msg}}
              </b-form-invalid-feedback>
            </b-form-group>
            <b-form-group>
              <b-form-input
                size="lg"
                class="form-control-alt"
                type="password"
                placeholder="密码"
                v-model="form.password.value"
              ></b-form-input>
              <b-form-invalid-feedback :state="form.password.valid">
                {{form.password.msg}}
              </b-form-invalid-feedback>
            </b-form-group>
            <b-button type="submit" variant="primary" class="pr-4"><TIcon icon="sign-in-alt" /> 登录</b-button>
          </b-form>
        </div>
      </TBlock>
    </div>
  </div>
</template>
<script>
import Vue from 'vue'
import valid from '@/utils/valid'
import {userLogin} from '@/api/auth'

export default {
  data() {
    return {
      form: {
        username: {
          value: '',
          require: '请输入用户名'
        },
        password: {
          value: '',
          require: '请输入密码',
          length: {
            max: 20,
            min: 3,
            msg: '请输入正确长度3到20位'
          }
        }
      }
    }
  },
  methods: {
    submitHandler() {
      const rst = valid.check(this.form)
      this.form = rst.values
      if (!rst.valid) {
        return false
      }
      userLogin({
        username: this.form.username.value,
        password: this.form.password.value
      }).then(res => {
        const _this = this
        Vue.ls.set('ACCESS-TOKEN', res.info.token)
        this.$dialog.success('登录成功', '', () => {
          _this.$router.push('/')
        })
      })
      return false

    }
  }
}
</script>