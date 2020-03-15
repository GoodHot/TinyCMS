<template>
<li :class="liClasses">
  <a :class="linkClasses" href="javascript:void(0)" @click="navClick">
    <t-icon class="nav-main-link-icon" v-if="icon" :icon="icon" :pack="iconPack" />
    <span class="nav-main-link-name">{{title}}</span>
    <b-badge class="nav-main-link-badge" v-if="badgeText" pill :variant="badgeTheme">{{ badgeText }}</b-badge>
  </a>
  <ul class="nav-main-submenu" v-if="child && child.length > 0">
    <t-navigation-item
      v-for="(item, index) in child"
      :key="index"
      :active="item.active"
      :href="item.href"
      :to="item.to"
      :title="item.title"
      :badgeText="item.badgeText"
      :badgeTheme="item.badgeTheme"
      :icon="item.icon"
      :iconPack="item.iconPack"
      :child="item.child"
      :showEvent="showEvent"
      :horizontal="horizontal"
      @onopen="onopenHandler"
    >
    </t-navigation-item>
  </ul>
</li>
</template>
<script>
import {PropTypes} from '@/utils/types'

export default {
  props: {
    horizontal: PropTypes.Boolean.def(true),
    active: PropTypes.Boolean,
    href: PropTypes.String,
    to: PropTypes.String,
    title: PropTypes.String,
    badgeText: PropTypes.String,
    badgeTheme: PropTypes.String.def('success'),
    icon: PropTypes.String,
    iconPack: PropTypes.String.def('fa'),
    child: PropTypes.Array,
    showEvent: PropTypes.String.def('hover')
  },
  data () {
    return {
      isOpen: false
    }
  },
  computed: {
    liClasses () {
      return [
        'nav-main-item',
        {
          'open': this.isOpen || this.active
        }
      ]
    },
    linkClasses () {
      return [
        'nav-main-link',
        {
          'active': this.active,
          'nav-main-link-submenu': this.child && this.child.length > 0
        }
      ]
    }
  },
  mounted () {
    this.isOpen = this.active
  },
  methods: {
    navClick () {
      if (!this.horizontal || (this.showEvent === 'click' && this.child && this.child.length > 0)) {
        this.isOpen = !this.isOpen
        if (!this.isOpen) {
          this.active = false
        }
        this.$emit("onopen", this)
        return
      }
      if (this.href) {
        window.location = this.href
      } else {
        this.$router.push(this.to)
      }
    },
    $close () {
      this.active = false
      this.isOpen = false
    }
  }
}
</script>