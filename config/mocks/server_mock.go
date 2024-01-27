package mocks

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

var (
	mockedserver = mockserver{
		mocks: make(map[string]*Mock),
	}
)

type mockserver struct {
	enable      bool
	servermutex sync.Mutex
	mocks       map[string]*Mock
}

func StartMockServer() {
	mockedserver.servermutex.Lock()
	defer mockedserver.servermutex.Unlock()
	mockedserver.enable = true
}

func StopMockServer() {
	mockedserver.servermutex.Lock()
	defer mockedserver.servermutex.Unlock()
	mockedserver.enable = false
}

func AddMock(mock Mock) {
	mockedserver.servermutex.Lock()
	defer mockedserver.servermutex.Unlock()
	key := mockedserver.getMockKey(mock.Method, mock.URL, mock.RequestBody)

	mockedserver.mocks[key] = &mock
}

func (m *mockserver) cleanBody(body string) string {
	nbody := strings.TrimSpace(body)
	if nbody == "" {
		return ""
	}

	nbody = strings.Replace(nbody, "\t", "", -1)
	nbody = strings.Replace(nbody, "\n", "", -1)
	return nbody
}

func (m *mockserver) getMockKey(method, url, body string) string {
	hash := md5.New()
	hash.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hash.Sum(nil))
	return key
}

func FlushMock() {
	mockedserver.servermutex.Lock()
	defer mockedserver.servermutex.Unlock()
	mockedserver.mocks = make(map[string]*Mock)
}

func GetMock(method, url, body string) *Mock {
	if !mockedserver.enable {
		return nil
	}

	if mock := mockedserver.mocks[mockedserver.getMockKey(method, url, body)]; mock != nil {
		return mock
	}

	return &Mock{
		Error: fmt.Errorf("no mock matched for %s - %s", method, url),
	}

}
