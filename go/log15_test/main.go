package main

import (
	"os"

	log "github.com/inconshreveable/log15"
)

func main() {
	logger := log.New("context", "key/value")

	logger.Info("Info msg", "first", 10, "second", 20, "third", 30)
	logger.Debug("Debug msg", "first", 10, "second", 20, "third", 30)
	logger.Warn("Warn msg", "first", 10, "second", 20, "third", 30)
	logger.Error("Error msg", "first", 10, "second", 20, "third", 30)
	logger.Crit("Crit msg", "first", 10, "second", 20, "third", 30)

	logger.SetHandler(log.MultiHandler(
		log.StreamHandler(os.Stderr, log.LogfmtFormat()),
		log.LvlFilterHandler(
			log.LvlError,
			log.Must.FileHandler("log.json", log.JsonFormat()))))

	logger.Info("Info msg", "first", 10, "second", 20, "third", 30)
	logger.Debug("Debug msg", "first", 10, "second", 20, "third", 30)
	logger.Warn("Warn msg", "first", 10, "second", 20, "third", 30)
	logger.Error("Error msg", "first", 10, "second", 20, "third", 30)
	logger.Crit("Crit msg", "first", 10, "second", 20, "third", 30)

}
