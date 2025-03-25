package main

import (
	"reflect"
	"testing"
)

func TestPathItemsToString(t *testing.T) {
	tests := []struct {
		name  string
		items []PathResponseItem
		want  []string
	}{
		{
			name: "single item",
			items: []PathResponseItem{
				{Name: "path1", Source: "source1", SourceOnDemand: true},
			},
			want: []string{`{"name": "path1","source": "source1","sourceOnDemand": true}`},
		},
		{
			name: "multiple items",
			items: []PathResponseItem{
				{Name: "path1", Source: "source1", SourceOnDemand: true},
				{Name: "path2", Source: "source2", SourceOnDemand: false},
			},
			want: []string{
				`{"name": "path1","source": "source1","sourceOnDemand": true}`,
				`{"name": "path2","source": "source2","sourceOnDemand": false}`,
			},
		},
		{
			name:  "no items",
			items: []PathResponseItem{},
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.want) == 0 {
				if got := pathItemsToString(tt.items); len(got) != 0 {
					t.Errorf("pathItemsToString() = %v, want %v", got, tt.want)
				}
				return
			}

			if got := pathItemsToString(tt.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathItemsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
