<h1 align="center">✅ To Go List CLI</h1>

<img src="https://github.com/Dedo-Finger2/togo-cli/blob/master/public/images/cover.png?raw=true" />

<p>
  <img src="https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff" />
  <img src="https://badgen.net/badge/icon/terminal?icon=terminal&label" />
  <img src="https://badgen.net/badge/icon/windows?icon=windows&label" />
  <img src="https://img.shields.io/badge/License-MIT-blue.svg" />
</p>

<h3 align="center">✅ A <strong>simple</strong> to do list <strong>CLI</strong> written in <strong>Go</strong>! ✅</h3> 

---

## 📝 About

This CLI application is designed to manage a simple task list directly from the terminal. 
The application enables users to create, manage, and track a to do list with a range of functionalities, all through straightforward terminal commands.

## 🔑 Key Features

| Feature         | Description                                                                                                   |
|-----------------|---------------------------------------------------------------------------------------------------------------|
| **Create a To Do List** | Users can initialize a to do list by specifying a name. The application will create and store this list in a CSV file, saved in a dedicated folder called "ToGoLists" within the user's Documents directory. |
| **Add Tasks**   | Once a shopping list is created, users can add tasks to it. Each task is identified by a unique numeric ID, a title, and a creation date. The system will display relative timestamps like "created a few seconds ago." |
| **Manage Tasks** | Users have the ability to:                                                                                   |
| **Complete Tasks** | Mark tasks as completed.                                                                                     |
| **Uncomplete Tasks** | Revert completed tasks to their previous state.                                                             |
| **Remove Tasks** | Delete tasks from the list.                                                                                     |
| **View Tasks**  | Tasks can be displayed in a tabular format within the terminal, showing both completed and pending tasks.      |

## 🌳 Requirements

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

## ❓ Build it from source

```bash
git clone https://github.com/Dedo-Finger2/togo-cli.git
```

```bash
cd togo-cli
```

```bash
make build-linux OR make build-windows
```

Follow the next instrutions in the Usage section.

## 🔨 Usage

### Linux

After dowloading the binary (or after building it from source) use the following command.

```bash
sudo mv togo /usr/local/bin/
```

### Windows

After dowloading the binary (or after building it from source) move the `togo.exe` file into any folder that is already in the PATH of your machine.
Usually the paths: `C:\Program Files\Togo\` or `C:\Users\<YOUR_USER>\AppData\Local\Programs\Togo\` are okay to use.

### Commands

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

## 📱 Contact me

- LinkedIn: https://www.linkedin.com/in/antonio-mauricio-4645832b3/
- Instagram: https://www.instagram.com/antonioalmeida2003/
- E-mail: antonioimportant@gmail.com
