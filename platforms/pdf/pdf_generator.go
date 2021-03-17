package pdf

import (
	"fmt"
	"log"
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

// GetThePdf function
func GetThePdf(fileDirectory string) string {
	pdfg, erra := wkhtmltopdf.NewPDFGenerator()
	if erra != nil {
		fmt.Println("Error While Generating the Pdf ")
		return ""
	}
	pdfg.Dpi.Set(30)
	pdfg.ImageDpi.Set(30)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	homepath := os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY") + "pdf/" + helper.GenerateRandomString(5, helper.CHARACTERS) + ".pdf"
	page := wkhtmltopdf.NewPage(fileDirectory)
	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)
	pdfg.AddPage(page)
	pdfCreationError := pdfg.Create()
	if pdfCreationError != nil {
		println(pdfCreationError.Error())
		fmt.Println("Error while Creating the pdf file ")
		return ""
	}
	// Generating Random Name to Be Output Name
	writingError := pdfg.WriteFile(homepath)
	if writingError != nil {
		log.Println("Error While Writing the File to The Directory Name ", homepath)
		return ""
	}
	os.Remove(fileDirectory)
	return homepath
}
