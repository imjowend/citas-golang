package database

import (
	"context"

	"github.com/Imjowend/citas-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=user password=password dbname=citas port=5432"

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Cita{})

	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) Close() error {
	postgreSQL, err := repo.db.DB()
	if err != nil {
		return err
	}
	return postgreSQL.Close()
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	err := repo.db.Create(user)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *PostgresRepository) InsertCita(ctx context.Context, cita *models.Cita) error {
	err := repo.db.Create(cita)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	err := repo.db.Find(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return user, nil
}

func (repo *PostgresRepository) GetCitaByDNI(ctx context.Context, id string) (*models.Cita, error) {
	var cita *models.Cita
	err := repo.db.Find(&cita)
	if err.Error != nil {
		return nil, err.Error
	}
	return cita, nil
}

func (repo *PostgresRepository) DeleteCitaByDNI(ctx context.Context, id string, userId string) error {
	var cita *models.Cita
	err := repo.db.Delete(&cita)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (repo *PostgresRepository) UpdateCitaByDNI(ctx context.Context, cita *models.Cita, userId string) error {
	// Logica para update
	return err
}

func (repo *PostgresRepository) GetCitas(ctx context.Context, page uint64) ([]*models.Cita, error) {
	var citas []*models.Cita
	err := repo.db.Find(&citas)
	if err.Error != nil {
		return nil, err.Error
	}
	return citas, nil
}
