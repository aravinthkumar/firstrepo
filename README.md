## Requirement

Need an application which can track daily recording of habits.

Project the yearly/quarterly view of habits.

**Plan**


| Field Names | Description                            |
| ----------- | -------------------------------------- |
| ID          | ID for every different plan            |
| Name        | Name of Plan E.g Plan A,Plan B, Plan C |



**Plan Item**

| Field Names | Description                      |
| ----------- | -------------------------------- |
| Plan        | Plan ID                          |
| Item        | Habit Items for the Plan         |
| Name        | Name of the Habit                |
| StartTime   | Planned Start Time of the Habit  |
| EndTime     | Planned    End Time of the Habit |


**Overall Habit**


| Field Names              | Description                     |
| ------------------------ | ------------------------------- |
| Date                     |                                 |
| CompletedItems           | Number of Completed Habit Items |
| PlannedItems             | Number of Planned Habit Items   |
| Percentage of Completion | Completed/Planned * 100         |



Each Habit would need the following information

**Habit Items**



| Field Names | Description             |
| ----------- | ----------------------- |
| ID          | Habit Items             |
| Name        | Habit Name              |
| Date        | Date of the Habit       |
| StartTime   | Start Time of the Habit |
| EndTime     | End Time of the Habit   |
| IsCompleted | Is Habit Completed      |
| Comment     | Comment section         |




## Techology Stack

Golang : API Development

MySQL/Postgre : DB

UI : Yet to confirm

Docker & K8s 

Pipeline : GitHub Actions/ CircleCI



