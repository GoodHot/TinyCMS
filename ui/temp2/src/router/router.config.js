
export const routerMap = [
  {
    path: '/',
    name: 'index',
    component: () => import('@/layouts/BaseLayout'),
    children: [
      {
        path: '/',
        name: 'index',
        component: () => import('@/pages/index'),
      },
      {
        path: '/elements/icon',
        name: 'icon',
        component: () => import('@/pages/elements/icon'),
      },
      {
        path: '/elements/block',
        name: 'block',
        component: () => import('@/pages/elements/block'),
      },
      {
        path: '/elements/navigation',
        name: 'navigation',
        component: () => import('@/pages/elements/navigation'),
      }
    ]
  }
]