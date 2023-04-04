import { createRouter, createWebHashHistory } from 'vue-router';

const routes = [
  {
    name: 'notFound',
    path: '/:path(.*)+',
    redirect: {
      name: 'home',
    },
  },
  {
    name: 'user',
    path: '/user',
    component: () => import('./view/user'),
    meta: {
      title: '会员中心',
    },
  },
  {
    name: 'name',
    path: '/name',
    component: () => import('./view/name'),
    meta: {
      title: '购物车',
    },
  },
  {
    name: 'goods',
    path: '/goods',
    component: () => import('./view/goods'),
    meta: {
      title: '商品详情',
    },
  },
  {
    name: 'home',
    path: '/home',
    component: () => import('./view/home'),
    meta: {
      title: '主页',
    },
  },
  {
    name: 'chattgpt',
    path: '/chattgpt',
    component: () => import('./view/chattgpt'),
    meta: {
      title: 'chattgpt',
    },
  },
  {
    name: 'creativity',
    path: '/creativity',
    component: () => import('./view/creativity'),
    meta: {
      title: '创意写作',
    },
  },
];

const router = createRouter({
  routes,
  history: createWebHashHistory(),
});

router.beforeEach((to, from, next) => {
  const title = to.meta && to.meta.title;
  if (title) {
    document.title = title;
  }
  next();
});

export { router };
