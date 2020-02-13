<template>
  <div class="content">
    <TNavigation :data="navData" />
    <div class="block d-none d-lg-block">
        <div class="block-content">
            <router-view/>
        </div>
    </div>
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
            }
          ]
        },
        {
          title: 'Book',
          icon: 'book',
          to: '/'
        }
      ]
    }
  }
}
</script>