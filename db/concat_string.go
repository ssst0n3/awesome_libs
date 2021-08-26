package util

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
)


/*
Usage:

type Result struct {
	Data ConcatString `gorm:"type:text[]" json:"data"`
}

query := "select GROUP_CONCAT(data) AS data"
*/

type ConcatString []string

func (a *ConcatString) Scan(src interface{}) error {
	switch src := src.(type) {
	case []byte:
		return a.scanBytes(src)
	case string:
		return a.scanBytes([]byte(src))
	case nil:
		*a = nil
		return nil
	}

	return fmt.Errorf("cannot convert %T to ConcatString", src)
}

func (a *ConcatString) scanBytes(src []byte) error {
	elems := bytes.Split(src, []byte(","))
	if *a != nil && len(elems) == 0 {
		*a = (*a)[:0]
	} else {
		b := make(ConcatString, len(elems))
		for i, v := range elems {
			if b[i] = string(v); v == nil {
				return fmt.Errorf("parsing array element index %d: cannot convert nil to string", i)
			}
		}
		*a = b
	}
	return nil
}

func (a ConcatString) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return strings.Join(a, ","), nil
}
