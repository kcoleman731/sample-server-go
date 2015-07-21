package controller

import (
	"database/sql"
	"net/http"
)

type FeatureController struct {
	DB *sql.DB
}

func (controller *FeatureController) Post(request *http.Request) {

}

func (controller *FeatureController) Get(request *http.Request) {

}

func (controller *FeatureController) Update(request *http.Request) {

}

func (controller *FeatureController) Delete(request *http.Request) {

}
