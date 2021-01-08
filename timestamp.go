package edtf

import (
	"strconv"
	"strings"
	"time"
)

type Timestamp struct {
	timestamp int64
}

func NewTimestampWithTime(t *time.Time) *Timestamp {
	return &Timestamp{t.Unix()}
}

func (ts *Timestamp) Time() *time.Time {

	t := time.Unix(ts.Unix(), 0)
	t = t.UTC()

	return &t
}

func (ts *Timestamp) Unix() int64 {
	return ts.timestamp
}

func (ts *Timestamp) UnmarshalJSON(b []byte) error {

	s := strings.Trim(string(b), `"`)
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		return err
	}

	*ts = Timestamp{i}
	return nil
}

func (ts Timestamp) MarshalJSON() ([]byte, error) {
	str_ts := strconv.FormatInt(ts.timestamp, 10)
	return []byte(str_ts), nil
}
