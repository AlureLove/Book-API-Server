package config

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

// define application configuration

type Config struct {
	App   *App   `json:"app"`
	MySQL *MySQL `json:"mysql"`
}

func (c *Config) String() string {
	v, _ := json.Marshal(c)
	return string(v)
}

type App struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (a *App) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
	Debug    bool   `json:"debug"`

	lock sync.Mutex
	db   *gorm.DB
}

func (c *MySQL) DB() *gorm.DB {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		if c.Debug {
			db = db.Debug()
		}

		c.db = db
	}

	return c.db
}
