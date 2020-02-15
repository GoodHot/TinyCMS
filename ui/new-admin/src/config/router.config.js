
export const routerMap = [
  {
    path: '/',
    name: 'hello',
    component: () => import('@/layouts/BaseLayout'),
    children: [
      {
        path: '/',
        name: 'hello',
        component: () => import('@/pages/components/index'),
      },
      {
        path: '/content',
        name: 'content',
        component: () => import('@/pages/content'),
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
        path: '/modal',
        name: 'modal',
        component: () => import('@/pages/components/modal')
      },
      {
        path: '/table',
        name: 'table',
        component: () => import('@/pages/components/table')
      }
    ]
  }
]