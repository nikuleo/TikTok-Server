package service

import (
	"TikTokServer/cache"
	message "TikTokServer/idl/gen"
	"TikTokServer/model"
	"TikTokServer/pkg/errorcode"
)

func CommentAction(authID, videoID, actionType int64, commentText string, commentID int64) (*message.DouyinCommentActionResponse, error) {
	var err error
	if actionType == 1 {
		err = model.CreateComment(authID, videoID, commentText)
	}
	if actionType == 2 {
		err = model.DeleteComment(commentID, videoID)
	}

	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	// TODO:  通知消息队列删除缓存
	return &message.DouyinCommentActionResponse{}, nil
}

func CommentList(userID, videoID int64) (*message.DouyinCommentListResponse, error) {
	// 从缓存中获取评论列表
	commentsMessage, err := cache.GetVideoCommentFromCache(videoID)
	if err != nil {
		errCode := errorcode.ErrHttpCache
		errCode.SetError(err)
		return nil, errCode
	}
	resp := &message.DouyinCommentListResponse{}
	if commentsMessage != nil {
		resp.CommentList = commentsMessage
		return resp, nil
	}
	// 缓存未命中，从数据库中获取后存入缓存
	comments, err := model.GetCommentList(videoID)
	if err != nil {
		errCode := errorcode.ErrHttpDatabase
		errCode.SetError(err)
		return nil, errCode
	}
	commentsMessage = PackCommentList(comments, userID)
	resp.CommentList = commentsMessage
	err = cache.SetVideoCommentToCache(videoID, commentsMessage)
	if err != nil {
		errCode := errorcode.ErrHttpCache
		errCode.SetError(err)
		return nil, errCode
	}
	return resp, nil
}

func PackCommentList(comments []*model.Comment, userID int64) []*message.Comment {
	commentList := make([]*message.Comment, len(comments))

	for i, c := range comments {
		comment := &message.Comment{
			Id:         int64(c.ID),
			Content:    c.Content,
			CreateDate: c.CreatedAt.Format("01-02"),
			User:       PackUserInfo(&c.User),
		}
		commentList[i] = comment
	}
	return commentList
}
