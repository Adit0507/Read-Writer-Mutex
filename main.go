package main

import "sync"

type ReadWriteMutex struct {
	readersCounter int        // count no. of reader goroutines in critcal section
	readersLock    sync.Mutex //synchronizing readers access
	globalLock     sync.Mutex //locks writer access
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.readersLock.Lock() // only one goroutine is allowed at one time
	rw.readersCounter++

	if rw.readersCounter == 1 { // if a reader gr. is the 1st one, it attempts to lock the globalLock
		rw.globalLock.Lock()
	}

	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.globalLock.Lock() // locking for writer access
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.readersLock.Lock()
	rw.readersCounter-- // reader gr. decrements by 1

	if rw.readersCounter == 0 { // if reader gr. is the last one, unlocks the global lock
		rw.globalLock.Unlock()
	}

	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() { // writer gr. finishing its critical section unlocks the global lock
	rw.globalLock.Unlock()
}
