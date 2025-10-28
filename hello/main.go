package main

import (
	"fmt"
	"time"

	"example.com/common"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type SampleInput struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func main() {
	// common module のメッセージ
	fmt.Println(common.Message())

	// zapロガーの初期化（軽く使うだけ）
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("3333333", zap.String("id", uuid.NewString()))

	// validator を1回だけ動かす
	v := validator.New()
	input := SampleInput{Name: "rei", Email: "rei@example.com"}
	if err := v.Struct(input); err != nil {
		logger.Error("validation error", zap.Error(err))
	} else {
		logger.Info("validation ok")
	}

	// redisクライアントを「作るだけ」（繋ぎはしない）
	_ = redis.Options{
		Addr: "localhost:6379",
		DB:   7,
	}

	// aws設定を「読もうとするだけ」（本当にアクセスまではしないのでエラーにはなるが型参照で十分）
	_, _ = config.LoadDefaultConfig(ctxNoop{})
}

// ctxNoop is a minimal context-like type to satisfy the compiler for this demo
type ctxNoop struct{}

func (ctxNoop) Deadline() (deadline time.Time, ok bool) { return time.Time{}, false }
func (ctxNoop) Done() <-chan struct{}                   { return nil }
func (ctxNoop) Err() error                              { return nil }
func (ctxNoop) Value(key any) any                       { return nil }
