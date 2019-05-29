package code

type errorCode struct {
	ERROR        int
	NotFound     int
	LoginError   int
	LoginTimeout int
	TokeNotValid int
}


var ErrorCode = errorCode {
	ERROR:        1,
	NotFound:     404,
	LoginError:   1000,
	LoginTimeout: 1001,
	TokeNotValid: 10001,
}
