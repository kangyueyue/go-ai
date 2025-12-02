package aihelper

import (
	"context"
	"fmt"
	"sync"
)

type ModelName string

const (
	OpenAI ModelName = "openai"
	Ollama ModelName = "ollama"
)

// ModelCreator 模型创造器
type ModelCreator func(ctx context.Context, config map[ModelName]interface{}) (AIModel, error)

// AIModelFactory 模型工厂
type AIModelFactory struct {
	ModelCreatorMap map[ModelName]ModelCreator
}

var (
	globalFactory *AIModelFactory
	factoryOnce   sync.Once
)

// GetGlobalFactory 创建模型工厂
func GetGlobalFactory() *AIModelFactory {
	factoryOnce.Do(func() {
		if globalFactory == nil {
			globalFactory = &AIModelFactory{
				ModelCreatorMap: make(map[ModelName]ModelCreator),
			}
		}
		globalFactory.registerCreators() // 注册
	})
	return globalFactory
}

// registerCreators 单个模型注册（初始化）
func (f *AIModelFactory) registerCreators() {
	f.ModelCreatorMap[OpenAI] = func(ctx context.Context, config map[ModelName]interface{}) (AIModel, error) {
		return NewOpenAIModel(ctx)
	}

	// TODO Ollama
	f.ModelCreatorMap[Ollama] = func(ctx context.Context, config map[ModelName]interface{}) (AIModel, error) {
		return nil, nil
	}
}

// CreateAIModel 创建模型
func (f *AIModelFactory) CreateAIModel(ctx context.Context,
	modelType ModelName, config map[ModelName]interface{},
) (AIModel, error) {
	creator, ok := f.ModelCreatorMap[modelType]
	if !ok {
		return nil, fmt.Errorf("unsupport model type:%s", modelType)
	}
	return creator(ctx, config)
}

// CreateAIHelper 创建ai helper
func (f *AIModelFactory) CreateAIHelper(ctx context.Context,
	modelType ModelName, sessionId string, config map[ModelName]interface{},
) (*AIHelper, error) {
	model, err := f.CreateAIModel(ctx, modelType, config)
	if err != nil {
		return nil, err
	}
	return NewAIHelper(model, sessionId), nil
}

// RegisterModel 注册模型
func (f *AIModelFactory) RegisterModel(modelType ModelName, creator ModelCreator) {
	f.ModelCreatorMap[modelType] = creator
}
