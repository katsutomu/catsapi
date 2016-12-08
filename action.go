package main

// MIKE 三毛猫
var MIKE = 1

// SearchCatAction 猫画像検索
// @Title 猫画像検索API
// @Description 種別IDを入力すると種別名を返却します
// @Param type query int true "猫ID"
// @Success 200 {object} string "種別名"
// @Resource /cat
// @Router /v1/cat/images [get]
func SearchCatAction(t int) string {
	if t == MIKE {
		return "三毛猫"
	}
	return "いないよ！"
}
