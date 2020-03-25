<template>
  <t-block rounded class="w-100 mb-0 overflow-hidden py-lg-5" theme="transparent">
    <div class="mb-2 text-center">
      <a class="link-fx font-w700 font-size-h1" href="javascript:void(0)">
        <span class="text-dark">Tiny</span>
        <span class="text-primary">CMS</span>
      </a>
      <p class="text-uppercase font-w700 font-size-sm text-muted">登 录</p>
    </div>
    <b-form @submit="onSubmit">
      <b-form-group>
        <b-form-input class="form-control-alt" v-model="form.username" placeholder="输入用户名"></b-form-input>
      </b-form-group>
      <b-form-group>
        <b-form-input class="form-control-alt" v-model="form.password" type="password" placeholder="输入登录密码"></b-form-input>
      </b-form-group>
      <b-form-group>
        <b-form-checkbox>
          下次自动登录
        </b-form-checkbox>
      </b-form-group>
      <b-form-group class="text-center">
        <b-overlay
          :show="loading"
          rounded
          opacity="0.6"
          spinner-small
          spinner-variant="primary"
          class="d-inline-block"
        >
          <b-button type="submit" variant="hero-primary" block>
            <t-icon icon="sign-in-alt" size="fw" class="mr-1" /> 登 录
          </b-button>
        </b-overlay>
      </b-form-group>
    </b-form>
  </t-block>
</template>
<script>
import bgImg from '@/assets/imgs/auth_bg.jpg'
import { ACCESS_TOKEN } from '@/utils/request'
import { userLogin } from '@/api/auth'
import Vue from 'vue'

export default {
  data () {
    return {
      bgImg,
      form: {
        username: 'admin',
        password: 'admin'
      },
      loading: false
    }
  },
  methods: {
    onSubmit (evt) {
      evt.preventDefault()
      this.loading = true
      userLogin(this.form).then(res => {
        this.loading = false
        Vue.ls.set(ACCESS_TOKEN, res.info.token)
        window.setTimeout(() => {
          this.$router.push('/')
        }, 1000)
      })
    }
  }
}
</script>
