<template>
  <t-input-drawer ref="userInput" placeholder="选择用户" :dataValue="choose.nickname" :size="size" :inputClass="inputClass" @onshow="onshowHandler" :readonly="true" @clear="clearHandler">
    <b-overlay :show="loading" rounded="sm">
      <div style="max-height: 200px; overflow-y: scroll">
        <ul class="t-category-drawer">
          <li v-for="user in users" :key="user.id" @click="chooseItem(user)">
            <img :src="assetsURL(user.avatar)" width="20" height="20" class="mr-2" />
            {{ user.nickname }}
          </li>
        </ul>
      </div>
    </b-overlay>
  </t-input-drawer>
</template>
<script>
import { PropTypes } from '@/utils/types'
import { getAllAdmin } from '@/api/admin'

export default {
  data () {
    return {
      loading: false,
      hasData: false,
      users: [],
      choose: {
        nickname: ''
      }
    }
  },
  props: {
    size: PropTypes.String,
    inputClass: PropTypes.String
  },
  methods: {
    chooseItem (item) {
      this.choose.nickname = item.nickname + ''
      this.$refs.userInput.hideDrawer()
      this.$emit('change', item.id)
    },
    calcPaddingLeft (level) {
      if (level === 0) {
        return ''
      }
      return `padding-left: ${1 + level}rem`
    },
    onshowHandler () {
      if (!this.hasData) {
        this.loadAdmin()
      }
    },
    loadAdmin () {
      getAllAdmin().then(res => {
        this.loading = false
        this.hasData = true
        this.users = res.admins
      })
    },
    clearHandler () {
      this.choose.nickname = ''
      this.$emit('change', 0)
    }
  }
}
</script>