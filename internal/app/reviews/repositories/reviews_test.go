package repositories

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"testing"
)

var configFile = flag.String("f", "reviews.yml", "set config file which viper will loading.")

func TestReviewsRepository_Get(t *testing.T) {
	flag.Parse()

	sto, err := CreateReviewRepository(*configFile)
	if err != nil {
		t.Fatalf("create reviews Repository error,%+v", err)
	}

	tests := []struct {
		name     string
		id       uint64
		expected int
	}{
		{"id=1", 1, 1},
		{"id=2", 2, 1},
		{"id=3", 3, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rs, err := sto.Query(test.id)
			if err != nil {
				t.Errorf("query reviews error:%+v", err)
			}

			assert.Equal(t, test.expected, len(rs))
		})
	}
}
