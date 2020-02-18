
export const routerMap = [
  {
    path: '/',
    name: 'hello',
    component: () => import('@/layouts/BaseLayout'),
    children: [
      {
        path: '/',
        name: 'hello',
        component: () => import('@/pages/index'),
      },
      {
        path: '/content',
        name: 'content',
        component: () => import('@/pages/content/content'),
      },
      {
        path: '/publish',
        name: 'publish',
        component: () => import('@/pages/publish/publish'),
      },
      {
        path: '/user',
        name: 'user',
        component: () => import('@/pages/user/user'),
      },
      {
        path: '/general',
        name: 'general',
        component: () => import('@/pages/setting/general'),
      },
      {
        path: '/component/icon',
        name: 'icon',
        component: () => import('@/pages/components/icon'),
      },
      {
        path: '/component/button',
        name: 'button',
        component: () => import('@/pages/components/button'),
      },
      {
        path: '/component/badge',
        name: 'badge',
        component: () => import('@/pages/components/badge')
      },
      {
        path: '/component/form',
        name: 'form',
        component: () => import('@/pages/components/form')
      },
      {
        path: '/component/block',
        name: 'block',
        component: () => import('@/pages/components/block')
      },
      {
        path: '/component/modal',
        name: 'modal',
        component: () => import('@/pages/components/modal')
      },
      {
        path: '/component/table',
        name: 'table',
        component: () => import('@/pages/components/table')
      },
      {
        path: '/component/popover',
        name: 'popover',
        component: () => import('@/pages/components/popover')
      },
      {
        path: '/component/checkbox',
        name: 'checkbox',
        component: () => import('@/pages/components/checkbox')
      },
      {
        path: '/component/select',
        name: 'select',
        component: () => import('@/pages/components/select')
      }
    ]
  }
]