package main

import (
	"bytes"
	"encoding/json"
	"time"
)

type PgTime struct {
	time.Time
}

func (self PgTime) MarshalJSON() ([]byte, error) {
	b := new(bytes.Buffer)
	jsonEncoder := json.NewEncoder(b)
	err := jsonEncoder.Encode(self.Format("2006-01-02 15:04:05"))
	return b.Bytes(), err
}

func (self *PgTime) UnmarshalJSON(data []byte) error {
	br := bytes.NewReader(data)
	dec := json.NewDecoder(br)

	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}

	now := time.Now()
	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, now.Location())
	if err != nil {
		return err
	}
	println(t.String())

	b, err := t.MarshalJSON()
	if err != nil {
		return err
	}

	err = self.Time.UnmarshalJSON(b)
	return err
}
