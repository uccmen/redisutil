# Example usage of redisutil

```go
package main

var redisInstance *util.RedisInstance

func init() {
  redisInstance = redisutil.NewRedis()
}

func isThingieExists(thingie string) bool {
  conn := redisInstance.DB().Get()
  defer conn.Close()

  thingieKey := "thingie-hash:" + thingie

  out, err := conn.Do("GET", thingieKey)
  if err != nil {
    return false
  }

  if out != nil {
    return true
  }

  return false
}

```
