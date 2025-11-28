package test

import (
	"testing"

	"github.com/kangyueyue/go-ai/config"
)

func TestInitConfig(t *testing.T){
	config := config.GetConfig()
	if config == nil{
		t.Error("config is nil")
		return 
	}
	t.Log(config.Email.AuthCode)
}