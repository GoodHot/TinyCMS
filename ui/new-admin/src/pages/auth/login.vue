<template>
  <div class="row justify-content-center">
    <div class="col-md-8 col-lg-6 col-xl-4">
      <TBlock title="登录" class="block-themed block-fx-shadow mb-0">
        <div class="p-sm-3 px-lg-4 py-lg-5">
          <h1 class="mb-2">TinyCMS</h1>
          <p>欢迎，请登录</p>
          <TForm>
            <TFormItem col="12" v-slot:default="item" >
              <TInput placeholder="用户名" :state="item.state" v-model="form.username" theme="alt" valid="" />
            </TFormItem>
            <TFormItem col="12" v-slot:default="item" >
              <TInput placeholder="密码" type="password" :state="item.state" v-model="form.password" theme="alt" />
            </TFormItem>
            <TFormItem col="12" >
              <TButton icon="sign-in-alt" class="pr-4" @click="loginHandler">登录</TButton>
            </TFormItem>
          </TForm>
        </div>
      </TBlock>
    </div>
  </div>
</template>
<script>
import Vue from 'vue'
import { userLogin } from '@/api/auth'
export default {
  data() {
    return {
      form: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    loginHandler() {
      userLogin({
        username: this.form.username,
        password: this.form.password
      }).then(res => {
        const _this = this
        Vue.ls.set('ACCESS-TOKEN', res.info.token)
        this.$dialog.success('登录成功', '', () => {
          _this.$router.push('/')
        })
      })
    }
  }
}
</script>