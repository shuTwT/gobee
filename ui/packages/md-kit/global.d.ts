import ProductCard from "./src/components/ProductCard.ce.vue"
declare module 'vue' {
  export interface GlobalComponents {
    'product-card':typeof ProductCard
  }
}

export {}