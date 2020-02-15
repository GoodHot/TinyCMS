<template>
  <div class="content">
    <TNavigation :data="navData" />
    <router-view/>
  </div>
</template>
<script>
export default {
  data() {
    return {
      navData: [
        {
          title: 'Home',
          icon: 'home',
          active: true,
          to: '/'
        },
        {
          title: '组件',
          icon: 'ad',
          badge: '3',
          children: [
            {
              title: 'Badge',
              icon: {
                icon: 'badge',
                pack: 'si'
              },
              to: '/badge'
            },
            {
              title: 'Button',
              icon: {
                icon: 'circle',
                pack: 'far'
              },
              to: '/button'
            },
            {
              title: 'Icon',
              icon: 'icons',
              badge: '5',
              to: '/icon'
            },
            {
              title: 'Form',
              icon: {
                icon: 'wpforms',
                pack: 'fab'
              },
              to: '/form'
            },
            {
              title: 'Block',
              icon: {
                icon: 'elementor',
                pack: 'fab'
              },
              to: '/block'
            },
            {
              title: 'Modal',
              icon: {
                icon: 'frame',
                pack: 'si'
              },
              to: '/modal'
            }
          ]
        }
      ]
    }
  }
}
</script>