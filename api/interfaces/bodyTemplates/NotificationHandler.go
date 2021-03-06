package bodyTemplates

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
	"github.com/Atelier-Developers/alethia/domain/entity/Skill"
	"time"
)

type NotificationResponseBody struct {
	Id         uint64    `json:"id"`
	ReceiverId uint64    `json:"receiver_id"`
	Created    time.Time `json:"created"`
	IsSeen     bool      `json:"is_seen"`
}

type NotificationBirthdayResponseBody struct {
	NotificationResponseBody
	Creator UserGetResponseBody `json:"creator"`
}

type NotificationChangeWorkResponseBody struct {
	NotificationResponseBody
	BackgroundHistory entity.BackgroundHistory `json:"background_history"`
	Creator           UserGetResponseBody      `json:"creator"`
}

type NotificationCommentResponseBody struct {
	NotificationResponseBody
	Comment Comment.Comment     `json:"comment"`
	Post    Post.Post           `json:"post"`
	Creator UserGetResponseBody `json:"creator"`
}

type NotificationEndorseResponseBody struct {
	NotificationResponseBody
	UserSkill Skill.Skill         `json:"user_skill"`
	Creator   UserGetResponseBody `json:"creator"`
}

type NotificationLikeComment struct {
	NotificationResponseBody
	Comment Comment.Comment     `json:"comment"`
	Creator UserGetResponseBody `json:"creator"`
}

type NotificationLikePost struct {
	NotificationResponseBody
	Post    Post.Post           `json:"post"`
	Creator UserGetResponseBody `json:"creator"`
}

type NotificationReply struct {
	NotificationResponseBody
	Comment        Comment.Comment     `json:"comment"`
	RepliedComment Comment.Comment     `json:"replied_comment"`
	Creator        UserGetResponseBody `json:"creator"`
}

type NotificationViewProfile struct {
	NotificationResponseBody
	Creator UserGetResponseBody `json:"creator"`
}
