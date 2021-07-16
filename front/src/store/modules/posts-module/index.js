import {API} from '@/api/api'
import axios from "axios";

const state = {
    posts: [],
};

const mutations = {
    setPosts(state, item) {
        state.posts = item;
    }
};

const actions = {
    getPosts(context,) {
        let posts = [
            {
                user: {
                    username: "Ali",
                },
                text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیلی شاخم خلاصه.',
                date: '11/1/1 12:12',
                // media: ''
            },
            {
                user: {
                    username: "Ali REZA",
                },
                text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیفتاب به صورت افقی تا ما تحت ادم میرفت. من خیلی شاخم خلاصه.',
                date: '11/1/1 12:12',
                // media: ''
            },
            {
                user: {
                    username: "ASQAR",
                },
                text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت. افتاب به صورت افقی تا ما تحت ادم میرفت. من خیلی شاخم خلاصه.',
                date: '11/1/1 12:12',
                // media: ''
            },
            {
                user: {
                    username: "gorbeye tak shakh",
                },
                text: 'امروز خیلی هوا کیری تحت ادم میرفتدم میرفت. من خیلی شامروز خیلی هوا کیری تحت ادم میرفتدم میرفت. من خیلی شامروز خیلی هوا کیری تحت ادم میرفتدم میرفت. من خیلی شامروز خیلی هوا کیری تحت ادم میرفتدم میرفت. من خیلی شاخم خلاصه.',
                date: '11/1/1 12:12',
                // media: ''
            },
            {
                user: {
                    username: "rangin kamune siah",
                },
                text: 'امروز خیلی هوا کیری بود. افتاب به صورت افقی تا ما تحت ادم میرفت تحت ادم میرفت تحت ادم میرفت تحت ادم میرفت. من خیلی شاخم خلاصه.',
                date: '11/1/1 12:12',
                // media: ''
            },
        ];
        context.commit('setPosts', posts);
    },
    async addPost(context, payload) {
        console.log("ADD POST");
        try {
            let res = await axios.post(API.addPost, payload);
            console.log(res)
        }
        catch(e){
            console.log(e);
        }
    }
};

const getters = {
    posts: (state) => state.posts,
};


export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters
}