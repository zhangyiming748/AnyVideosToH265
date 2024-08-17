package main

import (
	"AnyVideosToH265/file"
	"AnyVideosToH265/hevc"
	"AnyVideosToH265/log"
	"AnyVideosToH265/util"
	"fmt"
	"time"
)

func init() {
	log.SetLog()
}
func main() {
	root := "/data"
	files := file.GetVideoFile(file.GetAllFile(root))
	util.WriteByLine("/data/notH265.txt", files)
	for i := 9; i > 0; i-- {
		fmt.Printf("\r等待%d秒后开始转码", i)
		time.Sleep(1 * time.Second)
	}
	videos := util.ReadByLine("/data/notH265.txt")
	for _, video := range videos {
		fmt.Println(video)
		hevc.ProcessVideo2H265(video)
	}
}
