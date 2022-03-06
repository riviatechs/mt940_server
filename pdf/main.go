package pdf

import (
	"github.com/MalukiMuthusi/logger"
	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/riviatechs/mt940_server/cloudstore"
	"go.uber.org/zap"
)

func PDFGenerator() {
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		logger.Fatalf("PDFGenerator", zap.Error(err))
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtml.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	// Create a new input page from an URL
	page := wkhtml.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		logger.Fatalf("PDFGenerator", zap.Error(err))
	}

	c := cloudstore.CloudStore{}
	go c.UploadFile(pdfg.Buffer(), "first.pdf")
}
