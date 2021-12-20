package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateAvailableSeats(t *testing.T) {
	type args struct {
		countSeats int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "0 elements",
			args: args{
				countSeats: 0,
			},
			want: []int{},
		},
		{
			name: "5 elements",
			args: args{
				countSeats: 5,
			},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateAvailableSeats(tt.args.countSeats); !reflect.DeepEqual(fmt.Sprint(got), fmt.Sprint(tt.want)) {
				t.Errorf("GenerateAvailableSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}
