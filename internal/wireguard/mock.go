package wireguard

// import "fmt"

// // MockDB is a mock implementation of DeviceDB
// type MockDB struct {
// 	devices map[uint]*Device
// }

// // CreateDevice creates a new device
// func (db *MockDB) CreateDevice(device *Device) error {
// 	db.devices[device.ID] = device
// 	return nil
// }

// // ReadDevice reads a device by ID
// func (db *MockDB) ReadDevice(id uint) (*Device, error) {
// 	device, exists := db.devices[id]
// 	if !exists {
// 		return nil, fmt.Errorf("Device not found")
// 	}
// 	return device, nil
// }

// // UpdateDevice updates a device
// func (db *MockDB) UpdateDevice(device *Device) error {
// 	_, exists := db.devices[device.ID]
// 	if !exists {
// 		return fmt.Errorf("Device not found")
// 	}
// 	db.devices[device.ID] = device
// 	return nil
// }

// // DeleteDevice deletes a device by ID
// func (db *MockDB) DeleteDevice(id uint) error {
// 	_, exists := db.devices[id]
// 	if !exists {
// 		return fmt.Errorf("Device not found")
// 	}
// 	delete(db.devices, id)
// 	return nil
// }
