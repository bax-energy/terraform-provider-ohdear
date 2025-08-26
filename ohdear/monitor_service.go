package ohdear

import (
	"context"
	"fmt"
)

type MonitorService interface {
	Get(ctx context.Context, id int) (*Monitor, error)
	Create(ctx context.Context, req CreateMonitorRequest) (*Monitor, error)
	Update(ctx context.Context, id int, req UpdateMonitorRequest) (*Monitor, error)
	Delete(ctx context.Context, id int) error
}

type monitorService struct {
	client *Client
}

func (s *monitorService) Get(ctx context.Context, id int) (*Monitor, error) {
	var out Monitor
	if err := s.client.do(ctx, "GET", fmt.Sprintf("/monitors/%d", id), nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *monitorService) Create(ctx context.Context, req CreateMonitorRequest) (*Monitor, error) {
	var out Monitor
	if err := s.client.do(ctx, "POST", "/monitors", req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *monitorService) Update(ctx context.Context, id int, req UpdateMonitorRequest) (*Monitor, error) {
	var out Monitor
	if err := s.client.do(ctx, "PUT", fmt.Sprintf("/monitors/%d", id), req, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (s *monitorService) Delete(ctx context.Context, id int) error {
	return s.client.do(ctx, "DELETE", fmt.Sprintf("/monitors/%d", id), nil, nil)
}
