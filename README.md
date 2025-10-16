# go_task_tracker_cli

## Command Overview

### Display Tasks

| Command | Description |
|---------|-------------|
| `list` | Show all tasks regardless of status |
| `list [status]` | Show tasks with specific progress status<br/>**Status options:** `todo`, `in-progress`, `done` |

### Task Management

| Command | Description | Parameters |
|---------|-------------|------------|
| `add [task description]` | Add new task to the list | `task description` - Text description of the task |
| `update [task ID] [updated description]` | Update task description | `task ID` - Unique identifier of the task<br/>`updated description` - New task description |
| `delete [task ID]` | Remove task from the list | `task ID` - Unique identifier of the task to delete |

### Status Management

| Command | Description | Parameters |
|---------|-------------|------------|
| `mark-in-progress [task ID]` | Change task status to IN-PROGRESS | `task ID` - Unique identifier of the task |
| `mark-done [task ID]` | Change task status to DONE | `task ID` - Unique identifier of the task |

## Usage Examples

```bash
# Display all tasks
./go_task_tracker_cli list

# Display only todo tasks
./go_task_tracker_cli list todo

# Add a new task
./go_task_tracker_cli add "Complete project documentation"

./go_task_tracker_cli add "Release user authentication system"

./go_task_tracker_cli add "Create nanomachines son"

# Update an existing task
./go_task_tracker_cli update 2 "Finish user authentication system"

# Delete a task
./go_task_tracker_cli delete 1

# Change task status
./go_task_tracker_cli mark-in-progress 2
./go_task_tracker_cli mark-done 1