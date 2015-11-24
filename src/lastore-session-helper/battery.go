package main

import "dbus/com/deepin/daemon/power"
import log "github.com/cihub/seelog"

func (l *Lastore) MonitorBattery() {
	m, err := power.NewPower("com.deepin.daemon.Power", "/com/deepin/daemon/Power")
	if err != nil {
		log.Warnf("Failed MonitorBattery: %v\n", err)
	}

	m.BatteryPercentage.ConnectChanged(func() {
		percentage := m.BatteryPercentage.Get()
		log.Infof("Current battery percentage:%v (notified:%v)\n", percentage, l.notifiedBattery)
		if l.notifiedBattery {
			return
		}
		l.notifiedBattery = true

		if m.BatteryPercentage.Get() < 30.0 && l.SystemOnChanging {
			NotifyLowPower()
		}
	})
}
