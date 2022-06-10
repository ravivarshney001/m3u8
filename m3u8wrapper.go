package m3u8

import (
	"errors"
	"github.com/ravivarshney001/m3u8/dl"
)

const chanSize = 25

type DownloadMP4Response struct {
	Stats dl.Stats
}

func DownloadMP4(url, output, fileName string) (DownloadMP4Response, error) {
	var resp DownloadMP4Response
	if url == "" {
		return resp, errors.New("url is required")
	}
	if output == "" {
		return resp, errors.New("output is required")
	}
	downloader, err := dl.NewTask(output, url, fileName)
	if err != nil {
		return resp, err
	}
	if err := downloader.Start(chanSize); err != nil {
		return resp, err
	}
	resp.Stats = downloader.Stats
	return resp, nil
}
