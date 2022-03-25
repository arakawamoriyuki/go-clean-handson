package routers

import (
	"fmt"
	"main/models"
	"main/pkg/setting"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: テストケース毎のテストデータ挿入や初期化

func TestRoutersSetupRouter(t *testing.T) {
	setting.Setup("../conf/test.ini")
	models.Setup()
	router := SetupRouter()

	tests := []struct {
		Title          string
		RequestMethod  string
		RequestPath    string
		ResponceStatus int
		ResponceBody   string
	}{
		{
			Title:          "データがない場合空配列を返す",
			RequestMethod:  "GET",
			RequestPath:    "/api/todos",
			ResponceStatus: 200,
			ResponceBody:   "[]",
		},
		{
			Title:          "指定IDのデータがない場合status:400でエラーメッセージを返す",
			RequestMethod:  "GET",
			RequestPath:    "/api/todos/1",
			ResponceStatus: 400,
			ResponceBody:   "{\"message\":\"record not found\"}",
		},
	}

	for _, td := range tests {
		td := td
		title := fmt.Sprintf(
			"SetupRouter %s %s %s",
			td.RequestMethod,
			td.RequestPath,
			td.Title,
		)
		t.Run(title, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(td.RequestMethod, td.RequestPath, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, td.ResponceStatus)
			assert.Equal(t, w.Body.String(), td.ResponceBody)
		})
	}
}
