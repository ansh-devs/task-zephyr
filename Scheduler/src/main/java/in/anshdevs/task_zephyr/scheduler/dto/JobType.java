package in.anshdevs.task_zephyr.scheduler.dto;

public enum JobType {
    SEND_MAIL("SEND_MAIL"),
    RUN_COMMAND("RUN_COMMAND");
    JobType(String msg) {}
}
