package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeNoInput(t *testing.T) {
	buffer := new(bytes.Buffer)
	threshold := float64(0.0)
	limit := float64(0.0)
	inputs := []float64{}
	expected := "0.0\n"
	compute(buffer, threshold, limit, inputs)
	assert.Equal(t, expected, buffer.String())
}

func TestComputeOneInput(t *testing.T) {
	buffer := new(bytes.Buffer)
	threshold := float64(0.0)
	limit := float64(0.0)
	inputs := []float64{0.0}
	expected := "0.0\n0.0\n"
	compute(buffer, threshold, limit, inputs)
	assert.Equal(t, expected, buffer.String())
}

func TestComputeFirstSample(t *testing.T) {
	buffer := new(bytes.Buffer)
	threshold := float64(1000.0)
	limit := float64(1000000.0)
	inputs := []float64{19.0, 0.0, 1000, 1001.5, 20000, 25000000.0}
	expected := "0.0\n0.0\n0.0\n1.5\n19000.0\n980998.5\n1000000.0\n"
	compute(buffer, threshold, limit, inputs)
	assert.Equal(t, expected, buffer.String())
}

func TestComputeSecondSample(t *testing.T) {
	buffer := new(bytes.Buffer)
	threshold := float64(0.0)
	limit := float64(100.0)
	inputs := []float64{0.0, 10.0, 50.0, 50.0, 10.0, 20.0}
	expected := "0.0\n10.0\n50.0\n40.0\n0.0\n0.0\n100.0\n"
	compute(buffer, threshold, limit, inputs)
	assert.Equal(t, expected, buffer.String())
}
