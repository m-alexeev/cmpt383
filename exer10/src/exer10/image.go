package exer10

import(
	"math"
	"image"
	"image/png"
	"os"
	"fmt"
)


func withinRad(x int, y int, inner int, outer int) bool{
	quad := math.Pow((math.Abs(float64(x))-100),2) + math.Pow((math.Abs(float64(y))-100),2)
	// fmt.Println("%v, %v, %v", quad, math.Pow(float64(inner+5),2), math.Pow(float64(outer+5),2))
	//  
	if  quad >= math.Pow(float64(inner),2) && quad <= math.Pow(float64(outer),2){
		return true
	}
	return false
}

func closeFile(f *os.File ){
	err := f.Close()
	if err != nil{
		panic(err)
	}
}


func DrawCircle(outerRadius int, innerRadius int, outputFile string){
	img := image.NewRGBA(image.Rect(0,0,200,200))
	for y := 0; y < 200; y ++{
		for x:= 0; x <200; x ++{
			inside := withinRad(x,y ,innerRadius,outerRadius)
			for p :=0 ; p < 3; p++{
				if inside{
					img.Pix[y*200*4 + 4*x+p] = 0
				}else {
					img.Pix[y*200*4 + 4*x+p] = 255
				}
			}		
			img.Pix[y*200*4 + 4*x+3] = 255

		}
	}


	//Error checking was sourced from the golang docs
	outFile, err := os.Create(outputFile)
	if err != nil{
		fmt.Println(err)
	}
	if err := png.Encode(outFile,img); err !=  nil{
		panic(err)
	}
	
	defer closeFile(outFile)
	
}