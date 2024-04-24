# uugnet

A BBS-style telnet server for UUGulite.

<img src="https://i.ibb.co/rd6MFVk/image.jpg" alt="uugnet screenshot" width="400"/>

### Telnet Commands

| Command          | Description                    |
| ---------------- | ------------------------------ |
| `help [command]` | Lists commands and their usage |
| `exit`           | Disconnect from uugnet         |

If you would like to add a command, read [Telnet Commands](https://github.com/k4ben/uugnet/tree/master/internal/commands#telnet-commands) for more info.

## Getting Started

```bash
go build
./uugnet serve
```

### CLI Commands

| Command                     | Description           |
| --------------------------- | --------------------- |
| `uugnet serve`              | Run the uugnet server |
| `uugnet userlist`           | List users            |
| `uugnet useradd <username>` | Add user              |
