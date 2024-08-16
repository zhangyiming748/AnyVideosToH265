package log

import (
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
	"strings"
)

func SetLog() {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{"AnyVideosToH265.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}

	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
