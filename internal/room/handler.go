package room

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kirantiloh/gameroom/pkg/auth"
)

type RoomHandler struct {
	service RoomService
}

func NewHandler(service RoomService) RoomHandler {
	return RoomHandler{
		service: service,
	}
}

func (h RoomHandler) CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body != nil {
		defer r.Body.Close()
	}

  user := r.Context().Value("user").(auth.UserData)

  createRoomDto := &CreateRoomDto{}

	if err := json.NewDecoder(r.Body).Decode(createRoomDto); err != nil {
    fmt.Println(r.Body)
    w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
		})
		return
	}

  roomDto := &RoomDto{
    CreateRoomDto: *createRoomDto,
    LeaderID: user.ID,
  }

	uuid, err := h.service.CreateRoom(roomDto)
	if err != nil {
    println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Bad Request",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Room with id %s successfully created!", uuid),
		"roomId":  uuid,
	})
}

func (h RoomHandler) FindRoomByUUIDHandler(w http.ResponseWriter, r *http.Request) {
	uuid := chi.URLParam(r, "uuid")

	room, err := h.service.FindRoom(uuid)

	if err != nil {
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Room not found",
			})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]Room{
		"room": *room,
	})
}
