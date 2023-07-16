# Portal TRO #

## API Contract

You can find api contract doc. in [here](https://github.com/Nashrullah136/portal-tro/tree/main/app)

## Run
- Compile the applicaation first
```shell
go build
```
- To run application use this command
```shell
crm debug
```
- To specified the env location, you can use --env flag. example command
```shell
crm --env "C:\.env" debug
```
## List of avaliable command
```
  debug - to run and print the log to command prompt
  install - to install application as service
  remove - to uninstall application service
  start - to start the service
  stop - to stop the service
```
