package common

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/0xPolygon/polygon-sdk/types"
)

// Min returns the strictly lower number
func Min(a, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}

// Max returns the strictly bigger number
func Max(a, b uint64) uint64 {
	if a > b {
		return a
	}

	return b
}

func ConvertUnmarshalledInt(x interface{}) (int64, error) {
	switch tx := x.(type) {
	case float64:
		return roundFloat(tx), nil
	case string:
		v, err := types.ParseUint64orHex(&tx)
		if err != nil {
			return 0, err
		}

		return int64(v), nil
	default:
		return 0, errors.New("unsupported type for unmarshalled integer")
	}
}

func roundFloat(num float64) int64 {
	return int64(num + math.Copysign(0.5, num))
}

func ToFixedFloat(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))

	return float64(roundFloat(num*output)) / output
}

// SetupDataDir sets up the data directory and the corresponding sub-directories
func SetupDataDir(dataDir string, paths []string) error {
	if err := createDir(dataDir); err != nil {
		return fmt.Errorf("failed to create data dir: (%s): %w", dataDir, err)
	}

	for _, path := range paths {
		path := filepath.Join(dataDir, path)
		if err := createDir(path); err != nil {
			return fmt.Errorf("failed to create path: (%s): %w", path, err)
		}
	}

	return nil
}

// DirectoryExists checks if the directory at the specified path exists
func DirectoryExists(directoryPath string) bool {
	// Grab the absolute filepath
	pathAbs, err := filepath.Abs(directoryPath)
	if err != nil {
		return false
	}

	// Check if the directory exists, and that it's actually a directory if there is a hit
	if fileInfo, statErr := os.Stat(pathAbs); os.IsNotExist(statErr) || (fileInfo != nil && !fileInfo.IsDir()) {
		return false
	}

	return true
}

// createDir creates a file system directory if it doesn't exist
func createDir(path string) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
