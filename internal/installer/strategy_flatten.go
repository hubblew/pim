package installer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hubblew/pim/internal/utils"
	"github.com/spf13/afero"
)

type flattenStrategy struct {
	fs         afero.Fs
	outputPath string
}

var _ Strategy = (*flattenStrategy)(nil)

func NewFlattenStrategy(fs afero.Fs, path string) Strategy {
	return &flattenStrategy{
		fs:         fs,
		outputPath: path,
	}
}

func (s *flattenStrategy) Initialize(_ UserPrompter) error {
	if err := s.fs.RemoveAll(s.outputPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to delete output directory '%s': %w", s.outputPath, err)
	}

	if err := s.fs.MkdirAll(s.outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory '%s': %w", s.outputPath, err)
	}
	return nil
}

func (s *flattenStrategy) AddFile(srcPath, relativePath string) error {
	dstPath := filepath.Join(s.outputPath, filepath.Base(relativePath))
	return utils.CopyFile(s.fs, srcPath, dstPath)
}

func (s *flattenStrategy) Close() error {
	return nil
}
