package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alexey/firstApp/domain/models"
	"github.com/alexey/firstApp/domain/repository"
)

type UserRepository struct {
	db        *sql.DB
	modelUser *models.User
}

func NewUserRepository(db *sql.DB) repository.InterfaceUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	row := u.db.QueryRowContext(ctx, "SELECT id, name, email, password, status FROM users WHERE email = $1", email)

	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Status)
	if err != nil {
		return nil, fmt.Errorf("Ошибка получения user")
	}

	model, err := models.NewUser("bboy23@mail.ru", "87654321", "Alexey", nil, true)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return model, nil
}

func (u *UserRepository) Save(ctc context.Context, user *models.User) error {
	quiry := `insert into users (id, password, email, name, roles, status, isselected) 
                                        values ($1, $2, $3, $4, $5, $6, $7)
										on conflict (id)
										do update set 
										    name = excluded.name,
											password = excluded.password,
											status = excluded.starus
										returning id`

	err := u.db.QueryRowContext(ctc,
		quiry,
		user.ID,
		user.Password,
		user.Email,
		user.Name,
		user.Roles,
		user.Status,
		user.IsSelected).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("Не получилось сохранить пользователя %w", err)
	}
	return nil
}
