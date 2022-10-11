package common

import (
	"testing"
)

func TestValidName(t *testing.T) {
	testCases := map[string]struct {
		wantErr bool
		str     string
	}{
		"check name with valid name case 1": {
			wantErr: false,
			str:     "appdoamin",
		},
		"check name with valid name case 2": {
			wantErr: false,
			str:     "app-doamin",
		},
		"check name with invalid name case 1": {
			wantErr: true,
			str:     "app_domain",
		},
		"check name with invalid name case 2": {
			wantErr: true,
			str:     "app.domain",
		},
		"check name with invalid name case 3: max length overflow": {
			wantErr: true,
			str:     "abcdefghigklmnopqrstuvwxyz0123456789-abcdefghigklmnopqrstuvwxyz0123456789",
		},
		"check name with invalid name case 4: number start": {
			wantErr: true,
			str:     "1",
		},
		"check name with invalid name case 5: hyphen end": {
			wantErr: true,
			str:     "acbc-",
		},
	}

	for k, item := range testCases {
		gotErr := ValidName(item.str)
		switch item.wantErr {
		case true:
			if gotErr == nil {
				t.Errorf("%s failure, want got err, but got nil", k)
			}
		case false:
			if gotErr != nil {
				t.Errorf("%s failure, want got nil, but got err", k)
			}
		}
	}

}

func TestValidAppName(t *testing.T) {
	testCases := map[string]struct {
		wantErr bool
		str     string
	}{
		"check name with valid name case 1": {
			wantErr: false,
			str:     "appdoamin",
		},
		"check name with valid name case 2": {
			wantErr: false,
			str:     "app-doamin",
		},
		"check name with invalid name case 1": {
			wantErr: true,
			str:     "app_domain",
		},
		"check name with invalid name case 2": {
			wantErr: true,
			str:     "app.domain",
		},
		"check name with invalid name case 3: max length overflow": {
			wantErr: true,
			str:     "abcdefghigklmnopqrstuvwxyz0123456789-abcdefghi",
		},
	}

	for k, item := range testCases {
		gotErr := ValidAppName(item.str)
		switch item.wantErr {
		case true:
			if gotErr == nil {
				t.Errorf("%s failure, want got err, but got nil", k)
			}
		case false:
			if gotErr != nil {
				t.Errorf("%s failure, want got nil, but got err", k)
			}
		}
	}
}

func TestValidVersion(t *testing.T) {
	testCases := map[string]struct {
		wantErr bool
		str     string
	}{
		"check name with valid version case 1": {
			wantErr: false,
			str:     "1.0.0",
		},
		"check name with valid version case 2": {
			wantErr: false,
			str:     "1.0.0-rc.1",
		},
		"check name with invalid version case 1": {
			wantErr: true,
			str:     "1-0-0",
		},
		"check name with invalid version case 2": {
			wantErr: true,
			str:     "1",
		},
	}

	for k, item := range testCases {
		gotErr := ValidVersion(item.str)
		switch item.wantErr {
		case true:
			if gotErr == nil {
				t.Errorf("%s failure, want got err, but got nil", k)
			}
		case false:
			if gotErr != nil {
				t.Errorf("%s failure, want got nil, but got err", k)
			}
		}
	}
}
