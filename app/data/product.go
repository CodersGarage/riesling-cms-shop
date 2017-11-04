package data

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2"
	"riesling-cms-shop/app/conn"
	"math"
	"riesling-cms-shop/app/utils"
)

/**
 * := Coded with love by Sakib Sami on 3/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

const (
	ProductCollectionName        = "products"
	ProductStatusDraft           = 0
	ProductStatusPublished       = 1
	ProductPerPage               = 10
	ProductPaginationDefaultPage = 1
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
	DownloadLink     string   `json:"-"`
	TotalDownload    int      `json:"total_download"`
	ExtraInformation string   `json:"extra_information"`
	CreatedBy        string   `json:"-"`
	UpdatedBy        string   `json:"-"`
}

type ProductRequest struct {
	Hash             string   `json:"hash"`
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Price            float32  `json:"price"`
	Discount         float32  `json:"discount"`
	Status           int      `json:"status"`
	Tags             []string `json:"tags"`
	Previews         []string `json:"previews"`
	DownloadLink     string   `json:"download_link"`
	ExtraInformation string   `json:"extra_information"`
}

type ProductStat struct {
	ProductPerPage int `json:"per_page"`
	ProductOnPage  int `json:"products"`
	ProductTotal   int `json:"product_total"`
	PageTotal      int `json:"page_total"`
	PageOn         int `json:"page_on"`
}

type ProductResult struct {
	Products    []Product    `json:"products"`
	ProductStat *ProductStat `json:"stat,omitempty"`
}

func (pr *ProductRequest) ProcessCreate() Product {
	product := Product{}
	product.Hash = utils.GetUUID()
	product.Code = pr.Code
	product.Name = pr.Name
	product.Description = pr.Description
	product.Price = pr.Price
	product.Discount = pr.Discount
	product.Status = pr.Status
	product.Tags = pr.Tags
	product.Previews = pr.Previews
	product.Favourites = 0
	product.DownloadLink = pr.DownloadLink
	product.TotalDownload = 0
	product.ExtraInformation = pr.ExtraInformation
	return product
}

func (pr *ProductRequest) ProcessUpdate() (Product, bool) {
	product := Product{}
	if product.Find(pr.Hash) {
		product.Code = pr.Code
		product.Name = pr.Name
		product.Description = pr.Description
		product.Price = pr.Price
		product.Discount = pr.Discount
		product.Status = pr.Status
		product.Tags = pr.Tags
		product.Previews = pr.Previews
		product.DownloadLink = pr.DownloadLink
		product.ExtraInformation = pr.ExtraInformation
		return product, true
	}
	return product, false
}

func (p *Product) Save() bool {
	if err := conn.GetCollection(ProductCollectionName).Save(p); err == nil {
		return true
	}
	return false
}

func (p *Product) Update() (bool, *Product) {
	return p.Save(), p
}

func (p *Product) Delete(hash string) bool {
	if p.Find(hash) {
		err := conn.GetCollection(ProductCollectionName).DeleteDocument(p)
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
	collection := conn.GetCollection(ProductCollectionName)
	q := collection.Collection().Find(query).
		Skip(ProductPerPage * queryPage).
		Limit(ProductPerPage).
		Sort("-$natural")

	count, _ := q.Count()
	pagination := p.Paginate(query, count, pageNow)
	pagination.Current = pageNow
	return q, pagination
}

func (p *Product) Paginate(query bson.M, count, page int) *bongo.PaginationInfo {
	op, _ := conn.GetCollection(ProductCollectionName).Find(query).Paginate(ProductPerPage, page)
	info := &bongo.PaginationInfo{}
	totalPages := op.TotalPages
	if page < 1 {
		page = 1
	} else if page > totalPages {
		page = totalPages
	}

	info.TotalPages = totalPages
	info.PerPage = ProductPerPage
	info.Current = page
	info.TotalRecords = op.TotalRecords

	if info.Current < info.TotalPages {
		info.RecordsOnPage = info.PerPage
	} else {
		info.RecordsOnPage = int(math.Mod(float64(count), float64(ProductPerPage)))
		if info.RecordsOnPage == 0 && count > 0 {
			info.RecordsOnPage = ProductPerPage
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
	}, ProductPaginationDefaultPage)
	return it.Next(p)
}

func (p *Product) IsProductExists(code string) bool {
	it, _ := p.FindQuery(bson.M{
		"code": code,
	}, ProductPaginationDefaultPage)
	return it.Next(p)
}

func (p *Product) FindStat(info *bongo.PaginationInfo) *ProductStat {
	stat := ProductStat{}
	if info != nil {
		stat.PageOn = info.Current
		stat.ProductOnPage = info.RecordsOnPage
		stat.ProductPerPage = info.PerPage
		stat.ProductTotal = info.TotalRecords
		stat.PageTotal = info.TotalPages
	}
	return &stat
}

func (p *Product) IteratorToArray(it *mgo.Iter) []Product {
	products := []Product{}
	temp := Product{}
	for it.Next(&temp) {
		products = append(products, temp)
	}
	return products
}

func (p *Product) IteratorToArrayPublishedOnly(it *mgo.Iter) []Product {
	products := []Product{}
	temp := Product{}
	for it.Next(&temp) {
		if temp.Status == ProductStatusPublished {
			products = append(products, temp)
		}
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
		"status": ProductStatusPublished,
	}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStat(info)
	return result
}

func (p *Product) FindDrafts(pageNow int) ProductResult {
	q, info := p.ProductQuery(bson.M{
		"status": ProductStatusDraft,
	}, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArray(it)
	result.ProductStat = p.FindStat(info)
	return result
}

func (p *Product) Search(query bson.M, pageNow int) ProductResult {
	q, _ := p.ProductQuery(query, pageNow)
	it := q.Iter()
	result := ProductResult{}
	result.Products = p.IteratorToArrayPublishedOnly(it)
	return result
}
