
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
        path: '/icon',
        name: 'icon',
        component: () => import('@/pages/icon'),
      },
      {
        path: '/button',
        name: 'button',
        component: () => import('@/pages/button'),
      },
      {
        path: '/badge',
        name: 'badge',
        component: () => import('@/pages/badge')
      }
    ]
  }
]