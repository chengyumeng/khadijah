package config_test

import (
	"testing"

	"github.com/chengyumeng/khadijah/pkg/config"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSave(t *testing.T) {
	config.UserConfigDir = "/tmp"

	Convey("Test Config Save", t, func() {
		err := config.SetUser(&config.User{Username: "test"})

		Convey("When Save config", func() {
			So(err, ShouldBeNil)
			err = config.Save()
			So(err, ShouldBeNil)
		})
		err = config.LoadOption()
		Convey("When load option from config", func() {
			So(err, ShouldBeNil)
			So(config.GlobalOption.User.Username, ShouldEqual, "test")
		})

	})
}
