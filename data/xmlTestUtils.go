package data

import (
	"github.com/subchen/go-xmldom"
	"io/ioutil"
)

func fileContents2String(fileName string) (string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func getDomRootElementFromFile(xmlFileName string) (*xmldom.Node, error) {
	xml, err := fileContents2String(xmlFileName)
	if err != nil {
		return nil, err
	}
	doc := xmldom.Must(xmldom.ParseXML(xml))
	return doc.Root, nil

}

func getNumTestPlans(xmlFileName string) (int, error) {
	root, err := getDomRootElementFromFile(xmlFileName)
	if err != nil {
		return -1, err
	}

	// find all testPlans
	testPlans := root.Query("//TestPlan")
	return len(testPlans), nil
}

func getTestPlanNames(xmlFileName string) ([]string, error) {
	root, err := getDomRootElementFromFile(xmlFileName)
	if err != nil {
		return make([]string, 0, 0), err
	}
	// find all testPlans
	testPlans := root.Query("//TestPlan")
	planNames := make([]string, len(testPlans), len(testPlans))
	for i, c := range testPlans {
		planNames[i] = c.GetAttributeValue("testname")
	}
	return planNames, nil

}

func getTestPlanEnabledVals(xmlFileName string) ([]string, error) {
	root, err := getDomRootElementFromFile(xmlFileName)
	if err != nil {
		return make([]string, 0, 0), err
	}
	// find all testPlans
	testPlans := root.Query("//TestPlan")
	planEnabledVals := make([]string, len(testPlans), len(testPlans))
	for i, c := range testPlans {
		planEnabledVals[i] = c.GetAttributeValue("enabled")
	}
	return planEnabledVals, nil

}
