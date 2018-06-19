package api

import (
	"github.com/revel/revel"
	"appName/app/models"
	"github.com/davecgh/go-spew/spew"
)

type ArticleApi struct {
	Api
}


func (c ArticleApi) GetArticles() revel.Result {

	// articleモデルを利用
	articles := []models.Article{}

	// Idが降順になるように取得
	DB.Order("id desc").Find(&articles)

	if err := DB.Find(&articles).Error; err != nil {
		return c.SetMessage(err.Error()).HandleNotFoundError()
	}

	response := JsonResponse{}
	response.Response = articles // 結果を格納してあげる

	return c.RenderJSON(response)
}

func (c ArticleApi) GetArticle() revel.Result {
	// ルーティングで設定したurlに含まれる :id とかの部分はc.Params.Route.Getで取得
	id := c.Params.Route.Get("id")

	article := models.Article{}
	// DB.Firstの第二引数にそのidを渡してあげると第一引数のモデルからidが一致するデータを検索してくれます

	if err := DB.First(&article, id).Error; err != nil {
		return c.HandleNotFoundError()
	}

	response := JsonResponse{}
	response.Response = article

	return c.RenderJSON(response)
}

func (c ArticleApi) PostArticle() revel.Result {

	// articleモデルに値を格納
	//article := &models.Article{
	//	// x-www-form-urlencodeで飛んできたデータはc.Params.Form.Getで受け取れます
	//	Name:   c.Params.Form.Get("name"),
	//	Fetish: c.Params.Form.Get("fetish"),
	//}

	article := &models.Article{}
	c.Params.BindJSON(&article)

	spew.Dump(article)
	// DBで保存
	DB.Create(article)
	response := JsonResponse{}

	// この時点でarticleにはidが振られているのでそのまま返してあげます
	response.Response = article
	return c.RenderJSON(response)
}

func (c ArticleApi) PutArticle() revel.Result {
id := c.Params.Route.Get("id")
	article := &models.Article{}

	if err := DB.First(&article, id).Error; err != nil {
		return c.UnSpecifiedId().HandleBadRequestError()
	}
	// DB.First()で返ってきたデータの中身を入れ直す
	//article.Name = c.Params.Form.Get(" name")
	//article.Fetish = c.Params.Form.Get("fetish")
	c.Params.BindJSON(&article)


	// 入れ直したものをSave
	DB.Save(&article)

	response := JsonResponse{}
	response.Response = article

	return c.RenderJSON(response)
}

func (c ArticleApi) DeleteArticle() revel.Result {

	id := c.Params.Route.Get("id")

	article := models.Article{}
	// 第二引数入れたidに一致するデータを第一引数のモデルから削除

	if err := DB.First(&article, id).Error; err != nil {
		return c.UnSpecifiedId().HandleBadRequestError()
	}

	if err := DB.Delete(&article).Error; err != nil {
		return c.HandleBadRequestError()
	}

	response := JsonResponse{}
	response.Response = ("deleted. id: " + id)

	return c.RenderJSON(response)
}

func (c ArticleApi) SearchArticle() revel.Result {

	req := models.Article{}
	c.Params.BindJSON(&req)

	// 構造体のインスタンス化
	articles := []models.Article{}
	DB.Debug().Find(&articles,"name = ? or fetish = ?", req.Name, req.Fetish)

	// 指定した条件を元に複数のレコードを引っ張ってくる
	//if DB.Debug().Find(&articles,"name = ? or fetish = ?", req.Name, req.Fetish); len(articles) == 0 {
	//	return c.HandleNotFoundError()
	//}

	response := JsonResponse{}
	response.Response = articles

	return c.RenderJSON(response)
}
