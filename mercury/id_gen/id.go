package id_gen

import (
	"fmt"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMechineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMechineID, nil
}

func Init(mechineId uint16) (err error) {
	sonyMechineID = mechineId
	settings := sonyflake.Settings{}
	settings.MachineID = getMachineID
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("sony flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
