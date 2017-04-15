import Vue from 'vue/dist/vue'
import VueAjax from 'vue-resource/dist/vue-resource'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-default/index.css'
import login from './admin/login/login.vue'

Vue.use(ElementUI)
Vue.use(VueAjax)
new Vue({
  el: '#app',
  components: { 
    'hm-login':login,
    // 'hm-echarts':echarts 
  },
  // render: h => h(login)
})
