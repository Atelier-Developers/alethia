const API_PERFIX = 'http://localhost:7000' + '/';
console.log(process)

export const API = {
	signup: API_PERFIX + 'register',
	login: API_PERFIX + 'login',
	getSkills: API_PERFIX + 'users/skill',
	getLanguages: API_PERFIX + 'users/language',
	getBackground: API_PERFIX + 'users/background',
	getUser: API_PERFIX + 'users',
	addBackground: API_PERFIX + 'users/background',
	addLanguage: API_PERFIX + 'users/language',
	addSkill: API_PERFIX + 'users/skill',
	getAllLanguages: API_PERFIX + 'language',
	getAllSkills: API_PERFIX + 'skill',
	addPost: API_PERFIX + 'post',
	editUser: API_PERFIX + 'users',
}


// TODO
// LOGOUT
// add invite
// delete invite
// get user invites
// add friend
// delete freind

