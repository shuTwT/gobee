<script setup lang="ts">
const props = withDefaults(defineProps<{
    productId?:number|string
    name?:string
    price?:number
    img?:string
    description?:string
}>(),{
    name:'未知商品',
    price:0,
})

const formatedPrice = computed(() => {
    return props.price.toFixed(2)
})
</script>
<template>
    <div class="product-card">
        <div class="product-name">{{ props.name }}</div>
        <span class="product-price">{{ formatedPrice }}</span>
        <button class="buy-button">购买</button>
    </div>
</template>
<style scoped>
.product-card{
    padding: 10px 20px;
    background-color: #e0dbe1;
    border-radius: 4px;
}
.product-card .buy-button{
    padding: 4px 8px;
    background-color: #18a058;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}
</style>