package pkg

import (
	"fmt"
	"strings"

	"github.com/layer5io/meshkit/utils/csv"
)

var (
	componentSheetID = "985575256"
	shouldRegisterComp				 = "component" 
)

type ComponentCSV struct {
	ModelDisplayName string `json:"modelDisplayName"` // change to component
	Model string `json:"model"`
	Component string `json:"component"`
	// Registrant string `json:"registrant"`
	Category string `json:"category"`
	SubCategory string `json:"subCategory"`
	Link string `json:"ling"` // convert to source
	Website string `json:"website"`
	Docs string `json:"docs"`
	Shape string `json:"shape"`
	PrimaryColor string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
	StyleOverrides string `json:"styleOverrides"`
	Styles string `json:"styles"`
	ShapePolygonPoints string `json:"shapePolygonPoints"`
	DefaultData string `json:"defaultData"`
	Capabilities string `json:"capabilities"`
	LogoURL string `json:"logoURL"`
	SVGColor string `json:"svgColor"`
	SVGWhite string `json:"svgWhite"`
	SVGComplete string `json:"svgComplete"`
	PublishToRegistry string `json:"publishToRegistry"`
	ModelAnnotation string `json:"modelAnnotation"`// update to componentAnnotation
	AboutProject string `json:"aboutProject"`
	PageSubtTitle string `json:"pageSubtitle"`
	Feature1 string `json:"feature1"`
	Feature2 string `json:"feature2"`
	Feature3 string `json:"feature3"`
	HowItWorks string `json:"howItWorks"`
	HowItWorksDetails string `json:"howItWorksDetails"`
	StandardBlurb string `json:"standardBlurb"`
	Screenshots string `json:"screenshots"`
	FullPage string `json:"fullPage"`
}

type ComponentCSVHelper struct {
	SheetID string
	CSVPath string
	Components map[string][]ComponentCSV
}

func NewComponentCSVHelper(sheetURL string) (*ComponentCSVHelper, error) {
	sheetURL = sheetURL + "&gid=" + componentSheetID
	csvPath, err := DownloadCSV(sheetURL)
	if err != nil {
		return nil, err
	}

	return &ComponentCSVHelper{
		SheetID: componentSheetID,
		CSVPath: csvPath,
	}, nil
}

func getIndexForRegisterColForComp(cols []string) int {
	if shouldRegisterColIndex != -1 {
		return shouldRegisterColIndex
	}

	for index, col := range cols {
		if col == shouldRegisterComp {
			return index
		}
	}
	return shouldRegisterColIndex
}


func (mch *ComponentCSVHelper) ParseComponentsSheet(){
	ch := make(chan ComponentCSV, 1)
	errorChan := make(chan error, 1)
	csvReader, err := csv.NewCSVParser[ComponentCSV](mch.CSVPath, rowIndex, nil, func(columns []string, currentRow []string) bool {
		index := getIndexForRegisterColForComp(columns)
		idx := getIndexForRegisterColForComp(columns)
		flag := false
		if index != -1 && index < len(currentRow){
			shouldRegister := currentRow[index]
			flag = strings.ToLower(shouldRegister) == "true"
		}
		if idx != -1 && idx < len(currentRow){
			shouldRegister := currentRow[idx]
			flag = flag && strings.ToLower(shouldRegister) == ""
		}
		return flag
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		err := csvReader.Parse(ch, errorChan)
		if err != nil {
			fmt.Println(err)
		}
	}()
	for {
		select {

		case data := <-ch:
			if mch.Components[data.Model] == nil {
				mch.Components[data.Model] = make([]ComponentCSV, 0)
				return
			}
			mch.Components[data.Model] = append(mch.Components[data.Model], data)
		case err := <-errorChan:
			fmt.Println(err)

		case <-csvReader.Context.Done():
			return
		}
	}
}