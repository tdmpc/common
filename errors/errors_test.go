package errors

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBombAndRecover(t *testing.T) {
	Convey("test CustomError Bomb and Recover", t, func() {
		defer Recover(nil)

		instance := "test"
		_, err := strconv.Atoi(instance)
		if err != nil {
			Bomb(err, 500, "TestBombAndRecover convert %s to int error", instance)
		}
	})
}
