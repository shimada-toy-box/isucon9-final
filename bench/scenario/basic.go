package scenario

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chibiegg/isucon9-final/bench/isutrain"
)

// BasicScenario は基本的な予約フローのシナリオです
type BasicScenario struct {
	client *isutrain.Client
}

func NewBasicScenario(targetBaseURL string) (*BasicScenario, error) {
	client, err := isutrain.NewClient(targetBaseURL)
	if err != nil {
		return nil, err
	}

	return &BasicScenario{
		client: client,
	}, nil
}

func (s *BasicScenario) Run(ctx context.Context) error {
	err := s.client.Register(ctx, "hoge", "hoge")
	if err != nil {
		return fmt.Errorf("ユーザ登録ができません: %w", err)
	}

	err = s.client.Login(ctx, "hoge", "hoge")
	if err != nil {
		return fmt.Errorf("ユーザログインができません: %w", err)
	}

	stations, err := s.client.ListStations(ctx)
	if err != nil {
		return fmt.Errorf("駅一覧を取得できません: %w", err)
	}
	log.Println(stations)

	trains, err := s.client.SearchTrains(ctx, time.Now(), "東京", "大阪")
	if err != nil {
		return fmt.Errorf("列車検索ができません: %w", err)
	}
	log.Println(trains)

	seats, err := s.client.ListTrainSeats(ctx, "こだま", "96号")
	if err != nil {
		return fmt.Errorf("列車の座席座席列挙できません: %w", err)
	}
	log.Println(seats)

	reservation, err := s.client.Reserve(ctx)
	if err != nil {
		return fmt.Errorf("予約ができません: %w", err)
	}
	log.Println(reservation)

	err = s.client.CommitReservation(ctx, 1111)
	if err != nil {
		return fmt.Errorf("予約を確定できませんでした: %w", err)
	}

	reservations, err := s.client.ListReservations(ctx)
	if err != nil {
		return fmt.Errorf("予約を列挙できませんでした: %w", err)
	}
	log.Println(reservations)

	return nil
}
