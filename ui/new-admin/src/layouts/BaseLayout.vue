<template>
  <div class="content">
    <TNavigation :data="navData" />
    <router-view />
    <footer id="page-footer" class="bg-body-light">
      <div class="content py-3">
        <div class="row font-size-sm">
          <div class="col-sm-6 order-sm-2 py-1 text-center text-sm-right">
            Create by TinyCMS
          </div>
          <div class="col-sm-6 order-sm-1 py-1 text-center text-sm-left">
            <a class="font-w600" href="https://1.envato.market/AVD6j" target="_blank">TinyCMS 1.0</a> ©
            <span data-toggle="year-copy" class="js-year-copy-enabled">2020</span>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>
<script>
export default {
  data() {
    return {
      navData: [
        {
          title: "Home",
          icon: "home",
          active: true,
          to: "/"
        },
        {
          title: "内容管理",
          icon: {
            icon: 'book-open',
            pack: 'si'
          },
          to: '/content'
        },
        {
          title: "用户管理",
          icon: {
            icon: 'users',
            pack: 'si'
          },
          to: '/user'
        },
        {
          title: "组件",
          icon: "ad",
          badge: "3",
          children: [
            {
              title: "Badge",
              icon: {
                icon: "badge",
                pack: "si"
              },
              to: "/component/badge"
            },
            {
              title: "Button",
              icon: {
                icon: "circle",
                pack: "far"
              },
              to: "/component/button"
            },
            {
              title: "Icon",
              icon: "icons",
              badge: "5",
              to: "/component/icon"
            },
            {
              title: "Form",
              icon: {
                icon: "wpforms",
                pack: "fab"
              },
              to: "/component/form"
            },
            {
              title: "Block",
              icon: {
                icon: "elementor",
                pack: "fab"
              },
              to: "/component/block"
            },
            {
              title: "Modal",
              icon: {
                icon: "frame",
                pack: "si"
              },
              to: "/component/modal"
            },
            {
              title: "Table",
              icon: {
                icon: "table",
                pack: "fa"
              },
              to: "/component/table"
            },
            {
              title: "Popover",
              icon: {
                icon: "window-maximize",
                pack: "far"
              },
              to: "/component/popover"
            },
            {
              title: "Checkbox",
              icon: {
                icon: "check-square",
                pack: "far"
              },
              to: "/component/checkbox"
            }
          ]
        }
      ]
    };
  }
};
</script>