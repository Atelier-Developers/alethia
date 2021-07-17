import {API} from '@/api/api'
import axios from "axios";

const state = {
    conversations: [],
    messages: [],
};

const mutations = {
    setConversatoins(state, item) {
        state.conversations = item;
    },
    setMessages(state, item) {
        state.messages = item;
    },
};

const actions = {
    async getConversations(context,) {
        console.log("GET Conversations");
        try {
            let res = await axios.get(API.conversation);
            console.log(res);
            context.commit("setConversatoins", res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async getMessages(context, payload) {
        console.log("GET Messages");
        try {
            let res = await axios.get(API.conversation + '/' + payload + '/messages');
            console.log(res);
            context.commit("setConversatoins", res.data);
        } catch (e) {
            console.log(e);
        }
    },
    async updateConversation(context, payload) {
        console.log("update Conversations");
        try {
            let res = await axios.put(API.conversation, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async deleteConversation(context, payload) {
        console.log("delete Conversations");
        try {
            let res = await axios.delete(API.conversation, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async createConversation(context, payload) {
        console.log("create Conversations");
        try {
            let res = await axios.post(API.conversation, payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },
    async createMessage(context, payload) {
        console.log("delete Conversations");
        try {
            let res = await axios.post(API.conversation + '/message', payload);
            console.log(res);
        } catch (e) {
            console.log(e);
        }
    },


};

const getters = {
    conversations: (state) => state.conversations,
    messages: (state) => state.messages,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}