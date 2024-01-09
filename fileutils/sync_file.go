package fileutils

import (
	"os"
	"sync"
)

type SyncFile struct {
	file  *os.File
	mutex sync.Mutex
}

// NewSynchronizedFile synchronizes writing to a writer
func NewSyncFile(f *os.File) *SyncFile {
	sf := &SyncFile{file: f}
	return sf
}

func (sf *SyncFile) WriteString(s string) (int, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.file.WriteString(s)
}

func (sf *SyncFile) Write(b []byte) (int, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.file.Write(b)
}

func (sf *SyncFile) Close() error {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	return sf.file.Close()
}
