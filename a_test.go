package ygo_cards

import (
	"testing"
)

func Test_a(t *testing.T) {
	t.Logf("%v/%v",len(CardBag_test.ListIsValid()),len(CardBag_test.ListIsAll()))
}
