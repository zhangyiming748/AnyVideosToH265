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
	root := "F:\\分割完成\\淫语"
	files, _ := file.GetAllVideoFilesInDirNotHEVC(root)
	util.WriteByLine("notH265.txt", files)
	for i := 9; i > 0; i-- {
		fmt.Printf("\r等待%d秒后开始转码", i)
		time.Sleep(1 * time.Second)
	}
	videos := util.ReadByLine("notH265.txt")
	for _, video := range videos {
		fmt.Println(video)
		hevc.ProcessVideo2H265(video)
	}
}
