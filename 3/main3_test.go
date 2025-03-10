package main

import (
	"maps"
	"testing"
)

func TestAdd(t *testing.T) {
	m := make(map[string]int)
	Add(m, "Odin", 1)

	if val, exists := m["Odin"]; !exists || val != 1 {
		t.Errorf("Add failed: expected Odin -> 1, got %v -> %v", exists, val)
	}
}

func TestRemove(t *testing.T) {
	m := map[string]int{"Odin": 1, "Dva": 2}
	Remove(m, "Odin")

	if _, exists := m["Odin"]; exists {
		t.Errorf("Remove failed: key 'Odin' still exists")
	}
}

func TestCopy(t *testing.T) {
	original := map[string]int{"Odin": 1, "Dva": 2}
	copied := Copy(original)

	if !maps.Equal(original, copied) {
		t.Errorf("Copy failed: copied map does not match original")
	}

	original["Odin"] = 99
	if copied["Odin"] == 99 {
		t.Errorf("Copy failed: copied map is affected by changes to the original")
	}
}

func TestExists(t *testing.T) {
	m := map[string]int{"Odin": 1, "Dva": 2}

	if !Exists(m, "Odin") {
		t.Errorf("Exists failed: key 'Odin' should exist")
	}

	if Exists(m, "Freya") {
		t.Errorf("Exists failed: key 'Freya' should not exist")
	}
}

func TestGet(t *testing.T) {
	m := map[string]int{"Odin": 1, "Dva": 2}

	val, exists := Get(m, "Odin")
	if !exists || val != 1 {
		t.Errorf("Get failed: expected 1 and true, got %d and %v", val, exists)
	}

	val, exists = Get(m, "Freya")
	if exists || val != 0 {
		t.Errorf("Get failed: expected 0 and false, got %d and %v", val, exists)
	}
}
