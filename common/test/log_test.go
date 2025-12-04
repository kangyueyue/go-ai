package test

import (
	"fmt"
	"github.com/kangyueyue/go-ai/common/logger"
	"testing"
)

func TestGetProjectPath(t *testing.T) {
	// 获取项目路径
	projectPath := logger.GetProjectPath()
	fmt.Println(projectPath)
}
