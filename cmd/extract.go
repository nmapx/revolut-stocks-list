package cmd

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"github.com/piquette/finance-go/quote"
	"image/png"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

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
	extractCmd.PersistentFlags().StringVar(&whitelistTickers, "whitelist-tickers", "QWERTYUIOPASDFGHJKLZXCVBNM.", "Tickers whitelist eg. 'abcDEF123'")
	extractCmd.PersistentFlags().StringVar(&whitelistNames, "whitelist-names", "", "Companies names whitelist eg. 'abcDEF123'")
}

func extract(in []string, out string, languages []string, whT string, whN string) (csv *bytes.Buffer, errors *bytes.Buffer) {
	csv = new(bytes.Buffer)
	errors = new(bytes.Buffer)

	log.Println("Script initialized, starting")
	for _, element := range in {
		log.Printf("Loading image: %v", element)
		src, err := imaging.Open(element)
		if err != nil {
			panic(err)
		}

		bounds := src.Bounds()
		log.Println("Transforming the source image")
		tmp := imaging.CropAnchor(src, int(float64(bounds.Max.X)-float64(bounds.Max.X)/4.5), bounds.Max.Y, imaging.Right)
		tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)/2.5), bounds.Max.Y, imaging.Left)

		buff := new(bytes.Buffer)
		if err = png.Encode(buff, tmp); err != nil {
			fmt.Println("failed to create buffer", err)
		}

		log.Printf("Initializing OCR, setting languages %v", languages)
		client := gosseract.NewClient()
		defer client.Close()
		client.SetImageFromBytes(buff.Bytes())
		client.SetLanguage(languages...)

		log.Printf("Processing company names")
		if len(whN) > 0 {
			client.SetWhitelist(whN)
		}
		ocrText, _ := client.Text()
		ocrSliceNames := unify(strings.Split(ocrText, "\n"))

		log.Printf("Processing tickers")
		client.SetWhitelist(whT)
		ocrText, _ = client.Text()
		ocrSliceTickers := unify(strings.Split(ocrText, "\n"))

		if len(ocrSliceNames) != len(ocrSliceTickers) {
			panic("Data mismatch - please try again with different image")
		}

		log.Println("Preparing results")

		c := 0
		state := 0
		row := ""

		for i := 0; i < len(ocrSliceNames); i++ {
			if state == 0 {
				row = ocrSliceNames[i]
				state = 1
				continue
			} else if state == 1 {
				matchedNames, _ := regexp.MatchString(`^[A-Z\d]+[a-z]{2,}.*$`, ocrSliceNames[i])
				if matchedNames {
					row = row + "..." + ocrSliceNames[i]
					state = 2
					continue
				}
			}
			splittedName := strings.Split(row, " ")
			row = ocrSliceTickers[i] + "," + row


			q, err := quote.Get(ocrSliceTickers[i])
			if err != nil || q == nil || !strings.Contains(strings.ToLower(q.ShortName), strings.ToLower(splittedName[0])) {
				errors.WriteString(row + "\n")
			} else {
				csv.WriteString(row + "\n")
			}

			row = ""
			state = 0
			c++
		}

		log.Printf("Found %v records", c)
	}

	log.Println("Saving CSV file(s)")
	if err := ioutil.WriteFile(out, csv.Bytes(), 0644); err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile("./errors.csv", errors.Bytes(), 0644); err != nil {
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
