package cmd

import (
	"github.com/crerwin/distributedtranscoding/pkg/db"
	"github.com/crerwin/distributedtranscoding/pkg/dtc"
	"github.com/spf13/cobra"
)

var redisCmd = &cobra.Command{
	Use:   "redis",
	Short: "Interact with Redis",
}

var redisPingCmd = &cobra.Command{
	Use:   "ping",
	Short: "execute a PING against Redis",
	Run:   redisPingRun,
}

var redisInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the Redis database",
	Run:   redisInitRun,
}

var redisAddCmd = &cobra.Command{
	Use: "add",
	Run: redisAddRun,
}

func redisPingRun(cmd *cobra.Command, args []string) {
	redis := db.NewRedisClient()
	redis.Ping()
}

func redisInitRun(cmd *cobra.Command, args []string) {
	redis := db.NewRedisClient()
	redis.Initialize()
}

func redisAddRun(cmd *cobra.Command, args []string) {
	redis := db.NewRedisClient()
	i := new(dtc.Item)
	i.FileName = "A Cock and Bull Story (2005).mkv"
	i.Crop = "0:0:0:0"
	i.ForcedRate = "23.976"
	redis.AddToWorkQueue(i)
}
