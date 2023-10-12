import { createStore } from 'vuex'

export default createStore({
  state: {
    SERVERADDRESS: "http://localhost:8000",

  },
  getters: {
    getServerAddress(state) {
      return state.SERVERADDRESS
    },
  },
  mutations: {

  },
  actions: {

  }
})







