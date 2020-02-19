<template>
  <div class="row justify-content-center">
    <div class="col-md-8 col-lg-6 col-xl-4">
      <TBlock title="登录" class="block-themed block-fx-shadow mb-0">
        <div class="p-sm-3 px-lg-4 py-lg-5">
          <h1 class="mb-2">TinyCMS</h1>
          <p>欢迎，请登录</p>
          <TForm>
            <TFormItem col="12" v-slot:default="item" >
              <TInput placeholder="用户名" :state="item.state" theme="alt" />
            </TFormItem>
            <TFormItem col="12" v-slot:default="item" >
              <TInput placeholder="密码" type="password" :state="item.state" theme="alt" />
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
import { userLogin } from '@/api/auth'
export default {
  methods: {
    loginHandler() {
      userLogin({
        username: 'asdfasdf',
        password: 'ffffff'
      }).then(res => {
        console.log(res)
      })
    }
  }
}
</script>