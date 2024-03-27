package in.anshdevs.task_zephyr.scheduler.dto;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity
public class Job {
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    public String Id;

    public String Job_name;
    public String scheduled_at;
    public String scheduled_for;


    public Job() {}

    //! ---------- Getter & Setters ----------------


    public Job(String id, String job_name, String scheduled_at, String scheduled_for) {
        Id = id;
        Job_name = job_name;
        this.scheduled_at = scheduled_at;
        this.scheduled_for = scheduled_for;
    }

    public String getId() {
        return Id;
    }

    public void setId(String id) {
        Id = id;
    }

    public String getJob_name() {
        return Job_name;
    }

    public void setJob_name(String job_name) {
        Job_name = job_name;
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
}
