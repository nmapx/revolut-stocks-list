# Revolut Stocks List

Extract Revolut stocks list from the list screenshot(s) (use your mobile app for this).

I made this app since there is no official list of all available Revolut stocks.

**You can get a full list at the bottom of this page. It may not be up to date but I will do my best to keep it updated!**

## Technologies

- Go
- Docker with **compose**
- Tesseract OCR
- Make

## Requirements

- Docker >= 19.03 with **compose** >= 1.25
- Make >= 4.1 *(optional)*
> Note: use appropriate commands (found in Makefiles) without Make

## Setup guide

1. Clone the repository and navigate to the project root
2. Copy `.env.dist` file to `.env` and fill missing parameters
> Note: HOST_UID is your system user ID
3. Build docker image `make build`
4. Start container `make up` (development purposes only)
5. Run the script by executing appropriate command inside the container (see below)
6. Compile the app by executing `make -f Makefile.native build` inside the container

## Basic usage

### Docker

```bash
docker exec {CONTAINER_NAME} go run . extract
```

or

```bash
docker-compose -p {PROJECT_NAME} exec app go run . extract
```

### Binary

Binary not supported yet - use Docker environment.

## Advanced usage

### Input files (images/screenshots)

You can specify multiple input files by adding more parameters. By default it's processing 1 file (filepath `./input.jpg`). See example below.

```bash
... extract --input ./this/is/input/file_1.jpg --input ./this/is/input/file_2.jpg --input ./this/is/input/file_3.jpg
```

### Output file (CSV)

Script generates 1 output file (including all input files in it). By default it's `./output.csv` - you can change the filepath. See example below.

```bash
... extract --output ./this/is/output/file.csv
```

### Languages

By default only English is being used by the Tesseract library. In order to use also Polish and Russian you can add optional parameters. See example below.

```bash
... extract --lang eng --lang pol --lang rus
```

### Whitelist

By default the whitelist includes only A-Z characters (upper-case). You can change it by passing an optional parameter. See example below.

```bash
... extract --whitelist abcDEF123
```

## The list

