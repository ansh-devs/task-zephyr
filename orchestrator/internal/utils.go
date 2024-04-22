package internal

import "time"

func (s *Orchestrator) managePool() {
	s.Wg.Add(1)
	defer s.Wg.Done()
	tkr := time.NewTicker(time.Duration(s.MaxHealthCheckOverlook) * s.HealthCheckTTL)
	defer tkr.Stop()
	for {
		select {
		case <-tkr.C:
			s.cleanWorkerPool()

		case <-s.Ctx.Done():
			return
		}
	}
}

func (s *Orchestrator) cleanWorkerPool() {
	s.WorkerPoolMtx.Lock()
	defer s.WorkerPoolMtx.Unlock()
	for id, wrkr := range s.WorkerPool {
		if wrkr.HealthCheckMiss > s.MaxHealthCheckOverlook {
			_ = wrkr.Manager.Close()
			delete(s.WorkerPool, id)
			s.AcquirableWorkerIDsMtx.Lock()

			wrkravailable := len(s.WorkerPool)
			s.AcquirableWorkerIDs = make([]string, wrkravailable)
			for k, _ := range s.WorkerPool {
				s.AcquirableWorkerIDs = append(s.AcquirableWorkerIDs, k)
			}
			s.AcquirableWorkerIDsMtx.Unlock()
		} else {
			wrkr.HealthCheckMiss++
		}
	}
}

func (s *Orchestrator) ScrapeDatabaseForJobs() {

	tkr := time.NewTicker(time.Duration(time.Second * 15))
	defer tkr.Stop()

	for {
		select {
		case <-tkr.C:
			go s.AssignTaskToWorker()
		case <-s.Ctx.Done():
			return
		}
	}
}

func (s *Orchestrator) StartServer() error {
	return nil
}

func (s *Orchestrator) gracefulShutdown() error {
	return nil
}