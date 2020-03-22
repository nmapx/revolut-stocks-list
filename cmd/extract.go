package cmd

import (
	"bytes"
	"image"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"github.com/spf13/cobra"
)

var (
	input            []string
	output           string
	lang             []string
	whitelistTickers string
	whitelistNames   string
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract stocks tickers from the image.",
	Long:  `Check README for more details!`,
	Run: func(cmd *cobra.Command, args []string) {
		extract(input, output, lang, whitelistTickers, whitelistNames)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	var inputExample, langExample []string
	inputExample = append(inputExample, "./input.jpg")
	langExample = append(langExample, "eng")

	extractCmd.PersistentFlags().StringSliceVar(&input, "input", inputExample, "Input image eg. './input.jpg' - use multiple times for many files")
	extractCmd.PersistentFlags().StringVar(&output, "output", "./output.csv", "Output file eg. './output.csv'")
	extractCmd.PersistentFlags().StringSliceVar(&lang, "lang", langExample, "Languages used by the OCR - use multiple times for many languages")
	extractCmd.PersistentFlags().StringVar(&whitelistTickers, "whitelist-tickers", "QWERTYUIOPASDFGHJKLZXCVBNM", "Tickers whitelist eg. 'abcDEF123'")
	extractCmd.PersistentFlags().StringVar(&whitelistNames, "whitelist-names", "", "Companies names whitelist eg. 'abcDEF123'")
}

func extract(in []string, out string, languages []string, whT string, whN string) (csv *bytes.Buffer) {
	var err error
	var src image.Image
	var tmp *image.NRGBA
	var bounds image.Rectangle
	var client *gosseract.Client
	var i, c int
	var ocrText string
	var ocrSliceNames, ocrSliceTickers []string

	_, filename, _, _ := runtime.Caller(0)
	cacheDir := filepath.Dir(filename) + "/../cache"

	csv = new(bytes.Buffer)

	log.Println("Script initialized, starting")

	for _, element := range in {
		log.Printf("Loading image: %v", element)
		if src, err = imaging.Open(element); err != nil {
			panic(err)
		}

		bounds = src.Bounds()

		log.Println("Transforming the source image")
		tmp = imaging.CropAnchor(src, bounds.Max.X, int(float64(bounds.Max.Y)-float64(bounds.Max.X)/2.9), imaging.Bottom)
		tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)-float64(bounds.Max.X)/5.2), bounds.Max.Y, imaging.Right)
		tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)/3.3), bounds.Max.Y, imaging.Left)

		log.Println("Saving temporary file")
		if err = imaging.Save(tmp, cacheDir+"/tmp.jpg"); err != nil {
			panic(err)
		}

		log.Printf("Initializing OCR, setting languages %v", languages)
		client = gosseract.NewClient()
		defer client.Close()
		client.SetImage(cacheDir + "/tmp.jpg")
		client.SetLanguage(languages...)

		log.Printf("Processing company names")
		if len(whN) > 0 {
			client.SetWhitelist(whN)
		}
		ocrText, _ = client.Text()
		ocrSliceNames = unify(strings.Split(ocrText, "\n"))

		log.Printf("Processing tickers")
		client.SetWhitelist(whT)
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
	if err = ioutil.WriteFile(out, csv.Bytes(), 0644); err != nil {
		panic(err)
	}

	log.Println("Done")
	return
}

func unify(slice []string) (result []string) {
	for _, value := range slice {
		value = strings.ReplaceAll(value, ",", "")
		if len(value) == 0 {
			continue
		}
		result = append(result, value)
	}
	return
}
