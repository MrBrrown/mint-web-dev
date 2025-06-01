export async function fetchProducts() {
  const res = await fetch('/products/');
  if (!res.ok) {
    throw new Error('Ошибка загрузки продуктов');
  }

  const data = await res.json();

  return data.map(p => ({
    id: p.id,
    name: p.name,
    description: p.desc,
    image: 'https://img.lovepik.com/free-png/20210918/lovepik-light-bulb-png-image_400229766_wh1200.png',
    cost: p.price,
    attribs: p.attribs
  }));
}

export async function fetchProductById(id) {
  const res = await fetch(`/products/${id}`);
  if (!res.ok) {
    throw new Error(`Ошибка при получении товара с id=${id}`);
  }

  const p = await res.json();

  return {
    id: p.id,
    name: p.name,
    description: p.desc,
    image: 'https://img.lovepik.com/free-png/20210918/lovepik-light-bulb-png-image_400229766_wh1200.png',
    cost: p.price,
    attribs: p.attribs
  };
}
