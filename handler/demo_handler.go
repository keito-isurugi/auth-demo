package handler

import (
	"fmt"
	"net/http"
)

func RSADemo(w http.ResponseWriter, r *http.Request) {
	// 素数の選択
	p := 3
	q := 11

	// nとφ(n)の計算
	n := p * q               // n = 33
	phi := (p - 1) * (q - 1) // φ(n) = 20

	// 公開鍵の指数 e の選択
	e := 3

	// 秘密鍵の指数 d の計算 (eのモジュラ逆元)
	d := modInverse(e, phi)
	if d == -1 {
		http.Error(w, "モジュラ逆元が見つかりませんでした", http.StatusInternalServerError)
		return
	}

	// テキストで出力
	output := fmt.Sprintf(
		"公開鍵 (e, n): (%d, %d)\n"+
			"秘密鍵 (d, n): (%d, %d)\n"+
			"元のメッセージ: %d\n",
		e, n, d, n, 5)

	// メッセージ m の設定 (例: m = 5)
	m := 5

	// 公開鍵で暗号化: c = m^e mod n
	c := modExp(m, e, n)
	output += fmt.Sprintf("暗号化されたメッセージ: %d\n", c)

	// 秘密鍵で復号化: m = c^d mod n
	dec := modExp(c, d, n)
	output += fmt.Sprintf("復号化されたメッセージ: %d\n", dec)

	// 結果が元のメッセージと一致するか確認
	if dec == m {
		output += "復号化に成功しました!\n"
	} else {
		output += "復号化に失敗しました...\n"
	}

	// レスポンスに出力
	w.Write([]byte(output))
}

// モジュラ逆元を計算する関数
func modInverse(e, phi int) int {
	for d := 1; d < phi; d++ {
		if (e*d)%phi == 1 {
			return d
		}
	}
	return -1
}

// 繰り返し二乗法を用いたモジュラ演算の累乗計算
func modExp(base, exp, mod int) int {
	result := 1
	for exp > 0 {
		if exp%2 == 1 {
			result = (result * base) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}
	return result
}