As the result you will get a file with all recognized tickers. Company name is not supported yet (I'm adding it manually).

Below you can find already extracted list of all available stocks. I will update it from time to time.

| Ticker | Company name                                   | 
|-------|-------------------------------------------------| 
| FOXA  | Fox Corp Class A                                | 
| TWOU  | 2U Inc                                          | 
| MMM   | 3M Co                                           | 
| WUBA  | 58.com Inc                                      | 
| ABT   | Abbott Laboratories                             | 
| ABBV  | AbbVie Inc                                      | 
| ANF   | Abercrombie & Fitch Co.                         | 
| ACHN  | Achillion Pharmaceuticals, Inc.                 | 
| ATVI  | Activision Blizzard, Inc.                       | 
| ADBE  | Adobe Inc                                       | 
| AAP   | Advance Auto Parts, Inc.                        | 
| AMD   | Advanced Micro Devices, Inc.                    | 
| AES   | AES Corp                                        | 
| AFL   | AFLAC Incorporated                              | 
| A     | Agilent Technologies Inc                        | 
| AGNC  | AGNC Investment Corp                            | 
| AIG   | American International Group Inc                | 
| AL    | Air Lease Corp Class A                          | 
| APD   | Air Products & Chemicals, Inc.                  | 
| AKS   | AK Steel Holding Corporation                    | 
| AKAM  | Akamai Technologies, Inc.                       | 
| AA    | Alcoa Corp                                      | 
| ALXN  | Alexion Pharmaceuticals, Inc.                   | 
| BABA  | Alibaba Group Holding Ltd - ADR                 | 
| ALGN  | Align Technology, Inc.                          | 
| ALLT  | Allot Ltd                                       | 
| MDRX  | Allscripts Healthcare Solutions Inc             | 
| ALL   | Allstate Corp                                   | 
| ALLY  | Ally Financial Inc                              | 
| GOOGL | Alphabet Inc Class A                            | 
| GOOG  | Alphabet Inc Class C                            | 
| AABA  | --error--                                       | 
| ATUS  | Altice USA Inc                                  | 
| MO    | Altria Group Inc                                | 
| AMZN  | Amazon.com, Inc.                                | 
| ABEV  | AMBEV S A/S ADR                                 | 
| AMC   | AMC Entertainment Holdings Inc                  | 
| AAL   | American Airlines Group Inc                     | 
| AXL   | American Axle & Manufact. Holdings, Inc.        | 
| AEO   | American Eagle Outfitters                       | 
| AEP   | American Electric Power Company Inc             | 
| AXP   | American Express Company                        | 
| AMH   | American Homes 4 Rent Class A                   | 
| AMT   | American Tower Corp                             | 
| COLD  | AmeriCold Realty Trust                          | 
| AME   | AMETEK, Inc.                                    | 
| AMGN  | Amgen, Inc.                                     | 
| FOLD  | Amicus Therapeutics, Inc.                       | 
| AMRX  | Amneal Pharmaceuticals Inc                      | 
| APH   | Amphenol Corporation                            | 
| ADI   | Analog Devices, Inc.                            | 
| PLAN  | Anaplan Inc                                     | 
| ANGI  | ANGI Homeservices Inc                           | 
| AU    | AngloGold Ashanti Limited                       | 
| NLY   | ANNALY CAP MGMT/SH                              | 
| ANSS  | ANSYS, Inc.                                     | 
| AM    | Antero Midstream Corp                           | 
| AR    | Antero Resources Corp                           | 
| ANTM  | Anthem Inc                                      | 
| AOS   | A. O. Smith Corp                                | 
| APA   | Apache Corporation                              | 
| ARI   | Apollo Commercial Real Est. Finance Inc         | 
| AAPL  | Apple Inc.                                      | 
| AMAT  | Applied Materials, Inc.                         | 
| ATR   | AptarGroup, Inc.                                | 
| ARMK  | Aramark                                         | 
| ADM   | Archer Daniels Midland Co                       | 
| ARNC  | Arconic Inc                                     | 
| ARCC  | Ares Capital Corporation                        | 
| ANET  | Arista Networks Inc                             | 
| ARR   | ARMOUR Residential REIT, Inc.                   | 
| ARQL  | ArQule, Inc.                                    | 
| ASNA  | Ascena Retail Group Inc                         | 
| ASML  | ASML Holding NV                                 | 
| HOME  | At Home Group Inc                               | 
| T     | AT&T Inc.                                       | 
| ADSK  | Autodesk, Inc.                                  | 
| ATHM  | Autohome Inc                                    | 
| ADP   | Automatic Data Processing                       | 
| AZO   | AutoZone, Inc.                                  | 
| AVLR  | Avalara Inc                                     | 
| AVB   | AvalonBay Communities Inc                       | 
| EQH   | AXA Equitable Holdings Inc                      | 
| AXTA  | Axalta Coating Systems Ltd                      | 
| BTG   | B2Gold Corp.                                    | 
| BIDU  | Baidu Inc                                       | 
| BKR   | Baker Hughes Co                                 | 
| BLL   | Ball Corporation                                | 
| BBD   | Banco Bradesco SA                               | 
| BMA   | Banco Macro SA ADR Class B                      | 
| BSBR  | Banco Santander Brasil SA ADR                   | 
| BAC   | Bank of America Corp                            | 
| GOLD  | Barrick Gold Corp                               | 
| BHC   | Bausch Health Companies Inc                     | 
| BAX   | Baxter International Inc                        | 
| BBAR  | Banco Bbva Argentina SA                         | 
| BDX   | Becton Dickinson and Co                         | 
| BDC   | Belden Inc.                                     | 
| BRKB  | Berkshire Hathaway Inc. Class B                 | 
| BBY   | Best Buy Co Inc                                 | 
| BYND  | Beyond Meat Inc                                 | 
| BHP   | BHP Group Ltd                                   | 
| BILI  | Bilibili Inc - ADR                              | 
| BIO   | Bio-Rad Laboratories, Inc. Class A Common Stock | 
| BIIB  | Biogen Inc                                      | 
| BMRN  | BioMarin Pharmaceutical Inc.                    | 
| BITA  | Bitauto Hldg Ltd                                | 
| BJ    | BJs Wholesale Club Holdings Inc                 | 
| BB    | BlackBerry Ltd                                  | 
| BLK   | BlackRock, Inc.                                 | 
| BX    | Blackstone Group Inc                            | 
| BK    | Bank of New York Mellon Corp                    | 
| BA    | Boeing Co                                       | 
| BKNG  | Booking Holdings Inc                            | 
| BAH   | Booz Allen Hamilton Holding Corporation         | 
| BWA   | BorgWarner Inc.                                 | 
| BSX   | Boston Scientific Corporation                   | 
| BOX   | Box Inc                                         | 
| BRFS  | BRF S.A. (ADR) common stock                     | 
| BMY   | Bristol-Myers Squibb Co                         | 
| BMYRT | --error--                                       | 
| BRX   | Brixmor Property Group Inc                      | 
| AVGO  | Broadcom Inc                                    | 
| BAM   | Brookfield Asset Management Inc                 | 
| BIP   | Brookfield Infrastructure Partners L.P.         | 
| BEP   | Brookfield Renewable Partners LP                | 
| BFB   | --error--                                       | 
| BG    | Bunge Ltd                                       | 
| COG   | Cabot Oil & Gas Corporation                     | 
| CDNS  | Cadence Design Systems Inc                      | 
| CZR   | Caesars Entertainment Corporation Common Stock  | 
| CRC   | California Resources Corp                       | 
| CPE   | Callon Petroleum Company                        | 
| CAJ   | Canon Inc                                       | 
| COF   | Capital One Financial Corp.                     | 
| CAH   | Cardinal Health Inc                             | 
| CG    | Carlyle Group Inc                               | 
| CCL   | Carnival Corp                                   | 
| CRZO  | --error--                                       | 
| CARS  | Cars.com Inc                                    | 
| CAT   | Caterpillar Inc.                                | 
| CBL   | CBL & Associates Properties, Inc.               | 
| CBOE  | Cboe Global Markets Inc                         | 
| CBRE  | CBRE Group Inc                                  | 
| CELG  | CELGENE ORD                                     | 
| CX    | Cemex SAB de CV ADR                             | 
| CVE   | Cenovus Energy Inc                              | 
| CNC   | Centene Corp                                    | 
| CDEV  | Centennial Resource Development Inc             | 
| CNP   | CenterPoint Energy, Inc.                        | 
| CTL   | Centurylink Inc                                 | 
| CERN  | Cerner Corporation                              | 
| CF    | CF Industries Holdings, Inc.                    | 
| SCHW  | Charles Schwab Corporation Common Stock         | 
| CHTR  | Charter Communications Inc                      | 
| CHKP  | Check Point Software Technologies Ltd.          | 
| SM    | SM Energy Co                                    | 
| CC    | Chemours Co                                     | 
| LNG   | Cheniere Energy, Inc.                           | 
| CHK   | Chesapeake Energy Corporation                   | 
| CVX   | Chevron Corporation                             | 
| CHWY  | Chewy Inc                                       | 
| CHS   | Chico's FAS, Inc.                               | 
| CIM   | CHIMERA INVT CO/SH NEW                          | 
| CHL   | China Mobile Ltd.                               | 
| ZNH   | China Southern Airlines Co Ltd ADR Class H      | 
| CMG   | Chipotle Mexican Grill, Inc.                    | 
| CIEN  | Ciena Corporation                               | 
| C     | Citigroup Inc                                   | 
| XEC   | Cimarex Energy Co                               | 
| CTAS  | Cintas Corporation                              | 
| CSCO  | Cisco Systems, Inc.                             | 
| C     | Citigroup Inc                                   | 
| CFG   | Citizens Financial Group Inc                    | 
| CTXS  | Citrix Systems, Inc.                            | 
| CCO   | Clear Channel Outdoor Holdings Inc              | 
| CWEN  | Clearway Energy Inc Class C                     | 
| CLF   | Cleveland-Cliffs Inc                            | 
| CLDR  | Cloudera Inc                                    | 
| NET   | Cloudflare Inc                                  | 
| CLVS  | Clovis Oncology Inc                             | 
| CME   | CME Group Inc                                   | 
| CNX   | CNX Resources Corp                              | 
| KO    | Coca-Cola Co                                    | 
| CDE   | Coeur Mining Inc                                | 
| CGNX  | Cognex Corporation                              | 
| CTSH  | Cognizant Technology Solutions Corp             | 
| L     | Loews Corporation                               | 
| CLNY  | Colony Capital Inc                              | 
| CMCSA | Comcast Corporation                             | 
| CMA   | Comerica Incorporated                           | 
| COMM  | Commscope Holding Company Inc                   | 
| CYH   | Community Health Systems                        | 
| SBS   | Companhia de Saneamento Bsc DEDSP               | 
| CIG   | COMPANHIA ENERG/$.01 RED PFD SH                 | 
| SID   | Companhia Siderurgica Nacional                  | 
| BVN   | Compania de Minas Buenaventura SAA              | 
| CAG   | Conagra Brands Inc                              | 
| CX    | Cemex SAB de CV ADR                             | 
| CNDT  | Conduent Inc                                    | 
| CP    | Canadian Pacific Railway Ltd                    | 
| ED    | Consolidated Edison, Inc.                       | 
| STZ   | Constellation Brands, Inc. Class A              | 
| CLR   | Continental Resources, Inc.                     | 
| GLW   | Corning Incorporated                            | 
| CTVA  | Corteva Inc                                     | 
| COST  | Costco Wholesale Corporation                    | 
| COTY  | Coty Inc                                        | 
| BAP   | Credicorp Ltd.                                  | 
| CROX  | Crocs, Inc.                                     | 
| CRWD  | Crowdstrike Holdings Inc                        | 
| C     | Citigroup Inc                                   | 
| CSX   | CSX Corporation                                 | 
| CVS   | CVS Health Corp                                 | 
| CY    | Cypress Semiconductor Corporation               | 
| DAN   | Dana Inc                                        | 
| DHR   | Danaher Corporation                             | 
| DDOG  | Datadog Inc                                     | 
| DVA   | Davita Inc                                      | 
| DELL  | Dell Technologies Inc                           | 
| DAL   | Delta Air Lines, Inc.                           | 
| DNR   | Denbury Resources Inc.                          | 
| DVN   | Devon Energy Corp                               | 
| DO    | Diamond Offshore Drilling Inc                   | 
| DLR   | DIGITAL RLTY TR/SH                              | 
| APPS  | Digital Turbine Inc                             | 
| DFS   | Discover Financial Services                     | 
| DISCA | DISCOVERY COMMUNICATIONS INC. Common Stock      | 
| DISCK | Discovery Inc Series C                          | 
| DISH  | DISH Network Corp                               | 
| DHC   | Diversified Healthcare Trust                    | 
| DOCU  | Docusign Inc                                    | 
| DG    | Dollar General Corp.                            | 
| DLTR  | Dollar Tree, Inc.                               | 
| D     | Dominion Energy Inc                             | 
| DPZ   | Domino's Pizza, Inc.                            | 
| DHI   | D. R. Horton Inc                                | 
| DBX   | Dropbox Inc                                     | 
| DUK   | Duke Energy Corp                                | 
| DRE   | Duke Realty Corp                                | 
| DNKN  | Dunkin Brands Group Inc                         | 
| DD    | DuPont de Nemours Inc                           | 
| DXC   | DXC Technology Co                               | 
| ETFC  | E*TRADE Financial Corp                          | 
| EBAY  | eBay Inc                                        | 
| ECL   | Ecolab Inc.                                     | 
| EIX   | Edison International                            | 
| EW    | Edwards Lifesciences Corp                       | 
| ELAN  | Elanco Animal Health Inc                        | 
| EGO   | Eldorado Gold Corp                              | 
| EA    | Electronic Arts Inc.                            | 
| ESI   | Element Solutions Inc                           | 
| LY    | --error--                                       | 
| ERJ   | Embraer SA                                      | 
| EMR   | Emerson Electric Co.                            | 
| ECA   | Encana Corp                                     | 
| ET    | Energy Transfer LP Unit                         | 
| ENIA  | Enel Americas SA                                | 
| ENLC  | EnLink Midstream LLC                            | 
| EPD   | Enterprise Products Partners L.P.               | 
| EOG   | EOG Resources Inc                               | 
| EQT   | EQT Corporation                                 | 
| EFX   | Equifax Inc.                                    | 
| EQIX  | Equinix Inc                                     | 
| EQNR  | Equinor ASA                                     | 
| ETRN  | Equitrans Midstream Corp                        | 
| EQR   | Equity Residential                              | 
| EL    | Estee Lauder Companies Inc                      | 
| ETSY  | Etsy Inc                                        | 
| EB    | Eventbrite Inc                                  | 
| EVR   | Evercore Inc                                    | 
| XGN   | Exagen Inc                                      | 
| EXEL  | Exelixis, Inc.                                  | 
| EXC   | Exelon Corporation                              | 
| EXPE  | Expedia Group Inc                               | 
| STAY  | --error--                                       | 
| XOM   | Exxon Mobil Corporation                         | 
| FFIV  | F5 Networks, Inc.                               | 
| FB    | Facebook, Inc. Common Stock                     | 
| FDS   | FactSet Research Systems Inc.                   | 
| FAST  | Fastenal Company                                | 
| FDX   | FedEx Corporation                               | 
| RACE  | Ferrari NV                                      | 
| FCAU  | Fiat Chrysler Automobiles NV                    | 
| FITB  | Fifth Third Bancorp                             | 
| FEYE  | FireEye Inc                                     | 
| FHN   | First Horizon National Corp                     | 
| AG    | First Majestic Silver Corp.                     | 
| FSLR  | First Solar, Inc.                               | 
| FE    | FirstEnergy Corp.                               | 
| FIS   | Fidelity National Information Servcs Inc        | 
| FISV  | Fiserv Inc                                      | 
| FT    | Franklin Universal Trust                        | 
| FVE   | Five Star Senior Living Inc                     | 
| FVRR  | Fiverr International Ltd                        | 
| FLT   | FleetCor Technologies, Inc.                     | 
| FLEX  | Flex Ltd                                        | 
| FLR   | Fluor Corporation (NEW)                         | 
| FL    | Foot Locker, Inc.                               | 
| F     | Ford Motor Company                              | 
| FTNT  | Fortinet Inc                                    | 
| FTV   | Fortive Corp                                    | 
| BEN   | Franklin Resources, Inc.                        | 
| FCX   | Freeport-McMoRan Inc                            | 
| FREQ  | Frequency Therapeutics Inc                      | 
| FTR   | Frontier Communications Corp                    | 
| FSK   | FS KKR Capital Corp                             | 
| GME   | GameStop Corp.                                  | 
| GCI   | Gannett Co Inc                                  | 
| GPS   | Gap Inc                                         | 
| T     | AT&T Inc.                                       | 
| GDS   | GDS Holdings Ltd - ADR                          | 
| GD    | General Dynamics Corporation                    | 
| GE    | General Electric Company                        | 
| GIS   | General Mills, Inc.                             | 
| GM    | General Motors Company                          | 
| GNTX  | Gentex Corporation                              | 
| GNW   | Genworth Financial Inc                          | 
| GGB   | Gerdau SA ADR                                   | 
| GILD  | Gilead Sciences, Inc.                           | 
| GNL   | Global Net Lease Inc                            | 
| GLUU  | Glu Mobile Inc.                                 | 
| GDDY  | Godaddy Inc                                     | 
| GOL   | Gol Linhas Aereas Inteligentes SA               | 
| GFI   | Gold Fields Limited                             | 
| GS    | Goldman Sachs Group Inc                         | 
| GT    | Goodyear Tire & Rubber Co                       | 
| GPRO  | GoPro Inc                                       | 
| GPK   | Graphic Packaging Holding Company               | 
| GO    | Grocery Outlet Holding Corp                     | 
| GRPN  | Groupon Inc Common Stock                        | 
| GRUB  | GrubHub Inc                                     | 
| GGAL  | Grupo Financiero Galicia S.A.                   | 
| BSMX  | Banco Santander Mexico Sa Instcn De             | 
| SUPV  | Grupo Supervielle SA -ADR                       | 
| T     | AT&T Inc.                                       | 
| GES   | Guess?, Inc.                                    | 
| GPOR  | Gulfport Energy Corporation                     | 
| HRB   | H & R Block Inc                                 | 
| HAL   | Halliburton Company                             | 
| HBI   | Hanesbrands Inc.                                | 
| HOG   | Harley-Davidson Inc                             | 
| HMY   | Harmony Gold Mining Co.                         | 
| HIG   | Hartford Financial Services Group Inc           | 
| HAS   | Hasbro, Inc.                                    | 
| HCA   | HCA Healthcare Inc                              | 
| HDB   | HDFC Bank Limited                               | 
| HL    | Hecla Mining Company                            | 
| HEI   | Heico Corp                                      | 
| HSIC  | Henry Schein, Inc.                              | 
| HLF   | Herbalife Nutrition Ltd                         | 
| HTZ   | Hertz Global Holdings Inc                       | 
| HES   | Hess Corp.                                      | 
| HPE   | Hewlett Packard Enterprise Co                   | 
| HEXO  | Hexo Corp                                       | 
| HLT   | Hilton Hotels Corporation Common Stock          | 
| HGV   | Hilton Grand Vacations Inc                      | 
| HIMX  | Himax Technologies, Inc.                        | 
| HFC   | HollyFrontier Corp                              | 
| HOLX  | Hologic, Inc.                                   | 
| HD    | Home Depot Inc                                  | 
| HMC   | Honda Motor Co Ltd                              | 
| HON   | Honeywell International Inc.                    | 
| HRL   | Hormel Foods Corp                               | 
| HST   | Host Hotels and Resorts Inc                     | 
| HPQ   | HP Inc                                          | 
| HTHT  | Huazhu Group Ltd                                | 
| HUBS  | HubSpot Inc                                     | 
| HUM   | Humana Inc                                      | 
| HBAN  | Huntington Bancshares Incorporated              | 
| HUN   | Huntsman Corporation                            | 
| HCM   | HUTCHISON CHINA/S ADR                           | 
| HUYA  | HUYA Inc - ADR                                  | 
| H     | Hyatt Hotels Corporation                        | 
| IAG   | Iamgold Corp                                    | 
| IBM   | IBM Common Stock                                | 
| IBN   | ICICI Bank Ltd                                  | 
| IDXX  | IDEXX Laboratories, Inc.                        | 
| IGMS  | IGM Biosciences Inc                             | 
| W     | Wayfair Inc                                     | 
| ILMN  | Illumina, Inc.                                  | 
| IMGN  | ImmunoGen, Inc.                                 | 
| INCY  | Incyte Corporation                              | 
| INFN  | Infinera Corp.                                  | 
| INFY  | Infosys Ltd ADR                                 | 
| INTC  | Intel Corporation                               | 
| IBKR  | Interactive Brokers Group, Inc.                 | 
| ICE   | Intercontinental Exchange Inc                   | 
| IT    | Gartner Inc                                     | 
| P     | --error--                                       | 
| PG    | Procter & Gamble Co                             | 
| INTU  | Intuit Inc.                                     | 
| ISRG  | Intuitive Surgical, Inc.                        | 
| Z     | Zillow Group Inc Class C                        | 
| IVR   | Invesco Mortgage Capital Inc                    | 
| ISBC  | Investors Bancorp Inc                           | 
| NVTA  | InVitae Corp                                    | 
| INVH  | Invitation Homes Inc                            | 
| Q     | Qualitas Controladora SAB de CV                 | 
| IRBT  | iRobot Corporation                              | 
| IRM   | Iron Mountain Inc                               | 
| ITUB  | Itau Unibanco Holding SA ADR                    | 
| JNJ   | Johnson & Johnson                               | 
| JP    | Jupai Holdings Ltd                              | 
| JBHT  | J B Hunt Transport Services Inc                 | 
| JKHY  | Jack Henry & Associates, Inc.                   | 
| JD    | JD.Com Inc                                      | 
| JEF   | Jefferies Financial Group Inc                   | 
| JBLU  | JetBlue Airways Corporation                     | 
| JKS   | JinkoSolar Holding Co., Ltd                     | 
| DE    | Deere & Company                                 | 
| JPM   | JPMorgan Chase & Co.                            | 
| JMIA  | Jumia Technologies AG - ADR                     | 
| JNPR  | Juniper Networks, Inc.                          | 
| KAR   | KAR Auction Services Inc                        | 
| K     | Kellogg Company                                 | 
| KDP   | Keurig Dr Pepper Inc                            | 
| KEY   | KeyCorp                                         | 
| KEYS  | Keysight Technologies Inc                       | 
| KMB   | Kimberly Clark Corp                             | 
| KIM   | Kimco Realty Corp                               | 
| KMI   | Kinder Morgan Inc                               | 
| KGC   | Kinross Gold Corporation                        | 
| KKR   | KKR & Co Inc Class A                            | 
| KNX   | Knight-Swift Transportation Holdings Inc        | 
| KSS   | Kohl's Corporation                              | 
| KOS   | Kosmos Energy Ltd                               | 
| KHC   | Kraft Heinz Co                                  | 
| KTOS  | Kratos Defense & Security Solutions, Inc        | 
| KR    | Kroger Co                                       | 
| KT    | KT Corp                                         | 
| LB    | L Brands Inc                                    | 
| LRCX  | Lam Research Corporation                        | 
| LPI   | Laredo Petroleum Inc                            | 
| LS    | Middlefield Healthcare&Life Sci Divd Fnd        | 
| M     | Macy's Inc                                      | 
| LS    | Middlefield Healthcare&Life Sci Divd Fnd        | 
| LDOS  | Leidos Holdings Inc                             | 
| LEN   | Lennar Corporation                              | 
| LEVI  | Levi Strauss & Co.                              | 
| LXRX  | Lexicon Pharmaceuticals, Inc.                   | 
| LX    | LexinFintech Holdings Ltd - ADR                 | 
| LPL   | LG Display Co Ltd.                              | 
| FWONK | Liberty Media Formula One Series C              | 
| L     | Loews Corporation                               | 
| LTHM  | Livent Corp                                     | 
| LKQ   | LKQ Corporation                                 | 
| LMT   | Lockheed Martin Corporation                     | 
| LOMA  | Loma Negra Compania Indl Argentina SA           | 
| LW    | Lamb Weston Holdings Inc                        | 
| LTC   | LTC Properties Inc                              | 
| LK    | Luckin Coffee Inc - ADR                         | 
| L     | Loews Corporation                               | 
| LYFT  | LYFT Inc                                        | 
| M     | Macy's Inc                                      | 
| MRO   | Marathon Oil Corporation                        | 
| MPC   | Marathon Petroleum Corp                         | 
| MAR   | Marriott International Inc                      | 
| MMC   | Marsh & McLennan Companies, Inc.                | 
| MRVL  | Marvell Technology Group Ltd.                   | 
| MAS   | Masco Corp                                      | 
| MA    | Mastercard Inc                                  | 
| MTDR  | Matador Resources Co                            | 
| MTCH  | Match Group Inc                                 | 
| MAT   | Mattel Inc                                      | 
| MXIM  | Maxim Integrated Products Inc.                  | 
| MDR   | McDermott International Inc                     | 
| MCD   | Mcdonald's Corp                                 | 
| MUX   | McEwen Mining Inc                               | 
| MPW   | Medical Properties Trust, Inc.                  | 
| MDT   | Medtronic PLC                                   | 
| MLCO  | Melco Resorts & Entertainment Ltd               | 
| MELI  | Mercadolibre Inc                                | 
| MRK   | Merck & Co., Inc.                               | 
| MET   | Metlife Inc                                     | 
| MFA   | MFA FINL INC/SH                                 | 
| MTG   | MGIC Investment Corp.                           | 
| MGM   | MGM Resorts International                       | 
| MIK   | Michaels Companies Inc                          | 
| MCHP  | Microchip Technology Inc.                       | 
| MU    | Micron Technology, Inc.                         | 
| MSFT  | Microsoft Corporation                           | 
| MFG   | Mizuho Financial Group Inc.                     | 
| MBT   | Mobil'nye Telesistemy PAO                       | 
| MOMO  | Momo Inc                                        | 
| MDLZ  | MONDELEZ INTERNATIONAL INC Common Stock         | 
| MGI   | Moneygram International Inc                     | 
| MNST  | Monster Beverage Corp                           | 
| MCO   | Moody's Corporation                             | 
| MS    | Morgan Stanley                                  | 
| MORN  | Morningstar, Inc.                               | 
| MOS   | Mosaic Co                                       | 
| MSI   | Motorola Solutions Inc                          | 
| MPLX  | MPLX LP                                         | 
| MUFG  | Mitsubishi UFJ Financial Group Inc              | 
| MUR   | Murphy Oil Corporation                          | 
| MYL   | Mylan NV                                        | 
| NBR   | Nabors Industries Ltd.                          | 
| NDAQ  | Nasdaq Inc                                      | 
| NOV   | National-Oilwell Varco, Inc.                    | 
| NTCO  | Natura & Co                                     | 
| NAVI  | Navient Corp                                    | 
| NKTR  | Nektar Therapeutics                             | 
| NPTN  | NeoPhotonics Corp                               | 
| NTAP  | NetApp Inc.                                     | 
| NTES  | NetEase Inc                                     | 
| NFLX  | Netflix Inc                                     | 
| NBIX  | Neurocrine Biosciences, Inc.                    | 
| NBEV  | New Age Beverages Corp                          | 
| EDU   | New Oriental Education & Tech Grp               | 
| NRZ   | New Residential Investment Corp                 | 
| NYCB  | New York Community Bancorp, Inc.                | 
| NYMT  | NY MTG TR INC/SH                                | 
| NYT   | New York Times Co                               | 
| NWL   | Newell Brands Inc                               | 
| NEM   | Newmont Corporation                             | 
| NWSA  | News Corp Class A                               | 
| NEE   | NextEra Energy Inc                              | 
| NKE   | Nike Inc                                        | 
| NIO   | Nio Inc - ADR                                   | 
| NOAH  | Noah Holdings Limited                           | 
| NBL   | Noble Energy, Inc.                              | 
| NMR   | Nomura Holdings Inc                             | 
| JWN   | Nordstrom, Inc.                                 | 
| NSC   | Norfolk Southern Corp.                          | 
| NTRS  | Northern Trust Corporation                      | 
| NOC   | Northrop Grumman Corporation                    | 
| NLOK  | NortonLifeLock Inc                              | 
| NCLH  | Norwegian Cruise Line Holdings Ltd              | 
| NVCR  | Novocure Ltd                                    | 
| NRG   | NRG Energy Inc                                  | 
| NUE   | Nucor Corporation                               | 
| NTNX  | Nutanix Inc                                     | 
| NUVA  | NuVasive, Inc.                                  | 
| NVDA  | NVIDIA Corporation                              | 
| ORLY  | O'Reilly Automotive Inc                         | 
| OAS   | Oasis Petroleum Inc.                            | 
| OXY   | Occidental Petroleum Corporation                | 
| ODP   | Office Depot Inc                                | 
| OKTA  | Okta Inc                                        | 
| OLN   | Olin Corporation                                | 
| OMC   | Omnicom Group Inc.                              | 
| ON    | ON Semiconductor Corp                           | 
| OKE   | ONEOK, Inc.                                     | 
| OPK   | Opko Health Inc.                                | 
| ORCL  | Oracle Corporation                              | 
| ONVO  | Organovo Holdings Inc                           | 
| OSTK  | Overstock.com Inc                               | 
| PG    | Procter & Gamble Co                             | 
| PCAR  | PACCAR Inc                                      | 
| PD    | Pagerduty Inc                                   | 
| PANW  | Palo Alto Networks Inc                          | 
| PAM   | Pampa Energia S.A.                              | 
| PAAS  | Pan American Silver Corp.                       | 
| PH    | Parker-Hannifin Corp                            | 
| PE    | Parsley Energy Inc                              | 
| PTEN  | Patterson-UTI Energy, Inc.                      | 
| PAYX  | Paychex, Inc.                                   | 
| PYPL  | Paypal Holdings Inc                             | 
| PBF   | PBF Energy Inc                                  | 
| PTON  | Peloton Interactive Inc                         | 
| PBCT  | People's United Financial, Inc.                 | 
| PEP   | PepsiCo, Inc.                                   | 
| PBR   | PETROLEO BRASIL/ADR                             | 
| PFE   | Pfizer Inc.                                     | 
| PCG   | PG&E Corporation                                | 
| PM    | Philip Morris International Inc.                | 
| PSX   | Phillips 66                                     | 
| PDD   | Pinduoduo Inc                                   | 
| PINS  | Pinterest Inc                                   | 
| PXD   | Pioneer Natural Resources                       | 
| PBI   | Pitney Bowes Inc.                               | 
| PVTL  | --error--                                       | 
| PAA   | Plains All American Pipeline, L.P.              | 
| PLUG  | Plug Power Inc                                  | 
| PS    | Pluralsight Inc                                 | 
| PNC   | PNC Financial Services Group Inc                | 
| PPG   | PPG Industries, Inc.                            | 
| PPL   | PPL Corp                                        | 
| PBH   | Prestige Consumer Healthcare Inc                | 
| PRI   | Primerica, Inc.                                 | 
| PGR   | Progressive Corp                                | 
| PLD   | Prologis Inc                                    | 
| PFPT  | Proofpoint Inc                                  | 
| PRU   | Prudential Financial Inc                        | 
| PEG   | Public Service Enterprise Group Inc.            | 
| PSA   | Public Storage                                  | 
| PHM   | PulteGroup, Inc.                                | 
| PSTG  | Pure Storage Inc                                | 
| QTWO  | Q2 Holdings Inc                                 | 
| QEP   | QEP Resources Inc                               | 
| QCM   | --error--                                       | 
| QLS   | INDEXIQ ETF TR/IQ HEDGE LONG SHORT              | 
| PWR   | Quanta Services Inc                             | 
| QD    | Qudian Inc - ADR                                | 
| QRTEA | Qurate Retail Inc Series A                      | 
| RL    | Ralph Lauren Corp                               | 
| RRC   | Range Resources Corp.                           | 
| RTN   | Raytheon Company                                | 
| RLGY  | Realogy Holdings Corp                           | 
| REAL  | RealReal Inc                                    | 
| O     | Realty Income Corp                              | 
| REGN  | Regeneron Pharmaceuticals Inc                   | 
| RF    | Regions Financial Corp                          | 
| REGI  | Renewable Energy Group Inc                      | 
| RMD   | ResMed Inc.                                     | 
| QSR   | Restaurant Brands International Inc             | 
| RNG   | RingCentral Inc                                 | 
| RAD   | Rite Aid Corporation                            | 
| ROK   | Rockwell Automation                             | 
| ROKU  | Roku Inc                                        | 
| ROST  | Ross Stores, Inc.                               | 
| RY    | Royal Bank of Canada                            | 
| RCL   | Royal Caribbean Cruises Ltd                     | 
| RES   | RPC, Inc.                                       | 
| SPGI  | S&P Global Inc                                  | 
| CRM   | salesforce.com, inc.                            | 
| SBH   | Sally Beauty Holdings, Inc.                     | 
| SGMO  | Sangamo Therapeutics Inc                        | 
| SLB   | Schlumberger Limited.                           | 
| STX   | Seagate Technology PLC                          | 
| SRE   | Sempra Energy                                   | 
| SENS  | Senseonics Holdings Inc                         | 
| SHAK  | Shake Shack Inc                                 | 
| SHW   | Sherwin-Williams Co                             | 
| SHOP  | Shopify Inc                                     | 
| SBGL  | Sibanye Gold Ltd                                | 
| SPG   | Simon Property Group Inc                        | 
| SIRI  | Sirius XM Holdings Inc                          | 
| SKX   | Skechers USA Inc                                | 
| SWKS  | Skyworks Solutions Inc                          | 
| WORK  | Slack Technologies Inc                          | 
| SLM   | SLM Corp                                        | 
| SM    | SM Energy Co                                    | 
| SMAR  | Smartsheet Inc                                  | 
| SNAP  | Snap Inc                                        | 
| SOGO  | Sogou Inc                                       | 
| SNE   | Sony Corp                                       | 
| BID   | Bidenergy Ltd                                   | 
| SO    | Southern Co                                     | 
| SCCO  | Southern Copper Corp                            | 
| LUV   | Southwest Airlines Co                           | 
| SWN   | Southwestern Energy Company                     | 
| ONCE  | --error--                                       | 
| SPLK  | Splunk Inc                                      | 
| SPOT  | Spotify Technology SA                           | 
| S     | Sprint Corp                                     | 
| SFM   | Sprouts Farmers Market Inc                      | 
| SQ    | Square Inc                                      | 
| SRCI  | SRC Energy Inc                                  | 
| SSNC  | SS&C Technologies Holdings, Inc.                | 
| SWK   | Stanley Black & Decker, Inc.                    | 
| SBUX  | Starbucks Corporation                           | 
| TSG   | Stars Group Inc                                 | 
| STT   | State Street Corp                               | 
| STLD  | Steel Dynamics, Inc.                            | 
| SFIX  | Stitch Fix Inc                                  | 
| STNE  | StoneCo Ltd                                     | 
| SYK   | Stryker Corporation                             | 
| SMFG  | Sumitomo Mitsui Financial Grp, Inc.             | 
| SU    | Suncor Energy Inc.                              | 
| SPWR  | SunPower Corporation                            | 
| RUN   | Sunrun Inc                                      | 
| ST    | Sensata Technologies Holding PLC                | 
| SSS   | Stor-Age Property REIT Ltd                      | 
| SYF   | Synchrony Financial                             | 
| SNPS  | Synopsys, Inc.                                  | 
| SYY   | SYSCO Corporation                               | 
| TMUS  | T-Mobile Us Inc                                 | 
| TROW  | T. Rowe Price Group Inc                         | 
| TTWO  | TAKE-TWO INTERACTIVE SOFTWARE, INC Common Stock | 
| TLRD  | Tailored Brands Inc                             | 
| TSM   | Taiwan Semiconductor Mfg. Co. Ltd.              | 
| TAK   | Takeda Pharmaceutical Co Ltd                    | 
| TAL   | TAL Education Group                             | 
| SKT   | Tanger Factory Outlet Centers Inc.              | 
| TPR   | Tapestry Inc                                    | 
| TEDU  | Tarena International Inc                        | 
| TRGP  | Targa Resources Corp                            | 
| TGT   | Target Corporation                              | 
| TTM   | Tata Motors Limited ADR                         | 
| AMTD  | TD Ameritrade Holding Corp.                     | 
| FTI   | TechnipFMC PLC                                  | 
| TECK  | Teck Resources Ltd Class B                      | 
| TFX   | Teleflex Incorporated                           | 
| VIV   | Telefonica Brasil SA ADR                        | 
| TELL  | Tellurian Inc                                   | 
| TPX   | Tempur Sealy International Inc                  | 
| TME   | Tencent Music Entertainment Group - ADR         | 
| TER   | Teradyne, Inc.                                  | 
| TSLA  | Tesla Inc                                       | 
| TEVA  | Teva Pharmaceutical Industries Ltd              | 
| TXN   | Texas Instruments Incorporated                  | 
| DOW   | Dow Inc                                         | 
| GEO   | The GEO Group Inc                               | 
| TXMD  | TherapeuticsMD Inc                              | 
| TO    | Tower One Wireless Corp                         | 
| TIF   | Tiffany & Co.                                   | 
| TSU   | TIM Participacoes SA                            | 
| TIX   | --error--                                       | 
| T     | AT&T Inc.                                       | 
| TM    | Toyota Motor Corp                               | 
| TW    | Tradeweb Markets Inc                            | 
| TRXC  | Transenterix Inc                                | 
| TRV   | Travelers Companies Inc                         | 
| TCOM  | Trip.com Group Ltd                              | 
| TRIP  | Tripadvisor Inc Common Stock                    | 
| TGI   | Triumph Group Inc                               | 
| TFC   | Truist Financial Corp                           | 
| TRQ   | Turquoise Hill Resources Ltd                    | 
| TWLO  | Twilio Inc                                      | 
| TWTR  | Twitter Inc                                     | 
| TWO   | Two Harbors Investment Corp                     | 
| TSN   | Tyson Foods, Inc.                               | 
| UBER  | Uber Technologies Inc                           | 
| ULTA  | Ulta Beauty Inc                                 | 
| UCTT  | Ultra Clean Holdings Inc                        | 
| UGP   | Ultrapar Participacoes SA                       | 
| UAA   | Under Armour Inc Class A                        | 
| UA    | Under Armour Inc Class C                        | 
| UNP   | Union Pacific Corporation                       | 
| UAL   | United Airlines Holdings Inc                    | 
| MC    | Moelis & Co                                     | 
| X     | United States Steel Corporation                 | 
| UTX   | United Technologies Corporation                 | 
| UNH   | UnitedHealth Group Inc                          | 
| UNIT  | Uniti Group Inc                                 | 
| UNM   | Unum Group                                      | 
| TIGR  | Up Fintech Holding Ltd                          | 
| UPS   | United Parcel Service, Inc.                     | 
| URBN  | Urban Outfitters, Inc.                          | 
| USB   | U.S. Bancorp                                    | 
| SLCA  | U.S. Silica Holdings Inc                        | 
| UXIN  | Uxin Ltd                                        | 
| VALE  | Vale SA                                         | 
| VLO   | Valero Energy Corporation                       | 
| VGR   | Vector Group Ltd                                | 
| VEEV  | Veeva Systems Inc                               | 
| VTR   | Ventas, Inc.                                    | 
| VER   | Vereit Inc                                      | 
| VRSN  | Verisign, Inc.                                  | 
| VRSK  | Verisk Analytics, Inc.                          | 
| VZ    | Verizon Communications Inc.                     | 
| VRTX  | Vertex Pharmaceuticals Incorporated             | 
| VFC   | VF Corp                                         | 
| VIACA | ViacomCBS Inc Class A                           | 
| VIC   | Victory Mines Ltd                               | 
| VIPS  | Vipshop Holdings Ltd - ADR                      | 
| SPCE  | Virgin Galactic Holdings Inc Class A            | 
| V     | Visa Inc                                        | 
| VST   | Vistra Energy Corp                              | 
| MW    | --error--                                       | 
| VG    | Vonage Holdings Corp.                           | 
| WAB   | Westinghouse Air Brake Technologies Corp        | 
| WBA   | Walgreens Boots Alliance Inc                    | 
| WMT   | Walmart Inc                                     | 
| DIS   | Walt Disney Co                                  | 
| WM    | Waste Management, Inc.                          | 
| W     | Wayfair Inc                                     | 
| WB    | Weibo Corp                                      | 
| WFC   | Wells Fargo & Co                                | 
| WELL  | Welltower Inc                                   | 
| WEN   | Wendys Co                                       | 
| WC    | --error--                                       | 
| W     | Wayfair Inc                                     | 
| WRK   | Westrock Co                                     | 
| WEX   | WEX Inc                                         | 
| WY    | Weyerhaeuser Co                                 | 
| WLL   | Whiting Petroleum Corp                          | 
| WMB   | Williams Companies Inc                          | 
| WING  | Wingstop Inc                                    | 
| WIT   | Wipro Limited                                   | 
| WDAY  | Workday Inc                                     | 
| WWE   | World Wrestling Entertainment, Inc.             | 
| WPG   | Washington Prime Group Inc                      | 
| WPX   | WPX Energy Inc                                  | 
| WYNN  | Wynn Resorts, Limited                           | 
| XEL   | Xcel Energy Inc                                 | 
| XRX   | Xerox Holdings Corp                             | 
| XLNX  | Xilinx, Inc.                                    | 
| AUY   | Yamana Gold Inc.                                | 
| YPF   | YPF SA                                          | 
| YUM   | Yum! Brands, Inc.                               | 
| YUMC  | Yum China Holdings Inc                          | 
| YY    | JOYY Inc                                        | 
| ZAYO  | Zayo Group Holdings Inc                         | 
| ZBRA  | Zebra Technologies Corp.                        | 
| ZEN   | Zendesk Inc                                     | 
| A     | Agilent Technologies Inc                        | 
| ZBH   | Zimmer Biomet Holdings Inc                      | 
| ZION  | Zions Bancorporation NA                         | 
| ZTS   | Zoetis Inc                                      | 
| M     | Macy's Inc                                      | 
| ZS    | Zscaler Inc                                     | 
| ZTO   | ZTO Express (Cayman) Inc                        | 
| ZNGA  | Zynga Inc                                       | 
