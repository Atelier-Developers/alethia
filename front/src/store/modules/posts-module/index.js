import {API} from '@/api/api'
import axios from "axios";

const state = {
    myPosts: [],
    likedPosts: [],
    commentedPosts: [],
};

const mutations = {
    setMyPosts(state, item) {
        state.myPosts = item;
    },
    setLikedPosts(state, item) {
        state.likedPosts = item;
    },
    setCommentedPosts(state, item) {
        state.commentedPosts = item;
    }
};

const actions = {
    async getPosts(context,) {
        console.log("Get POSTs");
        try {
            let res = await axios.get(API.getPosts + '/postedByFriends');
            context.commit("setMyPosts", res.data);

            let res2 = await axios.get(API.getPosts + '/likedByFriends');
            context.commit("setLikedPosts", res2.data);

            let res3 = await axios.get(API.getPosts + '/commentedOnByFriends');
            context.commit("setCommentedPosts", res3.data);

            console.log(res)
            console.log(res2)
            console.log(res3)
        } catch (e) {
            console.log(e);
        }
    },
    async addPost(context, payload) {
        console.log("ADD POST");
        try {
            let res = await axios.post(API.addPost, payload);
            console.log(res)
        } catch (e) {
            console.log(e);
        }
    }
};

const getters = {
    posts: (state) => {
        let all = []
        if (state.commentedPosts != null) {
            state.commentedPosts.forEach((x) => x.type = 'CP')
            all.push(...state.commentedPosts)
        }
        if (state.likedPosts != null) {
            state.likedPosts.forEach((x) => x.type = 'LP')
            all.push(...state.likedPosts)
        }
        if (state.myPosts != null) {
            state.myPosts.forEach((x) => x.type = 'FP')
            all.push(...state.myPosts)
        }
        all.sort(function (a, b) {
            return (a.created < b.created) ? -1 : ((a.created > b.created) ? 1 : 0);
        });
        return all;
    },
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}