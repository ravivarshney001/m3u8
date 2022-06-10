package tests

import (
	"fmt"
	"github.com/ravivarshney001/m3u8"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://videos.classplusapp.com/b08bad9ff8d969639b2e43d5769342cc62b510c4345d2f7f153bec53be84fe35/xZTBtDYv/xZTBtDYv.m3u8?auth_key=1662814741-0-0-659e7e78a405ff2014c231467e628d74"
	if resp, err := m3u8.DownloadMP4(url, "tmp", "out.mp4"); err != nil {
		panic(err)
	} else {
		fmt.Println(resp.Stats.StartTime, resp.Stats.EndTime)
		fmt.Println(resp.Stats.SegCount, resp.Stats.FailedCount)
	}
}
