package app

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vaibhavchalse99/config"
	"go.uber.org/zap"
)

var (
	db     *sqlx.DB
	logger *zap.SugaredLogger
)

func Init() {
	InitLogger()
	err := initDB()
	if err != nil {
		panic(err)
	}
}

func InitLogger() {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	logger = zapLogger.Sugar()
}

func initDB() (err error) {
	dbconfig := config.Database()
	db, err = sqlx.Open(dbconfig.Driver(), dbconfig.ConnectionUrl())
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		return
	}
	db.SetMaxIdleConns(dbconfig.MaxPoolSize())
	db.SetMaxOpenConns(dbconfig.MaxOpenConn())
	db.SetConnMaxLifetime(time.Duration(dbconfig.DBMaxLifeTimeMins()) * time.Minute)

	return
}

func Close() {
	logger.Sync()
	db.Close()
}
