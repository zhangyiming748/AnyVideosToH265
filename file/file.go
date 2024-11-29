package file

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/h2non/filetype"
	"github.com/zhangyiming748/FastMediaInfo"
)

func GetAllVideoFilesInDirNotHEVC(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否是文件，如果是文件则将其绝对路径添加到files切片中
		if !info.IsDir() {
			if IsVideo(path) {
				mi := FastMediaInfo.GetStandMediaInfo(path)
				if mi.Video.Format == "HEVC" {
					log.Printf("跳过已经是h265的视频:%v\n", path)
				} else {
					files = append(files, path)
				}

			}

		}
		return nil
	})
	return files, err
}

/*
通过文件头判断是否为视频文件
*/
func GetVideoFile(s *[]string) []string {
	var nonSatisfyingVideos []string
	for _, path := range *s {
		if path != "" {
			if IsVideo(path) {
				//println(p.Video.CodecID) //hvc1
				//println(p.Video.Format)  //HEVC
				if GetNotH265ByMediainfo(path) {
					nonSatisfyingVideos = append(nonSatisfyingVideos, path)
					//satisfyingVideos = append(satisfyingVideos, path)
				} else {
					fmt.Printf("跳过HEVC的视频: %v\n", path)
					//fmt.Printf("不是HEVC的视频: %v\n", path)
					nonSatisfyingVideos = append(nonSatisfyingVideos, path)
				}
			}
		}
	}
	return nonSatisfyingVideos
}

/*
通过mediainfo判断是否为非hevc视频
非h265返回真
*/
func GetNotH265ByMediainfo(path string) bool {
	p := FastMediaInfo.GetStandMediaInfo(path)
	if p.Video.Format == "HEVC" {
		//fmt.Printf("HEVC的视频: %v\n", path)
		return false
	} else {
		//fmt.Printf("不是HEVC的视频: %v\n", path)
		return true
	}
}

/*
通过ffprob再次判断是否为非hevc视频
非h265返回真
*/
func GetNotH265ByFfprob(path string) bool {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "v:0", "-show_entries", "stream=codec_name", "-of", "default=noprint_wrappers=1:nokey=1", path)
	// ffprobe -v error -select_streams v:0 -show_entries stream=codec_name -of default=noprint_wrappers=1:nokey=1 "/mnt/e/video/Straplez/vp9/Laptop Light No Camera Plenty of Action 712752695.mp4"
	if output, _ := cmd.CombinedOutput(); string(output) == "hevc" {
		return false
	} else {
		return true
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
