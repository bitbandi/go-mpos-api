package mpos

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateTime(t *testing.T) {
	Convey("php timestamp", t, func() {
		Convey("not in struct", func() {
			dateStr := "2016-02-28 16:41:41"

			// create time
			tim, err := time.Parse(dateTimeFmt, dateStr)
			So(err, ShouldBeNil)

			to := TransactionTimestamp(tim)

			// marshal
			b, err := json.Marshal(&to)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, `"` + dateStr + `"`)

			// unmarshal
			err = json.Unmarshal([]byte(`"` + dateStr + `"`), &to)
			So(err, ShouldBeNil)
			So(to.String(), ShouldEqual, dateStr)
		})

		Convey("in struct", func() {
			jsonBytes := []byte(`{"name":"google","born":"2016-02-28 16:41:41"}`)
			var data = struct {
				Name string   `json:"name"`
				Born TransactionTimestamp `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			So(err, ShouldBeNil)
			So(data.Born.String(), ShouldEqual, "2016-02-28 16:41:41")

			// marshal again
			b, err := json.Marshal(&data)
			So(err, ShouldBeNil)
			So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
