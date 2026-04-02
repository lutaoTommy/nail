package logger

import log "github.com/jeanphorn/log4go"

func InitLogger() {
    // 加载 JSON 配置文件
    log.LoadConfiguration("logger/logger.json")
}