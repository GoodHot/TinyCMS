<template>
  <t-layout
    :sidebarShow="sidebarShow"
    :sidebarPosition="sidebarPosition"
    :sidebarMin="sidebarMin"
    :sidebarTheme="sidebarTheme"
    :sideVisiable="sideVisiable"
    :sideHoverMode="sideHoverMode"
    :headerMode="headerMode"
    :headerGlass="headerGlass"
    :headerTheme="headerTheme"
    :footerMode="footerMode"
    :contentMode="contentMode"
    @onSideHide="sideVisiable = false"
  >


    <template slot="nav">
      <div class="bg-header-dark">
        <div class="content-header bg-white-10">
          <a class="link-fx font-w600 font-size-lg text-white" href="index.html">
            <span class="smini-visible">
              <span class="text-white-75">T</span>
              <span class="text-white">in</span>
            </span>
            <span class="smini-hidden">
              <span class="text-white">Tiny</span>
              <span class="text-white-75">CMS</span>
              <span class="text-white font-size-base font-w400 ml-1">1.0</span>
            </span>
          </a>
          <div>
            <a class="js-class-toggle text-white-75" href="javascript:void(0)" @click="sidebarTheme = sidebarTheme === 'dark' ? 'light': 'dark'">
              <t-icon :icon="sidebarTheme === 'dark' ? 'toggle-on': 'toggle-off'" />
            </a>
            <a class="d-lg-none text-white ml-2" href="javascript:void(0)" >
              <t-icon icon="times-circle" />
            </a>
          </div>
        </div>
      </div>
      <div class="content-side content-side-full">
        <t-navigation :horizontal="false" rounded :data="navigation" :theme="sidebarTheme"></t-navigation>
      </div>
    </template>
    
    
    <template slot="header">
      <div class="content-header">
        <div>
          <b-button variant="dual" class="mr-1" @click="sidebarShow = !sidebarShow" ><t-icon size="fw" icon="bars" /></b-button>
          <b-button variant="dual"><t-icon size="fw" icon="search" /> <span class="ml-1 d-none d-sm-inline-block">Search</span></b-button>
        </div>
      </div>
      <div id="page-header-loader" class="overlay-header bg-primary-darker show" v-show="loadingShow">
        <div class="content-header">
            <div class="w-100 text-center">
              <t-icon icon="sun" class="fa-fw fa-spin text-white" />
              {{ loadingText }}
            </div>
        </div>
    </div>
    </template>


    <template slot="footer">
      <div class="content py-0">
        <div class="row font-size-sm">
            <div class="col-sm-6 order-sm-2 mb-1 mb-sm-0 text-center text-sm-right">
                Crafted with <i class="fa fa-heart text-danger"></i> by <a class="font-w600" href="https://1.envato.market/ydb" target="_blank">HotGood</a>
            </div>
            <div class="col-sm-6 order-sm-1 text-center text-sm-left">
                <a class="font-w600" href="https://1.envato.market/r6y" target="_blank">TinyCMS 1.0</a> &copy; <span data-toggle="year-copy"></span>
            </div>
        </div>
      </div>
    </template>


    <template slot="side">
      aaaaa
    </template>


    <div class="bg-body-light" v-if="title">
      <div class="content content-full">
        <h1 class="font-size-h2 mb-2">{{ title }}</h1>
        <h2 class="font-size-lg font-w400 text-muted mb-0">{{ description }}</h2>
      </div>
    </div>
    <div :class="!contentSide ? 'content' : ''">
      <router-view
        ref="routerView"
        @sidebarVisiable="onSidebarVisiable"
        @sidebarPosition="onSidebarPosition"
        @sidebarMin="onSidebarMin"
        @sidebarTheme="onSidebarTheme"
        @sideVisiable="onSideVisiable"
        @sideHoverMode="onSideHoverMode"
        @headerMode="onHeaderMode"
        @headerGlass="onHeaderGlass"
        @headerTheme="onHeaderTheme"
        @footerMode="onFooterMode"
        @contentMode="onContentMode"
        @onloading="onloading"
      />
    </div>
  </t-layout>
