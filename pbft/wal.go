package pbft

import (
	"csic/pbft/pbftpb"
	"os"
	"sync"
)

const (
	CURRENT_LOG_ID = "log_status_current_log_id"
	LOG_PREFIX = "logid_"
)

type WAL struct {
	dir 			string
	dirFile 		*os.File
	metadata 		[]byte
	state			pbftpb.HardState
	mu      		sync.Mutex
	enti    		uint64

	st				Storage
}

func CreateWal(dirpath string, metadata []byte) (*WAL, error) {
	// todo memorydb
	w := &WAL{
		dir:      dirpath,
		metadata: metadata,
		st:NewMemoryStorage(""),
	}

	return w, nil
}

func (w *WAL) Save(st pbftpb.HardState, ents []pbftpb.Entry) error {
	err := w.st.Append(ents)
	return err
}

func (w *WAL) Close() error {
	// todo write to disk
	// after write done, we send a signal to channel
	return nil
}