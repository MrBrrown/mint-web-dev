import { products } from './data/products';

export function fetchProducts() {
  return new Promise((resolve) => {
    setTimeout(() => resolve(products), 300);
  });
}

export function fetchProductById(id) {
  return fetchProducts().then(items => items.find(p => p.id === id));
}
