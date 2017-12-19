package main

import (
	"testing"
	"fmt"
)

func TestLayersCreation(t *testing.T) {
	layers := initLayers([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}, 0)

	checkResultInt(t, len(layers), 7)
}

func TestUpdateLayerScannerState(t *testing.T) {
	var layer *Layer

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 0)
	checkResultInt(t, layer.scannerPosition, 1)

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 1)
	checkResultInt(t, layer.scannerPosition, 2)

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 2)
	checkResultInt(t, layer.scannerPosition, 3)

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 3)
	checkResultInt(t, layer.scannerPosition, 2)

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 14)
	checkResultInt(t, layer.scannerPosition, 3)

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 15)
	checkResultInt(t, layer.scannerPosition, 2)

	layer = &Layer{depth:4}
	updateLayerScannerState(layer, 16)
	checkResultInt(t, layer.scannerPosition, 1)
}

func TestFirst(t *testing.T) {
	severity, _ := solveFirst([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	}, 0)

	checkResultInt(t, severity, 24)
}

func TestSecond(t *testing.T) {
	delayTick := solveSecond([]string{
		"0: 3",
		"1: 2",
		"4: 4",
		"6: 4",
	})

	checkResultInt(t, delayTick, 10)
}

//
//helper functions
//
func checkResultString(t *testing.T, actualResult string, requiredResult string) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}

func checkResultInt(t *testing.T, actualResult int, requiredResult int) {
	t.Helper()

	if actualResult != requiredResult {
		t.Error(fmt.Printf("required value must be %+v, but: %+v\n", requiredResult, actualResult))
	}
}
