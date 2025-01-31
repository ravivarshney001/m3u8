package m3u8

import (
	"errors"

	"github.com/ravivarshney001/m3u8/dl"
)

const chanSize = 25

func DownloadMP4(url, output, fileName string) error {
	if url == "" {
		return errors.New("url is required")
	}
	if output == "" {
		return errors.New("output is required")
	}
	downloader, err := dl.NewTask(output, url, fileName)
	if err != nil {
		return err
	}
	if err := downloader.Start(chanSize); err != nil {
		return err
	}
	return nil
}
