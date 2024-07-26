package repositories

import "google.golang.org/protobuf/types/known/timestamppb"

type ProtoBook struct {
	BookID        int64
	Title         string
	ISBN          string
	AuthorID      int64
	CategoryID    int64
	PublishedDate *timestamppb.Timestamp
	Description   string
	CreatedAt     *timestamppb.Timestamp
	UpdatedAt     *timestamppb.Timestamp
}
