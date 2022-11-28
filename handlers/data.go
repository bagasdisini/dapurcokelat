package handlers

import (
	datadto "app/dto"
	dto "app/dto/status"
	"app/models"
	"app/repositories"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
)

type handlerData struct {
	DataRepository repositories.DataRepository
}

func HandlerData(DataRepository repositories.DataRepository) *handlerData {
	return &handlerData{DataRepository}
}

var Message = ""

func (h *handlerData) ShowData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := datadto.DataResponse{
		Message: Message,
	}

	post := models.Data{
		Message: request.Message,
	}

	post, _ = h.DataRepository.ShowData(post)

	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events-ws", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	conn.WriteMessages(kafka.Message{Value: []byte(request.Message)})

	batch := conn.ReadBatch(1e3, 1e9)
	bytes := make([]byte, 1e3)

	for {
		_, err := batch.Read(bytes)
		if err != nil {
			break
		}
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: http.StatusOK, Data: post}
	json.NewEncoder(w).Encode(response)
}
