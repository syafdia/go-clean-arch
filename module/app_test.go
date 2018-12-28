package module

import (
	"testing"
)

func TestNewAppModule(t *testing.T) {
	tests := []struct {
		name string
		want *AppModule
	}{
		{name: "test 1 should success creating same instance", want: NewAppModule()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAppModule(); got != tt.want {
				t.Errorf("NewAppModule() = %v, want %v", got, tt.want)
			}
		})
	}
}
