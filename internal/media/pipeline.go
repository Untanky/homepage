package media

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"mime"
	"strings"

	"github.com/gen2brain/avif"
	"github.com/gen2brain/webp"
	"github.com/nfnt/resize"
)

type ImageType string

const (
	PNG  ImageType = "PNG"
	JPEG ImageType = "JPEG"
	WebP ImageType = "WEBP"
	AVIF ImageType = "AVIF"
)

var ErrUnknownImageType = errors.New("unknown image type")

type ImagePipeline struct {
	Original   image.Image
	MediaTypes []ImageType
	Widths     []uint
}

func (pipeline ImagePipeline) Run() ([]Version, error) {
	versions := make([]Version, 0, len(pipeline.Widths)*len(pipeline.MediaTypes))

	for _, width := range pipeline.Widths {
		scaledImage := scaleImage(pipeline.Original, width)
		height := uint(scaledImage.Bounds().Dy())

		for _, mediaType := range pipeline.MediaTypes {
			buffer := bytes.NewBuffer(nil)

			err := encodeImage(mediaType, scaledImage, buffer)
			if err != nil {
				return nil, fmt.Errorf("failed to encode image as '%s': %w", strings.ToLower(string(mediaType)), err)
			}

			mimetype := mime.TypeByExtension(fmt.Sprintf(".%s", mediaType))

			versions = append(versions, Version{
				VersionMetadata: VersionMetadata{
					Mimetype: mimetype,
					Width:    width,
					Height:   height,
				},
				Size:    buffer.Len(),
				Content: bytes.NewReader(buffer.Bytes()),
			})
		}
	}

	return versions, nil
}

func scaleImage(img image.Image, width uint) image.Image {
	imgWidth := uint(img.Bounds().Dx())

	if imgWidth == width {
		return img
	}

	return resize.Resize(width, 0, img, resize.Lanczos2)
}

func encodeImage(resultType ImageType, img image.Image, writer io.Writer) error {
	switch resultType {
	case PNG:
		return png.Encode(writer, img)
	case JPEG:
		return jpeg.Encode(writer, img, &jpeg.Options{
			Quality: 90,
		})
	case WebP:
		return webp.Encode(writer, img, webp.Options{
			Quality: 80,
			Method:  6,
		})
	case AVIF:
		return avif.Encode(writer, img, avif.Options{
			Quality:      60,
			QualityAlpha: 60,
			Speed:        10,
		})
	default:
		return ErrUnknownImageType
	}
}
