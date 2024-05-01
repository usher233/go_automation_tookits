package main

import (
    "fmt"
    "log"

    "github.com/shirou/gopsutil/disk"
)

func main() {
    // Get disk usage information
    diskUsage, err := disk.Usage("/")
    if err != nil {
        log.Fatalf("Error getting disk usage: %v", err)
    }

    // Calculate usage percentages
    usedPercent := int(diskUsage.UsedPercent)
    freePercent := 100 - usedPercent

    // Print disk usage information
    fmt.Println("Disk Usage Information:")
    fmt.Printf("Used: %d%%\n", usedPercent)
    fmt.Printf("Free: %d%%\n", freePercent)
	fmt.Printf("Total: %.2f GB\n", float64(diskUsage.Total)/1e9)
	fmt.Printf("Used: %.2f GB\n", float64(diskUsage.Used)/1e9)
	fmt.Printf("Free: %.2f GB\n", float64(diskUsage.Free)/1e9)
}