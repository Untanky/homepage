package media

// import (
// 	"bytes"
// 	"errors"
// 	"fmt"
// 	"image"
// 	"image/jpeg"
// 	"image/png"
// 	"io"
// 	"mime"
// 	"strings"
//
// 	"github.com/gen2brain/avif"
// 	"github.com/gen2brain/webp"
// 	"github.com/nfnt/resize"
// )
//
// type ImageType string
//
// const (
// 	PNG  ImageType = "PNG"
// 	JPEG ImageType = "JPEG"
// 	WebP ImageType = "WEBP"
// 	AVIF ImageType = "AVIF"
// )
//
// var ErrUnknownImageType = errors.New("unknown image type")
//
// type ImagePipeline struct {
// 	Original      image.Image
// 	OriginalScale float64
// 	Scales        []float64
// 	MediaTypes    []ImageType
// }
//
// func (pipeline ImagePipeline) Run() ([]AssetVersion, error) {
// 	versions := make([]AssetVersion, 0, len(pipeline.Scales)*len(pipeline.MediaTypes))
//
// 	for _, scale := range pipeline.Scales {
// 		scaledImage := scaleImage(scale, pipeline.OriginalScale, pipeline.Original)
//
// 		for _, mediaType := range pipeline.MediaTypes {
// 			buffer := bytes.NewBuffer(nil)
//
// 			err := encodeImage(mediaType, scaledImage, buffer)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to encode image as '%s': %w", strings.ToLower(string(mediaType)), err)
// 			}
//
// 			mimetype := mime.TypeByExtension(fmt.Sprintf(".%s", mediaType))
//
// 			versions = append(versions, AssetVersion{
// 				Scale:     scale,
// 				MediaType: mimetype,
// 				Buffer:    buffer,
// 			})
// 		}
// 	}
//
// 	return versions, nil
// }
//
// func scaleImage(scale, originalScale float64, img image.Image) image.Image {
// 	normalizedScale := scale / originalScale
//
// 	if normalizedScale == 1 {
// 		return img
// 	}
//
// 	width := uint(float64(img.Bounds().Dx()) * normalizedScale)
// 	return resize.Resize(width, 0, img, resize.Lanczos2)
// }
//
// func encodeImage(resultType ImageType, img image.Image, writer io.Writer) error {
// 	switch resultType {
// 	case PNG:
// 		return png.Encode(writer, img)
// 	case JPEG:
// 		return jpeg.Encode(writer, img, &jpeg.Options{
// 			Quality: 90,
// 		})
// 	case WebP:
// 		return webp.Encode(writer, img, webp.Options{
// 			Quality: 80,
// 			Method:  6,
// 		})
// 	case AVIF:
// 		return avif.Encode(writer, img, avif.Options{
// 			Quality:      80,
// 			QualityAlpha: 80,
// 			Speed:        10,
// 		})
// 	default:
// 		return ErrUnknownImageType
// 	}
// }
