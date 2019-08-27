package repositories

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
)

var configFile = flag.String("f", "ratings.yml", "set config file which viper will loading.")

func TestRatingsRepository_Get(t *testing.T) {
	flag.Parse()

	sto, err := CreateRatingRepository(*configFile)
	if err != nil {
		t.Fatalf("create product Repository error,%+v", err)
	}

	tests := []struct {
		name     string
		id       uint64
		expected bool
	}{
		{"1+1", 1, true},
		{"2+3", 2, true},
		{"4+5", 3, true},
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
