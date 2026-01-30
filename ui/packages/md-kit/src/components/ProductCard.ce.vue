<script setup lang="ts">
import { computed } from 'vue';

const props = withDefaults(defineProps<{
    productId?: number | string
    name?: string
    price?: number
    img?: string
    description?: string
}>(), {
    name: '未知商品',
    price: 0,
})

const cssVars = {
    '--product-card-bg': '#ffffff', /* 卡片背景色 */
    '--product-card-border': '#f5f5f5', /* 卡片边框色 */
    '--product-card-shadow': 'rgba(0,0,0,0.08)', /* 卡片默认阴影 */
    '--product-card-shadow-hover': 'rgba(0,0,0,0.12)', /* 卡片hover阴影 */
    '--product-card-text-main': '#333333', /* 主文字色（商品名） */
    '--product-card-text-secondary': '#999999', /* 次要文字色（简介/原价/销量） */
    '--product-card-text-ellipsis': '#f5f5f5', /* 分割线色 */
    '--product-card-price-primary': '#ff4d4f', /* 现价主色 */
    '--product-card-tag-hot': '#fa8c16', /* 热销标签色 */
    '--product-card-tag-new': '#52c41a', /* 新品标签色 */
    '--product-card-tag-discount': '#ff4d4f', /* 特惠标签色（默认） */
    '--product-card-btn-bg': '#ff4d4f', /* 按钮背景色 */
    '--product-card-btn-bg-hover': '#ff383f', /* 按钮hover背景色 */
    '--product-card-btn-text': '#ffffff', /* 按钮文字色 */
    '--product-card-tag-text': '#ffffff', /* 标签文字色 */
}

const formatedPrice = computed(() => {
    return props.price.toFixed(2)
})
</script>
<template>
    <div class="product-card" :style="cssVars">
        <div class="product-img-box">
            <span class="product-tag hot">热销</span>
            <img src="https://picsum.photos/320/200" alt="商品图片" class="product-img">
        </div>
        <div class="product-info">
            <h3 class="product-name">轻奢质感保温杯 316不锈钢便携水杯</h3>
            <p class="product-desc">长效保温保冷8小时，食品级316不锈钢，防摔防滑杯底，多色可选</p>
            <div class="product-price">
                <span class="price-current">¥89</span>
                <span class="price-original">¥129</span>
            </div>
            <div class="product-footer">
                <span class="product-sales">已售2.3万+</span>
                <button class="product-btn">立即购买</button>
            </div>
        </div>
    </div>
</template>
<style scoped>
/* 卡片容器 */
.product-card {
    width: 320px;
    margin: 20px auto;
    border-radius: 12px;
    overflow: hidden;
    box-shadow: 0 4px 12px var(--product-card-shadow);
    border: 1px solid var(--product-card-border);
    transition: all 0.3s ease;
    background: var(--product-card-bg);
}

.product-card:hover {
    transform: translateY(-6px);
    box-shadow: 0 8px 20px var(--product-card-shadow-hover);
}

/* 商品图片容器 */
.product-img-box {
    width: 100%;
    height: 200px;
    overflow: hidden;
    position: relative;
}

.product-img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.5s ease;
}

.product-card:hover .product-img {
    transform: scale(1.05);
}

/* 商品标签 */
.product-tag {
    position: absolute;
    top: 10px;
    left: 10px;
    padding: 3px 10px;
    border-radius: 20px;
    font-size: 12px;
    color: var(--product-card-tag-text);
    background: var(--product-card-tag-discount);
}

.product-tag.hot {
    background: var(--product-card-tag-hot);
}

.product-tag.new {
    background: var(--product-card-tag-new);
}

/* 商品信息区域 */
.product-info {
    padding: 18px;
}

/* 商品名称 */
.product-name {
    font-size: 16px;
    font-weight: 600;
    color: var(--product-card-text-main);
    line-height: 1.4;
    margin-bottom: 8px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

/* 商品简介 */
.product-desc {
    font-size: 13px;
    color: var(--product-card-text-secondary);
    line-height: 1.5;
    margin-bottom: 12px;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
}

/* 价格区域 */
.product-price {
    display: flex;
    align-items: center;
    gap: 8px;
}

.price-current {
    font-size: 20px;
    font-weight: 700;
    color: var(--product-card-price-primary);
}

.price-original {
    font-size: 13px;
    color: var(--product-card-text-secondary);
    text-decoration: line-through;
}

/* 底部按钮/销量 */
.product-footer {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 15px;
    padding-top: 15px;
    border-top: 1px solid var(--product-card-text-ellipsis);
}

.product-sales {
    font-size: 12px;
    color: var(--product-card-text-secondary);
}

.product-btn {
    padding: 6px 15px;
    border-radius: 20px;
    font-size: 14px;
    color: var(--product-card-btn-text);
    background: var(--product-card-btn-bg);
    border: none;
    cursor: pointer;
    transition: background 0.3s ease;
}

.product-btn:hover {
    background: var(--product-card-btn-bg-hover);
}

/* 响应式适配 */
@media (max-width: 768px) {
    .product-card {
        width: 90%;
        max-width: 320px;
    }
}
</style>