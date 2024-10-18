package view

import (
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

func JWTAuthPage(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl := template.HTML(`
			<h1>ログインが必要なページ アクセス成功！(JWT認証)</h1>
		`)
		t := template.Must(template.New("a").Parse(`<div>{{ . }}</div>`))
		t.Execute(w, tpl)
	}
}
