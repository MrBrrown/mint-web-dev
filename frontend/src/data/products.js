export const products = Array.from({ length: 20 }, (_, i) => ({
  id: i + 1,
  name: `Лампочка ${i + 1}`,
  image: 'https://img.lovepik.com/free-png/20210918/lovepik-light-bulb-png-image_400229766_wh1200.png',
  description: `Подробное описание лампочки ${i + 1}. Энергосберегающая LED-лампа, срок службы 25000 часов.`,
  cost: Math.floor(100 + Math.random() * 500),
  wattage: [5, 7, 10, 15][Math.floor(Math.random() * 4)]
}));
