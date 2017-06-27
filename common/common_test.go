package common

import (
	"image"
	"os"
	"testing"
)

func loadImage(filePath string) (image.Image, error) {
	// Open the file image
	fImage, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	// Ensure that the image file will be closed
	defer fImage.Close()

	// Convert it to an image "object"
	img, _, err := image.Decode(fImage)

	if err != nil {
		return nil, err
	}

	return img, nil
}

func TestGetSize(t *testing.T) {
	// Table tests
	var tTable = []struct {
		path   string
		width  int
		height int
	}{
		{"../dataset/test/1.png", 200, 200},
		{"../dataset/test/2.png", 6, 6},
		{"../dataset/test/3.png", 256, 256},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := loadImage(pair.path)
		width, height := GetSize(img)
		if width != pair.width {
			t.Error(
				"Expected: ", pair.width,
				"Received: ", width,
			)
		}
		if height != pair.height {
			t.Error(
				"Expected: ", pair.height,
				"Received: ", height,
			)
		}
	}
}

func TestIsGrayscale(t *testing.T) {
	// Table tests
	var tTable = []struct {
		path string
		res  bool
	}{
		{"../dataset/test/1.png", true},
		{"../dataset/test/2.png", true},
		{"../dataset/test/3.png", false},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		img, _ := loadImage(pair.path)
		res := IsGrayscale(img)
		if res != pair.res {
			t.Error(
				"Expected: ", pair.res,
				"Received: ", res,
			)
		}
	}
}

func TestCheckInputData(t *testing.T) {
	// Image is not in grayscale
	var images []image.Image
	img, err := loadImage("../dataset/test/3.png")
	if err != nil {
		t.Error(err)
	}
	images = append(images, img)
	err = CheckInputData(images)
	if err == nil {
		t.Error("Expected: Image is not in grayscale. Received: nil")
	}
	images = nil

	// Images have different sizes
	var paths []string
	paths = append(paths, "../dataset/test/1.png")
	paths = append(paths, "../dataset/test/2.png")

	for index := 0; index < len(paths); index++ {
		img, err := loadImage(paths[index])
		if err != nil {
			t.Error(err)
		}
		images = append(images, img)
	}
	err = CheckInputData(images)
	if err == nil {
		t.Error("Expected: Images have different sizes. Received: nil")
	}
	images = nil

	// No error
	img, err = loadImage("../dataset/test/1.png")
	if err != nil {
		t.Error(err)
	}
	images = append(images, img)
	err = CheckInputData(images)
	if err != nil {
		t.Error("Expected: nil. Received: ", err)
	}
}

func TestGetBinary(t *testing.T) {
	// Table tests
	var tTable = []struct {
		value     uint8
		threshold uint8
		result    string
	}{
		{120, 120, "1"},
		{214, 190, "1"},
		{150, 240, "0"},
	}

	// Test with all values in the table
	for _, pair := range tTable {
		result := GetBinary(pair.value, pair.threshold)
		if result != pair.result {
			t.Error(
				"Expected: ", pair.result,
				"Received: ", result,
			)
		}
	}
}

func TestGetPixels(t *testing.T) {
	img, err := loadImage("../dataset/test/2.png")
	if err != nil {
		t.Error(err)
	}
	pixels := GetPixels(img)

	var expectedPixels [][]uint8
	expectedPixels = append(expectedPixels, []uint8{  0, 255,   0, 255,   0, 255})
	expectedPixels = append(expectedPixels, []uint8{255, 255, 255, 255, 255,   0})
	expectedPixels = append(expectedPixels, []uint8{  0, 255, 255,   0, 255, 255})
	expectedPixels = append(expectedPixels, []uint8{255, 255,   0, 255, 255,   0})
	expectedPixels = append(expectedPixels, []uint8{  0, 255, 255, 255, 255, 255})
	expectedPixels = append(expectedPixels, []uint8{255,   0, 255,   0, 255,   0})

	if len(pixels) == len(expectedPixels) {
		for row := 0; row < len(pixels); row++ {
			for col := 0; col < len(pixels[0]); col++ {
				if pixels[row][col] != expectedPixels[row][col] {
					t.Error(
						"Expected value : ", expectedPixels[row][col],
						"Received value : ", pixels[row][col],
					)
				}
			}
		}
	} else {
		t.Error("Slices have different sizes")
	}
}
