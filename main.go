package main

import (
	"context"
	"fmt"
	"url-shortener/database/redis"
	"url-shortener/utils"
)

func main() {
	ctx := context.Background()
	redis.ConnectToRedis(ctx)
	fmt.Print(utils.Base62Converter(300000055300))
}
