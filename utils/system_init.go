package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-chat/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	REDIS *redis.Client
	DB    *gorm.DB
	CONF  *config.AppConf
)

func InitConfig(c *config.Config) *config.AppConf {
	if err := c.Viper().ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app: ", c.Viper().Get("app"))
	fmt.Println("config mysql: ", c.Viper().Get("mysql"))
	cfg := new(config.AppConf)
	if err := c.Viper().Unmarshal(cfg); err != nil {
		panic(err)
	}
	CONF = cfg
	return cfg
}

func InitMySQL(cfg *config.Mysql) {
	// 自定义日志模版
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(getMySQLDns(cfg)), &gorm.Config{Logger: newLogger})
}

func getMySQLDns(mysqlConfig *config.Mysql) string {
	return mysqlConfig.Username + ":" + mysqlConfig.Password + "@tcp(" + mysqlConfig.Url + ")/" + mysqlConfig.Database + "?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
}
func InitRedis(cfg *config.Redis) {
	REDIS = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConn,
	})
	//pong, err := REDIS.Ping().Result()
	//if err != nil {
	//	fmt.Println("init redis error: ", err)
	//} else {
	//	fmt.Println("init redis success: ", pong)
	//}
}

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish: ", msg)
	err = REDIS.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	subMsg := REDIS.Subscribe(ctx, channel)
	fmt.Println("Subscribe: ", ctx)
	msg, err := subMsg.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Subscribe: ", msg.Payload)
	return msg.Payload, err
}

//// NewLogger config logrus middleware
//func NewLogger(logConf *config.LogConf) (*log.LoggerW, error) {
//	logConf.TimeFormat = config.DateTimeFormat
//	logConf.Order = []string{
//		types.LogIpKey, types.LogHttpMethodKey,
//		types.LogHttpStatusKey, types.LogRequestPathKey,
//		types.LogRequestUrlKey, types.LogRequestCostKey,
//		types.LogRequestContentType, types.LogHttpContentLength,
//		types.LogResponseContentType, types.LogHttpResponseLength,
//		types.LogRecoverRequestKey, types.LogRecoverErrorKey,
//		types.LogRecoverStackKey, types.LogRequestIdKey}
//	logger, err := log.NewLogger(logConf)
//	if err != nil {
//		return nil, errors.Wrap(err, "load logger failed")
//	}
//	log.Set(logger.L())
//	return logger, nil
//}
