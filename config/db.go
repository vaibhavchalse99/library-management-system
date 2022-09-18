package config

import "fmt"

type dbConfig struct {
	driver            string
	host              string
	port              int
	name              string
	user              string
	password          string
	maxPoolSize       int
	maxOpenCons       int
	dbMaxLifeTimeMins int
}

func Database() dbConfig {
	return appConfig.db
}

func (c dbConfig) Driver() string {
	return c.driver
}

func (c dbConfig) ConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.user, c.password, c.host, c.port, c.name)
}

func (c dbConfig) MaxPoolSize() int {
	return c.maxPoolSize
}

func (c dbConfig) MaxOpenConn() int {
	return c.maxOpenCons
}

func (c dbConfig) DBMaxLifeTimeMins() int {
	return c.dbMaxLifeTimeMins
}

func getDatabaseConfig() dbConfig {
	return dbConfig{
		driver:            readEnvString("DB_DRIVER"),
		host:              readEnvString("DB_HOST"),
		port:              readEnvInt("DB_PORT"),
		name:              readEnvString("DB_NAME"),
		user:              readEnvString("DB_USER"),
		password:          readEnvString("DB_PASSWORD"),
		maxPoolSize:       readEnvInt("DB_MAX_POOL_SIZE"),
		maxOpenCons:       readEnvInt("DB_MAX_OPEN_CONS"),
		dbMaxLifeTimeMins: readEnvInt("DB_MAX_LIFE_TIME_MINS"),
	}
}
