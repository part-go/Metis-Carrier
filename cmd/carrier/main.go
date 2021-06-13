package main

import (
	"fmt"
	"github.com/RosettaFlow/Carrier-Go/cmd"
	dbcommand "github.com/RosettaFlow/Carrier-Go/cmd/carrier/db"
	"github.com/RosettaFlow/Carrier-Go/cmd/common"
	"github.com/RosettaFlow/Carrier-Go/common/flags"
	"github.com/RosettaFlow/Carrier-Go/common/logutil"
	"github.com/RosettaFlow/Carrier-Go/node"
	golog "github.com/ipfs/go-log/v2"
	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	runtimeDebug "runtime/debug"

	"github.com/urfave/cli/v2"
)

var (
	appFlags = []cli.Flag{
		flags.RPCHost,
		flags.RPCPort,
		flags.GRPCGatewayHost,
		flags.GRPCGatewayPort,
		// todo: more flags could be define here.
	}
)

func init() {
	appFlags = cmd.WrapFlags(appFlags)
}

func main() {
	app := cli.App{}
	app.Name = "carrier"
	app.Usage = "this is a carrier network implementation for Carrier Node"
	// set action func.
	app.Action = startNode
	app.Version = common.Version()
	app.Commands = []*cli.Command {
		dbcommand.Commands,
	}
	app.Flags = appFlags

	app.Before = func(ctx *cli.Context) error {
		// Load flags from config file, if specified.
		if err := flags.LoadFlagsFromConfig(ctx, app.Flags); err != nil {
			return err
		}

		format := ctx.String(flags.LogFormat.Name)
		if format == "" {
			format = "text"
		}
		switch format {
		case "text":
			formatter := new(prefixed.TextFormatter)
			formatter.TimestampFormat = "2006-01-02 15:04:05"
			formatter.FullTimestamp = true
			// If persistent log files are written - we disable the log messages coloring because
			// the colors are ANSI codes and seen as gibberish in the log files.
			formatter.DisableColors = ctx.String(flags.LogFileName.Name) != ""
			logrus.SetFormatter(formatter)
		case "fluentd":
			f := joonix.NewFormatter()
			if err := joonix.DisableTimestampFormat(f); err != nil {
				panic(err)
			}
			logrus.SetFormatter(f)
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		default:
			return fmt.Errorf("unknown log format %s", format)
		}

		logFileName := ctx.String(flags.LogFileName.Name)
		if logFileName != "" {
			if err := logutil.ConfigurePersistentLogging(logFileName); err != nil {
				log.WithError(err).Error("Failed to configuring logging to disk.")
			}
		}
		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			log.Errorf("Runtime panic: %v\n%v", x, string(runtimeDebug.Stack()))
			panic(x)
		}
	}()

	if err := app.Run(os.Args); err != nil {
		log.Error(err.Error())
	}
}

func startNode(ctx *cli.Context) error {
	// todo: some logic could be added here
	if args := ctx.Args(); args.Len() > 0 {
		return fmt.Errorf("invalid command: %q", args.Get(0))
	}

	// setting log level.
	verbosity := ctx.String(flags.VerbosityFlag.Name)
	if verbosity == "" {
		verbosity = flags.VerbosityFlag.Value
	}
	level, err := logrus.ParseLevel(verbosity)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	if level == logrus.TraceLevel {
		// libp2p specific logging.（special）
		golog.SetAllLoggers(golog.LevelDebug)
	}

	// initial no and start.
	node, err := node.New(ctx)
	if err != nil {
		return err
	}
	node.Start()
	return nil
}
