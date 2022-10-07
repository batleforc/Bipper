package model

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"sort"
	"strings"
)

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

type StaticImageHandler struct {
}

// Move file fs.DirEntry to a standard type in order to support other typ

func (s *StaticImageHandler) GetImageType(data []byte) string {
	incipitStr := string(data)
	for magic, mime := range magicTable {
		if strings.HasPrefix(incipitStr, magic) {
			return mime
		}
	}
	return ""
}

func (s *StaticImageHandler) SaveImage(src io.Reader, fileName string) error {
	if strings.ToLower(os.Getenv("ASSET_HANDLER")) == "filesystem" {
		dst, err := os.Create(fmt.Sprintf("%s/%s", os.Getenv("ASSET_PATH"), fileName))
		if err != nil {
			return err
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		return err
	} else {
		return fmt.Errorf("asset handler not supported")
	}
}

func (s *StaticImageHandler) GetAllFile() ([]fs.DirEntry, error) {
	if strings.ToLower(os.Getenv("ASSET_HANDLER")) == "filesystem" {
		files, err := os.ReadDir(os.Getenv("ASSET_PATH"))
		if err != nil {
			return nil, err
		}
		var fileNames []fs.DirEntry
		for _, file := range files {
			if !file.IsDir() {
				fileNames = append(fileNames, file)
			}
		}
		return fileNames, nil
	} else {
		return nil, fmt.Errorf("asset handler not supported")
	}
}

func (s *StaticImageHandler) DeleteFile(fileName string) error {
	if strings.ToLower(os.Getenv("ASSET_HANDLER")) == "filesystem" {
		return os.Remove(fmt.Sprintf("%s/%s", os.Getenv("ASSET_PATH"), fileName))
	} else {
		return fmt.Errorf("asset handler not supported")
	}
}

func (s *StaticImageHandler) GetFile(fileName string) (io.ReadCloser, error) {
	if strings.ToLower(os.Getenv("ASSET_HANDLER")) == "filesystem" {
		return os.Open(fmt.Sprintf("%s/%s", os.Getenv("ASSET_PATH"), fileName))
	} else {
		return nil, fmt.Errorf("asset handler not supported")
	}
}

// Get All User's file
func (s *StaticImageHandler) GetAllUsersFile(username string) ([]fs.DirEntry, error) {
	allFile, err := s.GetAllFile()
	if err != nil {
		return nil, err
	}
	var userFile []fs.DirEntry
	for _, file := range allFile {
		if strings.Split(file.Name(), ".")[0] == username {
			userFile = append(userFile, file)
		}
	}
	return userFile, nil
}

func (s *StaticImageHandler) DeleteAllUserFile(username string) error {
	userFile, err := s.GetAllUsersFile(username)
	if err != nil {
		return err
	}
	for _, file := range userFile {
		err := s.DeleteFile(file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *StaticImageHandler) DeleteUserFile(username, fileName string) error {
	return s.DeleteFile(fmt.Sprintf("%s.%s", username, fileName))
}

func (s *StaticImageHandler) DeleteOldestFileIfNeeded(username string) error {
	userFile, err := s.GetAllUsersFile(username)
	if err != nil {
		return err
	}
	if len(userFile) > 4 {
		sort.Slice(userFile, func(i, j int) bool {
			iInfo, errI := userFile[i].Info()
			jInfo, errJ := userFile[j].Info()
			if errI != nil || errJ != nil {
				return false
			}
			return iInfo.ModTime().Before(jInfo.ModTime())
		})
		return s.DeleteFile(userFile[4].Name())
	}
	return nil
}
