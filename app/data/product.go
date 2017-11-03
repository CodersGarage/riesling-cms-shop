package data

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2"
	"math"
	"riesling-cms-shop/app/conn"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

const (
	PRODUCT_COLLECTION_NAME  = "products"
	PRODUCT_STATUS_DRAFT     = 0
	PRODUCT_STATUS_PUBLISHED = 1
	PRODUCT_PER_PAGE         = 10
)

type Product struct {
	bongo.DocumentBase        `json:"-",bson:",inline"`
	Hash             string   `json:"hash"`
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Price            float32  `json:"price"`
	Discount         float32  `json:"discount"`
	Status           int      `json:"status"`
	Tags             []string `json:"tags"`
	Previews         []string `json:"previews"`
	Favourites       int      `json:"favourites"`
	DownloadLink     string   `json:"download_link"`
	TotalDownload    int      `json:"total_download"`
	ExtraInformation string   `json:"extra_information"`
	CreatedBy        string   `json:"created_by"`
	UpdatedBy        string   `json:"updated_by"`
}

type ProductStat struct {
	ProductPerPage int `json:"per_page"`
	ProductOnPage  int `json:"products"`
	ProductTotal   int `json:"product_total"`
	PageTotal      int `json:"page_total"`
	PageOn         int `json:"page_on"`
}

type ProductResult struct {
	Products    []Product   `json:"products"`
	ProductStat ProductStat `json:"stat"`
}

func (p *Product) Save() bool {
	if err := conn.GetCollection(PRODUCT_COLLECTION_NAME).Save(p); err == nil {
		return true
	}
	return false
}

func (p *Product) Update(hash string) (bool, *Product) {
	savedProduct := Product{}
	if savedProduct.Find(hash) {

		return savedProduct.Save(), &savedProduct
	}
	return false, p
}

func (p *Product) Delete(hash string) bool {
	if p.Find(hash) {
		err := conn.GetCollection(PRODUCT_COLLECTION_NAME).DeleteDocument(p)
		return err == nil
	}
	return false
}

func ProductQuery(query bson.M, pageNow int) *mgo.Query {
	collection := conn.GetMGoCollection(PRODUCT_COLLECTION_NAME)
	return collection.Find(query).
		Skip(PRODUCT_PER_PAGE * pageNow).
		Limit(PRODUCT_PER_PAGE).
		Sort("createdAt")
}

func (p *Product) CountQuery(query bson.M, pageNow int) int {
	n, err := ProductQuery(query, pageNow).Count()
	if err != nil {
		return 0
	}
	return n
}

func (p *Product) FindQuery(query bson.M, pageNow int) *mgo.Iter {
	return ProductQuery(query, pageNow).Iter()
}

func (p *Product) Find(hash string) bool {
	it := p.FindQuery(bson.M{
		"hash": hash,
	}, 0)
	return it.Next(p)
}

func (p *Product) IsProductExists(code string) bool {
	it := p.FindQuery(bson.M{
		"code": code,
	}, 0)
	return it.Next(p)
}

func (p *Product) FindStatQuery(query *mgo.Query, pageNow int) ProductStat {
	total, err := conn.GetMGoCollection(PRODUCT_COLLECTION_NAME).Count()
	count, err := query.Count()
	stat := ProductStat{}
	if err != nil {
		return stat
	}
	totalPages := int(math.Ceil(float64(total) / float64(PRODUCT_PER_PAGE)))
	stat.ProductPerPage = PRODUCT_PER_PAGE
	if count > PRODUCT_PER_PAGE {
		stat.ProductOnPage = PRODUCT_PER_PAGE
	} else {
		stat.ProductOnPage = count
	}
	stat.PageTotal = totalPages
	stat.ProductTotal = total
	stat.PageOn = pageNow
	return stat
}

func (p *Product) FindStat(query bson.M, pageNow int) ProductStat {
	q := ProductQuery(query, pageNow)
	return p.FindStatQuery(q, pageNow)
}

func (p *Product) IteratorToArray(it *mgo.Iter) []Product {
	products := []Product{}
	temp := Product{}
	for it.Next(&temp) {
		products = append(products, temp)
	}
	return products
}

func (p *Product) FindAll(pageNow int) ProductResult {
	q := ProductQuery(bson.M{}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStatQuery(q, pageNow)
	return result
}

func (p *Product) FindPublished(pageNow int) ProductResult {
	q := ProductQuery(bson.M{
		"status": PRODUCT_STATUS_PUBLISHED,
	}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStatQuery(q, pageNow)
	return result
}

func (p *Product) FindDrafts(pageNow int) ProductResult {
	q := ProductQuery(bson.M{
		"status": PRODUCT_STATUS_DRAFT,
	}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStatQuery(q, pageNow)
	return result
}
