package in.anshdevs.task_zephyr.scheduler.dto;


public enum JobStatus {
    COMPLETED("COMPLETED"),
    FAILED("FAILED"),
    REGISTERED("REGISTERED"),
    STARTED("STARTED"),
    ;

    JobStatus(String completed) {}
}
