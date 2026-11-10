package twosum

import (
    "reflect"
    "testing"
)

func TestTwoSum(t *testing.T) {
    cases := []struct {
        name string
        nums []int
        target int
        want []int
    }{
        {"basic", []int{2, 7, 11, 15}, 9, []int{0, 1}},
        {"middle", []int{3, 2, 4}, 6, []int{1, 2}},
        {"duplicates", []int{3, 3}, 6, []int{0, 1}},
        {"no match", []int{1, 2, 3}, 10, nil},
        {"empty", []int{}, 5, nil},
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got := TwoSum(tc.nums, tc.target)
            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("got %v, want %v", got, tc.want)
            }
        })
    }
}
