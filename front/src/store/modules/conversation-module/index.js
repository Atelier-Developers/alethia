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
        let res = await axios.get(API.conversation);
        context.commit("setConversatoins", res.data);
    },
    async getMessages(context, payload) {
        let res = await axios.get(API.conversation + '/message/' + payload);
        context.commit("setMessages", res.data);
    },
    async updateConversation(context, payload) {
        await axios.put(API.conversation, payload);
    },
    async deleteConversation(context, payload) {
        await axios.delete(API.conversation, {
            data: payload
        });
    },
    async createConversation(context, payload) {
        await axios.post(API.conversation, payload);
    },
    async createMessage(context, payload) {
        await axios.post(API.conversation + '/message', payload);
    },


};

const getters = {
    conversations: (state) => state.conversations,
    messages: (state) => {
        if (state.messages === null) {
            return state.messages
        }
        state.messages.sort(function (a, b) {
            return (a.created > b.created) ? 1 : ((a.created < b.created) ? -1 : 0);
        });
        return state.messages
    },
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}