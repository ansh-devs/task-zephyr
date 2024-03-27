package in.anshdevs.task_zephyr.scheduler.repository;

import in.anshdevs.task_zephyr.scheduler.dto.Job;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface SchedulerRepository extends JpaRepository<Job,String> {}
