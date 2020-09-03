package main

import (
	"image"
	"os"
	
	"image/png"
	"io/ioutil"
	"path/filepath"
    
    "github.com/ThinkingLogic/jenks"
    
    "go.uber.org/zap"
)

var logger *zap.Logger

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("It's alive",
		zap.String("Vitality test:", "passed"),
	
	)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	items, _ := ioutil.ReadDir(".")
	var data []float64
	
    for _, item := range items {
        if !item.IsDir() {
        	var extension = filepath.Ext(item.Name())
        	if extension == ".png" {
				file, _ := os.Open(item.Name())
				defer file.Close()
				imgConf, _, _ := image.DecodeConfig(file)
				width := imgConf.Width
				height := imgConf.Height
				
				file, _ = os.Open(item.Name())
				defer file.Close()
				img, _, _ := image.Decode(file)
				steps := 5
				alphaCount := 0
				for x := 0; x < width;  x = x + steps {
					for y := 0; y < height; y = y + steps {		
						_, _, _, alpha := img.At(x, y).RGBA()
						if int(alpha) > 0 {
							alphaCount++
						}
					}
				}	
				logger.Info("Alpha",
					zap.String("File: ", item.Name()),
					zap.Int("Count: ", int(alphaCount)),
				)
				data = append(data, float64(alphaCount))						
        	}
        }
    }
    allBreaks := jenks.AllNaturalBreaks(data, 4)
    
    for _,breaks := range allBreaks {
    	logger.Info("Natural breaks",
    		zap.Float64s("Breaks", breaks),
    	)
    }
    

    
}