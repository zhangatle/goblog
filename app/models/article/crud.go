package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/pagination"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"net/http"
)

// Get 通过 ID 获取文章
func Get(idStr string) (*Article, error) {
	var article *Article
	id := types.StringToUint64(idStr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// GetAll 获取全部文章
func GetAll(r *http.Request, perPage int) (*[]Article, pagination.ViewData, error) {

	db := model.DB.Model(Article{}).Order("created_at desc")
	_pager := pagination.New(r, db, route.Name2URL("home"), perPage)
	_viewData := _pager.Paging()

	var articles []Article

	_pager.Results(&articles)

	return &articles, _viewData, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (a *Article) Create() (err error) {
	result := model.DB.Create(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Update 更新文章
func (a *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

// Delete 删除文章
func (a *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&a)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

// GetByUserID 获取全部文章
func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("user_id = ?", uid).Preload("User").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
