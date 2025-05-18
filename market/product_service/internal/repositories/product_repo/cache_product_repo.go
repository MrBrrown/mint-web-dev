package productrepo

import (
	"sync"
	"time"
)

type CacheEntry[T any] struct {
	value   T
	expires time.Time
}

type CacheProductRepo struct {
	repo      ProductRepo
	mu        sync.RWMutex
	prodCache map[uint]CacheEntry[Product]
	listCache *CacheEntry[[]Product]
	ttl       time.Duration
}

func NewCacheRepo(r ProductRepo, ttl time.Duration) *CacheProductRepo {
	return &CacheProductRepo{
		repo:      r,
		prodCache: make(map[uint]CacheEntry[Product]),
		ttl:       ttl,
	}
}

func (r *CacheProductRepo) CreateProduct(product Product) (Product, error) {
	p, err := r.repo.CreateProduct(product)
	if err != nil {
		return Product{}, err
	}

	r.invalidateAll()
	return p, nil
}

func (r *CacheProductRepo) GetProduct(id uint) (Product, error) {
	r.mu.RLock()
	if e, ok := r.prodCache[id]; ok && time.Now().Before(e.expires) {
		r.mu.RUnlock()
		return e.value, nil
	}
	r.mu.RUnlock()

	product, err := r.repo.GetProduct(id)
	if err != nil {
		return Product{}, err
	}

	r.mu.Lock()
	r.prodCache[id] = CacheEntry[Product]{value: product, expires: time.Now().Add(r.ttl)}
	r.mu.Unlock()

	return product, nil
}

func (r *CacheProductRepo) UpdateProduct(id uint, product Product) (Product, error) {
	p, err := r.repo.UpdateProduct(id, product)
	if err != nil {
		return Product{}, err
	}

	r.invalidate(id)
	return p, nil
}

func (r *CacheProductRepo) DeleteProduct(id uint) error {
	err := r.repo.DeleteProduct(id)
	if err != nil {
		return err
	}

	r.invalidateAll()
	return nil
}

func (r *CacheProductRepo) GetProduts() ([]Product, error) {
	r.mu.RLock()
	if r.listCache != nil && time.Now().Before(r.listCache.expires) {
		vals := r.listCache.value
		r.mu.RUnlock()
		return vals, nil
	}
	r.mu.RUnlock()

	list, err := r.repo.GetProduts()
	if err != nil {
		return nil, err
	}

	r.mu.Lock()
	r.listCache = &CacheEntry[[]Product]{value: list, expires: time.Now().Add(r.ttl)}
	r.mu.Unlock()

	return list, nil
}

func (c *CacheProductRepo) invalidateAll() {
	c.mu.Lock()
	c.prodCache = make(map[uint]CacheEntry[Product])
	c.listCache = nil
	c.mu.Unlock()
}

func (c *CacheProductRepo) invalidate(id uint) {
	c.mu.Lock()
	delete(c.prodCache, id)
	c.listCache = nil
	c.mu.Unlock()
}
