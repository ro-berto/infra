package dumper

import (
	"bufio"
	"context"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseGoldenEyeJsonData(t *testing.T) {
	t.Parallel()

	Convey("Parse Data", t, func() {
		Convey("happy path", func() {
			file, err := os.Open("test.json")
			reader := bufio.NewReader(file)

			devices, err := parseGoldenEyeData(context.Background(), reader)
			So(err, ShouldEqual, nil)
			So(devices.Devices, ShouldNotEqual, nil)
		})
		Convey("parse for non existent file", func() {
			file, err := os.Open("test2.json")
			reader := bufio.NewReader(file)

			devices, err := parseGoldenEyeData(context.Background(), reader)
			So(err, ShouldNotEqual, nil)
			So(err.Error(), ShouldEqual, "unmarshal chunk failed while reading golden eye data for devices: invalid argument")
			So(devices, ShouldEqual, nil)
		})
	})
}
