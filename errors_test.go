package errors

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBombAndRecover(t *testing.T) {
	Convey("test Error Bomb and Recover", t, func() {
		defer Recover()

		instance := "test"
		_, err := strconv.Atoi(instance)
		if err != nil {
			Bomb(err, 500, "TestBombAndRecover convert test to int error", "字符串转数字错误")
		}
	})
}
