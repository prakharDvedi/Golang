package main

// Process represents a system process as a block of memory.
//
// MEMORY LAYOUT (on a 64-bit machine):
//
//	Address	  | Field      | Type      | Size | Padding
//	----------|------------|-----------|------|--------
//	0x00      | PID        | int32     | 4B   | 4B (Alignment for next field)
//	0x08      | Name       | string    | 16B  | 0
//	0x18      | IsElevated | bool      | 1B   | 7B (Alignment for struct size)
//
// Total Size: 32 bytes (optimized would be less if reordered)
type Process struct {
	PID        int32  // 4 bytes
	Name       string // 16 bytes (pointer + length)
	IsElevated bool   // 1 byte
}
