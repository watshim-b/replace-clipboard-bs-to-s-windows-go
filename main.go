package main

import (
	"strings"

	"github.com/atotto/clipboard"
)

func main() {
	// クリップボードにコピーされている文字列を読み込む
	cbStr, _ := clipboard.ReadAll()

	cbStr = strings.ReplaceAll(cbStr, "\\", "/")

	// C:が不要な場合は、この行を有効化すればよい
	// cbStrArray := strings.Split(cbStr, ":")
	// cbStr = strings.ReplaceAll(strings.Join(cbStrArray[1:], ""), "\\", "/")

	// 文字列をクリップボードに書き込む
	clipboard.WriteAll(cbStr)
}
