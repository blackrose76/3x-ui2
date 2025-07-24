package services

import (
	"fmt"
	"os"
	"os/exec"
)

type XrayService struct {
	cmd *exec.Cmd
}

func NewXrayService() *XrayService {
	return &XrayService{}
}

func (s *XrayService) Start() error {
	if s.cmd != nil {
		return fmt.Errorf("Xray is already running")
	}

	s.cmd = exec.Command("xray", "run", "-c", "config.json")
	s.cmd.Stdout = os.Stdout
	s.cmd.Stderr = os.Stderr

	err := s.cmd.Start()
	if err != nil {
		s.cmd = nil
		return err
	}

	return nil
}

func (s *XrayService) Stop() error {
	if s.cmd == nil {
		return fmt.Errorf("Xray is not running")
	}

	err := s.cmd.Process.Kill()
	if err != nil {
		return err
	}

	s.cmd = nil
	return nil
}

func (s *XrayService) Restart() error {
	if s.cmd != nil {
		err := s.Stop()
		if err != nil {
			return err
		}
	}

	return s.Start()
}
