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
	"strings"
	"syscall"
	"time"
)

const (
	minFanSpeed = 30  // Minimum fan speed, percent.
	maxFanSpeed = 100 // Maximum fan seed, percent.
	stopFanTemp = 45  // Temp to stop fan.
	maxFanTemp  = 75  // Temp to run fan on full speed.
)

func getCurrentTemp() int {
	out, execErr := exec.Command(
		"nvidia-smi", "--query-gpu=temperature.gpu", "--format=csv,noheader,nounits",
	).Output()
	if execErr != nil {
		log.Fatalf("Unable to get current temperature, error '%s'.\n", execErr)
	}
	raw := string(bytes.TrimSpace(out))
	result, err := strconv.ParseInt(raw, 10, 32)
	if err != nil {
		log.Fatalf("Invalid temperature value '%s', error '%s'.\n", raw, err)
	}
	return int(result)
}

func getFanSpeed(temp int) int {
	if temp < stopFanTemp {
		return 0
	}
	if temp > maxFanTemp {
		return maxFanSpeed
	}
	num := (temp - stopFanTemp) * (maxFanSpeed - minFanSpeed)
	base := minFanSpeed + num/minFanSpeed
	if num%minFanSpeed >= 15 {
		base++
	}
	return base
}

func setFanSpeed(speed int) {
	state := 1
	if speed == 0 {
		state = 0
	}
	if out, err := exec.Command(
		"nvidia-settings", "--no-config", "--no-write-config",
		fmt.Sprintf("--assign=GPUFanControlState=%d", state),
		fmt.Sprintf("--assign=GPUTargetFanSpeed=%d", speed),
	).CombinedOutput(); err != nil {
		log.Fatalf(
			"Unable to set fan speed, error '%s', out='%s'.\n",
			err, strings.TrimSpace(string(out)),
		)
	}
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for {
			temp := getCurrentTemp()
			speed := getFanSpeed(temp)
			setFanSpeed(speed)
			time.Sleep(15 * time.Second)
		}
	}()
	<-signals
}
