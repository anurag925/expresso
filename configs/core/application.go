package core

import (
	"expresso/configs/configurations"
	"expresso/configs/core/initializers"
	"fmt"

	"github.com/google/wire"
)

var DefaultApplicationSet = wire.NewSet(
	configurations.LoadConfigs,
	wire.Bind(new(configurations.Configuration), new(*configurations.Config)),
	initializers.ProvideDefaultFileLogger,
	wire.Bind(new(initializers.Logger), new(*initializers.FileLogger)),
	initializers.ProvideDefaultMySqlDB,
	wire.Bind(new(initializers.DB), new(*initializers.MySQL)),
	initializers.ProvideDefaultRedis,
	wire.Bind(new(initializers.Cache), new(*initializers.Redis)),
)

var DefaultBackendApplication = wire.NewSet(
	DefaultApplicationSet,
	initializers.ProvideDefaultHttpServer,
	wire.Bind(new(initializers.Server), new(*initializers.EchoServer)),
	initializers.DefaultMongoConnection,
	wire.Bind(new(initializers.Mongo), new(*initializers.MongoDB)),
	initializers.ProvideDefaultTaskSchedular,
	wire.Bind(new(initializers.AsyncTask), new(*initializers.TaskClient)),
	NewApplication,
	wire.Bind(new(App), new(*Application)),
)

var DefaultBackendAsyncWorker = wire.NewSet(
	DefaultApplicationSet,
	initializers.ProvideDefaultHttpServer,
	wire.Bind(new(initializers.Server), new(*initializers.EchoServer)),
	initializers.DefaultMongoConnection,
	wire.Bind(new(initializers.Mongo), new(*initializers.MongoDB)),
	initializers.ProvideDefaultTaskServer,
	wire.Bind(new(initializers.AsyncTask), new(*initializers.TaskClient)),
	NewApplication,
	wire.Bind(new(App), new(*Application)),
)

type App interface {
	Config() configurations.Configuration
	Server() initializers.Server
	DB() initializers.DB
	Cache() initializers.Cache
	Logger() initializers.Logger
	MongoDB() initializers.Mongo
	Tasks() initializers.AsyncTask
	Validator() initializers.Validator
}

type Application struct {
	configs   configurations.Configuration
	server    initializers.Server
	db        initializers.DB
	cache     initializers.Cache
	logger    initializers.Logger
	mongoDB   initializers.Mongo
	tasks     initializers.AsyncTask
	validator initializers.Validator
}

var _ App = (*Application)(nil)

func NewApplication(
	configs configurations.Configuration,
	logger initializers.Logger,
	db initializers.DB,
	cache initializers.Cache,
	server initializers.Server,
	mongoDB initializers.Mongo,
	tasks initializers.AsyncTask,
) *Application {
	fmt.Println(configs)
	return &Application{
		configs: configs,
		logger:  logger,
		db:      db,
		cache:   cache,
		server:  server,
		mongoDB: mongoDB,
		tasks:   tasks,
	}
}

func (a *Application) Config() configurations.Configuration {
	return a.configs
}

func (a *Application) Server() initializers.Server {
	return a.server
}

func (a *Application) DB() initializers.DB {
	return a.db
}

func (a *Application) Cache() initializers.Cache {
	return a.cache
}

func (a *Application) Logger() initializers.Logger {
	return a.logger
}

func (a *Application) MongoDB() initializers.Mongo {
	return a.mongoDB
}

func (a *Application) Tasks() initializers.AsyncTask {
	return a.tasks
}

func (a *Application) Validator() initializers.Validator {
	return a.validator
}
