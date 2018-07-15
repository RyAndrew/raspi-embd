package main

import (
	"fmt"

	"github.com/kidoman/embd"
)

//constants from https://github.com/adafruit/Adafruit_ADS1X15/blob/master/Adafruit_ADS1015.cpp

const ADS1015_I2C_ADDRESS byte = (0x48)

const ADS1015_REG_POINTER_CONVERT byte = (0x00)
const ADS1015_REG_POINTER_CONFIG byte = (0x01)

const ADS1015_REG_CONFIG_OS_MASK uint16 = (0x8000)
const ADS1015_REG_CONFIG_OS_SINGLE uint16 = (0x8000)  // Write: Set to start a single-conversion
const ADS1015_REG_CONFIG_OS_BUSY uint16 = (0x0000)    // Read: Bit = 0 when conversion is in progress
const ADS1015_REG_CONFIG_OS_NOTBUSY uint16 = (0x8000) // Read: Bit = 1 when device is not performing a conversion

const ADS1015_REG_CONFIG_MUX_MASK uint16 = (0x7000)
const ADS1015_REG_CONFIG_MUX_DIFF_0_1 uint16 = (0x0000) // Differential P = AIN0, N = AIN1 (default)
const ADS1015_REG_CONFIG_MUX_DIFF_0_3 uint16 = (0x1000) // Differential P = AIN0, N = AIN3
const ADS1015_REG_CONFIG_MUX_DIFF_1_3 uint16 = (0x2000) // Differential P = AIN1, N = AIN3
const ADS1015_REG_CONFIG_MUX_DIFF_2_3 uint16 = (0x3000) // Differential P = AIN2, N = AIN3
const ADS1015_REG_CONFIG_MUX_SINGLE_0 uint16 = (0x4000) // Single-ended AIN0
const ADS1015_REG_CONFIG_MUX_SINGLE_1 uint16 = (0x5000) // Single-ended AIN1
const ADS1015_REG_CONFIG_MUX_SINGLE_2 uint16 = (0x6000) // Single-ended AIN2
const ADS1015_REG_CONFIG_MUX_SINGLE_3 uint16 = (0x7000) // Single-ended AIN3

const ADS1015_REG_CONFIG_PGA_MASK uint16 = (0x0E00)
const ADS1015_REG_CONFIG_PGA_6_144V uint16 = (0x0000) // +/-6.144V range = Gain 2/3
const ADS1015_REG_CONFIG_PGA_4_096V uint16 = (0x0200) // +/-4.096V range = Gain 1
const ADS1015_REG_CONFIG_PGA_2_048V uint16 = (0x0400) // +/-2.048V range = Gain 2 (default)
const ADS1015_REG_CONFIG_PGA_1_024V uint16 = (0x0600) // +/-1.024V range = Gain 4
const ADS1015_REG_CONFIG_PGA_0_512V uint16 = (0x0800) // +/-0.512V range = Gain 8
const ADS1015_REG_CONFIG_PGA_0_256V uint16 = (0x0A00) // +/-0.256V range = Gain 16

const ADS1015_REG_CONFIG_MODE_MASK uint16 = (0x0100)
const ADS1015_REG_CONFIG_MODE_CONTIN uint16 = (0x0000) // Continuous conversion mode
const ADS1015_REG_CONFIG_MODE_SINGLE uint16 = (0x0100) // Power-down single-shot mode (default)

const ADS1015_REG_CONFIG_DR_MASK uint16 = (0x00E0)
const ADS1015_REG_CONFIG_DR_128SPS uint16 = (0x0000)  // 128 samples per second
const ADS1015_REG_CONFIG_DR_250SPS uint16 = (0x0020)  // 250 samples per second
const ADS1015_REG_CONFIG_DR_490SPS uint16 = (0x0040)  // 490 samples per second
const ADS1015_REG_CONFIG_DR_920SPS uint16 = (0x0060)  // 920 samples per second
const ADS1015_REG_CONFIG_DR_1600SPS uint16 = (0x0080) // 1600 samples per second (default)
const ADS1015_REG_CONFIG_DR_2400SPS uint16 = (0x00A0) // 2400 samples per second
const ADS1015_REG_CONFIG_DR_3300SPS uint16 = (0x00C0) // 3300 samples per second

