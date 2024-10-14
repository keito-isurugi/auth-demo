package view

import (
    "html/template"
	"net/http"
)

func ViewPasswordResetHandler(w http.ResponseWriter, r *http.Request) {
    // クエリパラメータを取得
	query := r.URL.Query()
	token := query.Get("token")

    tpl := template.HTML(`
        <h1>パスワードリセット画面</h1>
        <form method="post" action="http://localhost:8080/password_reset">
            <label for="user_id">ユーザーID: 1</label>
            <input type="hidden" name="user_id" value="1">
            <br>
            <label for="old_password">現在のパスーワード</label>
            <input type="text" name="old_password">
            <br>
            <label for="new_password">新しいパスーワード</label>
            <input type="text" name="new_password">
            <br>
            <label for="token">トークン</label>
            <input type="text" name="token" value="{{.Token}}">
            <br>
            <input type="submit" value="パスワードをリセットする">
        </form>
    `)

    	// データをテンプレートに渡す
	data := struct {
		Token string
	}{
		Token:  token,
	}

    t := template.Must(template.New("a").Parse(string(tpl)))
	t.Execute(w, data)
}