</template>
<script>
export default {
  mounted () {
    this.title = this.$refs.routerView.title
    this.description = this.$refs.routerView.description
    this.contentSide = this.$refs.routerView.contentSide
  },
  updated () {
    this.title = this.$refs.routerView.title
    this.description = this.$refs.routerView.description
    this.contentSide = this.$refs.routerView.contentSide
  },
  methods: {
    onSidebarVisiable (type) {
      this.sidebarShow = type
    },
    onSidebarPosition (pos) {
      this.sidebarPosition = pos
    },
    onSidebarMin (type) {
      this.sidebarMin = type
    },
    onSidebarTheme (theme) {
      this.sidebarTheme = theme
    },
    onSideVisiable (type) {
      this.sideVisiable = type
    },
    onSideHoverMode (type) {
      this.sideHoverMode = type
    },
    onHeaderMode (mode) {
      this.headerMode = mode
    },
    onHeaderGlass (type) {
      this.headerGlass = type
    },
    onHeaderTheme (theme) {
      this.headerTheme = theme
    },
    onFooterMode (mode) {
      this.footerMode = mode
    },
    onContentMode (mode) {
      this.contentMode = mode
    },
    onloading (state, text) {
      this.loadingShow = state
      this.loadingText = text
    }
  },
  data() {
    return {
      title: '',
      description: '',
      contentSide: false,
      sidebarShow: true,
      sidebarPosition: 'left',
      sidebarMin: false,
      sidebarTheme: 'light',
      sideVisiable: false,
      sideHoverMode: false,
      headerMode: 'fixed',
      headerGlass: false,
      headerTheme: 'dark',
      footerMode: 'static',
      contentMode: 'full',
      loadingShow: false,
      loadingText: '',
      navigation: [
        {
          title: '控制台',
          icon: 'cursor',
          iconPack: 'si',
          to: '/'
        },
        {
          title: 'CONTENT',
          type: 'heading'
        },
        {
          title: '发布文章',
          icon: 'note',
          iconPack: 'si',
          to: '/article/publish'
        },
        {
          title: '文章管理',
          icon: 'book-open',
          iconPack: 'si',
          to: '/article'
        },
        {
          title: '分类管理',
          icon: 'directions',
          iconPack: 'si',
          to: '/category/1'
        },
        {
          title: '标签管理',
          icon: 'tag',
          iconPack: 'si',
          to: '/tags/1'
        },
        {
          title: 'SETTINGS',
          type: 'heading'
        },
        {
          title: '自定义标签',
          icon: 'pin',
          iconPack: 'si',
          to: '/label'
        },
        {
          title: '皮肤管理',
          icon: 'magic-wand',
          iconPack: 'si',
          to: '/skin'
        },
        // {
        //   title: '资源管理',
        //   icon: 'paper-clip',
        //   iconPack: 'si',
        //   to: '/resource'
        // },
        {
          title: 'SYSTEM',
          type: 'heading'
        },
        {
          title: '系统设置',
          icon: 'equalizer',
          iconPack: 'si',
          to: '/setting'
        },
        // {
        //   title: '用户管理',
        //   icon: 'users',
        //   iconPack: 'si',
        //   to: '/admin'
        // }
        // {
        //   title: 'TOOL',
        //   type: 'heading'
        // },
        // {
        //   title: "控件",
        //   icon: "campground",
        //   active: true,
        //   child: [
        //     {
        //       title: "图标",
        //       to: "/elements/icon"
        //     },
        //     {
        //       title: "Block",
        //       to: "/elements/block"
        //     },
        //     {
        //       title: "Dropdown",
        //       to: "/elements/dropdown"
        //     }
        //   ]
        // },
        // {
        //   title: "组件",
        //   icon: "car-battery",
        //   child: [
        //     {
        //       title: "导航条",
        //       to: "/components/navigation"
        //     },
        //     {
        //       title: "布局",
        //       to: "/components/layout"
        //     },
        //     {
        //       title: "Nestable",
        //       to: "/components/nestable"
        //     }
        //   ]
        // }
      ]
    };
  }
};
</script>