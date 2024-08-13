<h1 align="center">To Go List CLI</h1>

<p align="center">
  A simple to do list CLI written in Go!
</p>

---

## About

This CLI application is designed to manage a simple task list directly from the terminal. 
The application enables users to create, manage, and track a to do list with a range of functionalities, all through straightforward terminal commands.

## Key Features

| Feature         | Description                                                                                                   |
|-----------------|---------------------------------------------------------------------------------------------------------------|
| **Create a To Do List** | Users can initialize a to do list by specifying a name. The application will create and store this list in a CSV file, saved in a dedicated folder called "ToGoLists" within the user's Documents directory. |
| **Add Tasks**   | Once a shopping list is created, users can add tasks to it. Each task is identified by a unique numeric ID, a title, and a creation date. The system will display relative timestamps like "created a few seconds ago." |
| **Manage Tasks** | Users have the ability to:                                                                                   |
| **Complete Tasks** | Mark tasks as completed.                                                                                     |
| **Uncomplete Tasks** | Revert completed tasks to their previous state.                                                             |
| **Remove Tasks** | Delete tasks from the list.                                                                                     |
| **View Tasks**  | Tasks can be displayed in a tabular format within the terminal, showing both completed and pending tasks.      |

## Requirements

### Functional

- [x] Ability to create and name a task list.
- [x] Add, remove, complete, and uncomplete tasks.
- [x] View tasks in a tabular format.

### Non-Functional

- [x] Data is stored in CSV format.
- [x] Files are saved in a "ToGoLists" folder.

### Business Rules

- [x] File names must be valid.
- [x] Tasks can only be managed if the file exists.
- [x] Tasks must be removed using their ID.
- [x] Only incomplete tasks can be completed, and only completed tasks can be uncompleted.

## Usage

First of all, go to Releases and download the executable of the project. After that move it to this path `COMMING SOON` and either restart your terminal or logout your session. After that you'll be able to run the commands with ease!

> Creates a new To Go tasklist at /home/{YOUR_NAME}/Documents/ToGoLists
```
togo create --name="My To Do List"
```

> Creates a new task in the task list created earlier
```
togo add --task="Walk the dog"
```

> Deletes the task with the providen ID
```
togo delete --id 1
```

> Completes the task with the providen ID
```
togo complete --id 1
```

> Incompletes the task with the providen ID
```
togo incomplete --id 1
```

> Lists all incompleted tasks
```
togo list
```

> Lists all tasks
```
togo list --all
```

> Lists all completed tasks
```
togo list --completed
```

## Contact me

- LinkedIn: https://www.linkedin.com/in/antonio-mauricio-4645832b3/
- Instagram: https://www.instagram.com/antonioalmeida2003/
- E-mail: antonioimportant@gmail.com
