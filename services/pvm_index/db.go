// (c) 2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pvm_index

import (
	"errors"

	"github.com/ava-labs/gecko/ids"
	"github.com/ava-labs/gecko/vms/components/codec"
	"github.com/gocraft/dbr"
	"github.com/gocraft/health"
)

const (
	// MaxSerializationLen is the maximum number of bytes a canonically
	// serialized tx can be stored as in the database.
	MaxSerializationLen = 16_384

	PaginationLimit = 500
)

var (
	ErrSerializationTooLong = errors.New("serialization is too long")
)

type DB struct {
	chainID ids.ID
	codec   codec.Codec
	stream  *health.Stream
	db      *dbr.Connection
}

// NewDBIndex creates a new DB for the given config
func NewDBIndex(stream *health.Stream, db *dbr.Connection, chainID ids.ID, codec codec.Codec) *DB {
	return &DB{
		stream:  stream,
		db:      db,
		chainID: chainID,
		codec:   codec,
	}
}

func (i *DB) newDBSession(name string) *dbr.Session {
	return i.db.NewSession(i.stream.NewJob(name))
}
