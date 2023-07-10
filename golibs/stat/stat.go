package stat

import (
	"os"
	"runtime"
	"time"

	"cqserver/golibs/util"
	client "github.com/influxdata/influxdb/client/v2"
	"github.com/shirou/gopsutil/process"
)

type StatConfig struct {
	Host          string                        // 主机名，用于提交时的tag
	ServerType    string                        // 服务器类型，如gs1，用于提交时的tag
	InfluxDbName  string                        // 数据库名字，注：目前udp提交情况下该参数无效
	InfluxDbAddr  string                        // 数据库地址
	AppInfoGetter func() map[string]interface{} // 获取应用层数据的接口
}

type Stat struct {
	util.DefaultModule

	tags   map[string]string
	conf   *StatConfig
	client client.Client
	done   chan struct{}
}

func NewStat(conf *StatConfig) *Stat {
	s := &Stat{
		conf: conf,
		tags: map[string]string{"host": conf.Host, "serverType": conf.ServerType},
		done: make(chan struct{}),
	}
	go s.worker()
	return s
}

func (this *Stat) Stop() {
	close(this.done)
}

func (this *Stat) getSysInfo() map[string]interface{} {
	var info = make(map[string]interface{})
	proc, _ := process.NewProcess(int32(os.Getpid()))
	if cpuInfo, err := proc.Times(); err == nil {
		info["cpu_num"] = runtime.NumCPU()
		info["cpu_user"] = cpuInfo.User
		info["cpu_sys"] = cpuInfo.System
		info["cpu_idle"] = cpuInfo.Idle
	}
	if memInfo, err := proc.MemoryInfo(); err == nil {
		info["mem_rss"] = int(memInfo.RSS) / (1024 * 1024)
		info["mem_vms"] = int(memInfo.VMS) / (1024 * 1024)
	}
	return info
}

func (this *Stat) addSysInfo(bp client.BatchPoints) {
	sysInfo := this.getSysInfo()
	if len(sysInfo) == 0 {
		return
	}
	pt, err := client.NewPoint("sysinfo", this.tags, sysInfo, time.Now())
	if err != nil {
		return
	}
	bp.AddPoint(pt)
}

func (this *Stat) addAppInfo(bp client.BatchPoints) {
	if this.conf.AppInfoGetter == nil {
		return
	}
	infos := this.conf.AppInfoGetter()
	if len(infos) == 0 {
		return
	}
	pt, err := client.NewPoint("appinfo", this.tags, infos, time.Now())
	if err != nil {
		return
	}
	bp.AddPoint(pt)
}

func (this *Stat) getBatchPoints() client.BatchPoints {
	bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  this.conf.InfluxDbName,
		Precision: "s",
	})
	this.addSysInfo(bp)
	this.addAppInfo(bp)
	return bp
}

func (this *Stat) worker() {
	ticker := time.NewTicker(time.Second)
loop:
	for {
		select {
		case <-ticker.C:
			if this.client == nil {
				if c, err := client.NewUDPClient(client.UDPConfig{Addr: this.conf.InfluxDbAddr}); err == nil {
					this.client = c
				} else {
					continue
				}
			}
			bp := this.getBatchPoints()
			if err := this.client.Write(bp); err != nil {
				this.client = nil
			}
		case <-this.done:
			break loop
		}
	}
	ticker.Stop()
}
