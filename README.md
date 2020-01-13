# Revolut Stocks List ![GitHub](https://img.shields.io/github/license/nmapx/revolut-stocks-list?style=flat) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/nmapx/revolut-stocks-list?style=flat) [![Go Report Card](https://goreportcard.com/badge/github.com/nmapx/revolut-stocks-list)](https://goreportcard.com/report/github.com/nmapx/revolut-stocks-list)

Extract Revolut stocks list from the list screenshot(s) ([example screenshot](./example-input.jpg) - use your mobile app for this).

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

First of all - check [example screenshot](./example-input.jpg). Your screenshot(s) must follow the pattern!

You can specify multiple input files by adding more parameters. By default it's processing 1 file (filepath `./input.jpg`) which can be very long (complete stock list in 1 screenshot (possible on some mobiles)). See example below.

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

## License

[MIT License](./LICENSE)

## The list

As the result you will get a file with all recognized tickers and companies names.

**OCR is still not the most accurate solution so you can find few mistakes there (about 2-3% misidentified tickers).**

Below you can find already extracted list of all available stocks. I will update it from time to time.

*Updated on: 13.01.2020*

*Number of rows: 796*

| Ticker | Company name     | 
|-------|-------------------| 
| FOXA  | 21st Century F    | 
| TWOU  | 2u                | 
| MMM   | am                | 
| WUBA  | 58.com            | 
| ABT   | Abbott Labs       | 
| ABBV  | AbbVie            | 
| ANF   | Abercrombie &     | 
| ACHN  | Achillion Pharn   | 
| ATVI  | Activision Blizz  | 
| ADBE  | Adobe             | 
| AAP   | Advance Auto \|   | 
| AMD   | Advanced Micr     | 
| AES   | AES Corporatio    | 
| AFL   | Aflac             | 
| A     | Agilent Techno    | 
| AGNC  | AGNC Investm      | 
| AIG   | AIG               | 
| AL    | Air Lease Corp    | 
| APD   | Air Products ar   | 
| AKS   | AK Steel Holdir   | 
| AKAM  | Akamai Technc     | 
| AA    | Alcoa             | 
| ALXN  | Alexion Pharm:    | 
| BABA  | Alibaba           | 
| ALGN  | Align Technolo    | 
| ALLT  | Allot Communi     | 
| MDRX  | Allscripts Heal   | 
| ALL   | Allstate          | 
| ALLY  | Ally Financial    | 
| GOOGL | Alphabet (Clas:   | 
| GOOG  | Alphabet (Clas:   | 
| AABA  | Altaba            | 
| ATUS  | Altice USA        | 
| MO    | Altria            | 
| AMZN  | Amazon            | 
| ABEV  | Ambev             | 
| AMC   | AMC               | 
| AAL   | American Airlin   | 
| AXL   | American Axle     | 
| AEO   | American Eagle    | 
| AEP   | American Elect    | 
| AXP   | American Expr     | 
| AMH   | American Hom      | 
| AMT   | American Towe     | 
| CLD   | Americold Real    | 
| AME   | Ametek            | 
| AMGN  | Amgen             | 
| FOLD  | Amicus Therap     | 
| AMRX  | Amneal Pharm:     | 
| APH   | Amphenol          | 
| ADI   | Analog Devices    | 
| PLAN  | Anaplan           | 
| ANGI  | ANGI Homeser      | 
| AU    | AngloGold Ash     | 
| NLY   | Annaly Capital    | 
| ANSS  | Ansys             | 
| AM    | Antero Midstre    | 
| AR    | Antero Resourc    | 
| ANTM  | Anthem            | 
| AOS   | AO Smith          | 
| APA   | Apache            | 
| ARI   | Apollo Comme:     | 
| AAPL  | Apple             | 
| AMAT  | Applied Materi    | 
| ATR   | AptarGroup        | 
| ARMK  | Aramark           | 
| ADM   | Archer-Daniels    | 
| ARNC  | Arconic           | 
| ARCC  | Ares Capital      | 
| ANET  | Arista Network    | 
| ARR   | ARMOUR Resic      | 
| ARQL  | ArQule            | 
| ASNA  | Ascena Retail     | 
| ASML  | ASML Holding      | 
| HOME  | At Home Grou      | 
| T     | ATRT              | 
| ADSK  | Autodesk          | 
| ATHM  | Autohome          | 
| ADP   | Automatic Datz    | 
| AZO   | AutoZone          | 
| AVLR  | Avalara           | 
| AVB   | AvalonBay Con     | 
| EQH   | AXA Equitable     | 
| AXTA  | Axalta Coating    | 
| BTG   | B2Gold Corpor:    | 
| BIDU  | Baidu             | 
| BKR   | Baker Hughes      | 
| BLL   | Ball Corporatio   | 
| BBD   | Banco Bradesc     | 
| BMA   | Banco Macro       | 
| BSBR  | Banco Santand     | 
| BAC   | Bank of Americ    | 
| GOLD  | Barrick Gold      | 
| BHC   | Bausch Health     | 
| BAX   | Baxter            | 
| BBAR  | BBVA Banco Fr     | 
| BDX   | Becton Dickins    | 
| BDC   | Belden            | 
| BRKB  | Berkshire Hath    | 
| BBY   | Best Buy          | 
| BYND  | Beyond Meat       | 
| BHP   | BHP Billiton      | 
| BILI  | Bilibili          | 
| BIO   | Bio-Rad Laborz    | 
| BIIB  | Biogen            | 
| BMRN  | Biomarin Pharr    | 
| BITA  | Bitauto Holding   | 
| BJ    | BJ's Wholesale    | 
| BB    | BlackBerry        | 
| BLK   | BlackRock         | 
| BX    | Blackstone        | 
| BK    | BNY Mellon        | 
| BA    | Boeing            | 
| BKNG  | Booking           | 
| BAH   | Booz Allen Han    | 
| BWA   | BorgWarner        | 
| BSX   | Boston Scienti    | 
| BOX   | Box               | 
| BRFS  | BRF               | 
| BMY   | Bristol-Myers S   | 
| BMYRT | Bristol-Myers S   | 
| BRX   | Brixmor Proper    | 
| AVGO  | Broadcom          | 
| BAM   | Brookfield Ass:   | 
| BIP   | Brookfield Infrz  | 
| BEP   | Brookfield Ren    | 
| BFB   | Brown Forman      | 
| BG    | Bunge             | 
| COG   | Cabot 0l & Gas    | 
| CDNS  | Cadence Desig     | 
| CZR   | Caesars Entert    | 
| CRC   | California Reso   | 
| CPE   | Callon Petrole    | 
| CAJ   | Canon             | 
| COF   | Capital One       | 
| CAH   | Cardinal Health   | 
| CG    | Carlyle Group     | 
| CCL   | Carnival          | 
| CRZO  | Carrizo Oil & Gz  | 
| CARS  | Cars.com          | 
| CAT   | Caterpillar       | 
| CBL   | CBL & Associal    | 
| CBOE  | CBOE Holdings     | 
| CBRE  | CBRE Group        | 
| CELG  | Celgene           | 
| CX    | CEMEX S.AB.       | 
| CVE   | Cenovus Energ     | 
| CNC   | Centene           | 
| CDEV  | Centennial Res    | 
| CNP   | CenterPoint En    | 
| CTL   | CenturyLink       | 
| CERN  | Cerner            | 
| CF    | CF Industries     | 
| SCHW  | Charles Schwa     | 
| CHTR  | Charter Comm      | 
| CHKP  | Check Point So    | 
| SM    | Chemical & Mir    | 
| CC    | Chemours Con      | 
| LNG   | Cheniere Energ    | 
| CHK   | Chesapeake Er     | 
| CVX   | Chevron           | 
| CHWY  | Chewy             | 
| CHS   | Chico's FAS       | 
| CIM   | Chimera Invest    | 
| CHL   | China Mobile      | 
| ZNH   | China Southerr    | 
| CMG   | Chipotle Mexic    | 
| CIEN  | Ciena Corporat    | 
| CX    | Cigna             | 
| XEC   | Cimarex Energ;    | 
| CTAS  | Cintas            | 
| CSCO  | Cisco             | 
| C     | Citigroup         | 
| CFG   | Citizens Financ   | 
| CTXS  | Citrix Systems    | 
| CCO   | Clear Channel (   | 
| CWEN  | Clearway Enerc    | 
| CLF   | Cliffs Natural R  | 
| CLDR  | Cloudera          | 
| NET   | Cloudflare        | 
| CLVS  | Clovis Oncolog    | 
| CME   | CME Group (Cl     | 
| CNX   | CNX Resource:     | 
| KO    | Coca-Cola         | 
| CDE   | Coeur Mining      | 
| CGNX  | Cognex Corpor     | 
| CTSH  | Cognizant Tect    | 
| CL    | Colgate-Palmo     | 
| CLNY  | Colony Capital    | 
| CMCSA | Comcast           | 
| CMA   | Comerica          | 
| COMM  | CommScope H       | 
| CYH   | Community He      | 
| SBS   | Companhia de      | 
| CIG   | Companhia Ene     | 
| SID   | Companhia Sid     | 
| BVN   | Compaiiia de      | 
| CAG   | ConAgra Foods     | 
| CX    | Concho Resour     | 
| CNDT  | Conduent          | 
| CP    | ConacoPhillips    | 
| ED    | Consolidated E    | 
| STZ   | Constellation B   | 
| CLR   | Continental Re:   | 
| GLW   | Corning           | 
| CTVA  | Corteva           | 
| COST  | Costco Wholes     | 
| COTY  | Coty              | 
| BAP   | Credicorp         | 
| CROX  | Crocs             | 
| CRWD  | CrowdStrike       | 
| CCX   | Crown Castle      | 
| CSX   | csx               | 
| CVS   | CVS Health        | 
| CY    | Cypress Semic     | 
| DAN   | Dana Holding      | 
| DHR   | Danaher           | 
| DDOG  | Datadog           | 
| DVA   | DaVita HealthC    | 
| DELL  | Dell Technolog    | 
| DAL   | Delta Air Lines   | 
| DNR   | Denbury Resou     | 
| DVN   | Devon Energy      | 
| DO    | Diamond Offsh     | 
| DLR   | Digital Realty T  | 
| APPS  | Digital Turbine   | 
| DFS   | Discover Finan    | 
| DISCA | Discovery Com     | 
| DISCK | Discovery Com     | 
| DISH  | DISH Network      | 
| DHC   | Diversified Hea   | 
| DOCU  | DocuSign          | 
| DG    | Dollar General    | 
| DLTR  | Dollar Tree       | 
| D     | Dominion Ener     | 
| DPZ   | Domina's Pizza    | 
| DHI   | DR Horton         | 
| DBX   | Dropbox           | 
| DUK   | Duke Energy       | 
| DRE   | Duke Realty       | 
| DNKN  | Dunkin' Brands    | 
| DD    | DuPont            | 
| DXC   | DXC Technolog     | 
| ETFC  | E*Trade           | 
| EBAY  | eBay              | 
| ECL   | Ecolab            | 
| EIX   | Edison Internat   | 
| EW    | Edwards Lifesc    | 
| ELAN  | Elanco Animal     | 
| EGO   | Eldorado Gold (   | 
| EA    | Electronic Arts   | 
| ESI   | Element Soluti    | 
| LLY   | Eli Lilly         | 
| ERJ   | Embraer           | 
| EMR   | Emerson           | 
| ECA   | Encana Corpor     | 
| ET    | Energy Transfe    | 
| ENIA  | Enersis Americ    | 
| ENLC  | EnLink Midstre    | 
| EPD   | Enterprise Proc   | 
| EOG   | EOG Resources     | 
| EQT   | EQT Corporatio    | 
| EFX   | Equifax           | 
| EQIX  | Equinix           | 
| EQNR  | Equinor           | 
| ETRN  | Equitrans Mids    | 
| EQR   | Equity Residen    | 
| EL    | Estee Lauder      | 
| ETSY  | Etsy              | 
| EB    | Eventbrite        | 
| EVR   | Evercore          | 
| XGN   | Exagen            | 
| EXEL  | Exelixis          | 
| EXC   | Exelon            | 
| EXPE  | Expedia           | 
| STAY  | Extended Stay.    | 
| XOM   | Exxon Mobil       | 
| FFIV  | F5 Networks       | 
| FB    | Facebook          | 
| FDS   | FactSet Resear    | 
| FAST  | Fastenal Comp     | 
| FDX   | FedEx             | 
| RACE  | Ferrari           | 
| FCAU  | Fiat Chrysler     | 
| FITB  | Fifth Third Ban   | 
| FEYE  | FireEye           | 
| FHN   | First Horizon N   | 
| AG    | First Majestic ¢  | 
| FSLR  | First Solar       | 
| FE    | FirstEnergy       | 
| FIS   | FIS               | 
| FISV  | Fiserv            | 
| FIT   | Fitbit            | 
| FVE   | Five Star Senio   | 
| FVRR  | Fiverr            | 
| FLT   | FleetCor Techn    | 
| FLEX  | Flextronics Inte  | 
| FLR   | Fluor Corporati   | 
| FL    | Foot Locker       | 
| F     | Ford              | 
| FTNT  | Fortinet          | 
| FTV   | Fortive           | 
| BEN   | Franklin Resou    | 
| FCX   | Freeport          | 
| FREQ  | Frequency The     | 
| FTR   | Frontier Comm     | 
| FSK   | FS KKR Capital    | 
| GME   | GameStop          | 
| GCI   | Gannett Co        | 
| GPS   | GAP               | 
| T     | Gartner           | 
| GDS   | GDS Holdings      | 
| GD    | General Dynam     | 
| GE    | General Electri   | 
| GIS   | General Mills     | 
| GM    | General Motors    | 
| GNTX  | Gentex            | 
| GNW   | Genworth Fina     | 
| GGB   | Gerdau S.A.       | 
| GILD  | Gilead Science    | 
| GNL   | Global Net Lea:   | 
| GLUU  | Glu Mobile        | 
| GDDY  | GoDaddy           | 
| GOL   | Gol Intelligent / | 
| GFI   | Gold Fields       | 
| GS    | Goldman Sach      | 
| GT    | Goodyear          | 
| GPRO  | GoPro             | 
| GPK   | Graphic Packa     | 
| GO    | Grocery Outlet    | 
| GRPN  | Groupon           | 
| GRUB  | GrubHub           | 
| GGAL  | Grupo Financie    | 
| BSMX  | Grupo Financie    | 
| SUPV  | Grupo Supervie    | 
| TV    | Grupo Televisa    | 
| GES   | Guess             | 
| GPOR  | Gulfport Energ;   | 
| HRB   | H&R Block         | 
| HAL   | Halliburton       | 
| HBI   | Hanesbrands       | 
| HOG   | Harley-Davids     | 
| HMY   | Harmony Gold      | 
| HIG   | Hartford Finan    | 
| HAS   | Hasbro            | 
| HCA   | HCA Healthcar     | 
| HDB   | HDFC Bank.        | 
| HL    | Hecla Mining C    | 
| HEI   | Heico             | 
| HSIC  | Henry Schein      | 
| HLF   | Herbalife         | 
| HTZ   | Hertz             | 
| HES   | Hess Corporati    | 
| HPE   | Hewlett Packar    | 
| HEXO  | Hexo Corp         | 
| HLT   | Hilton            | 
| HGV   | Hilton Grand V:   | 
| HIMX  | Himax Technol     | 
| HFC   | HollyFrontier     | 
| HOLX  | Hologic           | 
| HD    | Home Depot        | 
| HMC   | Honda Motor       | 
| HON   | Honeywell         | 
| HRL   | Hormel Foods      | 
| HST   | Host Hotels & \|  | 
| HPQ   | HP                | 
| HTHT  | Huazhu Group      | 
| HUBS  | HubSpot           | 
| HUM   | Humana            | 
| HBAN  | Huntington Bar    | 
| HUN   | Huntsman Cory     | 
| HCM   | Hutchison Chir    | 
| HUYA  | HUYA              | 
| H     | Hyatt Hotels      | 
| IAG   | lamgold Corp      | 
| IBM   | IBM               | 
| IBN   | ICICI Bank        | 
| IDXX  | IDEXX Laboratc    | 
| IGMS  | IGM Biosciencs    | 
| T     | Hllinois Tool Wo  | 
| ILMN  | llumina           | 
| IMGN  | ImmunoGen         | 
| INCY  | Incyte            | 
| INFN  | Infinera Corpor   | 
| INFY  | Infosys           | 
| INTC  | Intel             | 
| IBKR  | Interactive Brol  | 
| ICE   | Intercontinenta   | 
| IGT   | International G   | 
| P     | International P:  | 
| IPG   | Interpublic Gro   | 
| INTU  | Intuit            | 
| ISRG  | Intuitive Surgic  | 
| Z     | Invesco           | 
| IVR   | Invesco Mortg:    | 
| ISBC  | Investors Banc    | 
| NVTA  | Invitae Corpora   | 
| INVH  | Invitation Hom    | 
| Q     | iQIvI             | 
| IRBT  | iRobot            | 
| IRM   | Iron Mountain     | 
| ITUB  | Itad Unibanco     | 
| JNJ   | J&J               | 
| JP    | J. C. Penney C    | 
| JBHT  | J.B. Hunt         | 
| JKHY  | Jack Henry & A    | 
| JD    | JD.com            | 
| JEF   | Jefferies Finan   | 
| JBLU  | JetBlue Airway    | 
| JKS   | Jinkosolar Hol    | 
| DE    | John Deere        | 
| JPM   | JPMorgan          | 
| JMIA  | Jumia             | 
| JNPR  | Juniper Networ    | 
| KAR   | KAR Auction S     | 
| K     | Kellogg's         | 
| KDP   | Keurig Dr Pepp    | 
| KEY   | KeyCorp.          | 
| KEYS  | Keysight Techr    | 
| KMB   | Kimberly-Clark    | 
| KIM   | Kimco Realty C    | 
| KMI   | Kinder Morgan     | 
| KGC   | Kinross Gold C    | 
| KKR   | KKR & Co          | 
| KNX   | Knight-Swift Tr   | 
| KSS   | Kohl's            | 
| KOS   | Kosmos Energ;     | 
| KHC   | Kraft Heinz       | 
| KTOS  | Kratos            | 
| KR    | Kroger Co         | 
| KT    | KT                | 
| LB    | L Brands          | 
| LRCX  | Lam Research      | 
| LPI   | Laredo Petrole:   | 
| VS    | Las Vegas San     | 
| LM    | LATAM Airlines    | 
| LS    | Lattice Semico    | 
| LDOS  | Leidos Holding    | 
| LEN   | Lennar            | 
| LEVI  | Levi Strauss &    | 
| LXRX  | Lexicon Pharm     | 
| LX    | LexinFintech H    | 
| LPL   | LG Display        | 
| FWONK | Liberty Media ¢   | 
| LYV   | Live Nation Ent   | 
| LTHM  | Livent            | 
| LKQ   | LKQ               | 
| LMT   | Lockheed Mart     | 
| LOMA  | Loma Negra Cc     | 
| LW    | Lowe's            | 
| L     | LTC Properties    | 
| LK    | Luckin Coffee     | 
| LULY  | Lululemon         | 
| LYFT  | Lyft              | 
| M     | Macy's            | 
| MRO   | Marathon Oil C    | 
| MPC   | Marathon Petrc    | 
| MAR   | Marriott          | 
| MMC   | Marsh & McLer     | 
| MRVL  | Marvell Techne    | 
| MAS   | Masco Corporz     | 
| MA    | Mastercard        | 
| MTDR  | Matador Resou     | 
| MTCH  | Match Group       | 
| MAT   | Mattel            | 
| MXIM  | Maxim Integrat    | 
| MDR   | McDermott Inte    | 
| MCD   | McDonald's        | 
| MUX   | McEwen Minin      | 
| MPW   | Medical Proper    | 
| MDT   | Medtronic         | 
| MLCO  | Melco Crown E     | 
| MELI  | MercadoLibre      | 
| MRK   | Merck             | 
| MET   | Metlife           | 
| MFA   | MFA Financial     | 
| MTG   | MGIC Investme     | 
| MGM   | MGM Resorts       | 
| MIK   | Michaels Com      | 
| MCHP  | Microchip Tec     | 
| MU    | Micron Techno     | 
| MSFT  | Microsoft         | 
| MFG   | Mizuho Financ     | 
| MBT   | Mobile TeleSys    | 
| MOMO  | Momo              | 
| MDLZ  | Mondelez          | 
| MGI   | MoneyGram         | 
| MNST  | Monster           | 
| MCO   | Moody's           | 
| MS    | Morgan Stanle     | 
| MORN  | Morningstar       | 
| MOS   | Mosaic Compa      | 
| MSI   | Motorola Soluti   | 
| MPLX  | MPLX LP           | 
| MUFG  | MUFJ Financia     | 
| MUR   | Murphy Oil Cor    | 
| MYL   | Mylan NV          | 
| NBR   | Nabors Industr    | 
| NDAQ  | Nasdag OMX G      | 
| NOV   | National ilwel    | 
| NTCO  | Natura Holding    | 
| NAVI  | Navient Corpor    | 
| NKTR  | Nektar Therape    | 
| NPTN  | NeoPhotonics      | 
| NTAP  | NetApp            | 
| NTES  | NetEase           | 
| NFLX  | Netflix           | 
| NBIX  | Neurocrine Bio    | 
| NBEV  | New Age Bever     | 
| EDU   | New Oriental E    | 
| NRZ   | New Residenti     | 
| NYCB  | New York Com      | 
| NYMT  | New York Mort     | 
| NYT   | New York Time     | 
| NWL   | Newell Rubberr    | 
| NEM   | Newmont Mini      | 
| NWSA  | News Corporat     | 
| NEE   | NextEra Energy    | 
| NKE   | Nike              | 
| NI    | NIO               | 
| NOAH  | Noah Holdings     | 
| NBL   | Noble Energy      | 
| NMR   | Nomura            | 
| JWN   | Nordstrom         | 
| NSC   | Norfolk Southe    | 
| NTRS  | Northern Trust    | 
| NOC   | Northrop Grum     | 
| NLOK  | NortonLifeLocl    | 
| NCLH  | Norwegian Cru     | 
| NVCR  | NovoCure          | 
| NRG   | NRG Energy        | 
| NUE   | Nucor             | 
| NTNX  | Nutanix           | 
| NUVA  | NuVasive          | 
| NVDA  | Nvidia            | 
| ORLY  | O'Reilly Auto P   | 
| AS    | Oasis Petroleu    | 
| OX    | Occidental        | 
| P     | Office Depot      | 
| OKTA  | Okta              | 
| OLN   | Olin              | 
| M     | Omnicom Grou      | 
| ON    | ON Semicondu      | 
| OKE   | Oneok             | 
| PK    | Opko Health       | 
| ORCL  | Oracle            | 
| ONVO  | Organovo Hold     | 
| STK   | Overstock.com     | 
| PG    | P&G               | 
| PCAR  | Paccar            | 
| PD    | Pagerduty         | 
| PANW  | Palo Alto Netw    | 
| PAM   | Pampa Energia     | 
| PAAS  | Pan American      | 
| PH    | Parker-Hannifi    | 
| PE    | Parsley Energy    | 
| PTEN  | Patterson-UTI \|  | 
| PAYX  | Paychex           | 
| PYPL  | PayPal            | 
| PBF   | PBF Energy        | 
| PTON  | Peloton Interac   | 
| PBCT  | People's Unitec   | 
| PEP   | PepsiCo           | 
| PBR   | Petroleo Brasil   | 
| PFE   | Pfizer            | 
| PCG   | PGAE Corporat     | 
| PM    | Philip Morris     | 
| PSX   | Phillips 66       | 
| PDD   | Pinduoduo Inc     | 
| PINS  | Pinterest         | 
| PXD   | Pioneer Natura    | 
| PBI   | Pitney Bowes      | 
| PVTL  | Pivotal Softwar   | 
| PAA   | Plains All Amer   | 
| PLUG  | Plug Power        | 
| PS    | Pluralsight       | 
| PNC   | PNC Financial     | 
| PPG   | PPG Industries    | 
| PPL   | PPL Corporatio    | 
| PBH   | Prestige Brand    | 
| PRI   | Primerica         | 
| PGR   | Progressive       | 
| PLD   | Prologis          | 
| PFPT  | Proofpoint        | 
| PRU   | Prudential Fina   | 
| PEG   | Public Service\|  | 
| PSA   | Public Storage    | 
| PHM   | PulteGroup        | 
| PSTG  | Pure Storage      | 
| QTWO  | Q2 Holdings       | 
| QP    | QEP Resources     | 
| QCOM  | Qualcomm          | 
| L     | Qualys            | 
| PWR   | Quanta Service    | 
| QD    | Qudian            | 
| QRTEA | Qurate Retail     | 
| RL    | Ralph Lauren      | 
| RRC   | Range Resourc     | 
| RTN   | Raytheon Com      | 
| RLGY  | Realogy Holdin    | 
| REAL  | RealReal          | 
| O     | Realty Income     | 
| REGN  | Regeneron Pha     | 
| RF    | Regions Financ    | 
| REGI  | Renewable Ene     | 
| RMD   | ResMed            | 
| QSR   | Restaurant Bra    | 
| RNG   | RingCentral       | 
| RAD   | Rite Aid          | 
| ROK   | Rockwell Autor    | 
| ROKU  | Roku              | 
| ROST  | Ross Stores       | 
| RY    | Royal Bank of (   | 
| RCL   | Royal Caribbea    | 
| RES   | RPC               | 
| SPGI  | S                 | 
| CRM   | Salesforce        | 
| SBH   | Sally Beauty      | 
| SGMO  | Sangamo Ther.     | 
| SLB   | Schlumberger      | 
| STX   | Seagate Techn     | 
| SRE   | Sempra Energy     | 
| SENS  | Senseonics Ho     | 
| SHAK  | Shake Shack       | 
| SHW   | Sherwin-Wiliar    | 
| SHOP  | Shopify           | 
| SBGL  | Sibanye Gold L    | 
| SPG   | Simon Propert)    | 
| SIRI  | Sirius XM         | 
| SKX   | Skechers          | 
| SWKS  | Skyworks Solu     | 
| WORK  | Slack             | 
| SLM   | SLM Corporati     | 
| M     | SM Energy Con     | 
| SMAR  | Smartsheet        | 
| SNAP  | Snap              | 
| SGO   | Sogou             | 
| SNE   | Sony              | 
| BID   | Sotheby's         | 
| SU    | Southern Comy     | 
| SO    | Southem Copp      | 
| LV    | Southwest Airl    | 
| SWN   | Southwestern \|   | 
| ONCE  | Spark Therapet    | 
| SPLK  | Splunk            | 
| SPOT  | Spotify           | 
| S     | Sprint Corporat   | 
| SFM   | Sprouts Farme     | 
| SU    | Square            | 
| SRCI  | SRC Energy        | 
| SSNC  | SS&C Technolc     | 
| SWK   | Stanley Black &   | 
| SBUX  | Starbucks         | 
| TSG   | Stars Group       | 
| STT   | State Street      | 
| STLD  | Steel Dynamic:    | 
| SFIX  | Stitch Fix        | 
| STNE  | StoneCo           | 
| SYK   | Stryker           | 
| SMFG  | Sumitomo Mits     | 
| SU    | Suncor Energy     | 
| SPWR  | SunPower          | 
| RUN   | Sunrun            | 
| ST    | SunTrust Bank     | 
| SSS   | Sutter Rock Ca    | 
| SYF   | Synchrony Finz    | 
| SNPS  | Synopsys          | 
| SYY   | Sysco             | 
| TMUS  | T-Mobile          | 
| TROW  | T. Rowe Price     | 
| WO    | T2 Interactive    | 
| TLRD  | Tailored Brands   | 
| TSM   | Taiwan Semico     | 
| TAK   | Takeda Pharme     | 
| TAL   | Tal Education ¢   | 
| SKT   | Tanger Factory    | 
| TPR   | Tapestry          | 
| TEDU  | Tarena Internat   | 
| TRGP  | Targa Resource    | 
| TT    | Target            | 
| TTM   | Tata Motors       | 
| AMTD  | TD Ameritrade     | 
| FTI   | TechnipFMC pl     | 
| TECK  | Teck Resource     | 
| TFX   | Teleflex          | 
| VIV   | Telefonica Bras   | 
| TELL  | Tellurian         | 
| X     | Tempur Sealy      | 
| TME   | Tencent Music     | 
| TER   | Teradyne          | 
| TSLA  | Tesla             | 
| TEVA  | Teva Pharmace     | 
| XN    | Texas Instrume    | 
| DOW   | The Dow Chem      | 
| GEO   | The GEO Groug     | 
| TXMD  | TherapeuticsM     | 
| O     | Thermo Fisher     | 
| TIF   | Tiffany & Co      | 
| TSU   | TIM Participag    | 
| X     | TJX Companie:     | 
| TD    | Toronto-Domin     | 
| TM    | Toyota            | 
| TW    | Tradeweb Mark     | 
| TRXC  | TransEnterix      | 
| TRV   | Travelers         | 
| TCOM  | Trip.com          | 
| TRIP  | TripAdvisor       | 
| TI    | Triumph           | 
| TFC   | Truist Financia   | 
| TRQ   | Turquoise Hil     | 
| TWLO  | Twilio            | 
| TWTR  | Twitter           | 
| TWO   | Two Harbors In    | 
| TSN   | Tyson Foods       | 
| UBER  | Uber              | 
| ULTA  | Ulta Beauty       | 
| TT    | Ultra Clean Hol   | 
| UGP   | Ultrapar Partici  | 
| UAA   | Under Armour      | 
| A     | Under Armour (    | 
| UNP   | Union Pacific     | 
| VAL   | United Contine    | 
| UM    | United Microel    | 
| X     | United States €   | 
| UTX   | United Technol    | 
| UNH   | UnitedHealth      | 
| UNIT  | Uniti Group       | 
| UNM   | Unum Group        | 
| TIGR  | UP Fintech        | 
| UPS   | upPs              | 
| URBN  | Urban Outfitter   | 
| US    | US Bancorp        | 
| SLCA  | US Silica Holdi   | 
| UXIN  | Uxin              | 
| VALE  | Vale              | 
| VL    | Valero            | 
| VGR   | Vector Group      | 
| VEEV  | Veeva System      | 
| VTR   | Ventas            | 
| VER   | Vereit            | 
| VRSN  | Verisign          | 
| VRSK  | Verisk Analytic   | 
| VZ    | Verizon           | 
| VRTX  | Vertex Pharma     | 
| VFC   | VF Corporation    | 
| VIACA | ViacomCBS         | 
| VI    | VICI Properties   | 
| VIPS  | Vipshop Holdir    | 
| SPCE  | Virgin Galactic   | 
| V     | Visa              | 
| VST   | Vistra Energy     | 
| MW    | VMware            | 
| V     | Vonage Holdin     | 
| WAB   | WABTEC Corpc      | 
| WBA   | Walgreens Boo     | 
| WM    | Walmart           | 
| DIS   | Walt Disney       | 
| WM    | Waste Manage      | 
| W     | Wayfair           | 
| W     | Weibo             | 
| WFC   | wells Fargo       | 
| WELL  | Welltower         | 
| WEN   | Wendy's Comp      | 
| WDC   | Western Digita    | 
| W     | Western Union     | 
| WRK   | WestRock          | 
| WEX   | WEX               | 
| WY    | Weyerhaeuser      | 
| WLL   | Whiting Petrole   | 
| WMB   | Williams Comp     | 
| WING  | Wingstop          | 
| WIT   | Wipro             | 
| WDAY  | Workday           | 
| WWE   | World Wrestling   | 
| WPG   | WP Glimcher       | 
| WPX   | WPX Energy        | 
| WYNN  | Wynn Resorts      | 
| XEL   | Xcel Energy       | 
| XRX   | Xerox             | 
| XLNX  | Xilinx            | 
| AUY   | Yamana Gold       | 
| YPF   | YPF               | 
| YUM   | Yum!              | 
| YUMC  | Yum!              | 
| YY    | YY                | 
| ZAYO  | Zayo Group Ho!    | 
| ZBRA  | Zebra Technolc    | 
| ZEN   | Zendesk           | 
| Z     | Zillow Group      | 
| ZBH   | Zimmer Biome      | 
| ZION  | Zions Bancorpc    | 
| ZTS   | Zoetis            | 
| ZM    | Zoom              | 
| ZS    | Zscaler           | 
| ZTO   | ZTO Express       | 
| ZNGA  | Zynga             |
