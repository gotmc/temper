// Copyright (c) 2016–2024 The temper developers. All rights reserved.
// Project site: https://github.com/gotmc/temper
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package temper

import (
	"fmt"
	"log"

	libusb "github.com/gotmc/libusb/v2"
)

const (
	vendorID   = 0x0c45
	productID  = 0x7401
	interface1 = 0x00
	interface2 = 0x01
)

type temper struct {
	Timeout          int
	Device           *libusb.Device
	DeviceDescriptor *libusb.Descriptor
	DeviceHandle     *libusb.DeviceHandle
	ConfigDescriptor *libusb.ConfigDescriptor
	BulkEndpoint     *libusb.EndpointDescriptor
}

// NewViaSN creates a new temper instance by searching through the list of USB
// devices for the given serial number.
func NewViaSN(ctx *libusb.Context, sn string) (*temper, error) {
	var d temper
	usbDevices, err := ctx.GetDeviceList()
	if err != nil {
		return &d, fmt.Errorf("Error getting USB device list: %s", err)
	}
	// Search through the USB devices looking for serial number
	for _, usbDevice := range usbDevices {
		usbDeviceDescriptor, err := usbDevice.GetDeviceDescriptor()
		if err != nil {
			return &d, fmt.Errorf("Error getting device descriptor: %s", err)
		}
		// Check the VendorID and Product ID. If those don't match what we're
		// looking for, then there's no reason to open the device and read its S/N.
		if usbDeviceDescriptor.VendorID == vendorID &&
			usbDeviceDescriptor.ProductID == productID {
			// Found a TEMPer
			usbDeviceHandle, err := usbDevice.Open()
			if err != nil {
				return &d, fmt.Errorf("Error getting device handle: %s", err)
			}
			serialNum, err := usbDeviceHandle.GetStringDescriptorASCII(
				usbDeviceDescriptor.SerialNumberIndex)
			if err != nil {
				return &d, fmt.Errorf("Error reading S/N: %s", err)
			}
			if serialNum == sn {
				log.Printf("Found S/N %s. Creating device", sn)
				return create(usbDevice, usbDeviceHandle)
			}
			usbDeviceHandle.Close()
		}
	}
	// Close the list of devices
	return &d, fmt.Errorf("Couldn't find USB-1608FS-Plus S/N %s.", sn)
}

// GetFirstDevice creates a new instance of a temper using the first
// TEMPer found in the USB context.
func GetFirstDevice(ctx *libusb.Context) (*temper, error) {
	var d temper
	dev, dh, err := ctx.OpenDeviceWithVendorProduct(vendorID, productID)
	if err != nil {
		return &d, fmt.Errorf("Error opening the TEMPer using the VendorID and ProductID, %s", err)
	}
	return create(dev, dh)
}

func create(dev *libusb.Device, dh *libusb.DeviceHandle) (*temper, error) {
	var d temper
	err := dh.ClaimInterface(0)
	if err != nil {
		return &d, fmt.Errorf("Error claiming the bulk interface %s", err)
	}
	d.Device = dev
	d.DeviceHandle = dh
	deviceDescriptor, err := d.Device.GetDeviceDescriptor()
	if err != nil {
		return &d, fmt.Errorf("Error getting device descriptor %s", err)
	}
	d.DeviceDescriptor = deviceDescriptor
	configDescriptor, err := d.Device.GetActiveConfigDescriptor()
	if err != nil {
		return &d, fmt.Errorf("Error getting active config descriptor. %s", err)
	}
	d.ConfigDescriptor = configDescriptor
	firstDescriptor := configDescriptor.SupportedInterfaces[0].InterfaceDescriptors[0]
	d.BulkEndpoint = firstDescriptor.EndpointDescriptors[0]
	return &d, nil
}

// Close closes the USB interface and context to the TEMPer
func (d *temper) Close() error {
	// Release the interface and close up shop
	err := d.DeviceHandle.ReleaseInterface(0)
	if err != nil {
		return fmt.Errorf("Error releasing interface %s", err)
	}
	d.DeviceHandle.Close()
	return nil
}

// SendCommandToDevice sends the given command and data to the device and
// returns the number of bytes received and whether or not an error was
// received.
func (d *temper) SendCommandToDevice(cmd command, data []byte) (int, error) {
	if data == nil {
		data = []byte{0}
	}
	requestType := libusb.BitmapRequestType(
		libusb.HostToDevice, libusb.Vendor, libusb.DeviceRecipient)
	bytesReceived, err := d.DeviceHandle.ControlTransfer(
		requestType, byte(cmd), 0x0, 0x0, data, len(data), d.Timeout)
	if err != nil {
		return bytesReceived, fmt.Errorf("Error sending command '%s' to device: %s", cmd, err)
	}
	return bytesReceived, nil
}

func (d *temper) ReadCommandFromDevice(cmd command, data []byte) (int, error) {
	if data == nil {
		data = []byte{0}
	}
	requestType := libusb.BitmapRequestType(
		libusb.DeviceToHost, libusb.Vendor, libusb.DeviceRecipient)
	bytesReceived, err := d.DeviceHandle.ControlTransfer(
		requestType, byte(cmd), 0x0, 0x0, data, len(data), d.Timeout)
	if err != nil {
		return bytesReceived, fmt.Errorf("Error reading command '%s' from device: %s", cmd, err)
	}
	return bytesReceived, nil
}

func (d *temper) Read(p []byte) (n int, err error) {
	return d.DeviceHandle.BulkTransfer(
		d.BulkEndpoint.EndpointAddress,
		p,
		len(p),
		d.Timeout,
	)
}
