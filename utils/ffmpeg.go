package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
)

// GetSnapshot 生成视频缩略图并保存（作为封面）
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg_go.Input(videoPath).
		Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	err = imaging.Save(img, snapshotPath+".jpeg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".jpeg"
	return
}
