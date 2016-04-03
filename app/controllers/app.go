package controllers

import (
	
	//"fmt"
	"github.com/revel/revel"
	"JsonAPI/app/model"
	"net/http"
	"github.com/shwoodard/jsonapi"
	"time"


	)

type MyHtml []interface{}


type App struct {
	*revel.Controller
}


func testBlogForCreate(i int) *model.Blog {
	return &model.Blog{
		Id:        1 * i,
		Title:     "Title 1",
		CreatedAt: time.Now(),
		Posts: []*model.Post{
			&model.Post{
				Id:    1 * i,
				Title: "Foo",
				Body:  "Bar",
				Comments: []*model.Comment{
					&model.Comment{
						Id:   1 * i,
						Body: "foo",
					},
					&model.Comment{
						Id:   2 * i,
						Body: "bar",
					},
				},
			},
			&model.Post{
				Id:    2 * i,
				Title: "Fuubar",
				Body:  "Bas",
				Comments: []*model.Comment{
					&model.Comment{
						Id:   1 * i,
						Body: "foo",
					},
					&model.Comment{
						Id:   3 * i,
						Body: "bas",
					},
				},
			},
		},
		CurrentPost: &model.Post{
			Id:    1 * i,
			Title: "Foo",
			Body:  "Bar",
			Comments: []*model.Comment{
				&model.Comment{
					Id:   1 * i,
					Body: "foo",
				},
				&model.Comment{
					Id:   2 * i,
					Body: "bar",
				},
			},
		},
	}
}

func (r MyHtml) Apply(req *revel.Request, resp *revel.Response) {

	//jsonapiRuntime := jsonapi.NewRuntime().Instrument("blogs.list")
   
	resp.WriteHeader(http.StatusOK, "text/html")

	 if err := jsonapi.MarshalManyPayload(resp.Out, r); err != nil {
		http.Error(resp.Out, err.Error(), 500)
	}

	//resp.Out.Write([]byte(json))
	
}

func (c App) Index() revel.Result {
   
   
   blogs := make([]interface{},0,10)

   for i := 0; i < 10; i += 1 {
		blogs = append(blogs, testBlogForCreate(i))
	}

  
	return MyHtml(blogs)
}
