// Package dog manages Dogs - Deacon's helper workers for infrastructure tasks.
// Dogs are reusable workers with multi-rig worktrees, managed by the Deacon.
// Unlike polecats (single-rig, ephemeral sessions), dogs handle cross-rig infrastructure work.
package dog

import (
	"time"
)

// State represents a dog's operational state.
type State string

const (
	// StateIdle means the dog is available for work.
	StateIdle State = "idle"
	// StateWorking means the dog is executing a task.
	StateWorking State = "working"
)

// Dog represents a Deacon helper worker.
type Dog struct {
	Name          string            `json:"name"`
	State         State             `json:"state"`
	Path          string            `json:"path"`
	Worktrees     map[string]string `json:"worktrees,omitempty"`
	LastActive    time.Time         `json:"last_active"`
	Work          string            `json:"work,omitempty"`
	WorkStartedAt *time.Time        `json:"work_started_at,omitempty"` // When work was assigned (nil if idle)
	CreatedAt     time.Time         `json:"created_at"`
}

// DogState is the persistent state stored in .dog.json.
type DogState struct {
	Name          string            `json:"name"`
	State         State             `json:"state"`
	LastActive    time.Time         `json:"last_active"`
	Work          string            `json:"work,omitempty"`            // Current work assignment
	WorkStartedAt *time.Time        `json:"work_started_at,omitempty"` // When work was assigned (nil if idle)
	Worktrees     map[string]string `json:"worktrees,omitempty"`       // Rig -> path (for verification)
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}
