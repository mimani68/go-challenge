package repository

import (
	"context"
	"fmt"

	"app.ir/config"
	"app.ir/internal/data/db"
	"app.ir/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SegmentRepository struct {
	cfg config.Config
	db  db.Database
}

func CreateSegmentRepository(cfg config.Config, db db.Database) *SegmentRepository {
	return &SegmentRepository{cfg, db}
}

func (r *SegmentRepository) GetSegmentsStream(ctx context.Context, changed func([]*SegmentRepository)) {
	stream, err := r.db.Segments.Watch(ctx, mongo.Pipeline{})
	if err != nil {
		fmt.Println(err)
	}
	defer func(stream *mongo.ChangeStream, ctx context.Context) {
		err := stream.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(stream, ctx)
	for {
		select {
		case <-ctx.Done():
			err := stream.Close(ctx)
			if err != nil {
				fmt.Println(err)
			}
			return
		default:
			var data []*SegmentRepository
			ok := stream.Next(ctx)
			if !ok {
				err := stream.Err()
				if err != nil {
					println(err)
					return
				}
			}
			if ok {
				err := stream.Decode(data)
				if err != nil {
					println(err)
					return
				}
				changed(data)
			}
		}
	}
}

func (r *SegmentRepository) fetch(ctx context.Context, filter *bson.M, opts *options.FindOptions) ([]*SegmentRepository, error) {
	cursor, err := r.db.Segments.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	}(cursor, ctx)
	data := make([]*SegmentRepository, 0)
	if err = cursor.All(ctx, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (r *SegmentRepository) GetSegments(ctx context.Context, skip int64, limit int64) ([]*SegmentRepository, error) {
	filter, opts := &bson.M{}, options.Find().SetSort(bson.M{"title": 1}).SetSkip(skip).SetLimit(limit)
	data, err := r.fetch(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *SegmentRepository) GetSegmentById(ctx context.Context, id string) (*SegmentRepository, error) {
	data := new(SegmentRepository)
	if err := r.db.Segments.FindOne(ctx, bson.M{"_id": bson.M{"$eq": id}}).Decode(&data); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *SegmentRepository) CreateSegment(ctx context.Context, segment domain.Segment) (string, error) {
	segment.Id = primitive.NewObjectID().Hex()
	_, err := r.db.Segments.InsertOne(ctx, segment)
	if err != nil {
		return "", err
	}
	return segment.Id, nil
}

func (r *SegmentRepository) CountUserInSegment(ctx context.Context, segment string) (int32, error) {
	// number, err := r.db.Segments.Find(ctx, segment)
	number, err := r.db.Segments.CountDocuments(ctx, bson.M{"segment": segment})
	if err != nil {
		return 0, err
	}
	return int32(number), nil
}
