package cmd

import (
    "log"
    "strings"
    "image"
    "image/color"
    "bytes"
    "io/ioutil"

    "github.com/spf13/cobra"
    "github.com/otiai10/gosseract/v2"
    "github.com/disintegration/imaging"
)

var (
    input []string
    output string
    lang []string
    whitelist string
)

var extractCmd = &cobra.Command{
	Use: "extract",
    Short: "",
    Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
        var err error
        var src image.Image
        var tmp *image.NRGBA
        var bounds image.Rectangle
        var client *gosseract.Client
        var i, c int
        var ocrText string
        var ocrSlice []string
        var tickerMode bool

        csv := new(bytes.Buffer)

        for _, element := range input {
            if src, err = imaging.Open(element); err != nil {
                panic(err)
            }

            bounds = src.Bounds()

            tmp = imaging.New(0, 0, color.NRGBA{0, 0, 0, 0})
            tmp = imaging.CropAnchor(src, bounds.Max.X, int(float64(bounds.Max.Y)-float64(bounds.Max.X)/2.4), imaging.Bottom)
            tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)-float64(bounds.Max.X)/4.8), bounds.Max.Y, imaging.Right)
            tmp = imaging.CropAnchor(tmp, int(float64(bounds.Max.X)/7.2), bounds.Max.Y, imaging.Left)

            if err = imaging.Save(tmp, "./cache/tmp.jpg"); err != nil {
                panic(err)
            }
            
            client = gosseract.NewClient()
            defer client.Close()
            client.SetImage("./cache/tmp.jpg")
            client.SetLanguage(lang...)
            client.SetWhitelist(whitelist)

            ocrText, _ = client.Text()
            ocrSlice = strings.Split(ocrText, "\n")

            tickerMode = false
            c = 0
            for i = 0; i < len(ocrSlice); i++ {
                if len(ocrSlice[i]) == 0 {
                    continue
                }
                if tickerMode == false {
                    tickerMode = true
                    continue
                }
                csv.WriteString(ocrSlice[i] + "\n")
                tickerMode = false
                c++
            }

            log.Printf("Found %v tickers", c)
        }

        if err = ioutil.WriteFile(output, csv.Bytes(), 0644); err != nil {
            panic(err)
        }
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
