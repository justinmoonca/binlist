package binlist_test

import (
	"github.com/justinmoonca/binlist"
	"testing"
)

func TestGetBinInfo(t *testing.T) {
	info, err := binlist.GetBinInfo("4242424242424242", "http://your-ip-proxy-url")
	if err != nil {
		t.Error("get bin info error: ", err)
		return
	}
	t.Log("get bin info success: ", info)
}
