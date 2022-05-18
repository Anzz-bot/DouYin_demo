/*
 * @Author: alexander.huang
 * @Date:   2022-05-18 21:13:30
 * @Last Modified by: alexander.huang
 * @Last Modified time: 2022-05-18 21:13:30
 */
package bootstrap

import (
	"github.com/Anzz-bot/DouYin_demo/app/models"
	"github.com/Anzz-bot/DouYin_demo/global"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// use lumberjack instead of gorm logger
func getGormLogWriter() logger.Writer {
	var writer io.Writer

	// judge if open logger
	if global.App.Config.Database.EnableFileLogWriter {
		//init writer
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.RootDir + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		// gorm writer to console
		writer = os.Stdout
	}
	return log.New(writer, "\r\n", log.LstdFlags)

}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,                          // slow SQL threshold
		LogLevel:                  logMode,                                         // logmode
		IgnoreRecordNotFoundError: false,                                           // ignore ErrRecordNotFound err
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter, // refuse colorful print
	})

}

func InitializeDB() *gorm.DB {
	//the project is driven by MySQL and can continue to iterate in the later stage
	switch global.App.Config.Database.Driver {
	case "mysql":
		return initMysqlGorm()
	default:
		return initMysqlGorm()
	}
}

func initMysqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database

	if dbConfig.Database == "" {
		return nil
	}

	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) + ")/" +
		dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"

	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string default size
		DisableDatetimePrecision:  true,  // disable datetime precision, which is not supported in databases before MySQL 5.6
		DontSupportRenameIndex:    true,  // when renaming an index, delete and create a new one. Databases before MySQL 5.7 and MariaDB do not support renaming an index
		DontSupportRenameColumn:   true,  // when renaming an index, delete and create a new one. Databases before MySQL 5.7 and MariaDB do not support renaming an index
		SkipInitializeWithVersion: false, // Auto configuration based on version
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,            // disable automatic creation of foreign key constraints
		Logger:                                   getGormLogger(), // judge Logger
	}); err != nil {
		global.App.Log.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
		initMysqlTables(db)
		return db
	}

}

// init tables
func initMysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.App.Log.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
