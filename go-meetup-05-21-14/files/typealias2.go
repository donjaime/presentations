type DeviceCollection []*device

func (c DeviceCollection) Halt() {
  for _, device := range c {
    device.Halt()
  }
}
