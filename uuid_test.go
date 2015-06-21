package uuid

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func BenchmarkV1(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V1()
	}
}

func BenchmarkV1Hex(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V1().Hex()
	}
}

func BenchmarkV3(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V3(NamespaceOID, "opensource@goanywhere.io")
	}
}

func BenchmarkV3Hex(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V3(NamespaceOID, "opensource@goanywhere.io").Hex()
	}
}

func BenchmarkV4(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V4()
	}
}

func BenchmarkV4Hex(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V4().Hex()
	}
}

func BenchmarkV5(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V5(NamespaceOID, "opensource@goanywhere.io")
	}
}

func BenchmarkV5Hex(b *testing.B) {
	for index := 0; index < b.N; index++ {
		V5(NamespaceOID, "opensource@goanywhere.io").Hex()
	}
}

func TestV3(t *testing.T) {
	//""f2107fc9-aea6-3bf0-9ad8-3bef1b5f808b
	uuid := V3(NamespaceOID, "test@example.com")
	str := "70cd6896-ecb5-3388-85ca-384edc3f3e66"
	Convey("UUID Version 3 test", t, func() {
		So(uuid.Version(), ShouldEqual, 3)
		So(uuid.String(), ShouldEqual, str)
	})
}

func TestV5(t *testing.T) {
	//""f2107fc9-aea6-3bf0-9ad8-3bef1b5f808b
	uuid := V5(NamespaceOID, "test@example.com")
	str := "067f23a9-76a5-5585-b119-32402a120978"
	Convey("UUID Version 5 test", t, func() {
		So(uuid.Version(), ShouldEqual, 5)
		So(uuid.String(), ShouldEqual, str)
	})
}
