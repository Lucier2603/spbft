package pbft

import (
	"csic/pbft/pbftpb"
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/protobuf/proto"
	"os"
	"strconv"
	"sync"
)

var ErrOutOfBound = errors.New("request entry out of range.")



type Storage interface {
	// InitialState returns the saved HardState and ConfState information.
	InitialState() (pbftpb.HardState, pbftpb.ConfState, error)
	// Entries returns a slice of log entries in the range [lo,hi).
	// MaxSize limits the total size of the log entries returned, but
	// Entries returns at least one entry if any.
	Entries(lo, hi, maxSize uint64) ([]pbftpb.Entry, error)
	Entry(idx uint64) (pbftpb.Entry, error)
	Append([]pbftpb.Entry) error


	// Term returns the term of entry i, which must be in the range
	// [FirstIndex()-1, LastIndex()]. The term of the entry before
	// FirstIndex is retained for matching purposes even though the
	// rest of that entry may not be available.
	Term(i uint64) (uint64, error)
	// LastIndex returns the index of the last entry in the log.
	LastIndex() (uint64, error)
	// FirstIndex returns the index of the first log entry that is
	// possibly available via Entries (older entries have been incorporated
	// into the latest Snapshot; if storage only contains the dummy entry the
	// first log entry is not available).
	FirstIndex() (uint64, error)
	Snapshot() (Snapshot, error)
}

type MemoryStorage struct {
	db				Db
	mu				*sync.Mutex
}

func NewMemoryStorage(dir string) Storage {
	return &MemoryStorage{
		db:NewMemoryDb(dir),
		mu:&sync.Mutex{},
	}
}

func (st *MemoryStorage) InitialState() (pbftpb.HardState, pbftpb.ConfState, error) {
	// todo load from disk
	hs := &pbftpb.HardState{
		View:1,
		Vote:1,
		Commit:1,
	}
	cs := &pbftpb.ConfState{
		Nodes:[]uint64{},
		Learners:[]uint64{},
	}
	return *hs, *cs, nil
}

func (st *MemoryStorage) Entries(lo, hi, maxSize uint64) ([]pbftpb.Entry, error) {
	es := make([]pbftpb.Entry, 0)
	for i := lo; i <= hi; i++ {
		e, err := st.Entry(i)
		if err == nil {
			es = append(es, e)
		} else {
			return es, err
		}
	}
	return es, nil
}

func (st *MemoryStorage) Entry(idx uint64) (pbftpb.Entry, error) {
	e := pbftpb.Entry{}

	b := st.db.Get("_entry_" + strconv.FormatUint(idx, 10))
	if b == nil {
		return e, ErrMarshalJson
	}

	err := proto.Unmarshal(b, &e)
	return e, err
}

func (st *MemoryStorage) Append(ents []pbftpb.Entry) error {
	st.mu.Lock()
	defer st.mu.Unlock()

	idx, _ := st.LastIndex()

	for _, e := range ents {
		// todo 暂时使用json
		//b, err := proto.Marshal(&e)
		b, err := json.Marshal(e)
		if err != nil {
			return err
		}

		st.db.Put("_entry_" + strconv.FormatUint(e.Index, 10), b)

		if idx < e.Index {
			idx = e.Index
		}
	}

	// update last index
	st.db.PutString("_last_index", strconv.FormatUint(idx, 10))

	return nil
}

func (st *MemoryStorage) Term(i uint64) (uint64, error) {
	tstr := st.db.GetString("_term")
	if len(tstr) == 0 {
		return 1, nil
	}
	return strconv.ParseUint(tstr, 10, 64)
}

func (st *MemoryStorage) LastIndex() (uint64, error) {
	istr := st.db.GetString("_last_index")
	if len(istr) == 0 {
		return 1, nil
	}
	return strconv.ParseUint(istr, 10, 64)
}

func (st *MemoryStorage) FirstIndex() (uint64, error) {
	return 0, nil
}

func (st *MemoryStorage) Snapshot() (Snapshot, error) {
	s := &Snapshot{}
	return *s, nil
}





// DB
type Db interface {
	Put(k string, v []byte)
	Get(k string) []byte
	PutString(k string, v string)
	GetString(k string) string
	Delete(k string)
}

type MemoryDb struct {
	data			map[string][]byte

	writer			*os.File
}

func NewMemoryDb(dir string) *MemoryDb {
	db := &MemoryDb{
		data: make(map[string][]byte),
	}

	fmt.Println(dir)
	w, err := os.OpenFile(dir, os.O_WRONLY|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	db.writer = w

	return db
}

func (mst *MemoryDb) Put(k string, v []byte) {
	mst.data[k] = v

	mst.writer.Write(v)
	mst.writer.Write([]byte("\n"))

}

func (mst *MemoryDb) Get(k string) []byte {
	return mst.data[k]
}

func (mst *MemoryDb) PutString(k string, v string) {
	mst.data[k] = []byte(v)

	mst.writer.Write([]byte(v))
	mst.writer.Write([]byte("\n"))
}

func (mst *MemoryDb) GetString(k string) string {
	b := mst.Get(k)
	if b != nil {
		return string(b)
	}
	return ""
}

func (mst *MemoryDb) Delete(k string) {
	mst.data[k] = nil
}
