package main
 
import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main(){
unzip("/home/manojk/workspace/project/cb-gopay-apr/billers/billers.zip", "billers/")
}
 

func unzip(archive, target string) {
	zipReader, _ := zip.OpenReader(archive)
	for _, file := range zipReader.Reader.File {
 
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()
 
		targetDir := target
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)
 
		if file.FileInfo().IsDir() {
			log.Println("Directory Created:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			log.Println("File extracted:", file.Name)
 
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()
 
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
