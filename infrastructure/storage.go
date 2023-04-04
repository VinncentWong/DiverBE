package infrastructure

import (
	"fmt"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
)

var storage *cloudinary.Cloudinary

func ConnectToStorage() {
	_storage, _ := cloudinary.NewFromURL(fmt.Sprintf("cloudinary://%s:%s@%s",
		os.Getenv("STORAGE_KEY"),
		os.Getenv("STORAGE_SECRET"),
		os.Getenv("STORAGE_CLOUD_NAME")),
	)
	storage = _storage
}

func GetStorage() *cloudinary.Cloudinary {
	return storage
}
