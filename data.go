package main

import (
	_ "embed"
	"encoding/json"
	"log"
)

//go:embed reasons.json
var reasonsFile []byte

var reasons []string

func loadReasons() {
	err := json.Unmarshal(reasonsFile, &reasons)
	if err != nil {
		log.Fatalf("failed to load reasons: %v", err)
	}
	log.Printf("loaded %d reasons", len(reasons))
}