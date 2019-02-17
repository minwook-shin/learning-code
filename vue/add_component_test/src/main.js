import Vue from 'vue'
import App from './App.vue'
import App2 from './App2.vue'


import NewGlobalWorld from './components/NewGlobalWorld.vue';

Vue.config.productionTip = false

Vue.component('NewGlobalWorld', NewGlobalWorld);

new Vue({
  render: h => h(App),
}).$mount('#app')

new Vue({
  render: h => h(App2),
}).$mount('#app2')
