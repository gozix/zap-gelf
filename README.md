# GoZix Zap Gelf

## Dependencies

* [viper](https://github.com/gozix/viper)
* [zap](https://github.com/gozix/zap)

## Configuration example

```json
{
  "zap": {
    "cores": {
      "gelf": {
        "type": "gelf",
        "addr": "127.0.0.1:12001",
        "level": "debug"
      }
    }
  }
}
```
