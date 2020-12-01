package util

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	jpegconv converter = &jpegConverter{}
	pngconv  converter = &pngConverter{}
	gifconv  converter = &gifConverter{}
)

func Jpeg2Png(filepath string) error {
	newpath := replaceExt(filepath, "jpeg", "png")
	i, jerr := decode(jpegconv, filepath)
	if jerr != nil {
		return jerr
	}
	if perr := encode(pngconv, i, newpath); perr != nil {
		return perr
	}
	if rerr := os.Remove(filepath); rerr != nil {
		return rerr
	}
	return nil
}
func Jpeg2Gif(filepath string) error {
	newpath := replaceExt(filepath, "jpeg", "gif")
	i, jerr := decode(jpegconv, filepath)
	if jerr != nil {
		return jerr
	}
	if perr := encode(gifconv, i, newpath); perr != nil {
		return perr
	}
	if rerr := os.Remove(filepath); rerr != nil {
		return rerr
	}
	return nil
}
func Png2Jpeg(filepath string) error {
	newpath := replaceExt(filepath, "png", "jpeg")
	i, jerr := decode(pngconv, filepath)
	if jerr != nil {
		return jerr
	}
	if perr := encode(jpegconv, i, newpath); perr != nil {
		return perr
	}
	if rerr := os.Remove(filepath); rerr != nil {
		return rerr
	}
	return nil
}
func Png2Gif(filepath string) error {
	newpath := replaceExt(filepath, "png", "gif")
	i, jerr := decode(pngconv, filepath)
	if jerr != nil {
		return jerr
	}
	if perr := encode(gifconv, i, newpath); perr != nil {
		return perr
	}
	if rerr := os.Remove(filepath); rerr != nil {
		return rerr
	}
	return nil
}
func decode(c converter, path string) (image.Image, error) {
	f, openErr := os.Open(path)
	if openErr != nil {
		return nil, openErr
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	img, encErr := jpegconv.Decode(f)
	if encErr != nil {
		return nil, encErr
	}
	return img, nil
}

func encode(c converter, i image.Image, path string) error {
	f2, createErr := os.Create(path)
	if createErr != nil {
		return createErr
	}
	if err := c.Encode(f2, i); err != nil {
		return err
	}
	return nil
}
func replaceExt(filePath, from, to string) string {
	ext := filepath.Ext(filePath)
	if len(from) > 0 && ext != from {
		return filePath
	}
	return filePath[:len(filePath)-len(ext)] + to
}

type converter interface {
	Encode(w io.Writer, i image.Image) error
	Decode(r io.Reader) (image.Image, error)
}

type jpegConverter struct {
}

func (c jpegConverter) Encode(w io.Writer, i image.Image) error {
	return jpeg.Encode(w, i, nil)
}
func (c jpegConverter) Decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

type pngConverter struct {
}

func (c pngConverter) Encode(w io.Writer, i image.Image) error {
	return png.Encode(w, i)
}
func (c pngConverter) Decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

type gifConverter struct {
}

func (c gifConverter) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, nil)
}
func (c gifConverter) Decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}
