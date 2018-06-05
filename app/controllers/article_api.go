package controllers

import (
	"github.com/revel/revel"
)

type ArticleApi struct {
	*revel.Controller
}


func (c ArticleApi) GetArticles() revel.Result {

	response := JsonResponse{}
	response.Response = "all articles"

	return c.RenderJSON(response)
}

func (c ArticleApi) GetArticle() revel.Result {

	response := JsonResponse{}
	response.Response = "single article"

	return c.RenderJSON(response)
}

func (c ArticleApi) PostArticle() revel.Result {

	response := JsonResponse{}
	response.Response = "post article"

	return c.RenderJSON(response)
}

func (c ArticleApi) PutArticle() revel.Result {

	response := JsonResponse{}
	response.Response = "put article"

	return c.RenderJSON(response)
}

func (c ArticleApi) DeleteArticle() revel.Result {

	response := JsonResponse{}
	response.Response = "delete article"

	return c.RenderJSON(response)
}
