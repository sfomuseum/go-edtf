package edtf

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {

	edtf_json := `{
    "start": {
      "edtf": "-1990",
      "lower": {
        "datetime": "-1990-01-01T00:00:00Z",
        "timestamp": -124965504000,
        "ymd": {
          "year": -1990,
          "month": 1,
          "day": 1
        },
        "precision": 64
      },
      "upper": {
        "datetime": "-1990-12-31T23:59:59Z",
        "timestamp": -124933968001,
        "ymd": {
          "year": -1990,
          "month": 12,
          "day": 31
        },
        "precision": 64
      }
    },
    "end": {
      "edtf": "0400",
      "lower": {
        "datetime": "0400-01-01T00:00:00Z",
        "timestamp": -49544438400,
        "ymd": {
          "year": 400,
          "month": 1,
          "day": 1
        },
        "precision": 64
      },
      "upper": {
        "datetime": "0400-12-31T23:59:59Z",
        "timestamp": -49512816001,
        "ymd": {
          "year": 400,
          "month": 12,
          "day": 31
        },
        "precision": 64
      }
    },
    "edtf": "-1990/0400",
    "level": 0,
    "feature": "Time Interval"
  }`

	var d *EDTFDate

	err := json.Unmarshal([]byte(edtf_json), &d)

	if err != nil {
		t.Logf("Failed to unmarshal JSON, %v", err)
		t.Fail()
		return
	}

	lower, err := d.Lower()

	if err != nil {
		t.Logf("Expected lower time.Time but got none, %v", err)
		t.Fail()
		return
	}

	upper, err := d.Upper()

	if err != nil {
		t.Logf("Expected upper time.Time but got none, %v", err)
		t.Fail()
		return
	}

	if lower.Unix() != -124965504000 {
		t.Logf("Unexpected lower time.Time Unix timestamp")
		t.Fail()
		return
	}

	if upper.Unix() != -49512816001 {
		t.Logf("Unexpected lower time.Time Unix timestamp")
		t.Fail()
		return
	}

}
