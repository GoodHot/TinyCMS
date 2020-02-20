<template>
  <div class="row justify-content-center">
    <div class="col-md-8 col-lg-6 col-xl-4">
      <TBlock title="登录" subtitle="TinyCMS" class="block-themed block-fx-shadow mb-0">
        <div class="p-sm-3 px-lg-4 py-lg-5">
          <h1 class="mb-2">TinyCMS</h1>
          <p>欢迎，请登录</p>
          <b-form>
            <b-form-group>
              <b-form-input
                size="lg"
                class="form-control-alt"
                type="email"
                placeholder="用户名"
              ></b-form-input>
            </b-form-group>
            <b-form-group>
              <b-form-input
                size="lg"
                class="form-control-alt"
                type="email"
                placeholder="密码"
              ></b-form-input>
            </b-form-group>
            <b-button type="submit" variant="primary" class="pr-4"><TIcon icon="sign-in-alt" /> 登录</b-button>
          </b-form>
        </div>
      </TBlock>
    </div>
  </div>
</template>
<script>
// import Vue from 'vue'
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
      // userLogin({
      //   username: this.form.username,
      //   password: this.form.password
      // }).then(res => {
      //   const _this = this
      //   Vue.ls.set('ACCESS-TOKEN', res.info.token)
      //   this.$dialog.success('登录成功', '', () => {
      //     _this.$router.push('/')
      //   })
      // })
    }
  }
}
</script>