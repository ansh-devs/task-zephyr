package in.anshdevs.task_zephyr.scheduler.dto;

import com.fasterxml.jackson.annotation.*;
import jakarta.persistence.*;

@Entity
@Table(name = "jobs")
public class Job {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    @Column(name = "id")
    @JsonIgnore
    private String Id;
    @Column(name = "name")
    private String Jobname;
    @Column(name = "status")
    private String Status;



    @Column(name = "type")
    private String JobType;
    @Column(name = "scheduled_at")
    private String scheduled_at;
    @Column(name = "scheduled_for")
    private String scheduled_for;

    public Job(String id, String job_name, String jobType, String scheduled_at, String scheduled_for) {
        Id = id;
        Jobname = job_name;
        JobType = jobType;
        this.scheduled_at = scheduled_at;
        this.scheduled_for = scheduled_for;
    }

    public Job() {}

    public String getId() {
        return Id;
    }

    public void setId(String id) {
        Id = id;
    }
    @JsonGetter("name")
    public String getJob_name() {
        return Jobname;
    }
    @JsonSetter("name")
    public void setJob_name(String job_name) {
        Jobname = job_name;
    }

    public String getScheduled_at() {
        return scheduled_at;
    }

    public void setScheduled_at(String scheduled_at) {
        this.scheduled_at = scheduled_at;
    }

    public String getScheduled_for() {
        return scheduled_for;
    }

    public void setScheduled_for(String scheduled_for) {
        this.scheduled_for = scheduled_for;
    }

    public String getStatus() {
        return Status;
    }

    public void setStatus(String status) {
        Status = status;
    }
    public String getJobType() {
        return JobType;
    }

    public void setJobType(String jobType) {
        JobType = jobType;
    }
}
