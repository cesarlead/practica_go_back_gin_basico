package repository

import (
	"context"
	"time"

	"github.com/cesarlead/practica_go_back_gin_basico/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresUserRepo implementa UserRepository usando pgxpool.Pool.
type PostgresUserRepo struct {
	pool *pgxpool.Pool
}

// NewPostgresUserRepo inyecta el pool de conexiones.
func NewPostgresUserRepo(pool *pgxpool.Pool) *PostgresUserRepo {
	return &PostgresUserRepo{pool: pool}
}

func (r *PostgresUserRepo) FindAll() ([]*domain.User, error) {
	const q = `
        SELECT id, name, email, created_at, updated_at
        FROM users
        ORDER BY id`
	rows, err := r.pool.Query(context.Background(), q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		u := new(domain.User)
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *PostgresUserRepo) FindByID(id int) (*domain.User, error) {
	const q = `
        SELECT id, name, email, created_at, updated_at
        FROM users
        WHERE id = $1`
	row := r.pool.QueryRow(context.Background(), q, id)
	u := new(domain.User)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (r *PostgresUserRepo) Save(u *domain.User) (int, error) {
	const q = `
      INSERT INTO users (name, email, created_at, updated_at)
      VALUES ($1,$2,$3,$4)
      RETURNING id`
	now := time.Now().UTC()
	var id int
	err := r.pool.QueryRow(
		context.Background(),
		q,
		u.Name, u.Email, now, now,
	).Scan(&id)
	return id, err
}

func (r *PostgresUserRepo) Update(u *domain.User) error {
	const q = `
      UPDATE users
      SET name=$1, email=$2, updated_at=$3
      WHERE id=$4`
	u.UpdatedAt = time.Now().UTC()
	cmd, err := r.pool.Exec(
		context.Background(),
		q,
		u.Name, u.Email, u.UpdatedAt, u.ID,
	)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *PostgresUserRepo) Delete(id int) error {
	const q = `DELETE FROM users WHERE id=$1`
	cmd, err := r.pool.Exec(context.Background(), q, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}
