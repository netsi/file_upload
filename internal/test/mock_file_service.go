package test

type MockFileService struct {
	expectedGetFileBytes []byte
	expectedGetFileFound bool
	expectedStoreFileErr error
}

func NewMockFileService() *MockFileService {
	return &MockFileService{}
}

func (m *MockFileService) WithExpectedGetFileResponse(bytes []byte, found bool) *MockFileService {
	m.expectedGetFileBytes = bytes
	m.expectedGetFileFound = found
	return m
}

func (m *MockFileService) WithExpectedStoreFileResponse(err error) *MockFileService {
	m.expectedStoreFileErr = err
	return m
}

func (m *MockFileService) GetFile(_ string) ([]byte, bool) {
	return m.expectedGetFileBytes, m.expectedGetFileFound
}

func (m *MockFileService) StoreFile(_ string, _ []byte) error {
	return m.expectedStoreFileErr
}
