package view

import (
    "html/template"
	"net/http"
)

func ViewRequestPasswordResetHandler(w http.ResponseWriter, r *http.Request) {
    tpl := template.HTML(`
        <form method="post" action="http://localhost:8080/request_password_reset">
            <label for="id">ユーザーID: 1</label>
            <input type="hidden" name="user_id" value="1">
            <input type="submit" value="パスワードをリセットのリクエスト">
        </form>
    `)
    t := template.Must(template.New("a").Parse(`<div>{{ . }}</div>`))
    t.Execute(w, tpl)
}
