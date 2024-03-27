package in.anshdevs.task_zephyr.scheduler.service;

import in.anshdevs.task_zephyr.scheduler.dto.Job;
import in.anshdevs.task_zephyr.scheduler.repository.SchedulerRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class SchedulerService {
    @Autowired
    private SchedulerRepository repo;
    public List<Job> ListAllJobs(){
        return repo.findAll();
    }


}
