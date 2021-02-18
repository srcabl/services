package proto

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	pb "github.com/srcabl/protos/shared"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type dbAuditFields interface {
	CreatedByUUIDString() string
	CreatedAtUnixInt() int64
	UpdatedByUUIDNullString() sql.NullString
	UpdatedAtUnixNullInt() sql.NullInt64
}

// DBAuditFieldsToGRPC transforms database audit field values to grpc auditfields
func DBAuditFieldsToGRPC(af dbAuditFields) (*pb.AuditFields, error) {
	var createdBy, updatedBy []byte
	var createdAt, updatedAt *timestamppb.Timestamp
	cb, err := uuid.FromString(af.CreatedByUUIDString())
	if err != nil {
		errors.Wrapf(err, "failed to get uuid of created by: %s", af.CreatedByUUIDString())
	}
	createdBy = cb.Bytes()
	ct := time.Unix(af.CreatedAtUnixInt(), 0)
	createdAt, err = ptypes.TimestampProto(ct)
	if err != nil {
		errors.Wrapf(err, "failed to get timestamp for created at: %v", af.CreatedAtUnixInt())
	}
	if af.UpdatedByUUIDNullString().Valid {
		ub, err := uuid.FromString(af.UpdatedByUUIDNullString().String)
		if err != nil {
			errors.Wrapf(err, "failed to get uuid of updated by: %s", af.UpdatedByUUIDNullString().String)
		}
		updatedBy = ub.Bytes()
	}
	if af.UpdatedAtUnixNullInt().Valid {
		ut := time.Unix(af.UpdatedAtUnixNullInt().Int64, 0)
		updatedAt, err = ptypes.TimestampProto(ut)
		if err != nil {
			errors.Wrapf(err, "failed to get timestamp for updated at: %v", af.UpdatedAtUnixNullInt().Int64)
		}
	}
	return &pb.AuditFields{
		CreatedAt:     createdAt,
		CreatedByUuid: createdBy,
		UpdatedAt:     updatedAt,
		UpdatedByUuid: updatedBy,
	}, nil
}
