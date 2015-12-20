package splay

import (
	"time"
)

// Global constansts
const keyMaxLength int = 255
const commandInsertNode string = "insert"
const commandGetNode string = "get"
const commandRemoveNode string = "remove"
const memoryCheckFrequency = 1 * time.Second

// Trigger cache eviction when at % of limit
const memoryUsageThreshold byte = 80
