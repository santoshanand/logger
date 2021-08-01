# Go Logger

Install dependancy 

`go get github.com/santoshanand/logger`

Initialize console logger

```golang
logger.InitLogger(false)
```

Initialize file logger
```golang
logger.InitLogger(true)
```

Above line will create a `log.log` file and will add all the logs

function to use logger

`logger.Log("Welcome simple log (info) logger")`
`logger.Info("Welcome info logger")`