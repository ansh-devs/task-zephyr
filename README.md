
# **Task Zephyr**

![Thumbnail](https://gist.githubusercontent.com/ansh-devs/bf8ecc2edc0d57a7dfa15a1019bb59e0/raw/d46c7cabd38c8abe922016e467ab64a68a6bcd48/task-zephyr.png)

![GitHub commit activity](https://img.shields.io/github/commit-activity/t/ansh-devs/task-zephyr?style=for-the-badge&labelColor=%23000000&color=%23C1F0EC)

![GitHub last commit](https://img.shields.io/github/last-commit/ansh-devs/task-zephyr?style=for-the-badge&labelColor=%23000000&color=%23C1F0EC)

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/ansh-devs/task-zephyr/testchecks.yml?branch=main&event=push&style=for-the-badge&labelColor=%23000000&color=%23C1F0EC)

* Task Zephyr is a distributed task scheduler designed to execute various types of tasks asynchronously across multiple workers or machines. It provides a robust and scalable solution for offloading time-consuming or resource-intensive tasks from the main application, ensuring smooth and efficient operation.
### *Currently, task-zephyr supports two types of tasks:*

* ***Mail Task:*** This task is responsible for sending email notifications or transactional emails asynchronously. It leverages a dedicated worker process to handle the email delivery process, ensuring that the main application remains responsive and unaffected by potential delays or failures in email delivery.
* ***Shell Command Task:*** This task allows executing arbitrary shell commands or scripts in the background. It can be used for various purposes, such as running data processing jobs, executing system maintenance tasks, or triggering external processes without blocking the main application.
### The architecture of task-zephyr is designed to be highly scalable and distributed. It consists of multiple components:
* ***Scheduler:*** A reliable and persistent message queue that stores and manages the tasks to be executed. This queue acts as a buffer, ensuring that tasks are not lost even if workers go offline or experience failures.
* ***Orchestrator:*** The central component responsible for scheduling and distributing tasks to available workers based on their capabilities and load.
* ***Worker Nodes:*** These are the execution units that actually perform the tasks assigned by the scheduler. Worker nodes can be dynamically added or removed from the system, allowing for horizontal scaling as the workload increases.
  _Monitoring and Logging: Comprehensive monitoring and logging mechanisms are in place to track the status of tasks, detect failures, and provide insights into the overall performance of the system._