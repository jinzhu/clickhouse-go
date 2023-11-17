// Licensed to ClickHouse, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. ClickHouse, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by make codegen DO NOT EDIT.
// source: lib/column/codegen/array.tpl

package column

import (
	"database/sql"
	"github.com/ClickHouse/ch-go/proto"
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/shopspring/decimal"
	"math/big"
	"net"
	"net/netip"
	"time"
)

// appendRowPlain is a reflection-free realisation of append for plain arrays.
func (col *Array) appendRowPlain(v any) error {
	switch tv := v.(type) {
	case []float32:
		return appendRowPlain(col, tv)
	case []*float32:
		return appendNullableRowPlain(col, tv)
	case []float64:
		return appendRowPlain(col, tv)
	case []*float64:
		return appendNullableRowPlain(col, tv)
	case []int8:
		return appendRowPlain(col, tv)
	case []*int8:
		return appendNullableRowPlain(col, tv)
	case []int16:
		return appendRowPlain(col, tv)
	case []*int16:
		return appendNullableRowPlain(col, tv)
	case []int32:
		return appendRowPlain(col, tv)
	case []*int32:
		return appendNullableRowPlain(col, tv)
	case []int64:
		return appendRowPlain(col, tv)
	case []*int64:
		return appendNullableRowPlain(col, tv)
	case []uint8:
		return appendRowPlain(col, tv)
	case []*uint8:
		return appendNullableRowPlain(col, tv)
	case []uint16:
		return appendRowPlain(col, tv)
	case []*uint16:
		return appendNullableRowPlain(col, tv)
	case []uint32:
		return appendRowPlain(col, tv)
	case []*uint32:
		return appendNullableRowPlain(col, tv)
	case []uint64:
		return appendRowPlain(col, tv)
	case []*uint64:
		return appendNullableRowPlain(col, tv)
	case []string:
		return appendRowPlain(col, tv)
	case []*string:
		return appendNullableRowPlain(col, tv)
	case [][]byte:
		return appendRowPlain(col, tv)
	case []*[]byte:
		return appendNullableRowPlain(col, tv)
	case []sql.NullString:
		return appendRowPlain(col, tv)
	case []*sql.NullString:
		return appendNullableRowPlain(col, tv)
	case []int:
		return appendRowPlain(col, tv)
	case []*int:
		return appendNullableRowPlain(col, tv)
	case []uint:
		return appendRowPlain(col, tv)
	case []*uint:
		return appendNullableRowPlain(col, tv)
	case []big.Int:
		return appendRowPlain(col, tv)
	case []*big.Int:
		return appendNullableRowPlain(col, tv)
	case []decimal.Decimal:
		return appendRowPlain(col, tv)
	case []*decimal.Decimal:
		return appendNullableRowPlain(col, tv)
	case []bool:
		return appendRowPlain(col, tv)
	case []*bool:
		return appendNullableRowPlain(col, tv)
	case []sql.NullBool:
		return appendRowPlain(col, tv)
	case []*sql.NullBool:
		return appendNullableRowPlain(col, tv)
	case []time.Time:
		return appendRowPlain(col, tv)
	case []*time.Time:
		return appendNullableRowPlain(col, tv)
	case []sql.NullTime:
		return appendRowPlain(col, tv)
	case []*sql.NullTime:
		return appendNullableRowPlain(col, tv)
	case []uuid.UUID:
		return appendRowPlain(col, tv)
	case []*uuid.UUID:
		return appendNullableRowPlain(col, tv)
	case []netip.Addr:
		return appendRowPlain(col, tv)
	case []*netip.Addr:
		return appendNullableRowPlain(col, tv)
	case []net.IP:
		return appendRowPlain(col, tv)
	case []*net.IP:
		return appendNullableRowPlain(col, tv)
	case []proto.IPv6:
		return appendRowPlain(col, tv)
	case []*proto.IPv6:
		return appendNullableRowPlain(col, tv)
	case [][16]byte:
		return appendRowPlain(col, tv)
	case []*[16]byte:
		return appendNullableRowPlain(col, tv)
	case []orb.MultiPolygon:
		return appendRowPlain(col, tv)
	case []*orb.MultiPolygon:
		return appendNullableRowPlain(col, tv)
	case []orb.Point:
		return appendRowPlain(col, tv)
	case []*orb.Point:
		return appendNullableRowPlain(col, tv)
	case []orb.Polygon:
		return appendRowPlain(col, tv)
	case []*orb.Polygon:
		return appendNullableRowPlain(col, tv)
	case []orb.Ring:
		return appendRowPlain(col, tv)
	case []*orb.Ring:
		return appendNullableRowPlain(col, tv)
	default:
		return col.appendRowDefault(v)
	}
}
