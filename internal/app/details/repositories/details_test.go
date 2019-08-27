package repositories

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
)

var configFile = flag.String("f", "details.yml", "set config file which viper will loading.")

func TestDetailsRepository_Get(t *testing.T) {
	flag.Parse()

	sto, err := CreateDetailRepository(*configFile)
	if err != nil {
		t.Fatalf("create product Repository error,%+v", err)
	}

	tests := []struct {
		name     string
		id       uint64
		expected bool
	}{
		{"id=1", 1, true},
		{"id=2", 2, true},
		{"id=3", 3, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := sto.Get(test.id)

			if test.expected {
				assert.NoError(t, err )
			}else {
				assert.Error(t, err)
			}
		})
	}
}
