package usecase

import (
	"booking/internal/booking/entity"
	"booking/internal/booking/kafka"
	"booking/internal/booking/repository"
	"booking/internal/booking/server/consumer/dto"
	"booking/internal/booking/worker"
	"booking/pkg/metrics"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
)

type BookingUC struct {
	l *zap.SugaredLogger
	r *repository.Repo
	k *kafka.Producer
}

func NewBookingUC(l *zap.SugaredLogger, r *repository.Repo, k *kafka.Producer) *BookingUC {
	return &BookingUC{
		l: l,
		r: r,
		k: k,
	}
}

func (b *BookingUC) BookRoom() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Param("id")
		if id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not to be empty"})
			return
		}
		randNum := rand.Intn(10000)
		reqInt, err := b.r.BookRoom(id, randNum)
		roomBusyError := errors.New("all rooms are busy")
		if errors.Is(roomBusyError, err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "all rooms are busy"})
			return
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "cannot book room"})
			return
		}

		msg := dto.BookCode{Code: fmt.Sprintf("%d", randNum)}

		bm, err := json.Marshal(&msg)
		if err != nil {
			b.l.Errorf("Failed to marshal BookCode: %s", err)
			return
		}
		b.k.ProduceMessage(bm)
		metrics.HttpBookTotal.Inc()
		ctx.JSON(http.StatusOK, gin.H{"message": "room is booked, verify your book", "requestID": reqInt})
	}
}

func (b *BookingUC) ConfirmBook() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bookReq entity.BookRequest
		if err := ctx.ShouldBindJSON(&bookReq); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err := b.r.ConfirmBook(&bookReq)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		worker.Jobs <- bookReq.Id
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	}
}

func (b *BookingUC) CreateHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var hotel entity.Hotel
		if err := ctx.ShouldBindJSON(&hotel); err != nil {
			b.l.Infof("cannot unmarshall hotel with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with get hotel"})
			return
		}
		id, err := b.r.CreateHotel(context.Background(), &hotel)
		if err != nil {
			b.l.Infof("cannot unmarshall hotel with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error with create hotel"})
			return
		}
		ctx.JSON(http.StatusOK, id)
	}
}

func (b *BookingUC) GetHotels() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hotels, err := b.r.GetHotels(context.Background())
		if err != nil {
			b.l.Infof("cannot get hotels with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error with get hotels"})
			return
		}
		ctx.JSON(http.StatusOK, hotels)
	}
}

func (b *BookingUC) UpdateHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			b.l.Info("user get empty id")
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "id should not be empty"})
			return
		}
		var hotel entity.Hotel
		if err := ctx.ShouldBindJSON(&hotel); err != nil {
			b.l.Infof("cannot unmarshall hotel with err:%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "error with take hotel"})
			return
		}
		err := b.r.UpdateHotel(context.Background(), id, &hotel)
		if err != nil {
			b.l.Infof("cannot update hotel with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error with update hotel"})
			return
		}
		ctx.JSON(http.StatusOK, id)
	}
}

func (b *BookingUC) DeleteHotel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			b.l.Info("user get empty id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not be empty"})
			return
		}
		err := b.r.DeleteHotel(context.Background(), id)
		if err != nil {
			b.l.Infof("cannot delete hotel with err:%v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error with delete hotel"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": id})
	}
}

func (b *BookingUC) GetHotelById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			b.l.Info("user get empty id")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "id should not be empty"})
			return
		}
		hotel, err := b.r.GetHotelById(context.Background(), id)
		if err != nil {
			b.l.Infof("cannot get hotel by with id %s with err:%v", id, err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error with get hotel by id"})
			return
		}
		ctx.JSON(http.StatusOK, hotel)
	}
}
