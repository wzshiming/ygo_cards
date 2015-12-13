package ygo_cards

import ygo "github.com/wzshiming/ygo_core"

var CardBag_test = ygo.NewCardVersion()

func init() {
	// 第1期 341张
	d1_1(CardBag_test)
	d1_2(CardBag_test)
	// 第2期 635张
	d2_1(CardBag_test)
	d2_2(CardBag_test)
	d2_3(CardBag_test)

}
