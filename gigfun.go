/*
	gigfun - Kludge for a Gigabyte videocard.

	Copyright (C) 2022 Vadim Kuznetsov <vimusov@gmail.com>

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	StartCoolingTemp = 70
	TempProbesCount  = 5
)

func getCurrentTemp() uint64 {
	out, err := exec.Command(
		"nvidia-smi",
		"--query-gpu=temperature.gpu",
		"--format=csv,noheader,nounits",
	).Output()
	if err != nil {
		log.Fatalf("Unable to get current temperature, error '%s'.\n", err)
	}
	raw := string(bytes.Trim(out, "\n"))
	result, err := strconv.ParseUint(raw, 10, 32)
	if err != nil {
		log.Fatalf("Invalid temperature value '%s', error '%s'.\n", raw, err)
	}
	return result
}

func getAvgTemp() uint64 {
	var avgTemp uint64 = 0
	for i := 1; i <= TempProbesCount; i++ {
		avgTemp += getCurrentTemp()
		time.Sleep(1 * time.Second)
	}
	return avgTemp / TempProbesCount
}

func setFanSpeed(isManual uint8) {
	if out, err := exec.Command(
		"nvidia-settings",
		"--no-config", "--no-write-config",
		fmt.Sprintf("--assign=GPUFanControlState=%d", isManual),
		"--assign=GPUTargetFanSpeed=100",
	).CombinedOutput(); err != nil {
		log.Fatalf("Unable to set fan speed, error '%s', stdout='%s'.\n", err, out)
	}
}

func regulateTemp() {
	wasStarted := false
	wasStopped := false
	for {
		temp := getAvgTemp()
		if temp <= StartCoolingTemp {
			if !wasStopped {
				setFanSpeed(0)
				wasStopped = true
			}
			wasStarted = false
			continue
		}
		if !wasStarted {
			setFanSpeed(1)
			wasStarted = true
			wasStopped = false
		}
	}
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go regulateTemp()
	<-signals
}
