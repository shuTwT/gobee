import { defineCustomElement } from 'vue'
import ProductCard from './components/ProductCard.ce.vue'

const ProductCardElement = defineCustomElement(ProductCard)

export {
    ProductCardElement
}

export function register() {
  if (typeof window != undefined) {
    
    if(!customElements.get('product-card') ){
      customElements.define('product-card', ProductCardElement)
    }
    
  }else {
    console.log("window is undefined")
  }
}
