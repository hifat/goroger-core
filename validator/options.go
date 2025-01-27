package vl

type VlTest struct {
	ErrMessage string
}

type VlTestOption = func(test *VlTest)

func Message(msg string) func(test *VlTest) {
	return func(test *VlTest) {
		test.ErrMessage = msg
	}
}
