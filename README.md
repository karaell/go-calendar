# 📅 Calendar Application

A calendar application for managing events and reminders for added events

## 🚀 Features

- **Event Management**: create, edit, remove events
- **Priority System**: low, medium, high priority
- **Smart Reminders**: configurable notifications with time intervals
- **Data Persistence**: automatic saving in JSON or ZIP format
- **Interactive CLI**: command hints and convenient interface
- **Logging**: detailed logging of all operations
- **Command History**: saving and restoring input history

## 📦 Installation and Setup

### Prerequisites

- Go 1.24.5
- Git

### Installation

```bash
# Clone repository
git clone <repository-url>
cd go-calendar

# Install dependencies
go mod download

# Run application
go run main.go
```

## 📋 Usage

### Main Commands

#### Adding an Event
```bash
add "Event Title" "date and time" "priority"
```

**Examples:**
```bash
add "Team Meeting" "2024-12-25 14:00" "low"
add "Project Presentation" "2024-12-26 10:00" "medium"
add "Project Deadline" "2024-12-24 18:00" "high"
```

#### Viewing Events
```bash
list
```

#### Updating an Event
```bash
update "event_id" "new_title" "new_date" "new_priority"
```

**Example:**
```bash
update "abc123" "Updated Title" "2024-12-27 15:00" "high"
```

#### Deleting an Event
```bash
remove "event_id"
```

#### Managing Reminders
```bash
# Add reminder
add_event_reminder "event_id" "message" "reminder_time"

# Remove reminder
remove_event_reminder "event_id"
```

**Example:**
```bash
add_event_reminder "abc123" "Don't forget about the meeting!" "2024-12-25 13:45"
```

#### Other Commands
```bash
help    # Show help
log     # Show input history
exit    # Exit application
```

### Data Formats

#### Priorities
- `low` - low
- `medium` - medium
- `high` - high

#### Dates and Times
Various formats are supported:
- `2024-12-25 14:00`
- `2024-12-25T14:00:00`
- `Dec 25, 2024 2:00 PM`
- `25/12/2024 14:00`

#### Event Titles
- Length: 3-50 characters
- Allowed characters: letters, numbers, spaces, commas, hyphens, periods

## 🔧 Configuration

### Data Files

The application creates the following files:

- `calendar.json` - main calendar file with events
- `calendar_log.json` - command history
- `app.log` - application log

### Switching Storage Type

In `main.go` you can change the storage type:

```go
// JSON storage (default)
s := storage.CreateJsonStorage("calendar.json")

// ZIP storage
s := storage.CreateZipStorage("calendar.zip")
```

## 🧪 Testing

### Running Tests

```bash
# Tests for all packages
go test ./...
```

## 📊 Logging

The application maintains detailed logging of all operations:

- **INFO** - successful operations
- **WARNING** - warnings
- **ERROR** - errors

Logs are saved to the `app.log` file with timestamps and context.

## 🔒 Security

- Validation of all input data
- Safe file system operations
- Protection against incorrect dates
- Data integrity verification

## 🐛 Debugging

### Enabling Detailed Logging

Logs contain detailed information about all operations, which helps with debugging.

### Checking Data Files

If you encounter issues, you can check the contents of files:
- `calendar.json` - event structure
- `calendar_log.json` - command history
- `app.log` - system logs

---

**Note**: This is a feature-rich calendar application designed for managing events and reminders. All data is automatically saved and restored when restarting.
