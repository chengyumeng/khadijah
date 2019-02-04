package table

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInitTable(t *testing.T) {
	Convey("Test Table Init", t, func() {
		Convey("Horizontal Table", func() {
			t := NewTable(Horizontal)
			So(t.IsEmpty(), ShouldBeTrue)
			err := t.AddRow([]string{"a", "b", "c"})
			So(err, ShouldBeError)
			t.SetHeaders([]string{"h1", "h2"})
			err = t.AddRow([]string{"a", "b", "c"})
			So(err, ShouldBeNil)
			So(t.IsEmpty(), ShouldBeFalse)
		})

	})
}
