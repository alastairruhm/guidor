package client

import (
	"fmt"
	"net/http"

	"github.com/alastairruhm/guidor/src/schema"

	"github.com/alastairruhm/guidor/client/context"
)

const (
	databasesBasePath = "/api/databases"

	// ActionInProgress is an in progress action status
	ActionInProgress = "in-progress"

	//ActionCompleted is a completed action status
	ActionCompleted = "completed"
)

// DatabasesService handles communction with action related methods of the
type DatabasesService interface {
	// List(context.Context, *ListOptions) ([]Instance, *Response, error)
	// Get(context.Context, int) (*Instance, *Response, error)
	Register(context.Context, schema.Database) (*schema.DatabaseResult, *Response, error)
}

// DatabasesServiceOp handles communition with the image action related methods of the
type DatabasesServiceOp struct {
	client *Client
}

var _ DatabasesService = &DatabasesServiceOp{}

type databasesRoot struct {
	Databases []schema.DatabaseResult `json:"databases"`
}

type databaseRoot struct {
	Database *schema.DatabaseResult `json:"database"`
}

// List all instances
// func (s *InstancesServiceOp) List(ctx context.Context, opt *ListOptions) ([]Instance, *Response, error) {
// 	path := instancesBasePath
// 	path, err := addOptions(path, opt)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	root := new(instancesRoot)
// 	resp, err := s.client.Do(ctx, req, root)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return root.Instances, resp, err
// }

// Get an instance by ID.
// func (s *InstancesServiceOp) Get(ctx context.Context, id int) (*Instance, *Response, error) {
// 	if id < 1 {
// 		return nil, nil, NewArgError("id", "cannot be less than 1")
// 	}

// 	path := fmt.Sprintf("%s/%d", instancesBasePath, id)
// 	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	root := new(instanceRoot)
// 	resp, err := s.client.Do(ctx, req, root)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return root.Instance, resp, err
// }

// Register an database by ID.
func (s *DatabasesServiceOp) Register(ctx context.Context, d schema.Database) (*schema.DatabaseResult, *Response, error) {

	path := fmt.Sprintf("%s", databasesBasePath)
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, &d)
	if err != nil {
		return nil, nil, err
	}

	root := new(databaseRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Database, resp, err
}
