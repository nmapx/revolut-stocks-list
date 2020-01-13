package cmd

import (
	"bytes"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"github.com/spf13/cobra"
)

var (
	input     []string
	output    string
	lang      []string
	whitelist string
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract stocks tickers from the image.",
	Long:  `Check README for more details!`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var src image.Image
		var tmp *image.NRGBA
		var bounds image.Rectangle
		var client *gosseract.Client
		var i, c int
		var ocrText string
		var ocrSliceNames, ocrSliceTickers []string

		csv := new(bytes.Buffer)

		log.Println("Script initialized, starting")

		for _, element := range input {
			log.Printf("Loading image: %v", element)
			if src, err = imaging.Open(element); err != nil {
				panic(err)
			}

			bounds = src.Bounds()

			log.Println("Transforming the source image")
			tmp = imaging.New(0, 0, color.NRGBA{0, 0, 0, 0})
			tmp = imaging.CropAnchor(src, bounds.Max.X, int(float64(bounds.Max.Y)-float64(bounds.Max.X)/2.4), imaging.Bottom)
			tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)-float64(bounds.Max.X)/4.7), bounds.Max.Y, imaging.Right)
			tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)/3.3), bounds.Max.Y, imaging.Left)

			log.Println("Saving temporary file")
			if err = imaging.Save(tmp, "./cache/tmp.jpg"); err != nil {
				panic(err)
			}

			log.Printf("Initializing OCR, setting languages %v", lang)
			client = gosseract.NewClient()
			defer client.Close()
			client.SetImage("./cache/tmp.jpg")
			client.SetLanguage(lang...)

			log.Printf("Processing company names")
			ocrText, _ = client.Text()
			ocrSliceNames = unify(strings.Split(ocrText, "\n"))

			log.Printf("Processing tickers")
			client.SetWhitelist(whitelist)
			ocrText, _ = client.Text()
			ocrSliceTickers = unify(strings.Split(ocrText, "\n"))

			if len(ocrSliceNames) != len(ocrSliceTickers) {
				panic("Data mismatch - please try again with different image")
			}

			log.Println("Preparing results")
			c = 0
			for i = 0; i < len(ocrSliceTickers); i += 2 {
				csv.WriteString(ocrSliceTickers[i+1] + "," + ocrSliceNames[i] + "\n")
				c++
			}

			log.Printf("Found %v records", c)
		}

		log.Println("Saving CSV file")
		if err = ioutil.WriteFile(output, csv.Bytes(), 0644); err != nil {
			panic(err)
		}

		log.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	inputExample := make([]string, 1)
	inputExample[0] = "./input.jpg"

	langExample := make([]string, 1)
	langExample[0] = "eng"

	extractCmd.PersistentFlags().StringSliceVar(&input, "input", inputExample, "Input image eg. './input.jpg' - use multiple times for many files")
	extractCmd.PersistentFlags().StringVar(&output, "output", "./output.csv", "Output file eg. './output.csv'")
	extractCmd.PersistentFlags().StringSliceVar(&lang, "lang", langExample, "Languages used by the OCR - use multiple times for many languages")
	extractCmd.PersistentFlags().StringVar(&whitelist, "whitelist", "QWERTYUIOPASDFGHJKLZXCVBNM", "Whitelist eg. 'abcDEF'")
}

func unify(slice []string) []string {
	var result []string
	for _, value := range slice {
		if len(value) == 0 {
			continue
		}
		result = append(result, strings.ReplaceAll(value, ",", ""))
	}
	return result
}
