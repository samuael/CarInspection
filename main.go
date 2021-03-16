package main

import (
	"github.com/samuael/Project/CarInspection/pkg/constants/model"
	"github.com/samuael/Project/CarInspection/platforms/helper"
)

func main() {
	insup := &model.InspectionUpdate{}
	jres := helper.MarshalThis(insup)
	println(string(jres))
}
