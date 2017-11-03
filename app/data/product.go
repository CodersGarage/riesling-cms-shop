package data

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2"
	"riesling-cms-shop/app/conn"
	"math"
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

func (p *Product) ProductQuery(query bson.M, pageNow int) (*mgo.Query, *bongo.PaginationInfo) {
	if pageNow <= 0 {
		pageNow = 1
	}
	queryPage := pageNow
	if queryPage > 0 {
		queryPage--
	}
	collection := conn.GetCollection(PRODUCT_COLLECTION_NAME)
	q := collection.Collection().Find(query).
		Skip(PRODUCT_PER_PAGE * queryPage).
		Limit(PRODUCT_PER_PAGE).
		Sort("-$natural")

	count, _ := q.Count()
	pagination := p.Paginate(query, count, pageNow)
	pagination.Current = pageNow
	return q, pagination
}

func (p *Product) Paginate(query bson.M, count, page int) *bongo.PaginationInfo {
	op, _ := conn.GetCollection(PRODUCT_COLLECTION_NAME).Find(query).Paginate(PRODUCT_PER_PAGE, page)
	info := &bongo.PaginationInfo{}
	totalPages := op.TotalPages
	if page < 1 {
		page = 1
	} else if page > totalPages {
		page = totalPages
	}

	info.TotalPages = totalPages
	info.PerPage = PRODUCT_PER_PAGE
	info.Current = page
	info.TotalRecords = op.TotalRecords

	if info.Current < info.TotalPages {
		info.RecordsOnPage = info.PerPage
	} else {
		info.RecordsOnPage = int(math.Mod(float64(count), float64(PRODUCT_PER_PAGE)))
		if info.RecordsOnPage == 0 && count > 0 {
			info.RecordsOnPage = PRODUCT_PER_PAGE
		}
	}
	return info
}

func (p *Product) FindQuery(query bson.M, pageNow int) (*mgo.Iter, *bongo.PaginationInfo) {
	q, info := p.ProductQuery(query, pageNow)
	return q.Iter(), info
}

func (p *Product) Find(hash string) bool {
	it, _ := p.FindQuery(bson.M{
		"hash": hash,
	}, 0)
	return it.Next(p)
}

func (p *Product) IsProductExists(code string) bool {
	it, _ := p.FindQuery(bson.M{
		"code": code,
	}, 0)
	return it.Next(p)
}

func (p *Product) FindStat(info *bongo.PaginationInfo) ProductStat {
	stat := ProductStat{}
	if info != nil {
		stat.PageOn = info.Current
		stat.ProductOnPage = info.RecordsOnPage
		stat.ProductPerPage = info.PerPage
		stat.ProductTotal = info.TotalRecords
		stat.PageTotal = info.TotalPages
	}
	return stat
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
	q, info := p.ProductQuery(bson.M{}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.ProductStat = p.FindStat(info)
	result.Products = p.IteratorToArray(it)
	return result
}

func (p *Product) FindPublished(pageNow int) ProductResult {
	q, info := p.ProductQuery(bson.M{
		"status": PRODUCT_STATUS_PUBLISHED,
	}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStat(info)
	return result
}

func (p *Product) FindDrafts(pageNow int) ProductResult {
	q, info := p.ProductQuery(bson.M{
		"status": PRODUCT_STATUS_DRAFT,
	}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStat(info)
	return result
}
