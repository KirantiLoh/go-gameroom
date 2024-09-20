package users

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Get(email string) (*User, error)
	Create(userDto *UserDto) error
	Update(user *User) error
}

type userRepo struct {
	db   *sqlx.DB
	psql sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) UserRepository {
	return &userRepo{
		db:   db,
		psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (u *userRepo) Get(email string) (*User, error) {
	user := &User{}

  sql, args, _ := u.psql.Select("*").From("users").Where(sq.Eq{"email": email}).ToSql()
  if err := u.db.Get(user, sql, args...); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userRepo) Create(userDto *UserDto) error {

	if _, err := u.psql.Insert("users").Columns("email", "name", "password", "is_verified").Values(userDto.Email, userDto.Name, userDto.Password, false).RunWith(u.db).Exec(); err != nil {
		return err
	}

	return nil
}

func (u *userRepo) Update(user *User) error {
	if _, err := u.psql.Update("users").Set("is_verified", user.IsVerified).Set("name", user.Name).Where(sq.Eq{"id": user.ID}).RunWith(u.db).Exec(); err != nil {
		return err
	}

	return nil
}
