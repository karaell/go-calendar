package storage

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

type ZipStorage struct {
	*Storage
}

func CreateZipStorage(filename string) *ZipStorage {
	return &ZipStorage{
		&Storage{
			filename: filename,
		},
	}
}

func (z *ZipStorage) Save(data []byte) error {
	f, err := os.Create(z.GetFilename())
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSaveStorage, err)
	}

	defer func() {
		err = f.Close()

		if err != nil {
			fmt.Println(fmt.Errorf("%w %s: %w", ErrCloseFile, z.GetFilename(), err))
		}
	}()

	zw := zip.NewWriter(f)
	defer func() {
		err = zw.Close()

		if err != nil {
			fmt.Println(fmt.Errorf("%w: %w", ErrCloseZip, err))
		}
	}()

	zf, err := zw.Create("data")
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSaveStorage, err)
	}

	_, err = zf.Write(data)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrSaveStorage, err)
	}

	return nil
}

func (z *ZipStorage) Load() ([]byte, error) {
	zr, err := zip.OpenReader(z.GetFilename())
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrLoadStorage, err)
	}
	defer func() {
		err = zr.Close()

		if err != nil {
			fmt.Println(fmt.Errorf("%w: %w", ErrCloseZip, err))
		}
	}()

	if len(zr.File) == 0 {
		return nil, fmt.Errorf("%w: %w", ErrLoadStorage, ErrEmptyZip)
	}

	file := zr.File[0]

	fc, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrLoadStorage, err)
	}
	defer func() {
		err = fc.Close()

		if err != nil {
			fmt.Println(fmt.Errorf("%w %s: %w", ErrCloseFile, z.GetFilename(), err))
		}
	}()

	return io.ReadAll(fc)
}
