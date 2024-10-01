package services

import (
	"challenge/internal/services/strategies"
	"errors"
)

type Strategy interface {
	Process(bucket, key string) error
}

type StrategyContext struct {
	strategy Strategy
}

func NewStrategyContext(fileType string) *StrategyContext {
	var strategy Strategy
	switch fileType {
	case ".json":
		strategy = &strategies.JSONProcessor{}
	case ".csv":
		strategy = &strategies.CSVProcessor{}
	default:
		strategy = nil
	}
	return &StrategyContext{strategy: strategy}
}

func (c *StrategyContext) Execute(bucket, key string) error {
	if c.strategy == nil {
		return errors.New("estrategia no definida para este tipo de archivo")
	}
	return c.strategy.Process(bucket, key)
}
