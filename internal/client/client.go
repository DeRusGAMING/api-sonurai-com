package client

import (
	"api/internal/model"
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WallpaperClient interface {
	Get(ctx context.Context, id string) (*model.Image, error)
	GetByOldID(ctx context.Context, id int) (*model.Image, error)
	List(ctx context.Context, q ListQuery) (*model.ListResponse, error)
	ListByTag(ctx context.Context, tag string) (*model.ListResponse, error)
}

func NewClient(collection string, firestore *firestore.Client) *Client {
	return &Client{
		collection: collection,
		firestore:  firestore,
	}
}

type Client struct {
	collection string
	firestore  *firestore.Client
}

type ListQuery struct {
	Limit          int
	StartAfterDate int
	StartAfterID   string
	Reverse        bool
}

func (c *Client) Get(ctx context.Context, id string) (*model.Image, error) {
	doc, err := c.firestore.Collection(c.collection).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, nil
		}
		return nil, err
	}

	image := new(model.Image)
	if err := jsonToInterface(doc.Data(), &image); err != nil {
		return nil, err
	}

	return image, nil
}

func (c *Client) GetByOldID(ctx context.Context, id int) (*model.Image, error) {
	iter := c.firestore.Collection(c.collection).Where("oldId", "==", id).Documents(ctx)

	doc, err := iter.Next()
	if err != nil {
		if status.Code(err) == codes.NotFound || err == iterator.Done {
			return nil, nil
		}
		return nil, err
	}

	image := new(model.Image)
	if err := jsonToInterface(doc.Data(), &image); err != nil {
		return nil, err
	}

	return image, nil
}

func (c *Client) List(ctx context.Context, q ListQuery) (*model.ListResponse, error) {
	query := c.firestore.Collection(c.collection).Limit(q.Limit)

	if q.Reverse {
		query = query.
			OrderBy("date", firestore.Asc).
			OrderBy("id", firestore.Desc)

	} else {
		query = query.
			OrderBy("date", firestore.Desc).
			OrderBy("id", firestore.Asc)
	}

	if q.StartAfterDate != 0 && q.StartAfterID != "" {
		query = query.StartAfter(q.StartAfterDate, q.StartAfterID)
	}

	dsnap, err := query.Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var res model.ListResponse
	for _, v := range dsnap {
		var wallpaper model.ImageBasic
		if err := jsonToInterface(v.Data(), &wallpaper); err != nil {
			return nil, err
		}
		res.Data = append(res.Data, wallpaper)
	}

	if q.Reverse {
		reverseImages(res.Data)
	}

	return &res, nil
}

func (c *Client) ListByTag(ctx context.Context, tag string) (*model.ListResponse, error) {
	dsnap, err := c.firestore.Collection(c.collection).
		Limit(36).
		OrderBy(fmt.Sprintf("tags.%s", tag), firestore.Desc).
		Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var res model.ListResponse
	for _, v := range dsnap {
		var wallpaper model.ImageBasic
		if err := jsonToInterface(v.Data(), &wallpaper); err != nil {
			return nil, err
		}
		res.Data = append(res.Data, wallpaper)
	}

	return &res, nil
}

func jsonToInterface(in map[string]interface{}, out interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, out); err != nil {
		return err
	}
	return nil
}

func reverseImages(a []model.ImageBasic) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}