package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/signintech/gopdf"
)

const (
	PAGE_WIDTH    float64 = 21.0
	PAGE_HEIGHT   float64 = 29.7
	CELL_WIDTH    float64 = 4.0
	FONT_SIZE     int     = 10
	LINE_HEIGHT   float64 = 0.6
	MARGIN_BOTTOM float64 = 1.0

	MARGIN_TOP float64 = 1.0
)

var Students = []struct {
	SrNo   string
	USN    string
	Name   string
	Login  string
	Logout string
}{
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"2", "4AL21EC053", "Michael Smith", "09:00 AM", "01:00 PM"},
	{"3", "4AL21EC053", "zwexdcrtfvgybhunijmko,lp.ctfvgybhunjkm", "10:15 AM", "02:00 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrboooyyy", "10:15 AM", "02:00 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"1", "4AL21EC053", "John Doe", "08:30 AM", "12:30 PM"},
	{"4", "4AL21EC053", "AlexandertfygijoefnobfyuagHQRIJWEIUAHQJOChrbjjjjjjjjjjjjjoooyyy", "10:15 AM", "02:00 PM"},
}

func OuterBorderSection(pdf *gopdf.GoPdf) {
	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetLineWidth(0.05)
	pdf.Line(1, 1, 20, 1)
	pdf.Line(1, 1, 1, 28.7)
	pdf.Line(1, 28.7, 20, 28.7)
	pdf.Line(20, 1, 20, 28.7)
}

func findMainHeaderCordinates1(pdf *gopdf.GoPdf, spacing float64, text string) (float64, float64, error) {
	textWidth, err := pdf.MeasureTextWidth(text)

	if err != nil {
		return 0.0, 0.0, err
	}

	return (PAGE_WIDTH / 2) - (textWidth / 2), pdf.GetY() + spacing, nil
}
func findMainHeaderCordinates2(pdf *gopdf.GoPdf, spacing float64, text string) (float64, float64, error) {
	textWidth, err := pdf.MeasureTextWidth(text)

	if err != nil {
		return 0.0, 0.0, err
	}

	return ((PAGE_WIDTH / 2) - 3.05) - (textWidth / 2), pdf.GetY() + spacing, nil
}

