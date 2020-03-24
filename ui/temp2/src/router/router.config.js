
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
        path: '/article',
        name: 'ArticleList',
        component: () => import('@/pages/article/list'),
      },
      {
        path: '/category',
        name: 'Category',
        component: () => import('@/pages/article/category'),
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
        path: '/components/navigation',
        name: 'navigation',
        component: () => import('@/pages/components/navigation'),
      },
      {
        path: '/components/layout',
        name: 'layout',
        component: () => import('@/pages/components/layout'),
      }
    ]
  }
]