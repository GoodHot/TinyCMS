<template>
  <t-input-drawer ref="userInput" placeholder="选择用户" :clearVariant="clearVariant" :dataValue="choose.nickname" :size="size" :inputClass="inputClass" @onshow="onshowHandler" :readonly="true" @clear="clearHandler">
    <b-overlay :show="loading" rounded="sm">
      <div style="max-height: 200px; overflow-y: scroll">
        <ul class="t-category-drawer">
          <li v-for="user in users" :key="user.id" @click="chooseItem(user)">
            <img :src="assetsURL(user.avatar)" width="20" height="20" class="mr-2" />
            {{ user.nickname }}
            <span v-if="user.is_self">(我)</span>
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
  model: {
    prop: 'value',
    event: 'value.change'
  },
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
    inputClass: PropTypes.String,
    clearVariant: PropTypes.String.def('light'),
    value: PropTypes.Number
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
      if (this.users && this.users.length > 0) {
        return
      }
      getAllAdmin().then(res => {
        this.loading = false
        this.hasData = true
        this.users = res.admins
        this.setValue()
      })
    },
    clearHandler () {
      this.choose.nickname = ''
      this.$emit('change', 0)
    },
    setValue () {
      this.users.map(u => {
        if (u.id === this.value) {
          this.choose.nickname = u.nickname + ''
        }
        console.log(u)
      })
    }
  },
  watch: {
    value () {
      this.loadAdmin()
    }
  }
}
</script>