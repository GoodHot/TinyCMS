<template>
<li :class="liClasses">
  <span v-if="type === 'heading'">{{ title }}</span>
  <a :class="linkClasses" href="javascript:void(0)" @click="navClick" v-else>
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
      :blockIsHead="blockIsHead"
      @onopen="onopenHandler"
    >
    </t-navigation-item>
  </ul>
</li>
</template>
<script>
import {PropTypes} from '@/utils/types'

export default {
  model: {
    prop: 'active',
    event: 'activeChange'
  },
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
    showEvent: PropTypes.String.def('hover'),
    blockIsHead: PropTypes.Boolean,
    type: PropTypes.String,
  },
  data () {
    return {
      isOpen: false
    }
  },
  computed: {
    liClasses () {
      return [
        {
          'open': this.isOpen,
          'nav-main-heading': this.type === 'heading',
          'nav-main-item': this.type !== 'heading'
        }
      ]
    },
    linkClasses () {
      return [
        'nav-main-link',
        {
          'active': this.isOpen,
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
      if (this.child && this.child.length > 0) {
        if (!this.horizontal || this.showEvent === 'click') {
          this.isOpen = !this.isOpen
          this.$emit("onopen", this)
          return
        }
      } else {
        if (this.href) {
          window.location = this.href
        } else {
          this.$router.push(this.to)
        }
      }
    },
    onopenHandler () {

    },
    $close () {
      this.isOpen = false
    }
  }
}
</script>