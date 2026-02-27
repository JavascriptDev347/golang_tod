package config

import (
	"fmt"
	"log"
	"os"
    "to-do/models"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	
)

var DB *gorm.DB

func ConnectDB() {
	cfg := AppConfig
	// **DSN** (Data Source Name) — PostgreSQL ga ulanish uchun connection string
	var dsn string

	// Render DATABASE_URL beradi
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		dsn = dbURL
	} else {
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
			cfg.DBSSLMode,
		)
	}

	//postgres.Open(dsn) — DSN ni PostgreSQL driveriga beradi. gorm.Open() — ulanishni ochadi va *gorm.DB qaytaradi.
	//  logger.Info — har bir SQL query ni terminаlgа chiqaradi, development da qulay.
	// Production da logger.Silent qilish tavsiya etiladi. log.Fatalf — xato bo'lsa dasturni darhol to'xtatadi.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Database ga ulanishda xatolik: %v", err)
	}

	// Connection pool sozlamalari
	// GORM ichida Go ning standart database/sql paketi yashiringan. db.DB()
	// orqali shu pastki qatlamdagi *sql.DB ga murojaat qilamiz — connection pool ni boshqarish uchun kerak.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("SQL DB instance olishda xatolik: %v", err)
	}

	//Connection Pool — DB ga har safar yangi ulanish ochish qimmat, shuning uchun ulanishlar qayta ishlatiladi.
	// SetMaxOpenConns(10) — bir vaqtda maksimum 10 ta ochiq ulanish. 11-so'rov kelsa, bitta bo'shashini kutadi.
	// SetMaxIdleConns(5) — so'rov bo'lmaganda ham 5 ta ulanish "tayyor holda" turadi, tezroq javob berish uchun.
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)

	DB = db
	log.Println("✅ Database ga muvaffaqiyatli ulandi!")

	// Auto migrate
	runMigrations()
}

func runMigrations() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Todo{},
	)
	if err != nil {
		log.Fatalf("Migration xatolik: %v", err)
	}
	log.Println("✅ Migration muvaffaqiyatli bajarildi!")
}
