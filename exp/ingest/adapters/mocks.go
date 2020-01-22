package adapters

import (
	"github.com/stellar/go/exp/ingest/io"
	"github.com/stellar/go/xdr"
	"github.com/stretchr/testify/mock"
)

type MockHistoryArchiveAdapter struct {
	mock.Mock
}

func (m *MockHistoryArchiveAdapter) GetLatestLedgerSequence() (uint32, error) {
	args := m.Called()
	return args.Get(0).(uint32), args.Error(1)
}

func (m *MockHistoryArchiveAdapter) BucketListHash(sequence uint32) (xdr.Hash, error) {
	args := m.Called(sequence)
	return args.Get(0).(xdr.Hash), args.Error(1)
}

func (m *MockHistoryArchiveAdapter) GetState(
	sequence uint32, tempSet io.TempSet, maxStreamRetries int,
) (io.StateReader, error) {
	args := m.Called(sequence, tempSet, maxStreamRetries)
	return args.Get(0).(io.StateReader), args.Error(1)
}

func (m *MockHistoryArchiveAdapter) GetLedger(sequence uint32) (io.ArchiveLedgerReader, error) {
	args := m.Called(sequence)
	return args.Get(0).(io.ArchiveLedgerReader), args.Error(1)
}