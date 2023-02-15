package initializers

import (
	"context"
	"crypto/tls"
	"database/sql"
	"expresso/configs/configurations"
	"fmt"
	"os"

	"github.com/anurag925/mongoboiler"
	"github.com/go-redis/redis"
	"github.com/google/wire"
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ConfigSet = wire.NewSet(
	configurations.LoadConfigs,
	wire.Bind(new(configurations.Configuration), new(*configurations.Config)),
)

var DefaultLoggerSet = wire.NewSet(
	ConfigSet,
	ProvideDefaultFileLogger,
	wire.Bind(new(Logger), new(*FileLogger)),
)

func ProvideDefaultFileLogger(config configurations.Configuration) (*FileLogger, error) {
	zaplogger, err := defaultZapLogger(config)
	if err != nil {
		return nil, err
	}
	sugared := zaplogger.Sugar()
	return NewFileLogger(zaplogger, sugared), nil
}

func defaultZapLogger(config configurations.Configuration) (*zap.Logger, error) {
	path := config.App().Dir + "/logs/expresso.log"
	c := zap.NewProductionEncoderConfig()
	c.EncodeTime = zapcore.ISO8601TimeEncoder
	c.EncodeLevel = zapcore.LowercaseLevelEncoder
	fileEncoder := zapcore.NewJSONEncoder(c)
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.InfoLevel
	core := zapcore.NewTee(zapcore.NewCore(fileEncoder, writer, defaultLogLevel))
	if config.Env() == configurations.EnvDevelopment {
		c.EncodeLevel = zapcore.CapitalColorLevelEncoder
		defaultLogLevel := zapcore.DebugLevel
		consoleEncoder := zapcore.NewConsoleEncoder(c)
		core = zapcore.NewTee(
			zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		)
	}
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)), err
}

var DefaultHttpServerSet = wire.NewSet(
	ConfigSet,
	ProvideDefaultHttpServer,
	wire.Bind(new(Server), new(*EchoServer)),
)

func ProvideDefaultHttpServer(config configurations.Configuration) (*EchoServer, error) {
	httpServer := NewHttpServer(echo.New())
	server := NewEchoServer(httpServer)
	if config.Env() == configurations.EnvStaging {
		// server.server.HideBanner = true
		// server.server.HidePort = true
		server.server.Logger.SetLevel(log.INFO)
	} else {
		server.server.Debug = true
		server.server.Logger.SetLevel(log.DEBUG)
	}
	// handler.HTTPErrorHandler = dto.CustomHttpErrorHandler
	// handler.Validator = init.app.Validator()
	// server.Start(config.App().Port)
	return server, nil
}

var DefaultMySqlSet = wire.NewSet(
	ConfigSet,
	ProvideDefaultMySqlDB,
	wire.Bind(new(DB), new(*MySQL)),
)

func ProvideDefaultMySqlDB(config configurations.Configuration) (*MySQL, error) {
	db, err := defaultSqlDB(config)
	if err != nil {
		return nil, err
	}
	mysql := NewMySQL(db)
	if err = mysql.Connect(); err != nil {
		return nil, err
	}
	boil.SetDB(mysql.db)
	return mysql, nil
}

func defaultSqlDB(config configurations.Configuration) (*sql.DB, error) {
	// db url
	dbUrl := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		config.App().DBUserName,
		config.App().DBPassword,
		config.App().DBHost,
		config.App().DBPort,
		config.App().DBName,
	)
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

var DefaultRedisSet = wire.NewSet(
	ConfigSet,
	ProvideDefaultRedis,
	wire.Bind(new(Cache), new(*Redis)),
)

func ProvideDefaultRedis(config configurations.Configuration) (*Redis, error) {
	defaultRedisClient, err := defaultRedisConnection(config)
	if err != nil {
		return nil, err
	}
	if err = defaultRedisClient.Ping().Err(); err != nil {
		return nil, err
	}
	redis := NewRedis(NewCacheClient(defaultRedisClient))
	if err = redis.Connect(); err != nil {
		return nil, err
	}
	return redis, nil
}

func defaultRedisConnection(config configurations.Configuration) (*redis.Client, error) {
	options := &redis.Options{
		Addr:       fmt.Sprintf("%s:%d", config.App().RedisHost, config.App().RedisPort),
		Password:   config.App().RedisPassword,
		DB:         config.App().RedisDB,
		MaxRetries: config.App().RedisMaxRetry,
		PoolSize:   config.App().RedisPoolSize,
		TLSConfig:  &tls.Config{},
	}
	if config.Env() == configurations.EnvDevelopment {
		options.TLSConfig = nil
	}
	return redis.NewClient(options), nil
}

// default mongo db provider
func DefaultMongoConnection(config configurations.Configuration) (*MongoDB, error) {
	ctx := context.TODO()
	url := config.App().MongoDBUrl
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	return NewMongoDB(ctx, url, mongoboiler.New(client, config.App().MongoDBName, ctx)), nil
}

// Default Provider for task
func ProvideDefaultTaskSchedular(cfg configurations.Configuration) (*TaskClient, error) {
	conn := asynq.RedisClientOpt{
		Addr:      fmt.Sprintf("%s:%d", cfg.App().RedisHost, cfg.App().RedisPort),
		Password:  cfg.App().RedisPassword,
		PoolSize:  cfg.App().RedisPoolSize,
		DB:        cfg.App().RedisDB,
		TLSConfig: &tls.Config{},
	}
	if cfg.Env() == configurations.EnvDevelopment {
		conn.TLSConfig = nil

	}
	task := NewTaskClient(conn)
	return task, nil
}

// Default Provider for task
func ProvideDefaultTaskServer(cfg configurations.Configuration, logger Logger) (*TaskClient, error) {
	conn := asynq.RedisClientOpt{
		Addr:      fmt.Sprintf("%s:%d", cfg.App().RedisHost, cfg.App().RedisPort),
		Password:  cfg.App().RedisPassword,
		PoolSize:  cfg.App().RedisPoolSize,
		DB:        cfg.App().RedisDB,
		TLSConfig: &tls.Config{},
	}
	opt := asynq.Config{
		Concurrency: 10,
		Logger:      logger,
		Queues: map[string]int{
			"high":    6,
			"default": 3,
			"low":     1,
		},
		LogLevel: asynq.InfoLevel,
	}
	if cfg.Env() == configurations.EnvDevelopment {
		conn.TLSConfig = nil
		opt.LogLevel = asynq.DebugLevel
	}
	task := NewTaskServer(conn, opt)
	return task, nil
}
