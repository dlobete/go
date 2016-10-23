// Copyright 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style license described in the
// LICENSE file.

package diag

import (
	"fmt"
	"time"

	"github.com/platinasystems/goes/optional/gpio"
)

func diagHost() {

	var pinstate bool
	var r string
	var d uint8
	fmt.Printf("\n%15s|%25s|%10s|%10s|%10s|%10s|%6s|%35s\n", "function", "parameter", "units", "value", "min", "max", "result", "description")
	fmt.Printf("---------------|-------------------------|----------|----------|----------|----------|------|-----------------------------------\n")

	/* diagTest: HOST_TO_BMC_INT_L, BMC_TO_HOST_INT_L, BMC_TO_HOST_NMI
	CPLD diag mode enables HOST_TO_BMC_INT_L = BMC_TO_HOST_INT_L | BMC_TO_HOST_NMI, toggle BMC driven signals and check host_to_bmc_int_l for correct state
	*/
	gpio.GpioSet("CPU_TO_MAIN_I2C_EN", true)
	gpio.GpioSet("BMC_TO_HOST_RST_L", false)
	diagI2cWriteOffsetByte(0x00, 0x77, 0x03, 0xFF)
	diagI2cWriteOffsetByte(0x00, 0x77, 0x07, 0xFE)

	gpio.GpioSet("BMC_TO_HOST_INT_L", false)
	gpio.GpioSet("BMC_TO_HOST_NMI", false)
	pinstate = gpio.GpioGet("HOST_TO_BMC_INT_L")
	if !pinstate {
		r = "pass"
	} else {
		r = "fail"
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "h2b_int=b2h_int|b2h_nmi", "-", pinstate, active_low_on_min, active_low_on_max, r, "h2b_int(0)=b2h_int(0)|b2h_nmi(0)")

	gpio.GpioSet("BMC_TO_HOST_INT_L", true)
	pinstate = gpio.GpioGet("HOST_TO_BMC_INT_L")
	if pinstate {
		r = "pass"
	} else {
		r = "fail"
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "h2b_int=b2h_int|b2h_nmi", "-", pinstate, active_low_off_min, active_low_off_max, r, "h2b_int(1)=b2h_int(1)|b2h_nmi(0)")
	gpio.GpioSet("BMC_TO_HOST_INT_L", false)

	gpio.GpioSet("BMC_TO_HOST_NMI", true)
	pinstate = gpio.GpioGet("HOST_TO_BMC_INT_L")
	if pinstate {
		r = "pass"
	} else {
		r = "fail"
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "h2b_int=b2h_int|b2h_nmi", "-", pinstate, active_low_off_min, active_low_off_max, r, "h2b_int(1)=b2h_int(0)|b2h_nmi(1)")
	gpio.GpioSet("BMC_TO_HOST_INT_L", true)

	diagI2cWriteOffsetByte(0x00, 0x77, 0x07, 0xFF)
	gpio.GpioSet("CPU_TO_MAIN_I2C_EN", false)
	gpio.GpioSet("BMC_TO_HOST_RST_L", true)

	/* diagTest: BMC_TO_HOST_RST_L
	Reset host and validate host behaves accordingly
	*/
	gpio.GpioSet("CPU_TO_MAIN_I2C_EN", true)
	gpio.GpioSet("BMC_TO_HOST_RST_L", false)
	_, d = diagI2cPing(0x00, 0x77, 0x01, 1)
	if (d & 0x10) == 0 {
		r = "pass"
		pinstate = false
	} else {
		r = "fail"
		pinstate = true
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "bmc_to_host_rst_l_on", "-", pinstate, active_low_on_min, active_low_on_max, r, "rst on, check rst=0 on cpu card")

	gpio.GpioSet("BMC_TO_HOST_RST_L", true)
	_, d = diagI2cPing(0x00, 0x77, 0x01, 1)
	if (d & 0x10) == 0x10 {
		r = "pass"
		pinstate = true
	} else {
		r = "fail"
		pinstate = false
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "bmc_to_host_rst_l_off", "-", pinstate, active_low_off_min, active_low_off_max, r, "rst off, check rst=1 on cpu card")
	gpio.GpioSet("CPU_TO_MAIN_I2C_EN", false)

	/* diagTest: HOST_TO_BMC_I2C_GPIO
	toggle HOST_TO_BMC_I2C_GPIO and validate bmc can detect the proper signal states
	*/
	gpio.GpioSet("CPU_TO_MAIN_I2C_EN", true)
	time.Sleep(50 * time.Millisecond)
	diagI2cWriteOffsetByte(0x00, 0x74, 0x02, 0xF7)
	diagI2cWriteOffsetByte(0x00, 0x74, 0x06, 0xF7)
	pinstate = gpio.GpioGet("HOST_TO_BMC_I2C_GPIO")
	if !pinstate {
		r = "pass"
	} else {
		r = "fail"
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "host_to_bmc_i2c_gpio_low", "-", pinstate, active_high_off_min, active_high_off_max, r, "check gpio is high")

	diagI2cWriteOffsetByte(0x00, 0x74, 0x02, 0xFF)
	pinstate = gpio.GpioGet("HOST_TO_BMC_I2C_GPIO")
	if pinstate {
		r = "pass"
	} else {
		r = "fail"
	}
	fmt.Printf("%15s|%25s|%10s|%10t|%10t|%10t|%6s|%35s\n", "host", "host_to_bmc_i2c_gpio_high", "-", pinstate, active_high_on_min, active_high_on_max, r, "check gpio is low")
	gpio.GpioSet("CPU_TO_MAIN_I2C_EN", false)

}
