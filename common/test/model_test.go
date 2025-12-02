package test

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kangyueyue/go-ai/common/aihelper"
	"path"
	"runtime"
	"testing"
)

func init() {
	_, current, _, _ := runtime.Caller(0)
	envPath := path.Dir(current+"/../../../") + "/.env"
	if err := godotenv.Load(envPath); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func TestNewOpenAIModel(t *testing.T) {
	model, err := aihelper.NewOpenAIModel(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(model.GetModelType())
}
