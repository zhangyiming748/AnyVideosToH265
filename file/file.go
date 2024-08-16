package file

import (
	"fmt"
	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastMediaInfo"
	"os"
	"os/exec"
	"strings"
)

func GetAllFile(root string) *[]string {
	//root := "/mnt/e/video"
	cmd := exec.Command("find", root, "-type", "f")
	output, err := cmd.CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	} else {
		//fmt.Println(string(output))
	}
	paths := strings.Split(string(output), "\n")
	return &paths
}

func GetVideoFile(s *[]string) {
	for _, path := range *s {
		if path != "" {
			if IsVideo(path) {
				//fmt.Printf("单独文件%v\n", path)
				p := FastMediaInfo.GetStandMediaInfo(path)
				//println(p.Video.CodecID) //hvc1
				//println(p.Video.Format)  //HEVC
				if p.Video.Format == "HEVC" {
					//fmt.Printf("满足条件的视频: %v\n", path)
					satisfyingVideos = append(satisfyingVideos, path)
				} else {
					//fmt.Printf("不满足条件的视频: %v\n", path)
					nonSatisfyingVideos = append(nonSatisfyingVideos, path)
				}
			}
		}
	}
}
func IsVideo(fp string) bool {
	file, _ := os.Open(fp)
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	if filetype.IsVideo(head) {
		fmt.Printf("File: %v is a video\n", fp)
		return true
	}
	return false
}

/*
二次判断视频编码
h265返回真
*/
func ffprob(s string) bool {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=codec_name", "-of", "default=noprint_wrappers=1:nokey=1", s)
	if output, _ := cmd.CombinedOutput(); string(output) == "hevc" {
		return true
	} else {
		return false
	}
}
