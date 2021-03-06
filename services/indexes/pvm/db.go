// (c) 2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package pvm

import (
	"errors"

	"github.com/ava-labs/gecko/utils/codec"
	"github.com/gocraft/dbr"
	"github.com/gocraft/health"
)

var (
	ErrSerializationTooLong = errors.New("serialization is too long")
)

type DB struct {
	networkID uint32
	chainID   string
	codec     codec.Codec
	stream    *health.Stream
	db        *dbr.Connection
}

// NewDBIndex creates a new DB for the given config
func NewDBIndex(stream *health.Stream, db *dbr.Connection, networkID uint32, chainID string, codec codec.Codec) *DB {
	return &DB{
		networkID: networkID,
		chainID:   chainID,
		codec:     codec,
		stream:    stream,
		db:        db,
	}
}

func (db *DB) newSession(name string) *dbr.Session {
	return db.db.NewSession(db.stream.NewJob(name))
}
