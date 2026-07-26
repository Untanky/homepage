package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"image"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
//
// 	"github.com/google/uuid"
// 	"github.com/untanky/homepage/internal/database"
// 	"github.com/untanky/homepage/internal/media"
// 	"github.com/untanky/homepage/internal/media/sql"
//
// 	_ "image/jpeg"
// 	_ "image/png"
//
// 	_ "github.com/gen2brain/avif"
// 	_ "github.com/gen2brain/webp"
// )
//
// var (
// 	scales      = floatSlice{}
// 	mediaTypes  = mediaTypeSlice{}
// 	source      = flag.String("source", "", "the source image to be used")
// 	path        = flag.String("path", "", "the path under which the media is available")
// 	sourceScale = flag.Float64("sourceScale", 1, "the scale of the source image")
// )
//
// type floatSlice []float64
//
// func (s *floatSlice) String() string {
// 	return fmt.Sprintf("%v", *s)
// }
//
// func (s *floatSlice) Set(value string) error {
// 	n, err := strconv.ParseFloat(value, 32)
// 	if err != nil {
// 		return fmt.Errorf("invalid float: %s", value)
// 	}
// 	*s = append(*s, n)
// 	return nil
// }
//
// type mediaTypeSlice []media.ImageType
//
// func (s *mediaTypeSlice) String() string {
// 	return fmt.Sprintf("%v", *s)
// }
//
// func (s *mediaTypeSlice) Set(value string) error {
// 	*s = append(*s, media.ImageType(strings.ToUpper(value)))
// 	return nil
// }
//
// func main() {
// 	flag.Var(&scales, "scale", "the desired scale (can be used multiple times)")
// 	flag.Var(&mediaTypes, "mediaType", "the desired media type (can be used multiple times)")
//
// 	flag.Parse()
//
// 	if err := run(context.Background()); err != nil {
// 		log.Printf("failed to convert images: %s", err.Error())
// 		os.Exit(1)
// 	}
// }
//
// func run(ctx context.Context) error {
// 	img, err := readImage()
// 	if err != nil {
// 		return err
// 	}
//
// 	pipeline := media.ImagePipeline{
// 		Original:      img,
// 		OriginalScale: *sourceScale,
// 		Scales:        scales,
// 		MediaTypes:    mediaTypes,
// 	}
//
// 	versions, err := pipeline.Run()
// 	if err != nil {
// 		return err
// 	}
//
// 	assetID, err := uuid.NewV7()
// 	if err != nil {
// 		return fmt.Errorf("generating uuid: %w", err)
// 	}
//
// 	asset := media.Asset{
// 		ID:       media.AssetID(assetID),
// 		Name:     *path,
// 		Width:    0,
// 		Height:   0,
// 		Versions: versions,
// 	}
//
// 	databaseClient, err := database.Setup(ctx, database.Config{
// 		Username: "postgres",
// 		Password: "postgres",
// 		Host:     "localhost",
// 		Port:     5432,
// 		Database: "postgres",
// 	})
// 	if err != nil {
// 		return fmt.Errorf("setting up database: %w", err)
// 	}
//
// 	repo := sql.NewMediaRepository(databaseClient)
//
// 	if err := repo.Create(ctx, asset); err != nil {
// 		return fmt.Errorf("saving asset: %w", err)
// 	}
//
// 	return nil
// }
//
// func readImage() (image.Image, error) {
// 	imageFile, err := os.Open(*source)
// 	if err != nil {
// 		return nil, fmt.Errorf("opening source image: %w", err)
// 	}
// 	defer imageFile.Close()
//
// 	img, _, err := image.Decode(imageFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("decode source image: %w", err)
// 	}
//
// 	return img, nil
// }
