package linq

import (
	"reflect"
	"testing"
)

type T int

func TestFrom(t *testing.T) {
	type args struct {
		s []T
	}
	tests := []struct {
		name string
		args args
		want *List[T]
	}{
		{
			name: "get list",
			args: args{
				s: []T{1, 2, 3, 4, 5},
			},
			want: &List[T]{
				slice: []T{1, 2, 3, 4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := From(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("From() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_First(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		filter []func(T, int) bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    T
		wantErr bool
	}{
		{
			name: "get first element",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: nil,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []T{},
			},
			args: args{
				filter: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "get first element with function",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 2
					},
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "element does not exist",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 10
					},
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			got, err := l.First(tt.args.filter...)
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_FirstOrDefault(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		filter []func(value T, index int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   T
	}{
		{
			name: "get first element",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: nil,
			},
			want: 1,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []T{},
			},
			args: args{
				filter: nil,
			},
			want: 0,
		},
		{
			name: "get first element with function",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 2
					},
				},
			},
			want: 2,
		},
		{
			name: "element does not exist",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 10
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			if got := l.FirstOrDefault(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_MustFirst(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		filter []func(value T, index int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   T
		raised bool
	}{
		{
			name: "get first element",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: nil,
			},
			want:   1,
			raised: false,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []T{},
			},
			args: args{
				filter: nil,
			},
			want:   0,
			raised: true,
		},
		{
			name: "get first element with function",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 2
					},
				},
			},
			want:   2,
			raised: false,
		},
		{
			name: "element does not exist",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 10
					},
				},
			},
			want:   0,
			raised: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			defer func() {
				err := recover()
				if (err != nil) != tt.raised {
					t.Errorf("MustFirst() panic = %v, raised %v", err, tt.raised)
				}
			}()
			if got := l.MustFirst(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Last(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		filter []func(value T, index int) bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    T
		wantErr bool
	}{
		{
			name: "get last element",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: nil,
			},
			want:    5,
			wantErr: false,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []T{},
			},
			args: args{
				filter: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "get last element with function",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 2
					},
				},
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "element does not exist",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 10
					},
				},
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			got, err := l.Last(tt.args.filter...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_LastOrDefault(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		filter []func(value T, index int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   T
	}{
		{
			name: "get last element",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: nil,
			},
			want: 5,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []T{},
			},
			args: args{
				filter: nil,
			},
			want: 0,
		},
		{
			name: "get last element with function",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 2
					},
				},
			},
			want: 2,
		},
		{
			name: "element does not exist",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 10
					},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			if got := l.LastOrDefault(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_MustLast(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		filter []func(value T, index int) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   T
		raised bool
	}{
		{
			name: "get last element",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: nil,
			},
			want:   5,
			raised: false,
		},
		{
			name: "empty slice",
			fields: fields{
				slice: []T{},
			},
			args: args{
				filter: nil,
			},
			want:   0,
			raised: true,
		},
		{
			name: "get last element with function",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 2
					},
				},
			},
			want:   2,
			raised: false,
		},
		{
			name: "element does not exist",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				filter: []func(T, int) bool{
					func(v T, i int) bool {
						return v == 10
					},
				},
			},
			want:   0,
			raised: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			defer func() {
				err := recover()
				if (err != nil) != tt.raised {
					t.Errorf("MustLast() panic = %v, raised %v", err, tt.raised)
				}
			}()
			if got := l.MustLast(tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_At(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    T
		wantErr bool
	}{
		{
			name: "get element by index",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 2,
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "index is lower than 0",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: -1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "index is over max elements",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 5,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			got, err := l.At(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("At() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("At() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_MustAt(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   T
		raised bool
	}{
		{
			name: "get element by index",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 2,
			},
			want:   3,
			raised: false,
		},
		{
			name: "index is lower than 0",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: -1,
			},
			want:   0,
			raised: true,
		},
		{
			name: "index is over max elements",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 5,
			},
			want:   0,
			raised: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			defer func() {
				err := recover()
				if (err != nil) != tt.raised {
					t.Errorf("MustAt() panic = %v, raised %v", err, tt.raised)
				}
			}()
			if got := l.MustAt(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_AtOrDefault(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   T
	}{
		{
			name: "get element by index",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 2,
			},
			want: 3,
		},
		{
			name: "index is lower than 0",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: -1,
			},
			want: 0,
		},
		{
			name: "index is over max elements",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			if got := l.AtOrDefault(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AtOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Skip(t *testing.T) {
	type fields struct {
		slice []T
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *List[T]
	}{
		{
			name: "get elements from first",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 0,
			},
			want: &List[T]{
				slice: []T{1, 2, 3, 4, 5},
			},
		},
		{
			name: "get elements from second",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 1,
			},
			want: &List[T]{
				slice: []T{2, 3, 4, 5},
			},
		},
		{
			name: "get elements from last",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 4,
			},
			want: &List[T]{
				slice: []T{5},
			},
		},
		{
			name: "get elements from negative index",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: -1,
			},
			want: &List[T]{
				slice: []T{1, 2, 3, 4, 5},
			},
		},
		{
			name: "get elements with index exceeded the maximum",
			fields: fields{
				slice: []T{1, 2, 3, 4, 5},
			},
			args: args{
				index: 5,
			},
			want: &List[T]{
				slice: []T{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := From(tt.fields.slice)
			if got := l.Skip(tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}
