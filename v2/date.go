package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type date struct {
	year  int
	month int
	day   int
}

func (d date) String() string {
	return fmt.Sprintf("%v/%v/%v", d.year, d.month, d.day)
}

func (d *date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	d.year, _ = strconv.Atoi(s[0:4])
	d.month, _ = strconv.Atoi(s[5:7])
	d.day, _ = strconv.Atoi(s[8:10])
	return nil
}
