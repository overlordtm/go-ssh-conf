package ssh_conf

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevinburke/ssh_config"
)

func init() {
	log.SetOutput(os.Stderr)
}

func Parse(sources []string) (cfg *ssh_config.Config, err error) {

	cfg = &ssh_config.Config{}

	for _, pth := range sources {
		partialErr := parse(cfg, os.Expand(pth, os.Getenv))
		if partialErr != nil {
			err = errors.Join(err, partialErr)
		}
	}

	return
}

func parse(targetCfg *ssh_config.Config, pth string) (err error) {

	stat, err := os.Stat(pth)
	if os.IsNotExist(err) {
		return fmt.Errorf("path %s does not exist", pth)
	}
	if err != nil {
		return fmt.Errorf("error stat %s: %w", pth, err)
	}

	if stat.IsDir() {
		return parseDir(targetCfg, pth)
	} else {
		cfg, err := parseFile(pth)
		targetCfg.Hosts = append(targetCfg.Hosts, cfg.Hosts...)
		return err
	}
}

func parseDir(targetCfg *ssh_config.Config, dir string) (err error) {

	err = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Mode()&os.ModeSymlink == os.ModeSymlink {
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}

			if stat, err := os.Stat(target); err != nil {
				return err
			} else {
				if stat.IsDir() {
					return parseDir(targetCfg, target)
				} else {
					cfg, err := parseFile(target)
					if err != nil {
						return errors.Join(err, fmt.Errorf("error parsing file %s: %w", target, err))
					}

					targetCfg.Hosts = append(targetCfg.Hosts, cfg.Hosts...)
					return nil
				}
			}
		}

		cfg, err := parseFile(path)
		if err != nil {
			return errors.Join(err, fmt.Errorf("error parsing file %s: %w", path, err))
		}

		targetCfg.Hosts = append(targetCfg.Hosts, cfg.Hosts...)
		return nil
	})

	return err
}

func parseFile(file string) (*ssh_config.Config, error) {
	if !strings.HasSuffix(file, ".conf") {
		return nil, fmt.Errorf("file %s is not a .conf file", file)
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ssh_config.Decode(f)
}
