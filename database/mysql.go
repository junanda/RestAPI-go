package database

import (
	"fmt"
	"log"

	"github.com/junanda/golang-aa/config"
	"github.com/junanda/golang-aa/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB interface {
	Database
	GetDb() *gorm.DB
}

type MysqlDatabase struct {
	cfg config.Config
	db  *gorm.DB
}

func NewMysqlDB(cfg config.Config) MysqlDB {
	return &MysqlDatabase{
		cfg: cfg,
	}
}

func (m *MysqlDatabase) Connect() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.cfg.User, m.cfg.Password, m.cfg.Host, m.cfg.Port, m.cfg.DBName)
	m.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed Connect to Database Mysql: ", err.Error())
	}

	log.Println("mysql connection success...")

	if err = m.db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to run Migration: ", err.Error())
	}

	log.Println("Migrate Database")
}

func (m *MysqlDatabase) Close() {
	dbc, err := m.db.DB()
	if err != nil {
		log.Fatal("error databases mysql:", err)
	}

	dbc.Close()
	log.Println("Mysql Connection close...")
}

func (m *MysqlDatabase) GetDb() *gorm.DB {
	return m.db
}
