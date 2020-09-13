import Vue from 'vue';
import VueRouter from 'vue-router';
import Dashboard from './components/Dashboard';
import AddForm from './components/AddForm';
import axios from 'axios';

let base_uri = process.env.VUE_APP_SERVER;
axios.defaults.baseURL = (process.env.VUE_APP_IS_PROD != 'true') ? 'http://localhost:8080/' : base_uri;
Vue.use(VueRouter);

export default new VueRouter({
    mode: 'history',
    routes: [
        { path: '/', component: Dashboard },
        { path: '/new', component: AddForm}
    ]
});