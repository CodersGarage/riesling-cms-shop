package api

import (
	"net/http"
	"riesling-cms-shop/app/data"
	"github.com/s4kibs4mi/govalidator"
	"riesling-cms-shop/app/utils"
	"strconv"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := data.Product{}
	rules := govalidator.MapData{
		"code": []string{"required", "between:1,50"},
		"name": []string{"required", "between:1,100"},
	}
	options := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &product,
	}
	vr := govalidator.New(options)
	err := vr.ValidateJSON()
	if len(err) == 0 {
		product.Hash = utils.GetUUID()
		product.Favourites = 0
		product.TotalDownload = 0
		if !product.IsProductExists(product.Code) {
			if product.Save() {
				resp := APIResponse{
					Code:    http.StatusOK,
					Message: "Product has been created.",
					Data:    product,
				}
				ServeAsJSON(resp, w)
				return
			}
			resp := APIResponse{
				Code:    http.StatusInternalServerError,
				Message: "Something went wrong.",
			}
			ServeAsJSON(resp, w)
			return
		}
		resp := APIResponse{
			Code:    http.StatusConflict,
			Message: "Product code already exists.",
		}
		ServeAsJSON(resp, w)
		return
	}
	resp := APIResponse{
		Code:  http.StatusUnprocessableEntity,
		Error: err,
	}
	ServeAsJSON(resp, w)
}

func FindProducts(w http.ResponseWriter, r *http.Request) {
	product := data.Product{}
	value := GetURLParam("page", r)
	page, err := strconv.Atoi(value)
	if err != nil {
		products := product.FindAll(0)
		resp := APIResponse{
			Code: http.StatusOK,
			Data: products,
		}
		ServeAsJSON(resp, w)
		return
	}
	products := product.FindAll(page)
	resp := APIResponse{
		Code: http.StatusOK,
		Data: products,
	}
	ServeAsJSON(resp, w)
	return
}

func FindProductsPublished(w http.ResponseWriter, r *http.Request) {
	product := data.Product{}
	value := GetURLParam("page", r)
	page, err := strconv.Atoi(value)
	if err != nil {
		products := product.FindPublished(0)
		resp := APIResponse{
			Code: http.StatusOK,
			Data: products,
		}
		ServeAsJSON(resp, w)
		return
	}
	products := product.FindPublished(page)
	resp := APIResponse{
		Code: http.StatusOK,
		Data: products,
	}
	ServeAsJSON(resp, w)
	return
}

func FindProductsDrafts(w http.ResponseWriter, r *http.Request) {
	product := data.Product{}
	value := GetURLParam("page", r)
	page, err := strconv.Atoi(value)
	if err != nil {
		products := product.FindDrafts(0)
		resp := APIResponse{
			Code: http.StatusOK,
			Data: products,
		}
		ServeAsJSON(resp, w)
		return
	}
	products := product.FindDrafts(page)
	resp := APIResponse{
		Code: http.StatusOK,
		Data: products,
	}
	ServeAsJSON(resp, w)
	return
}

func SearchProducts(w http.ResponseWriter, r *http.Request) {
	product := data.Product{}
	value := GetURLParam("page", r)
	page, err := strconv.Atoi(value)
	var query bson.M
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &query)
	if err != nil {
		fmt.Println(err)

		products := product.Search(query, 0)
		resp := APIResponse{
			Code: http.StatusOK,
			Data: products,
		}
		ServeAsJSON(resp, w)
		return
	}
	products := product.Search(query, page)
	resp := APIResponse{
		Code: http.StatusOK,
		Data: products,
	}
	ServeAsJSON(resp, w)
	return
}
