package scanner

import (
	"bytes"
	"errors"
	"fmt"
	"sync"

	"time"

	"github.com/zserge/hid"
)

const vendorIDofZebraScanner = 1504
const ibmHandHeldUsb = 4864

var (
	ErrScannerNotFound         = errors.New("scanner was not found")
	ErrUnexpectedInterfaceMode = errors.New("the interface must be configured to IBM Hand-held USB")
	ErrUnexpectedNoData        = errors.New("the scanned code has no data")
	ErrTooLongScannedCode      = errors.New("the byte length is bigger than 64 bytes")
	wg sync.WaitGroup
)

func findZebraScanner() hid.Device {
	var zebraScannerDevice hid.Device
	hid.UsbWalk(func(device hid.Device) {
		info := device.Info()

		fmt.Printf("%04x:%04x:%04x:%02x\n", info.Vendor, info.Product, info.Revision, info.Interface)

		if info.Vendor == vendorIDofZebraScanner {
			if info.Product == ibmHandHeldUsb {
				fmt.Println("Zebra scanner found")
				zebraScannerDevice = device
				return
			} else {
				fmt.Println("Interface mode must be 'IBM Hand-held USB'")
			}
		}
	})

	return zebraScannerDevice
}

func splitByNull(data []byte) ([]byte, error) {
	if len(data) == 0 {
		// No barcode has empty string.
		return nil, ErrUnexpectedNoData
	}

	firstNullIndex := bytes.Index(data, []byte{0})
	restBytes := data[1+firstNullIndex:]
	secondNullIndex := bytes.Index(restBytes, []byte{0})

	if secondNullIndex == -1 {
		return nil, ErrTooLongScannedCode
	}

	return restBytes[:secondNullIndex], nil
}

func scan(device hid.Device) {
	readBytes := 64
	infiniteWaiting := 0

	var recursive func()
	recursive = func() {
		rawBytes, err := device.Read(readBytes, time.Duration(infiniteWaiting))
		if err != nil {
			fmt.Printf("USB device was disconnected: [%s]", err)
			wg.Done()
			return
		}
		payload, err := splitByNull(rawBytes)
		if err != nil {
			fmt.Println("Payload error", err)
			wg.Done()
			return
		} else {
			fmt.Println(string(payload))
		}
		recursive()
	}
	recursive()
}

func Run() {
	fmt.Println("Scanner test START!!!")

	zebraScannerDevice := findZebraScanner()
	if zebraScannerDevice == nil {
		fmt.Println("No USB HID devices found")
		return
	}

	err := zebraScannerDevice.Open()
	if err != nil {
		fmt.Println("failed to open device", err)
		return
	}

	wg.Add(1)
	go scan(zebraScannerDevice)
	wg.Wait()

	fmt.Println("FINISH")
}
