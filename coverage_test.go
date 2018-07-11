package coverage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCOV(t *testing.T) {
	lcovParser := New("./test/lcov.info", LCOV)
	report, err := lcovParser.Parse()

	assert.NoError(t, err)
	assert.Equal(t, 67, report.CoveredBranches)
	assert.Equal(t, 1757, report.TotalBranches)
	assert.Equal(t, 412, report.CoveredFunctions)
	assert.Equal(t, 1695, report.TotalFunctions)
	assert.Equal(t, 2798, report.CoveredLines)
	assert.Equal(t, 6395, report.TotalLines)
}

func TestCobertura(t *testing.T) {
	coberturaParser := New("./test/coverage.xml", COBERTURA)
	report, err := coberturaParser.Parse()

	assert.NoError(t, err)
	assert.Equal(t, 8, report.TotalLines)
	assert.Equal(t, 6, report.CoveredLines)
	assert.Equal(t, 4, report.TotalBranches)
	assert.Equal(t, 2, report.CoveredBranches)
}
