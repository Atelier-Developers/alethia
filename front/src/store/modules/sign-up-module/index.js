import axios from 'axios'
import {API} from '@/api/api'

const state = {};

const mutations = {};

const actions = {
    async signup(context, payload) {
        console.log(payload)
        console.log(API.signup)
        let res = await axios.post(API.signup, payload);
        return res;
    }
};

const getters = {};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}