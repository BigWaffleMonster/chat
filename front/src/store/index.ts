import { createStore } from 'vuex'
import { wsModule } from './wsModule'

export default createStore({
  state: () => ({
    nickname: ''
  }),

  getters: {
    getNickname(state) {
      return state.nickname
    }
  },

  mutations: {
    setNickname(state, nick) {
      state.nickname = nick
    }
  },

  actions: {

  },

  modules: {
    ws: wsModule
  }
})
