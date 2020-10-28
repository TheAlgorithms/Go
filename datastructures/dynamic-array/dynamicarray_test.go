package dynamic_array

import (
	"reflect"
	"testing"
)

func Test_dynamicArray_add(t *testing.T) {

	tests := []struct {
		name     string
		adds     int
		size     int
		capacity int
	}{
		{
			name:     "zero",
			adds:     0,
			size:     0,
			capacity: 0,
		},
		{
			name:     "one",
			adds:     1,
			size:     1,
			capacity: 10,
		},
		{
			name:     "three",
			adds:     3,
			size:     3,
			capacity: 10,
		},
		{
			name:     "ten",
			adds:     10,
			size:     10,
			capacity: 10,
		},
		{
			name:     "twelve",
			adds:     12,
			size:     12,
			capacity: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := dynamicArray{}

			for i := 0; i < tt.adds; i++ {
				da.add(struct{}{})
			}
			if da.size != tt.size {
				t.Errorf("wrong size got: %v, want: %v", da.size, tt.size)
			}
			if da.capacity != tt.capacity {
				t.Errorf("wrong capacity got: %v, want: %v", da.capacity, tt.capacity)
			}
		})
	}
}

func Test_dynamicArray_get(t *testing.T) {

	tests := []struct {
		name    string
		index   int
		want    interface{}
		wantErr bool
	}{
		{
			name:    "first",
			index:   0,
			want:    1,
			wantErr: false,
		}, {
			name:    "second",
			index:   1,
			want:    2,
			wantErr: false,
		},
		{
			name:    "third",
			index:   2,
			want:    3,
			wantErr: false,
		},
		{
			name:    "underflow",
			index:   -1,
			wantErr: true,
		},
		{
			name:    "overflow",
			index:   3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &dynamicArray{
				size:        3,
				capacity:    3,
				elementData: []interface{}{1, 2, 3},
			}
			got, err := da.get(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("error expected? %v but got? %v", tt.wantErr, err == nil)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, but want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicArray_put(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		element interface{}
		data    []interface{}
		want    []interface{}
		wantErr bool
	}{
		{
			name:    "first",
			index:   0,
			element: 123,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{123, 2, 3},
			wantErr: false,
		},
		{
			name:    "last",
			index:   2,
			element: 123,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{1, 2, 123},
			wantErr: false,
		},
		{
			name:    "underflow",
			index:   -1,
			element: 123,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{1, 2, 3},
			wantErr: true,
		},
		{
			name:    "overflow",
			index:   10,
			element: 123,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{1, 2, 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &dynamicArray{
				size:        len(tt.data),
				capacity:    len(tt.data),
				elementData: tt.data,
			}
			err := da.put(tt.index, tt.element)
			if (err != nil) != tt.wantErr {
				t.Errorf("put() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(da.elementData, tt.want) {
				t.Errorf("put() data = %v, want %v", da.elementData, tt.want)
			}
		})
	}
}

func Test_dynamicArray_remove(t *testing.T) {
	tests := []struct {
		name    string
		index   int
		data    []interface{}
		want    []interface{}
		wantErr bool
	}{
		{
			name:    "first",
			index:   0,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{2, 3, nil},
			wantErr: false,
		},
		{
			name:    "last",
			index:   2,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{1, 2, nil},
			wantErr: false,
		},
		{
			name:    "underflow",
			index:   -1,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{1, 2, 3},
			wantErr: true,
		},
		{
			name:    "overflow",
			index:   10,
			data:    []interface{}{1, 2, 3},
			want:    []interface{}{1, 2, 3},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &dynamicArray{
				size:        len(tt.data),
				capacity:    len(tt.data),
				elementData: tt.data,
			}
			err := da.remove(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("remove() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(da.elementData, tt.want) {
				t.Errorf("remove() data = %v, want %v", da.elementData, tt.want)
			}
		})
	}
}

func Test_dynamicArray_getData(t *testing.T) {
	tests := []struct {
		name string
		data []interface{}
		want []interface{}
	}{
		{
			name: "nil",
			data: nil,
			want: nil,
		},
		{
			name: "ints",
			data: []interface{}{1, 2, 3},
			want: []interface{}{1, 2, 3},
		},
		{
			name: "strings",
			data: []interface{}{"a", "b", "c"},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "structs",
			data: []interface{}{struct{}{}, struct{}{}, struct{}{}},
			want: []interface{}{struct{}{}, struct{}{}, struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &dynamicArray{
				size:        len(tt.data),
				capacity:    len(tt.data),
				elementData: tt.data,
			}
			if got := da.getData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get data = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicArray_isEmpty(t *testing.T) {

	tests := []struct {
		name  string
		size  int
		empty bool
	}{
		{
			name:  "empty",
			size:  0,
			empty: true,
		},
		{
			name:  "not empty",
			size:  123,
			empty: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &dynamicArray{
				size: tt.size,
			}
			if got := da.isEmpty(); got != tt.empty {
				t.Errorf("isEmpty() = %v, want %v", got, tt.empty)
			}
		})
	}
}
