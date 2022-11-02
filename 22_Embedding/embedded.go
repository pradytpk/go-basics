package main

import "fmt"

const (
	Small = iota
	Medium
	Large
)

const (
	Ground = iota
	Air
)

type BeltSize int
type Shipping int

func (b BeltSize) String() string {
	return []string{"Small", "Medium", "Large"}[b]
}

func (s Shipping) String() string {
	return []string{"Ground", "Air"}[s]
}

type Conveyor interface {
	Convey() BeltSize
}

type Shipper interface {
	Ship() Shipping
}

type WarehouseAutomator interface {
	Conveyor
	Shipper
}

type SpamMail struct {
	amount int
}

func (s SpamMail) String() string {
	return "Spam mail"
}

func (s *SpamMail) Ship() Shipping {
	return Air
}

func (s *SpamMail) Convey() BeltSize {
	return Small
}

func automate(item WarehouseAutomator) {
	fmt.Printf("Convey %v on %v conveyor\n", item, item.Convey())
	fmt.Printf("Ship %v via %v\n", item, item.Ship())
}

type ToxicWaste struct {
	weight int
}

func (t *ToxicWaste) Ship() Shipping {
	return Ground
}

//--Summary:
//  Create a system monitoring dashboard using the existing dashboard component structures. Each array element in the components represent a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) Averagebandwidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}
func (b *CpuTemp) AverageCpuTemp() int {
	sum := 0
	for i := 0; i < len(b.temp); i++ {
		sum += int(b.temp[i])
	}
	return sum / len(b.temp)
}
func (b *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}

//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	mail := SpamMail{40000}
	automate(&mail)
	// Using embedded interfaces to ensure only items that
	// implement both Conveyor and Shipper can be automated
	// automate(&ToxicWaste{300}) // Won't work!
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	dash := Dashboard{
		BandwidthUsage: bandwidth,
		MemoryUsage:    memory,
		CpuTemp:        temp,
	}
	//* Print out a 5-second average from each component using promoted
	//  methods on the Dashboard
	fmt.Printf("Average bandwidth usage:%v\n", dash.Averagebandwidth())
	fmt.Printf("Average CpuTemp usage:%v\n", dash.AverageCpuTemp())
	fmt.Printf("Average memory usage:%v\n", dash.AverageMemoryUsage())
}
