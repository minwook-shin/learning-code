import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    text: "hello, world!",
    num: 10
  },
  getters: {
    text(state) {
      return state.text;
    },
    num(state) {
      return state.num
    }
  },
  mutations: {
    addNum(state, payload) {
      state.num += payload.value;

    }
  },
  actions: {
    addNum(context, payload) {
      context.commit('addNum', {
        value: payload.num
      });
    }
  }
});