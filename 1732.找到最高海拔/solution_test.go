package leetcode

import (
	"reflect"
	"testing"
)

// 1732.找到最高海拔
// https://leetcode-cn.com/problems/find-the-highest-altitude
func largestAltitude(gain []int) int {

}
func TestSolution(t *testing.T) {
	testCases := []struct {
		desc string
		want 
	}{
		{
            want: ,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			get := 
			if !reflect.DeepEqual(tC.want,get){
				t.Errorf("input: %+v get: %v\n",tC,get)
			}
		})
	}
}
