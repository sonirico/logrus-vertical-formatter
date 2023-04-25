# Vertical formatter for logrus logger

## Example

```go

func main() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(
		logrusverticalformatter.New(
			logrusverticalformatter.WithPadRight(" ", 40),
		),
	)

	logger := log.WithField("service", "orders")

	logger.Infoln("order placed successfully")
}


```

## Output

```
ts                                      2023-04-26 00:01:43.559701525 +0200 CEST m=+0.000046735
service                                 orders
message                                 order placed successfully
____________________________________________________________________________________________________


```