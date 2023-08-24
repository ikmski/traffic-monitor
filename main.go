package main

import (
	"fmt"
	"os"
	"time"
)

const (
	SIZE_KB uint64 = 1024
	SIZE_MB uint64 = 1024 * 1024
	SIZE_GB uint64 = 1024 * 1024 * 1024
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("interface name is required")
		os.Exit(1)
	}

	iface := os.Args[1]

	err := run(iface)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(iface string) error {

	ticker := time.NewTicker(1 * time.Second)

	var prev *Stat

	for range ticker.C {

		current, err := getStat(iface)
		if err != nil {
			return err
		}

		if prev == nil {
			prev = current
		}

		rxRate := (current.RxBytes - prev.RxBytes) * 8 // bps 単位に変換
		txRate := (current.TxBytes - prev.TxBytes) * 8 // bps 単位に変換

		output(iface, rxRate, txRate)

		prev = current
	}

	return nil
}

func output(iface string, rxRate uint64, txRate uint64) {

	// カーソルを一番上に移動するための制御シーケンス
	fmt.Print("\033[H\033[2J")

	fmt.Printf("Interface: %s\n", iface)
	fmt.Printf("  Rx: %s\n", rateStr(rxRate))
	fmt.Printf("  Tx: %s\n", rateStr(txRate))
}

func rateStr(rate uint64) string {

	if rate < SIZE_KB {
		return fmt.Sprintf("%d [bps]", rate)
	}

	if rate < SIZE_MB {
		s := float64(rate) / float64(SIZE_KB)
		return fmt.Sprintf("%.1f [Kbps]", s)
	}

	if rate < SIZE_GB {
		s := float64(rate) / float64(SIZE_MB)
		return fmt.Sprintf("%.1f [Mbps]", s)
	}

	s := float64(rate) / float64(SIZE_GB)
	return fmt.Sprintf("%.1f [Gbps]", s)
}
