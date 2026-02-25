package utils

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

// multipart.File — HTTP request dan kelgan fayl stream i. Cloudinary to'g'ridan to'g'ri stream dan o'qiydi, diskka saqlamaydi.
// multipart.FileHeader — fayl haqida meta ma'lumot: nomi, hajmi, content type.
func UploadImage(file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {

	// 1. Cloudinary instance yaratish
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)

	if err != nil {
		return "", err
	}

	// 2. Fayl nomidan extension olish
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowedExt := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExt[ext] {
		return "", fmt.Errorf("faqat jpg, jpeg, png, webp formatlar ruxsat etilgan")
	}

	// / 3. Fayl hajmini tekshirish (max 5MB)
	if fileHeader.Size > 5*1024*1024 {
		return "", fmt.Errorf("fayl hajmi 5MB dan oshmasligi kerak")
	}

	// 4. Cloudinary ga yuklash
	ctx := context.Background()
	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: folder,
	})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}

func DeleteImage(publicID string) error {
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

func ExtractPublicID(imageURL string) string {
	// https://res.cloudinary.com/demo/image/upload/v1234/categories/abc123.jpg
	// → categories/abc123

	parts := strings.Split(imageURL, "/upload/")
	if len(parts) < 2 {
		return ""
	}

	// v1234/ qismini olib tashlash
	withoutVersion := strings.SplitN(parts[1], "/", 2)
	if len(withoutVersion) < 2 {
		return ""
	}

	// .jpg extensionni olib tashlash
	return strings.TrimSuffix(withoutVersion[1], filepath.Ext(withoutVersion[1]))
}
