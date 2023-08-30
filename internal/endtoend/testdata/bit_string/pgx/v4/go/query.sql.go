// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package querytest

import (
	"context"
)

const selectTest = `-- name: SelectTest :many
SELECT v_bit_null, v_varbit_null, v_bit, v_varbit
from test_table
`

func (q *Queries) SelectTest(ctx context.Context) ([]TestTable, error) {
	rows, err := q.db.Query(ctx, selectTest)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TestTable
	for rows.Next() {
		var i TestTable
		if err := rows.Scan(
			&i.VBitNull,
			&i.VVarbitNull,
			&i.VBit,
			&i.VVarbit,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
