const API_PERFIX = 'http://localhost:7000' + '/';
console.log(process)

export const API = {
    signup: API_PERFIX + 'register',
    login: API_PERFIX + 'login',
    getSkills: API_PERFIX + 'users/skill',
    getLanguages: API_PERFIX + 'users/language',
    getBackground: API_PERFIX + 'users/background',
    getUser: API_PERFIX + 'users',
    userBackground: API_PERFIX + 'users/background',
    userLanguage: API_PERFIX + 'users/language',
    userSkill: API_PERFIX + 'users/skill',
    getAllLanguages: API_PERFIX + 'language',
    getAllSkills: API_PERFIX + 'skill',
    addPost: API_PERFIX + 'post',
    getPosts: API_PERFIX + 'users/post',
    editUser: API_PERFIX + 'users',
    getInvites: API_PERFIX + 'users/invite',
    getFriends: API_PERFIX + 'users/friend',
    getPost: API_PERFIX + 'post',
    likePost: API_PERFIX + 'post/like',
    commentPost: API_PERFIX + 'post/comment',
    replyComment: API_PERFIX + 'post/comment/reply',
    likeComment: API_PERFIX + 'post/comment/like',
    repost: API_PERFIX + 'post',
    conversation: API_PERFIX + 'conversation',
    notif: API_PERFIX + 'users/notification',
    mutual: API_PERFIX + 'users/mutualConnection'

}


// TODO
// LOGOUT
// add invite
// delete invite
// get user invites
// add friend
// delete freind

