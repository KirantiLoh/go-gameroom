package room


type RoomService interface {
  FindRoom(uuid string) (*Room, error)
  CreateRoom(roomDto *RoomDto) (string, error)
}

type roomServiceImpl struct {
  repo RoomRepository
}

func NewService(repo RoomRepository) RoomService {
  return &roomServiceImpl{
    repo: repo,
  }
}

func (r *roomServiceImpl) FindRoom(uuid string) (*Room, error) {
  return r.repo.Get(uuid)
}

func (r *roomServiceImpl) CreateRoom(roomDto *RoomDto) (string, error) {
  return r.repo.Create(roomDto)
}
