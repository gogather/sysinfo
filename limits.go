package sysinfo

// UNLIMITED unlimited
const (
	UNLIMITED = -1
)

// Limits process limits
type Limits struct {
	MaxCPUTime          int64
	MaxFileSize         int64
	MaxDataSize         int64
	MaxStackSize        int64
	MaxCoreFileSize     int64
	MaxResidentSet      int64
	MaxProcesses        int64
	MaxOpenFiles        int64
	MaxLockedMemory     int64
	MaxAddressSpace     int64
	MaxFileLocks        int64
	MaxPendingSignals   int64
	MaxMsgqueueSize     int64
	MaxNicePriority     int64
	MaxRealtimePriority int64
	MaxRealtimeTimeout  int64
}
