package pkg

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/AnshKumar200/Harbor/config"
)

func CreateImage(reader io.ReadCloser, image ImageId) error {
	rootDir := filepath.Join(config.DefaultImageStoreRootDir, image.Digest)
	prs, err := Exist(rootDir)
	if err != nil {
		return err
	}
	if prs {
		return fmt.Errorf("image <%s> already exists", rootDir)
	}

	if err := os.MkdirAll(rootDir, 0700); err != nil {
		return err
	}

	if err := Untar(rootDir+"/rootfs", reader); err != nil {
		return err
	}

	source, err := os.Create(filepath.Join(rootDir, "source"))
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(source)
	if err := encoder.Encode(image); err != nil {
		return err
	}

	log.Printf("image <%s> stored at <%s>", image.Digest, rootDir)
	return nil
}

func Untar(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)
		switch header.Typeflag {

		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(f, tr); err != nil {
				return err
			}
			f.Close()
		}
	}
}
