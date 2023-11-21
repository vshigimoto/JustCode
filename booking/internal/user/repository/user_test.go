package repository

import (
	"booking/internal/user/entity"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestRepo_GetByLoginUser(t *testing.T) {
	user := entity.User{
		Id:       1,
		Name:     "dimka",
		Email:    "123@mail.ru",
		Login:    "dimka",
		Password: "123",
	}
	id := "dimka"
	url := "localhost:9235/api/user/v1/user/" + id
	res, err := http.Get(url)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		t.Fatal("error with req")
	}
	httpClient := http.Client{}
	res, err = httpClient.Do(req)
	if err != nil {
		t.Fatalf("client making http request err: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("failed to read response body err: %w", err)
	}
	var resp entity.User
	json.Unmarshal(body, &resp)
	if resp != user {
		t.Fatalf("not equal")
	}
}
