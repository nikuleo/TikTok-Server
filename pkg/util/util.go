package util

import (
	"strconv"
	"time"
)

func GetCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}

// 实现 string 切片转 int64 切片工具函数
func ConvtStrSliceToInt64Slice(strSlice []string) ([]int64, error) {
	intSlice := make([]int64, 0, len(strSlice))

	for _, str := range strSlice {
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}

		intSlice = append(intSlice, num)
	}

	return intSlice, nil
}

func I64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
