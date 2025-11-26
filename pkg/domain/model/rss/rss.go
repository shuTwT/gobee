package rss

import "encoding/xml"

// --------------------------
// 定义 RSS 2.0 结构体（映射 XML 标签）
// --------------------------
// RSS2 对应根节点 <rss>
type RSS2 struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

// Channel 对应 <channel> 节点（RSS 核心容器）
type Channel struct {
	Title       string `xml:"title"`       // 频道标题
	Link        string `xml:"link"`        // 频道链接
	Description string `xml:"description"` // 频道描述
	PubDate     string `xml:"pubDate"`     // 频道发布时间
	Items       []Item `xml:"item"`        // 文章列表
}

// Item 对应 <item> 节点（单篇文章）
type Item struct {
	Title       string `xml:"title"`       // 文章标题
	Link        string `xml:"link"`        // 文章链接
	Description string `xml:"description"` // 文章摘要
	PubDate     string `xml:"pubDate"`     // 文章发布时间
	Author      string `xml:"author"`      // 文章作者（可选）
	GUID        string `xml:"guid"`        // 唯一标识（可选）
}

// --------------------------
// 定义 Atom 1.0 结构体（兼容）
// --------------------------
// AtomFeed 对应根节点 <feed>
type AtomFeed struct {
	XMLName xml.Name    `xml:"feed"`
	Xmlns   string      `xml:"xmlns,attr"` // Atom 命名空间
	Title   string      `xml:"title"`
	Link    []AtomLink  `xml:"link"` // 可能有多个 link（href 为目标链接）
	Updated string      `xml:"updated"`
	Entries []AtomEntry `xml:"entry"`
}

// AtomLink 对应 <link> 节点（Atom 链接）
type AtomLink struct {
	Href string `xml:"href,attr"`
	Rel  string `xml:"rel,attr"` // 类型（如 alternate 为主要链接）
}

// AtomEntry 对应 <entry> 节点（Atom 单篇文章）
type AtomEntry struct {
	Title   string     `xml:"title"`
	Link    []AtomLink `xml:"link"`
	Updated string     `xml:"updated"`
	Summary string     `xml:"summary"` // 摘要（对应 RSS 的 description）
	Author  AtomAuthor `xml:"author"`
	ID      string     `xml:"id"` // 唯一标识
}

type AtomAuthor struct {
	Name string `xml:"name"`
}
