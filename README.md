# uugnet

A BBS-style telnet server for UUGers.

![image](https://i.ibb.co/rd6MFVk/image.jpg)

### Telnet Commands

| Command          | Description                    |
| ---------------- | ------------------------------ |
| `help [command]` | Lists commands and their usage |
| `exit`           | Disconnect from uugnet.        |

## Getting Started

```bash
go build
./uugnet serve
```

### CLI Commands

| Command                     | Description            |
| --------------------------- | ---------------------- |
| `uugnet serve`              | Run the uugnet server. |
| `uugnet userlist`           | List users.            |
| `uugnet useradd <username>` | Add user.              |
