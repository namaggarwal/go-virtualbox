package virtualbox

import (
	"fmt"
	"os/exec"
)

// IVBoxManage ...
type IVBoxManage interface {
	CreateVM(name string, osType string, baseFolder string, shouldRegister bool) error
	StartVM(name string) error
	AddStorageCtl(vmName string, name string, ctlType string, controller string) error
	AttachStorage(vmName string, controllerName string, port int32, device int32, storageType string, medium string) error
	CreateMedium(mediumType string, filePath string, size int32, format string) error
}

// VBoxManage ...
type vBoxManage struct {
}

func (m *vBoxManage) CreateVM(name string, osType string, baseFolder string, shouldRegister bool) error {
	cmd := exec.Command("VBoxManage")
	cmd.Args = append(cmd.Args, "createvm")
	cmd.Args = append(cmd.Args, "--name")
	cmd.Args = append(cmd.Args, name)
	cmd.Args = append(cmd.Args, "--ostype")
	cmd.Args = append(cmd.Args, osType)
	cmd.Args = append(cmd.Args, "--basefolder")
	cmd.Args = append(cmd.Args, baseFolder)
	if shouldRegister {
		cmd.Args = append(cmd.Args, "--register")
	}
	_, err := m.execute(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (m *vBoxManage) StartVM(name string) error {
	cmd := exec.Command("VBoxManage")
	cmd.Args = append(cmd.Args, "startvm")
	cmd.Args = append(cmd.Args, name)
	_, err := m.execute(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (m *vBoxManage) AddStorageCtl(vmName string, name string, ctlType string, controller string) error {
	cmd := exec.Command("VBoxManage")
	cmd.Args = append(cmd.Args, "storagectl")
	cmd.Args = append(cmd.Args, vmName)
	cmd.Args = append(cmd.Args, "--name")
	cmd.Args = append(cmd.Args, name)
	cmd.Args = append(cmd.Args, "--add")
	cmd.Args = append(cmd.Args, ctlType)
	cmd.Args = append(cmd.Args, "--controller")
	cmd.Args = append(cmd.Args, controller)
	_, err := m.execute(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (m *vBoxManage) AttachStorage(vmName string, controllerName string, port int32, device int32, storageType string, medium string) error {
	cmd := exec.Command("VBoxManage")
	cmd.Args = append(cmd.Args, "storageattach")
	cmd.Args = append(cmd.Args, vmName)
	cmd.Args = append(cmd.Args, "--storagectl")
	cmd.Args = append(cmd.Args, controllerName)
	cmd.Args = append(cmd.Args, "--port")
	cmd.Args = append(cmd.Args, fmt.Sprintf("%d", port))
	cmd.Args = append(cmd.Args, "--device")
	cmd.Args = append(cmd.Args, fmt.Sprintf("%d", device))
	cmd.Args = append(cmd.Args, "--type")
	cmd.Args = append(cmd.Args, storageType)
	cmd.Args = append(cmd.Args, "--medium")
	cmd.Args = append(cmd.Args, medium)
	_, err := m.execute(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (m *vBoxManage) CreateMedium(mediumType string, filePath string, size int32, format string) error {
	cmd := exec.Command("VBoxManage")
	cmd.Args = append(cmd.Args, "createmedium")
	cmd.Args = append(cmd.Args, mediumType)
	cmd.Args = append(cmd.Args, "--filename")
	cmd.Args = append(cmd.Args, filePath)
	cmd.Args = append(cmd.Args, "--size")
	cmd.Args = append(cmd.Args, fmt.Sprintf("%d", size))
	cmd.Args = append(cmd.Args, "--format")
	cmd.Args = append(cmd.Args, format)
	r, err := m.execute(cmd)
	if err != nil {
		println(fmt.Sprintf("%s", r))
		return err
	}
	return nil
}

func (m *vBoxManage) execute(cmd *exec.Cmd) ([]byte, error) {
	res, err := cmd.CombinedOutput()
	return res, err
}

// NewVBoxManage ...
func NewVBoxManage() IVBoxManage {
	return &vBoxManage{}
}