const ADS1015_REG_CONFIG_CMODE_MASK uint16 = (0x0010)
const ADS1015_REG_CONFIG_CMODE_TRAD uint16 = (0x0000)   // Traditional comparator with hysteresis (default)
const ADS1015_REG_CONFIG_CMODE_WINDOW uint16 = (0x0010) // Window comparator

const ADS1015_REG_CONFIG_CPOL_MASK uint16 = (0x0008)
const ADS1015_REG_CONFIG_CPOL_ACTVLOW uint16 = (0x0000) // ALERT/RDY pin is low when active (default)
const ADS1015_REG_CONFIG_CPOL_ACTVHI uint16 = (0x0008)  // ALERT/RDY pin is high when active

const ADS1015_REG_CONFIG_CLAT_MASK uint16 = (0x0004)   // Determines if ALERT/RDY pin latches once asserted
const ADS1015_REG_CONFIG_CLAT_NONLAT uint16 = (0x0000) // Non-latching comparator (default)
const ADS1015_REG_CONFIG_CLAT_LATCH uint16 = (0x0004)  // Latching comparator

const ADS1015_REG_CONFIG_CQUE_MASK uint16 = (0x0003)
const ADS1015_REG_CONFIG_CQUE_1CONV uint16 = (0x0000) // Assert ALERT/RDY after one conversions
const ADS1015_REG_CONFIG_CQUE_2CONV uint16 = (0x0001) // Assert ALERT/RDY after two conversions
const ADS1015_REG_CONFIG_CQUE_4CONV uint16 = (0x0002) // Assert ALERT/RDY after four conversions
const ADS1015_REG_CONFIG_CQUE_NONE uint16 = (0x0003)  // Disable the comparator and put ALERT/RDY in high state (default)

func initAdc(bus embd.I2CBus) {

	var config uint16
	config = ADS1015_REG_CONFIG_CQUE_NONE | // Disable the comparator (default val)
		ADS1015_REG_CONFIG_CLAT_NONLAT | // Non-latching (default val)
		ADS1015_REG_CONFIG_CPOL_ACTVLOW | // Alert/Rdy active low   (default val)
		ADS1015_REG_CONFIG_CMODE_TRAD | // Traditional comparator (default val)
		ADS1015_REG_CONFIG_DR_1600SPS // 1600 samples per second (default)

	//config |= ADS1015_REG_CONFIG_PGA_6_144V
	//ch1: 11101101110000
	//ch1: 1110110111
	//ch1: 951

	config |= ADS1015_REG_CONFIG_PGA_4_096V
	//ch1: 101100100000000
	//ch1: 22784

	config |= ADS1015_REG_CONFIG_MUX_SINGLE_0 //read from channel 0

	//config |= ADS1015_REG_CONFIG_OS_SINGLE //wake from sleep start a conversion
	config |= ADS1015_REG_CONFIG_MODE_CONTIN // Single-shot mode (default)

	err := bus.WriteWordToReg(ADS1015_I2C_ADDRESS, ADS1015_REG_POINTER_CONFIG, config)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}
func readAdcValue(bus embd.I2CBus) uint16 {

	var result uint16
	var err error
	result, err = bus.ReadWordFromReg(ADS1015_I2C_ADDRESS, ADS1015_REG_POINTER_CONVERT)
	if err != nil {
		fmt.Printf("%s", err)
		return 0
	}
	//fmt.Printf("ch1: %b\n", result)
	//shift 4 bits to read the upper 12 bits
	result = result >> 4
	fmt.Printf("ADC ch1: %#x\n", result)

	return result
	// fmt.Printf("ch1: %#x\n", result)
	// fmt.Printf("ch1: %b\n", result)
	// fmt.Printf("ch1: %d\n---\n", result)
}
