package hevc

import (
	"AnyVideosToH265/util"
	"github.com/zhangyiming748/FastMediaInfo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func ProcessVideo2H265(fp string) {
	mi := FastMediaInfo.GetStandMediaInfo(fp)
	FrameCount := mi.Video.FrameCount
	if mi.Video.CodecID == "hvc1" || mi.Video.CodecID == "vp09" {
		log.Println("跳过已经转码的视频")
		return
	}
	mp4 := strings.Replace(fp, filepath.Ext(fp), "hevc.mp4", 1)
	cmd := exec.Command("ffmpeg", "-i", fp, "-c:v", "libx265", "-tag:v", "hvc1",
		"-ac", "1", "-map_chapters", "-1", mp4)
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm64" {
		cmd = exec.Command("ffmpeg", "-i", fp, "-threads", "1", "-c:v", "libx265", "-tag:v", "hvc1", "-ac", "1", "-map_chapters", "-1", "-threads", "1", mp4)
	}
	log.Printf("生成的命令:%v\n", cmd.String())
	if err := util.ExecCommand(cmd, FrameCount); err != nil {
		return
	}
	log.Println("视频编码运行完成")
	if err := os.Remove(fp); err != nil {
		log.Printf("删除失败:%v\n", fp)
	} else {
		log.Printf("删除成功:%v\n", fp)
	}
}
