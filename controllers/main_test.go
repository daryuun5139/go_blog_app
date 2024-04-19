package controllers_test

import (
	"testing"

	"github.com/daryuun5139/go_blog_app/controllers"
	"github.com/daryuun5139/go_blog_app/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)

	m.Run()
}