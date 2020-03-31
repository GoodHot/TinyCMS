
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
        path: '/article/publish',
        name: 'ArticlePublish',
        component: () => import('@/pages/article/publish'),
      },
      {
        path: '/category/:page',
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
        path: '/elements/dropdown',
        name: 'dropdown',
        component: () => import('@/pages/elements/dropdown'),
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
      },
      {
        path: '/components/nestable',
        name: 'nestable',
        component: () => import('@/pages/components/nestable'),
      }
    ]
  },
  {
    path: '/auth',
    name: 'auth',
    component: () => import('@/layouts/AuthLayout'),
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import('@/pages/auth/login'),
      }
    ]
  }
]