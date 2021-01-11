# GoScp-cli

## Run

For run util after pull repository:

```shell
go run goScp-cli.go
```

## Build

For build project to binary file:

```shell
go build goScp-cli.go
./goScp-cli
```

## Help

Sub-commands of utility:

```shell
COMMANDS:
   init, i    Initialised structure for work
   login, l   Login to your account
   pull, p    complete a task on the list
   upload, u  options for task templates
   score, s   options for task templates
   help, h    Shows a list of commands or help for one command
```

### Init

Use this command immediately after downloading this utility. Here you can auto initialize your environment on Linux OS

```shell
goScp-cli.go init -h

OPTIONS:
   --dir value, -d value   Root dir for program (default: /home/ivan) [$HOME]
   --config FILE, -c FILE  Choose FILE to upload (default: .scp.yaml)
   --help, -h              show help (default: false)

```

### Login

For auth in system.

```shell
goScp-cli.go login -h

OPTIONS:
   --username id, -u id      choose task id for upload
   --password FILE, -p FILE  Choose FILE to upload
   --help, -h                show help (default: false)

```

### Pull

This command is needed to download files and other useful information about the task

```shell
goScp-cli.go pull -h

OPTIONS:
   --task id, -t id  choose task id for upload
   --help, -h        show help (default: false)

```

### Upload

Here you can upload your personal solution of task

```shell
goScp-cli.go upload -h 

OPTIONS:
   --task id, -t id      choose task id for upload
   --file FILE, -f FILE  Choose FILE to upload
   --help, -h            show help (default: false)

```

### Score

Here you can see your score like this

```shell
goScp-cli.go score

Username:   user
Score:      20
Challenges: 1/23
Place:      1/81

```

## Testing

```shell
LOGIN=<Login> PASSWORD=<Password> go test ./...
```

P.S. Some testing modules are not yet available