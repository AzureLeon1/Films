import Vue from 'vue'
import Router from 'vue-router'
import FilmList from '@/components/FilmList'
import FilmDetail from '@/components/FilmDetail'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'list',
      component: FilmList
    },
    {
      path: '/details',
      name: 'detail',
      component: FilmDetail
    },
  ]
})
