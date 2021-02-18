package proto

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const maxPageSize = 100

type listRequest interface {
	GetInitPageSize() uint32
	GetPageToken() string
}

// PaginationToken is a struct to represent the pagination
type PaginationToken struct {
	PageSize   uint
	PageNumber uint
}

// NewTokenFromRequest takes a list request and returns a pagination token
func NewTokenFromRequest(req listRequest) *PaginationToken {
	if req.GetPageToken() != "" {
		return detokenize(req.GetPageToken())
	}
	pageSize := uint(req.GetInitPageSize())
	if pageSize < 1 || pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	return &PaginationToken{
		PageSize:   pageSize,
		PageNumber: 0,
	}
}

func detokenize(token string) *PaginationToken {
	var t *PaginationToken
	gob.NewDecoder(base64.NewDecoder(base64.StdEncoding, strings.NewReader(token))).Decode(t)
	return t
}

// EncodeNextToken crates an encodes token for the next page
func (t *PaginationToken) EncodeNextToken() (string, error) {
	t.PageNumber++
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	if err := gob.NewEncoder(encoder).Encode(t); err != nil {
		return "", errors.Wrap(err, "failed to token ize pagination token")
	}
	encoder.Close()
	return buf.String(), nil
}

// ApplyToQuery will append the order and limit statments to the query
func (t *PaginationToken) ApplyToQuery(query, orderBy string) string {
	offset := t.PageNumber * t.PageSize
	return fmt.Sprintf("%sORDER BY %s LIMIT %d, %d", query, orderBy, offset, t.PageSize)
}
