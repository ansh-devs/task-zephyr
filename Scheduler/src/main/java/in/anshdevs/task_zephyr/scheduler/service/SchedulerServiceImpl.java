package in.anshdevs.task_zephyr.scheduler.service;

import in.anshdevs.task_zephyr.scheduler.dto.Job;
import in.anshdevs.task_zephyr.scheduler.repository.SchedulerRepository;
import org.springframework.stereotype.Service;

import java.time.Instant;
import java.util.List;
import java.util.Optional;

@Service
public class SchedulerServiceImpl implements SchedulerService {

    public final SchedulerRepository repo;

    public SchedulerServiceImpl(SchedulerRepository repo) {
        this.repo = repo;
    }

    @Override
    public Job CreateJob(Job job) {
        long currentTime=Instant.now().getEpochSecond();
        job.setScheduled_at(Long.toString(currentTime));
        return repo.save(job);
    }

    @Override
    public List<Job> ListAllJobs() {
         return repo.findAll();
    }

    @Override
    public Optional<Job> GetJobById(String id) {
        return repo.findById(id);
    }
}
