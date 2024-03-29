// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.
//go:build linux || darwin || freebsd || unix
// +build linux darwin freebsd unix

package signal

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	log "github.com/sirupsen/logrus"
)

// SetupSignalHandler setup signal handler for TiDB Server
func SetupSignalHandler(ignoreSighup bool, shudownFunc func(bool)) {
	usrDefSignalChan := make(chan os.Signal, 1)

	signal.Notify(usrDefSignalChan, syscall.SIGUSR1)
	go func() {
		buf := make([]byte, 1<<16)
		for {
			sig := <-usrDefSignalChan
			if sig == syscall.SIGUSR1 {
				stackLen := runtime.Stack(buf, true)
				log.Printf("\n=== Got signal [%s] to dump goroutine stack. ===\n%s\n=== Finished dumping goroutine stack. ===\n", sig, buf[:stackLen])
			}
		}
	}()

	closeSignalChan := make(chan os.Signal, 1)

	// 忽略信号 终端控制进程结束(终端连接断开)
	if ignoreSighup {
		signal.Ignore(syscall.SIGHUP)
		signal.Notify(closeSignalChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
	} else {
		signal.Notify(closeSignalChan,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT)
	}

	go func() {
		sig := <-closeSignalChan
		log.Infof("Got signal [%s] to exit.", sig)
		shudownFunc(sig == syscall.SIGQUIT)
	}()
}
