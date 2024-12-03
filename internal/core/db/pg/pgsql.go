package pg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	entTasks "simple-api/internal/pkg/tasks/ent"
	entUsers "simple-api/internal/pkg/users/ent"
	"time"
)

type Postgres struct {
	DB *gorm.DB
}

// NewPostgres initializes the Postgres instance with a configurable logger
func NewPostgres(dsn string, enableLogs bool) *Postgres {
	var newLogger logger.Interface
	if enableLogs {
		newLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)
	} else {
		newLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	err = db.AutoMigrate(&entUsers.User{}, &entTasks.Task{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping PostgreSQL: %v", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return &Postgres{DB: db}
}

// GetDB returns the GORM database instance
func (p *Postgres) GetDB() *gorm.DB {
	return p.DB
}

// WithTransaction executes a function within a transaction
func (p *Postgres) WithTransaction(fn func(tx *gorm.DB) error) error {
	tx := p.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
