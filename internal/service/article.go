package service

import "gin-blog/internal/model"
import "gin-blog/pkg/app"

type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof= 1 2"`
}

type ArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type ArticleListRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof= 1 2"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=2,max=10"`
	Desc          string `form:"desc" binding:"max=20"`
	Content       string `form:"content" binding:"min=1,max=1000"`
	CoverImageUrl string `form:"cover_image_url" binding:"min=1,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=1 2"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=2,max=10"`
	Desc          string `form:"desc" binding:"max=20"`
	Content       string `form:"content" binding:"min=1,max=1000"`
	CoverImageUrl string `form:"cover_image_url" binding:"min=1,max=100"`
	State         uint8  `form:"state" binding:"required,oneof=1 2"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.State)
}

func (svc *Service) GetArticle(param *ArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	return svc.dao.CreateArticle(param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.CreatedBy)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	return svc.dao.UpdateArticle(param.ID, param.Title, param.Desc, param.Content, param.CoverImageUrl, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
