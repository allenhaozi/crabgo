package common

import "testing"

func TestTargetNginxIngressAnnotation(t *testing.T) {
	testCases := map[string]struct {
		want bool
		flag string
	}{
		"test hit the ingress annotation case 1": {
			want: true,
			flag: "nginx.ingress.kubernetes.io/app-root",
		},
		"test hit the ingress annotation case 2": {
			want: true,
			flag: "nginx.ingress.kubernetes.io/app-root",
		},
		"test hit the ingress annotation case 3": {
			want: true,
			flag: "nginx.ingress.kubernetes.io/app-root",
		},
		"test miss the ingress annotation case 3": {
			want: false,
			flag: "ingress.kubernetes.io/app-root",
		},
	}

	for k, item := range testCases {
		got := TargetNginxIngressAnnotation(item.flag)
		if got != item.want {
			t.Errorf("%s failure, want:%t, got:%t", k, item.want, got)
		}
	}
}
