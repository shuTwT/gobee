package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// 商品
type Product struct {
	ent.Schema
}

func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255).
			Comment("商品名称"),
		field.String("description").
			Optional().
			MaxLen(1000).
			Comment("商品描述"),
		field.String("short_description").
			Optional().
			MaxLen(500).
			Comment("简短描述"),
		field.String("sku").
			Unique().
			NotEmpty().
			MaxLen(100).
			Comment("商品SKU"),
		field.Int("price").
			Default(0).
			Comment("商品价格(分)"),
		field.Int("original_price").
			Optional().
			Comment("原价(分)"),
		field.Int("cost_price").
			Optional().
			Comment("成本价(分)"),
		field.Int("stock").
			Default(0).
			Comment("库存数量"),
		field.Int("min_stock").
			Default(0).
			Comment("最低库存预警"),
		field.Int("sales").
			Default(0).
			Comment("销售数量"),
		field.Int("category_id").
			Optional().
			Comment("分类ID"),
		field.String("brand").
			Optional().
			MaxLen(100).
			Comment("品牌"),
		field.String("unit").
			Optional().
			MaxLen(20).
			Comment("单位"),
		field.Float("weight").
			Optional().
			Comment("重量(kg)"),
		field.Float("volume").
			Optional().
			Comment("体积(立方米)"),
		field.JSON("images", []string{}).
			Optional().
			Comment("商品图片列表"),
		field.JSON("attributes", map[string]interface{}{}).
			Optional().
			Comment("商品属性"),
		field.JSON("tags", []string{}).
			Optional().
			Comment("商品标签"),
		field.Bool("active").
			Default(true).
			Comment("是否上架"),
		field.Bool("featured").
			Default(false).
			Comment("是否推荐"),
		field.Bool("digital").
			Default(false).
			Comment("是否数字商品"),
		field.String("meta_title").
			Optional().
			MaxLen(255).
			Comment("SEO标题"),
		field.String("meta_description").
			Optional().
			MaxLen(500).
			Comment("SEO描述"),
		field.String("meta_keywords").
			Optional().
			MaxLen(255).
			Comment("SEO关键词"),
		field.Int("sort_order").
			Default(0).
			Comment("排序"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		// 可以添加与其他实体的关联，如分类、品牌等
	}
}
