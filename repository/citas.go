package repository

import (
	"context"

	"github.com/Imjowend/citas-golang/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	InsertCita(ctx context.Context, post *models.Cita) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetCitaByDNI(ctx context.Context, id string) (*models.Cita, error)
	DeleteCitaByDNI(ctx context.Context, id string, userId string) error
	UpdateCitaByDNI(ctx context.Context, post *models.Cita, userId string) error
	GetCitas(ctx context.Context, page uint64) ([]*models.Cita, error)
	Close() error
}

var implementation Repository

func SetRepository(repository Repository) {
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func InsertCita(ctx context.Context, cita *models.Cita) error {
	return implementation.InsertCita(ctx, cita)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func GetCitaByDNI(ctx context.Context, id string) (*models.Cita, error) {
	return implementation.GetCitaByDNI(ctx, id)
}

func DeleteCitaByDNI(ctx context.Context, id string, userId string) error {
	return implementation.DeleteCitaByDNI(ctx, id, userId)
}

func UpdateCitaByDNI(ctx context.Context, cita *models.Cita, userId string) error {
	return implementation.UpdateCitaByDNI(ctx, cita, userId)
}

func GetCitas(ctx context.Context, page uint64) ([]*models.Cita, error) {
	return implementation.GetCitas(ctx, page)
}

func Close() error {
	return implementation.Close()
}
