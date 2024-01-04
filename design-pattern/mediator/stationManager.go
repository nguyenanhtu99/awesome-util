package main

import "sync"

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
	lock           sync.Mutex
	wait           *sync.WaitGroup
}

func newStationManger(wG *sync.WaitGroup) *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           sync.Mutex{},
		wait:           wG,
	}
}

func (s *StationManager) registerArrive(t Train) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.isPlatformFree {
		s.isPlatformFree = false
		go t.permitArrival()
		return
	}

	s.trainQueue = append(s.trainQueue, t)
}

func (s *StationManager) notifyAboutDeparture() {
	s.lock.Lock()
	defer s.lock.Unlock()
	
	s.wait.Done()
	s.isPlatformFree = true

	if len(s.trainQueue) == 0 {
		return
	}

	nextTrain := s.trainQueue[0]
	s.trainQueue = s.trainQueue[1:]
	nextTrain.permitArrival()
}
