package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stat struct {
	RxBytes uint64
	TxBytes uint64
}

func getStat(ifName string) (*Stat, error) {

	file, err := os.Open("/proc/net/dev")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())

		if !strings.Contains(line, ":") {
			continue
		}

		fields := strings.Fields(line)
		interfaceName := strings.Split(fields[0], ":")[0]

		if interfaceName != ifName {
			continue
		}

		rxBytes, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			panic(err)
		}
		txBytes, err := strconv.ParseUint(fields[9], 10, 64)
		if err != nil {
			panic(err)
		}

		return &Stat{
			RxBytes: rxBytes,
			TxBytes: txBytes,
		}, nil
	}

	return nil, fmt.Errorf("network interface not found: %s", ifName)
}
