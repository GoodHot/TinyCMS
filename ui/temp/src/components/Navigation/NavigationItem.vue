<template>
  <li class="nav-main-item">
    <a :class="classes" @click="router" href="javascript:void(0)">
      <TIcon class="nav-main-link-icon" :icon="icon.icon" :pack="icon.pack" />
      <span class="nav-main-link-name">{{ data.title }}</span>
      <b-badge class="nav-main-link-badge" pill :variant="badge.type" v-if="badge.text">{{ badge.text }}</b-badge>
    </a>
    <ul class="nav-main-submenu" v-if="data.children">
      <TNavigationItem v-for="item of data.children" :key="item.title + item.icon" :data="item" />
    </ul>
  </li>
</template>
<script>
export default {
  props: {
    data: {
      type: Object
    }
  },
  data() {
    return {
      icon: {
        icon: 'home',
        pack: 'fa'
      },
      badge: {
        type: 'primary',
        theme: 'pill',
        text: ''
      }
    }
  },
  mounted() {
    this.getIcon()
    this.getBadge()
  },
  computed: {
    classes() {
      return 'nav-main-link ' + (this.data.active ? 'active' : '' + (this.data.children ? 'nav-main-link-submenu' : ''))
    }
  },
  methods: {
    getIcon() {
      if (typeof(this.data.icon) === 'string') {
        this.icon.icon = this.data.icon
      } else {
        this.icon = this.data.icon
      }
    },
    getBadge() {
      if (!this.data.badge) {
        return
      }
      if (typeof(this.data.badge) === 'string') {
        this.badge.text = this.data.badge
      } else {
        this.badge = this.data.badge
      }
    },
    router() {
      if (this.data.to) {
        this.$router.push(this.data.to)
      } else if (this.data.href) {
        window.location.href = this.data.href
      }else {
        return false
      }
    }
  }
}
</script>