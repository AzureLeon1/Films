import Vue from 'vue'
import Router from 'vue-router'
import FilmList from '@/components/FilmList'
import FilmDetail from '@/components/FilmDetail'
import Statistics from '@/components/Statistics'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: '/films/',
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
    {
      path: '/statistics',
      name: 'statistics',
      component: Statistics
    }
  ]
})
