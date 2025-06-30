package inits

import (
	"log"
	"os"
	"path/filepath"
)

func makeLibsExist(inits Init) {
	// Check if the authPath and planPath exist.
	// If they don't exist, then we should create them
	// Not a fatal operation.
	paths := []string{
		inits.AuthPath,
		inits.PlanPath,
	}

	for _, path := range paths {
		createLib(path)
	}
}

func createLib(path string) {
	if err := createPath(path); err != nil {
		log.Printf("Error creating %s lib: %s", filepath.Base(path), err)
	}
}

func createPath(path string) error {
	if err := os.MkdirAll(path, 0755); err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}
