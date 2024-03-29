
export const routerMap = [
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
  },
  {
    path: '/',
    name: 'index',
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
        path: '/category',
        name: 'category',
        component: () => import('@/pages/content/category'),
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
        path: '/role',
        name: 'role',
        component: () => import('@/pages/user/role'),
      }
    ]
  }
]