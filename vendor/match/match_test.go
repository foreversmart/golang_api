package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchArr(t *testing.T) {
	assertion := assert.New(t)
	var data = []string{
		`ab`,
		`axb`,
		`axxb`,
		`axxxb`,
		`axxxbx`,
		`ac`,
		`xxxb`,
	}

	var res = []string{
		`ab`,
		`axb`,
		`axxb`,
		`axxxb`,
		`axxxbx`,
	}

	for i, v := range MatchArr(`ab`, data) {
		assertion.Equal(res[i], v)
	}
}

func TestMatchWeight(t *testing.T) {
	assertion := assert.New(t)

	var testData1 = []struct {
		pattern string
		content string
	}{
		{`ab`, `ab`}, // a*b
		{`a`, `ab`},
		{`ab`, `axb`},
		{`ab`, `axxb`},
		{`ab`, `axxxb`},
		{`ab`, `axxxxb`},
		{`ab`, `axxxxbx`},
		{`ab`, `axxxxbxx`},
	}

	for i, data := range testData1 {
		if i == 0 {
			continue
		}
		assertion.False(matchWeight(data.pattern, data.content) > matchWeight(testData1[i-1].pattern, testData1[i-1].content))
	}

	var testData2 = []struct {
		pattern string
		content string
	}{
		{`ab`, `ab`}, // *ab
		{`ab`, `abx`},
		{`ab`, `xab`},
		{`ab`, `xabx`},
		{`ab`, `xaxb`},
		{`ab`, `xaxbx`},
		{`ab`, `xxaxbx`},
	}

	for i, data := range testData2 {
		if i == 0 {
			continue
		}
		assertion.False(matchWeight(data.pattern, data.content) > matchWeight(testData2[i-1].pattern, testData2[i-1].content))
	}

}
