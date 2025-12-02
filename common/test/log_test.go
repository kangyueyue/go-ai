package test

import (
	"fmt"
	"github.com/kangyueyue/go-ai/common/log"
	"testing"
)

func TestGetProjectPath(t *testing.T) {
	// 获取项目路径
	projectPath := log.GetProjectPath()
	fmt.Println(projectPath)
}
