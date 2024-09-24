package room

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)



type RoomRepository interface {
  Get(uuid string) (*Room, error)
  Create(roomDto *RoomDto) (string, error)
}

type roomRepositoryImpl struct {
  db *sqlx.DB
  psql sq.StatementBuilderType
}

func NewRepository(db *sqlx.DB) RoomRepository {
  return &roomRepositoryImpl{
    db: db,
    psql: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
  }
}

func (r *roomRepositoryImpl) Get(uuid string) (*Room, error) {
  getRoom := r.psql.Select("*").From("rooms").Where(sq.Eq{"uuid": uuid})
  query, args, _ := getRoom.ToSql()

  room := &Room{}
  if err := r.db.Get(room, query, args...); err != nil {
    return nil, err
  }
  return room, nil
}

func (r *roomRepositoryImpl) Create(roomDto *RoomDto) (string, error) {
  createRoom := r.psql.Insert("rooms").Columns("name", "description", "leader_id").Values(roomDto.Name, roomDto.Description, roomDto.LeaderID).Suffix("RETURNING uuid")

  query, args, _ := createRoom.ToSql()

  var uuid string

  if err := r.db.QueryRow(query, args...).Scan(&uuid); err != nil {
    return "", err
  }


  return uuid, nil
}
