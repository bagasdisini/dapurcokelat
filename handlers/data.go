package handlers

import (
	dto "app/dto/status"
	"app/repositories"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type handlerData struct {
	DataRepository repositories.DataRepository
}

func HandlerData(DataRepository repositories.DataRepository) *handlerData {
	return &handlerData{DataRepository}
}

func (h *handlerData) ShowData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := (mux.Vars(r)["dataUser"])

	if id == "A" {
		conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)
		conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
		conn.WriteMessages(kafka.Message{Value: []byte(id)})

		message, _ := conn.ReadMessage(1e6)

		dataUser, err := h.DataRepository.ShowData(string(message.Value))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		conn2, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events", 0)
		conn2.SetWriteDeadline(time.Now().Add(time.Second * 10))
		conn2.WriteMessages(kafka.Message{Value: []byte(dataUser.Result)})

		message2, _ := conn2.ReadMessage(1e6)

		dataUser2, err := h.DataRepository.ShowData(string(message2.Value))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: http.StatusOK, Data: dataUser2}
		json.NewEncoder(w).Encode(response)
	}

	if id == "B" {
		conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events-1", 0)
		conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
		conn.WriteMessages(kafka.Message{Value: []byte(id)})

		message, _ := conn.ReadMessage(1e6)

		dataUser, err := h.DataRepository.ShowData(string(message.Value))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		conn2, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events-1", 0)
		conn2.SetWriteDeadline(time.Now().Add(time.Second * 10))
		conn2.WriteMessages(kafka.Message{Value: []byte(dataUser.Result)})

		message2, _ := conn2.ReadMessage(1e6)

		dataUser2, err := h.DataRepository.ShowData(string(message2.Value))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: http.StatusOK, Data: dataUser2}
		json.NewEncoder(w).Encode(response)
	}

	if id == "C" {
		conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events-2", 0)
		conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
		conn.WriteMessages(kafka.Message{Value: []byte(id)})

		message, _ := conn.ReadMessage(1e6)

		dataUser, err := h.DataRepository.ShowData(string(message.Value))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		conn2, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "quickstart-events-2", 0)
		conn2.SetWriteDeadline(time.Now().Add(time.Second * 10))
		conn2.WriteMessages(kafka.Message{Value: []byte(dataUser.Result)})

		message2, _ := conn2.ReadMessage(1e6)

		dataUser2, err := h.DataRepository.ShowData(string(message2.Value))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := dto.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusOK)
		response := dto.SuccessResult{Status: http.StatusOK, Data: dataUser2}
		json.NewEncoder(w).Encode(response)
	}

}
