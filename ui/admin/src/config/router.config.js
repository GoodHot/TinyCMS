// eslint-disable-next-line
import { UserLayout, BasicLayout, RouteView, BlankLayout, PageView } from '@/layouts'
import { bxAnaalyse } from '@/core/icons'

export const asyncRouterMap = [

  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: '首页' },
    redirect: '/dashboard/workplace',
    children: [
      {
        path: '/dashboard',
        name: 'dashboard',
        redirect: '/dashboard/workplace',
        component: RouteView,
        meta: { title: 'TinyCMS', keepAlive: true, icon: bxAnaalyse },
        children: [
          {
            path: '/dashboard/workplace',
            name: 'Workplace',
            component: () => import('@/views/dashboard/Workplace'),
            meta: { title: '工作台', keepAlive: true }
          }
        ]
      },
      {
        path: '/config',
        name: 'config',
        component: RouteView,
        meta: { title: '配置', keepAlive: true, icon: bxAnaalyse },
        children: [
          {
            path: '/config/dict',
            name: 'config',
            component: () => import('@/views/config/dict'),
            meta: { title: '字典设置', keepAlive: true }
          }
        ]
      },
      {
        path: '/content',
        name: 'content',
        component: RouteView,
        meta: { title: '内容', keepAlive: true, icon: bxAnaalyse },
        children: [
          {
            path: '/content/channel',
            name: 'content',
            component: () => import('@/views/content/channel'),
            meta: { title: '频道管理', keepAlive: true }
          },
          {
            path: '/config/publish',
            name: 'publish',
            component: () => import('@/views/content/publish'),
            meta: { title: '发布文章', keepAlive: true }
          },
          {
            path: '/config/article',
            name: 'article',
            component: () => import('@/views/content/article'),
            meta: { title: '文章列表', keepAlive: true }
          }
        ]
      }
    ]
  },
  {
    path: '*', redirect: '/404', hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import('@/views/user/Login')
      }
    ]
  },
  {
    path: '/404',
    component: () => import('@/views/exception/404')
  }

]
