package view

import (
	"html/template"
	"net/http"
)

func JWTLoginPage(w http.ResponseWriter, r *http.Request) {
    token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjkyNjIyMTAsInN1YiI6IjEyMzQifQ.XRANbDR5Ow12vKXDC62YPSb9BC1cRd7SBzHWD5ciuAk"

	tpl, err := template.New("protected").Parse(`
        <h1>ログインページ(JWT認証)</h1>
        <form method="post" action="http://localhost:8080/jwt_login">
            <label for="id">ID</label>
            <input type="text" name="id">
            <br>
            <label for="password">Password</label>
            <input type="test" name="password">
            <br>
            <input type="submit" label="submit">
        </form>

        <h1>Protected Page</h1>
		<p>JWT has been generated. Click the button below to send it in a request.</p>
		<a id="protected-link" href="http://localhost:8080/view/jwt_auth_page">Go to Protected Page</a>

        <script>
            // トークンをJavaScriptに渡す
            const token = "{{ .Token }}";
            document.cookie = "token=" + token + "; path=/";
        </script>
    `)
    if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// テンプレートにトークンを埋め込む
	err = tpl.Execute(w, struct{ Token string }{Token: token})
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
