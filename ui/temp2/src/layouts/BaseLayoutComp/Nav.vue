<template>
  <nav id="sidebar" aria-label="Main Navigation">
    <div class="bg-header-dark">
      <div class="content-header bg-white-10">
        <a class="link-fx font-w700 font-size-lg text-white" href="index.html">
          <span class="smini-visible">
            <span class="text-white-75">
              <i class="fab fa-wordpress"></i>
            </span>
          </span>
          <span class="smini-hidden">
            <span class="text-white">TinyCMS</span>
          </span>
        </a>
        <div>
          <a
            class="d-lg-none text-white ml-2"
            data-toggle="layout"
            data-action="sidebar_close"
            href="javascript:void(0)"
          >
            <i class="fa fa-times-circle"></i>
          </a>
        </div>
      </div>
    </div>
    <div class="content-side content-side-full">
      <t-navigation
        :horizontal="false"
        rounded
        :data="navigation"
      >
      </t-navigation>
    </div>
  </nav>
</template>
<script>
export default {
  data () {
    return {
      navigation: [
        {
          title: '控件',
          icon: 'campground',
          active: true,
          child: [
            {
              title: '图标',
              to: '/elements/icon'
            },
            {
              title: 'Block',
              to: '/elements/block'
            }
          ]
        },
        {
          title: '组件',
          icon: 'car-battery',
          child: [
            {
              title: '导航条',
              to: '/components/navigation'
            },
            {
              title: '布局',
              to: '/components/layout'
            }
          ]
        }
      ]
    }
  }
}
</script>