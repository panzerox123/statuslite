package main

import (
	"github.com/zcalusic/sysinfo"
)

type Status struct {
	cpuCount   uint
	cpuCores   uint
	cpuThreads uint
	cpuFreq    uint
	cpuVendor  string
	cpuModel   string
}

type NetworkStatus struct {
}

type HardwareSpecs struct {
}

type SoftwareSpecs struct {
}

type Users struct {
}

func NewStatus() Status {
	var s Status
	var si sysinfo.SysInfo

	si.GetSysInfo()

	s.cpuCount = si.CPU.Cpus
	s.cpuCores = si.CPU.Cores
	s.cpuThreads = si.CPU.Threads
	s.cpuFreq = si.CPU.Speed
	s.cpuModel = si.CPU.Model
	s.cpuVendor = si.CPU.Vendor

	return s
}

func (s *Status) ReturnHeartBeat(args *int, heartbeat *bool) error {
	*heartbeat = true
	return nil
}

func (s *Status) ReturnNetworkStatus(args *int, reply *NetworkStatus) error {
	return nil
}

func (s *Status) ReturnHardwareSpecifications(args *int, reply *HardwareSpecs) error {
	return nil
}

func (s *Status) ReturnSoftwareSpecifications(args *int, reply *SoftwareSpecs) error {
	return nil
}

func (s *Status) ReturnUsers(args *int, reply *Users) error {
	return nil
}
