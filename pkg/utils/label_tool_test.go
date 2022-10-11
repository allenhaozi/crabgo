package utils

import (
	"reflect"
	"testing"
)

func TestLabelTool_ParseQueryParams2Labels(t *testing.T) {
	type args struct {
		labels string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			args: args{
				labels: "aios.4pd.io/app=app-notebook-env-v-1-0-0,a=b",
			},
			want: map[string]string{
				"aios.4pd.io/app": "app-notebook-env-v-1-0-0",
				"a":               "b",
			},
			wantErr: false,
		},
		{
			args: args{
				labels: "aios.4pd.io/app=app-notebook-env-v-1-0-0",
			},
			want: map[string]string{
				"aios.4pd.io/app": "app-notebook-env-v-1-0-0",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := LabelTool{}
			got, err := lt.ParseQueryParams2Labels(tt.args.labels)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseQueryParams2Labels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseQueryParams2Labels() got = %v, want %v", got, tt.want)
			}
		})
	}
}
