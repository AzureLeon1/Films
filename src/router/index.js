import Vue from 'vue'
import Router from 'vue-router'
import FilmList from '@/components/FilmList'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'FilmList',
      component: FilmList
    }
  ]
})
