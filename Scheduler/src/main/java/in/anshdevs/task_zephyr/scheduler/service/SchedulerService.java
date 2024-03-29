package in.anshdevs.task_zephyr.scheduler.service;

import in.anshdevs.task_zephyr.scheduler.dto.Job;

import java.util.List;
import java.util.Optional;

public interface SchedulerService {
    public Job CreateJob(Job job);
    // ListAllJobs Lists all the jobs registered in the scheduler.
    public List<Job> ListAllJobs();
    // GetJobById fetch the job by the given id in the scheduler.
    public Optional<Job> GetJobById(String id);
}
