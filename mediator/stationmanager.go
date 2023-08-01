package mediator

type StationManager struct {
	Platformfree bool
	Trainqueue   []Train
}

func NewStationManager() *StationManager {
	return &StationManager{Platformfree: true}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.Platformfree {
		s.Platformfree = false
		return true
	}
	//if not free adding it to queue
	s.Trainqueue = append(s.Trainqueue, t)
	return false
}

func (s *StationManager) Notifyaboutdepature() {
	if !s.Platformfree {
		s.Platformfree = true
	}
	if len(s.Trainqueue) > 0 {
		firstTrainInQueue := s.Trainqueue[0]
		s.Trainqueue = s.Trainqueue[1:]
		firstTrainInQueue.Permitarrival()
	}
}
