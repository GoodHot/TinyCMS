
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
      }
    ]
  }
]