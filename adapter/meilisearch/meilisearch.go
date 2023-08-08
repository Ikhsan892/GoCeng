package meilisearch

import (
	"log"
	"os/exec"
)

type MeilisearchAdapter struct {
}

func NewMeilisearchAdapter() *MeilisearchAdapter {
	return &MeilisearchAdapter{}
}

func (m *MeilisearchAdapter) Init() error {
	output, err := exec.Command("./bin/meilisearch-windows-amd64.exe").Output()
	if err != nil {
		return err
	}

	log.Println(output)

	return nil
}