func AddResizedImage(pdf *gopdf.GoPdf, imgPath string, x, y, maxW, maxH float64) {

	file, err := os.Open(imgPath)
	if err != nil {
		log.Fatalf("Error opening image: %v", err)
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Fatalf("Error decoding image: %v", err)
	}
	imgW, imgH := float64(img.Width), float64(img.Height)

	scale := min(maxW/imgW, maxH/imgH)
	newW, newH := imgW*scale, imgH*scale

	if err := pdf.Image(imgPath, x, y, &gopdf.Rect{W: newW, H: newH}); err != nil {
		log.Fatalf("Error adding image: %v", err)
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func GeneratePdf() {
	pdf := gopdf.GoPdf{}

	pdf.Start(gopdf.Config{
		PageSize: *gopdf.PageSizeA4,
		Unit:     gopdf.UnitCM,
	})

	if err := pdf.AddTTFFont("bold-font", "./font-family/Roboto/static/Roboto-Bold.ttf"); err != nil {
		log.Fatal(err)
	}

	if err := pdf.AddTTFFont("light-font", "./font-family/Roboto/static/Roboto-Regular.ttf"); err != nil {
		log.Fatal(err)
	}

	pdf.AddHeader(func() {
		companyName := "VITHSUTRA TECHNOLOGIES Pvt. Ltd." // 5
		phone := "Phone: +919845849116"
		email := "Email: contact@vithsutra.com"
		web := "Web: www.vithsutra.com"

		OuterBorderSection(&pdf)

		if err := pdf.SetFont("bold-font", "", 17); err != nil {
			log.Fatal(err)
		}

		x, y, err := findMainHeaderCordinates1(&pdf, 2, companyName)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(companyName)

		if err := pdf.SetFont("light-font", "", 12); err != nil {
			log.Fatal(err)
		}

		x, y, err = findMainHeaderCordinates2(&pdf, 0.6, phone)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(phone)

		pdf.SetXY(x+5, y)
		pdf.Text(email)

		x, y, err = findMainHeaderCordinates1(&pdf, 0.6, web)
		if err != nil {
			log.Fatal(err)
		}

		pdf.SetXY(x, y)
		pdf.Text(web)

		pdf.SetStrokeColor(0, 0, 0)
		pdf.SetLineWidth(0.05)

		pdf.Line(1.9, 4, 19.1, 4)

		AddResizedImage(&pdf, "./finallogo.png", 1.4, 1.5, 3.1, 3.1)

	})

	GenerateStudentReport(&pdf, Students)

	err := pdf.WritePdf("output.pdf")
	if err != nil {
		log.Fatal("Failed to generate PDF:", err)
	}

}

func DisplayDate(pdf *gopdf.GoPdf) {

	if err := pdf.SetFont("bold-font", "", 14); err != nil {
		log.Fatal(err)
	}
	date := "12/03/2003"
	machineId := "VS242S45_"
	slotStatus := "morning"

	// x, y := 1.8, pdf.GetY()+1
	x, y := 1.8, 4.552777777777778
	fmt.Println(y)
	textX := x + 0.15
	textY := y + 0.6
	pdf.SetXY(textX, textY)
	pdf.Text("Date: ")

	if err := pdf.SetFont("light-font", "", 14); err != nil {
		log.Fatal(err)
	}
	textX = x + 1.4
	textY = y + 0.6
	pdf.SetXY(textX, textY)
	pdf.Text(date)

	if err := pdf.SetFont("bold-font", "", 14); err != nil {
		log.Fatal(err)
	}
	textX = x + 12
	textY = y + 0.6
	pdf.SetXY(textX, textY)
	pdf.Text("Machine ID: ")

	if err := pdf.SetFont("light-font", "", 14); err != nil {
		log.Fatal(err)
	}
	textX = x + 14.8
	textY = y + 0.6
	pdf.SetXY(textX, textY)
	pdf.Text(machineId)

	if err := pdf.SetFont("bold-font", "", 14); err != nil {
		log.Fatal(err)
	}

	x, y = 1.8, pdf.GetY()+0.8

	textX = x + 0.15
	textY = y + 0.3
	pdf.SetXY(textX, textY)
	pdf.Text("Slot: ")

	if err := pdf.SetFont("light-font", "", 14); err != nil {
		log.Fatal(err)
	}

	textX = x + 1.3
	textY = y + 0.3
	pdf.SetXY(textX, textY)
	pdf.Text(slotStatus)

}

func createTableForStudentsfirst(pdf *gopdf.GoPdf, isfirstPage bool) {

	if err := pdf.SetFont("bold-font", "", 14); err != nil {
		log.Fatal(err)
	}

	if isfirstPage {
		x, y := 1.0, pdf.GetY()+0.7

		pdf.SetStrokeColor(0, 0, 0)
		pdf.SetLineWidth(0.05)
		pdf.Line(1, 7, 20, 7)

		pdf.SetStrokeColor(0, 0, 0)
		pdf.SetLineWidth(0.05)
		pdf.Line(1, 8, 20, 8)

		lineX := x + 1.6
		lineY := y + 0.05

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		//first text
		textX := x + 0.15
		textY := y + 0.7
		pdf.SetXY(textX, textY)
		pdf.Text("Sr.No.")

		//second text
		textX = textX + 2.75
		pdf.SetXY(textX, textY)
		pdf.Text("USN")

		//second line
		lineX = (PAGE_WIDTH/2)/2 + 1
		lineY = y + 0.05

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		//third line
		lineX = PAGE_WIDTH/2 + 2
		lineY = y + 0.05

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		//third text
		textX = textX + 4.7
		pdf.SetXY(textX, textY)
		pdf.Text("NAME")

		//fourth line
		lineX = (PAGE_WIDTH / 2) + 5.65

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		// //fourth text
		textX = (PAGE_WIDTH / 2) + 3.15
		pdf.SetXY(textX, textY)
		pdf.Text("LOGIN")

		textX = (PAGE_WIDTH / 2) + 6.65
		pdf.SetXY(textX, textY)
		pdf.Text("LOGOUT")

	} else {
		x, y := 1.0, 4.7

		pdf.SetStrokeColor(0, 0, 0)
		pdf.SetLineWidth(0.05)
		pdf.Line(1, 4.7, 20, 4.7)

		pdf.SetStrokeColor(0, 0, 0)
		pdf.SetLineWidth(0.05)
		pdf.Line(1, 5.7, 20, 5.7)

		lineX := x + 1.6
		lineY := y + 0.05

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		//first text
		textX := x + 0.15
		textY := y + 0.7
		pdf.SetXY(textX, textY)
		pdf.Text("Sr.No.")

		//second text
		textX = textX + 2.75
		pdf.SetXY(textX, textY)
		pdf.Text("USN")

		//second line
		lineX = (PAGE_WIDTH/2)/2 + 1
		lineY = y + 0.05

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		//third line
		lineX = PAGE_WIDTH/2 + 2
		lineY = y + 0.05

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		//third text
		textX = textX + 4.7
		pdf.SetXY(textX, textY)
		pdf.Text("NAME")

		//fourth line
		lineX = (PAGE_WIDTH / 2) + 5.65

		pdf.Line(lineX, lineY, lineX, PAGE_HEIGHT-1)

		// //fourth text
		textX = (PAGE_WIDTH / 2) + 3.15
		pdf.SetXY(textX, textY)
		pdf.Text("LOGIN")

		textX = (PAGE_WIDTH / 2) + 6.65
		pdf.SetXY(textX, textY)
		pdf.Text("LOGOUT")

	}

}

func createStudentRow(pdf *gopdf.GoPdf, student struct {
	SrNo   string
	USN    string
	Name   string
	Login  string
	Logout string
}, startY float64) float64 {

	nameLines := splitText(student.Name, 27)
	lineCount := len(nameLines)
	rowHeight := 1.0 + (float64(lineCount-1) * 0.5)

	if err := pdf.SetFont("light-font", "", 12); err != nil {
		log.Fatal(err)
	}

	calculateY := func(baseY float64) float64 {
		return startY + (rowHeight-1)*0.3 + baseY
	}

	pdf.SetXY(1.6, calculateY(0.6))
	pdf.Text(student.SrNo)

	pdf.SetXY(3.2, calculateY(0.6))
	pdf.Text(student.USN)

	nameStartY := startY + 0.6
	for i, line := range nameLines {
		pdf.SetXY(6.7, nameStartY+(float64(i)*0.5))
		pdf.Text(line)
	}

	pdf.SetXY(13.4, calculateY(0.6))
	pdf.Text(student.Login)

	pdf.SetXY(17.2, calculateY(0.6))
	pdf.Text(student.Logout)

	pdf.SetStrokeColor(0, 0, 0)
	pdf.SetLineWidth(0.05)
	pdf.Line(1, startY+rowHeight, 20, startY+rowHeight)

	return startY + rowHeight
}

func splitText(text string, maxLen int) []string {
	var lines []string
	for len(text) > maxLen {
		lines = append(lines, text[:maxLen])
		text = text[maxLen:]
	}
	if len(text) > 0 {
		lines = append(lines, text)
	}
	return lines
}

func GenerateStudentReport(pdf *gopdf.GoPdf, students []struct {
	SrNo   string
	USN    string
	Name   string
	Login  string
	Logout string
}) {
	pdf.AddPage()
	DisplayDate(pdf)
	createTableForStudentsfirst(pdf, true)

	startY := pdf.GetY() + 0.4

	for i := 0; i < len(students); i++ {
		student := students[i]

		rowHeight := 1.0 + (float64(len(splitText(student.Name, 27))-1) * 0.5)

		if startY+rowHeight > PAGE_HEIGHT-MARGIN_BOTTOM {
			pdf.AddPage()
			startY = 5.7
			createTableForStudentsfirst(pdf, false)
		}

		startY = createStudentRow(pdf, student, startY)

	}
}

func main() {
	GeneratePdf()

}
