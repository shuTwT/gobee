package post

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"math/rand/v2"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/go-ego/gse"
	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/ent/category"
	"github.com/shuTwT/hoshikuzu/ent/post"
	"github.com/shuTwT/hoshikuzu/ent/tag"
	ai_service "github.com/shuTwT/hoshikuzu/internal/services/ai/chat"
	"github.com/shuTwT/hoshikuzu/pkg/cache"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
	"github.com/shuTwT/hoshikuzu/pkg/utils"
)

type PostService interface {
	QueryPostList(c context.Context, req model.PostListReq) ([]*ent.Post, error)
	QueryPostPage(c context.Context, req model.PostPageReq) ([]*ent.Post, int, error)
	QueryPostBySlug(c context.Context, slug string) (*ent.Post, error)
	QueryPostById(c context.Context, id int) (*ent.Post, error)
	CreatePost(c context.Context, title string, content string) (*ent.Post, error)
	UpdatePostContent(c context.Context, id int, content string, htmlContent *string, mdContent *string) (*ent.Post, error)
	UpdatePostSetting(c context.Context, id int, post model.PostUpdateReq) (*ent.Post, error)
	DeletePost(c context.Context, id int) error
	GetPostCount(c context.Context) (int, error)
	QueryPostMeta(c context.Context) ([]*model.PostMeta, error)
	GetPostMonthStats(c context.Context, req model.PostMonthStatsReq) ([]model.PostMonthStat, error)
	GetRandomPost(c context.Context) (*ent.Post, error)
	GetRandomPosts(c context.Context, limit int) ([]*ent.Post, error)
	GetRelatedPosts(c context.Context, id int, limit int) ([]*model.PostRelatedResp, error)
	SearchPosts(c context.Context, req model.PostSearchReq) ([]*model.PostSearchResp, int, error)
	PublishPost(c context.Context, id int) (*ent.Post, error)
	UnpublishPost(c context.Context, id int) (*ent.Post, error)
	PostCountByCategory(c context.Context, categoryID int) (int, error)
}

type PostServiceImpl struct {
	client    *ent.Client
	aiService ai_service.AIService
}

func NewPostServiceImpl(client *ent.Client, aiService ai_service.AIService) *PostServiceImpl {
	return &PostServiceImpl{client: client, aiService: aiService}
}

// postListFields 列表/分页查询仅选取的字段，排除 content、md_content、html_content 等正文大字段，
// 减小数据库传输量与内存占用，从而优化查询速度
var postListFields = []string{
	post.FieldID,
	post.FieldTitle,
	post.FieldSlug,
	post.FieldContentType,
	post.FieldStatus,
	post.FieldIsAutogenSummary,
	post.FieldIsVisible,
	post.FieldIsPinToTop,
	post.FieldIsAllowComment,
	post.FieldIsVisibleAfterComment,
	post.FieldIsVisibleAfterPay,
	post.FieldPrice,
	post.FieldPublishedAt,
	post.FieldViewCount,
	post.FieldCommentCount,
	post.FieldCover,
	post.FieldKeywords,
	post.FieldCopyright,
	post.FieldAuthor,
	post.FieldSummary,
	post.FieldCreatedAt,
}

func (s *PostServiceImpl) QueryPostList(c context.Context, req model.PostListReq) ([]*ent.Post, error) {
	query := s.client.Post.Query()

	if req.CategoryName != "" {
		query.Where(post.HasCategoriesWith(category.Name(req.CategoryName)))
	}

	if req.TagName != "" {
		query.Where(post.HasTagsWith(tag.Name(req.TagName)))
	}

	if req.Year != nil {
		startDate := time.Date(*req.Year, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(*req.Year+1, 1, 1, 0, 0, 0, 0, time.UTC)
		query.Where(post.PublishedAtGTE(startDate), post.PublishedAtLT(endDate))
	}

	if req.Month != nil && req.Year != nil {
		startDate := time.Date(*req.Year, time.Month(*req.Month), 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(*req.Year, time.Month(*req.Month)+1, 1, 0, 0, 0, 0, time.UTC)
		query.Where(post.PublishedAtGTE(startDate), post.PublishedAtLT(endDate))
	}

	if req.IsPinToTop != nil {
		query.Where(post.IsPinToTop(*req.IsPinToTop))
	}

	if req.Status != nil {
		query.Where(post.StatusEQ(post.Status(*req.Status)))
	}

	if req.IsVisible != nil {
		query.Where(post.IsVisible(*req.IsVisible))
	}

	query = query.
		Select(postListFields...).
		WithCategories().
		WithTags().
		Order(ent.Desc(post.FieldID))

	if req.Limit != nil {
		query.Limit(*req.Limit)
	}

	posts, err := query.All(c)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostServiceImpl) QueryPostPage(c context.Context, req model.PostPageReq) ([]*ent.Post, int, error) {
	query := s.client.Post.Query()

	if req.CategoryID != nil {
		query.Where(post.HasCategoriesWith(category.ID(*req.CategoryID)))
	}

	if req.CategoryName != "" {
		query.Where(post.HasCategoriesWith(category.Name(req.CategoryName)))
	}

	if req.TagID != nil {
		query.Where(post.HasTagsWith(tag.ID(*req.TagID)))
	}

	if req.TagName != "" {
		query.Where(post.HasTagsWith(tag.Name(req.TagName)))
	}

	if req.Title != "" {
		query.Where(post.TitleContains(req.Title))
	}

	// 创建时间区间过滤，前端日期范围结束值为所选最后一天的零点，加一天转为左闭右开区间
	if req.StartDate != nil {
		query.Where(post.CreatedAtGTE(req.StartDate.Time()))
	}

	if req.EndDate != nil {
		endDate := req.EndDate.Time().Add(24 * time.Hour)
		query.Where(post.CreatedAtLT(endDate))
	}

	if req.Year != nil {
		startDate := time.Date(*req.Year, 1, 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(*req.Year+1, 1, 1, 0, 0, 0, 0, time.UTC)
		query.Where(post.PublishedAtGTE(startDate), post.PublishedAtLT(endDate))
	}

	if req.Month != nil && req.Year != nil {
		startDate := time.Date(*req.Year, time.Month(*req.Month), 1, 0, 0, 0, 0, time.UTC)
		endDate := time.Date(*req.Year, time.Month(*req.Month)+1, 1, 0, 0, 0, 0, time.UTC)
		query.Where(post.PublishedAtGTE(startDate), post.PublishedAtLT(endDate))
	}

	if req.Status != nil {
		query.Where(post.StatusEQ(post.Status(*req.Status)))
	}

	if req.IsVisible != nil {
		query.Where(post.IsVisible(*req.IsVisible))
	}

	count, err := query.Count(c)
	if err != nil {
		return nil, 0, err
	}
	posts, err := query.
		Select(postListFields...).
		WithCategories().
		WithTags().
		Order(ent.Desc(post.FieldID)).
		Offset((req.Page - 1) * req.Size).
		Limit(req.Size).
		All(c)
	if err != nil {
		return nil, 0, err
	}
	return posts, count, nil
}

func (s *PostServiceImpl) CreatePost(c context.Context, title string, content string) (*ent.Post, error) {

	newPost, err := s.client.Post.Create().
		SetTitle(title).
		SetContent(content).
		SetHTMLContent(content).
		Save(c)
	slug, err := utils.GenerateSlug(title, newPost.CreatedAt.Unix())
	if err != nil {
		return newPost, err
	}
	newPost, err = s.client.Post.UpdateOneID(newPost.ID).
		SetSlug(slug).
		Save(c)
	return newPost, err
}

func (s *PostServiceImpl) QueryPostBySlug(c context.Context, slug string) (*ent.Post, error) {
	post, err := s.client.Post.Query().
		Where(post.Slug(slug)).
		WithCategories().
		WithTags().
		Only(c)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) QueryPostById(c context.Context, id int) (*ent.Post, error) {
	post, err := s.client.Post.Query().
		Where(post.ID(id)).
		WithCategories().
		WithTags().
		Only(c)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) UpdatePostContent(c context.Context, id int, content string, htmlContent *string, mdContent *string) (*ent.Post, error) {
	// html_content 是 HTML 内容的权威存储字段，调用方未显式传入时与 content 保持一致
	if htmlContent == nil {
		htmlContent = &content
	}
	newPost, err := s.client.Post.UpdateOneID(id).
		SetContent(content).
		SetNillableHTMLContent(htmlContent).
		SetNillableMdContent(mdContent).
		Save(c)
	return newPost, err
}

func (s *PostServiceImpl) UpdatePostSetting(c context.Context, id int, updateReq model.PostUpdateReq) (*ent.Post, error) {
	client := s.client
	var summary string
	needAsyncGen := false

	if updateReq.IsAutogenSummary {
		if strings.TrimSpace(updateReq.Summary) == "" {
			summary = "生成中..."
			needAsyncGen = true
		} else {
			summary = updateReq.Summary
		}
	} else {
		summary = updateReq.Summary
	}

	oldPost, err := client.Post.Query().Where(post.ID(id)).First(c)
	if err != nil {
		return nil, err
	}

	var slug string

	generatedSlug, err := utils.GenerateSlug(updateReq.Title, oldPost.CreatedAt.Unix())
	if err != nil {
		return nil, err
	}
	slug = generatedSlug

	newPost, err := client.Post.UpdateOneID(id).
		SetTitle(updateReq.Title).
		SetSlug(slug).
		SetCover(updateReq.Cover).
		SetKeywords(updateReq.Keywords).
		SetCopyright(updateReq.Copyright).
		SetAuthor(updateReq.Author).
		SetIsAutogenSummary(updateReq.IsAutogenSummary).
		SetIsVisible(updateReq.IsVisible).
		SetIsPinToTop(updateReq.IsPinToTop).
		SetIsAllowComment(updateReq.IsAllowComment).
		SetIsVisibleAfterComment(updateReq.IsVisibleAfterComment).
		SetIsVisibleAfterPay(updateReq.IsVisibleAfterPay).
		SetPrice(int(updateReq.Price * 100)).
		SetSummary(summary).
		AddCategoryIDs(updateReq.Categories...).
		AddTagIDs(updateReq.Tags...).
		Save(c)

	// 异步生成 AI 摘要
	if needAsyncGen && err == nil {
		go func(postID int, title string, content string) {
			ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
			defer cancel()

			generated, genErr := s.aiService.GenerateSummary(ctx, title, content)
			if genErr != nil {
				slog.Error("AI 摘要生成失败", "post_id", postID, "error", genErr.Error())
				if _, updateErr := s.client.Post.UpdateOneID(postID).
					SetSummary("生成失败").
					Save(context.Background()); updateErr != nil {
					slog.Error("更新摘要失败状态失败", "post_id", postID, "error", updateErr.Error())
				}
				return
			}

			if _, updateErr := s.client.Post.UpdateOneID(postID).
				SetSummary(generated).
				Save(context.Background()); updateErr != nil {
				slog.Error("保存 AI 生成摘要失败", "post_id", postID, "error", updateErr.Error())
			}
		}(id, updateReq.Title, updateReq.Content)
	}

	return newPost, err
}

func (s *PostServiceImpl) DeletePost(c context.Context, id int) error {
	return s.client.Post.DeleteOneID(id).Exec(c)
}

func (s *PostServiceImpl) GetPostCount(c context.Context) (int, error) {
	count, err := s.client.Post.Query().Count(c)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// QueryPostMeta 查询全部已发布可见文章的轻量元数据。
// 仅选取 slug、title、published_at 三列，避免加载正文等大字段导致响应 JSON 过大。
func (s *PostServiceImpl) QueryPostMeta(c context.Context) ([]*model.PostMeta, error) {
	posts, err := s.client.Post.Query().
		Where(
			post.StatusEQ("published"),
			post.IsVisible(true),
		).
		Select(post.FieldSlug, post.FieldTitle, post.FieldPublishedAt).
		Order(ent.Desc(post.FieldPublishedAt)).
		All(c)
	if err != nil {
		return nil, err
	}

	metas := make([]*model.PostMeta, 0, len(posts))
	for _, p := range posts {
		metas = append(metas, &model.PostMeta{
			Slug:        p.Slug,
			Title:       p.Title,
			PublishedAt: (*model.LocalTime)(p.PublishedAt),
		})
	}
	return metas, nil
}

func (s *PostServiceImpl) GetPostMonthStats(c context.Context, req model.PostMonthStatsReq) ([]model.PostMonthStat, error) {
	posts, err := s.client.Post.Query().Where(
		post.StatusEQ("published"),
		post.IsVisible(true),
	).Order(ent.Desc(post.FieldCreatedAt)).
		All(c)
	if err != nil {
		return nil, err
	}

	monthMap := make(map[string]int)
	for _, p := range posts {

		monthKey := fmt.Sprintf("%d-%02d", p.CreatedAt.Year(), p.CreatedAt.Month())
		monthMap[monthKey]++

	}

	var stats []model.PostMonthStat
	for month, count := range monthMap {
		stats = append(stats, model.PostMonthStat{
			Month: month,
			Count: count,
		})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Month > stats[j].Month
	})

	if req.Limit > 0 && req.Limit < len(stats) {
		stats = stats[:req.Limit]
	}

	return stats, nil
}

func (s *PostServiceImpl) GetRandomPost(c context.Context) (*ent.Post, error) {
	count, err := s.client.Post.Query().Where(
		post.StatusEQ("published"),
		post.IsVisible(true),
	).Count(c)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	offset := rand.IntN(count)

	post, err := s.client.Post.Query().Where(
		post.StatusEQ("published"),
		post.IsVisible(true),
	).WithCategories().
		WithTags().
		Order(ent.Asc(post.FieldID)).
		Offset(offset).
		Limit(1).
		First(c)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// GetRandomPosts 随机获取 N 篇已发布的可见文章
func (s *PostServiceImpl) GetRandomPosts(c context.Context, limit int) ([]*ent.Post, error) {
	if limit < 1 {
		limit = 1
	}
	ids, err := s.client.Post.Query().
		Where(post.StatusEQ("published"), post.IsVisible(true)).
		Order(ent.Asc(post.FieldID)).
		IDs(c)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return []*ent.Post{}, nil
	}
	if len(ids) > limit {
		rand.Shuffle(len(ids), func(i, j int) { ids[i], ids[j] = ids[j], ids[i] })
		ids = ids[:limit]
	}
	return s.client.Post.Query().
		Where(post.IDIn(ids...)).
		Select(postListFields...).
		WithCategories().
		WithTags().
		All(c)
}

// GetRelatedPosts 获取与指定文章相关的 N 篇推荐文章
//
// 评分规则：
//   - 标签匹配（0~100）：当前文章与候选文章标签集合的 Jaccard 相似度 ×100
//   - 标题相似（0~100）：标题按中文分词后 Jaccard 相似度 ×100
//   - 时间新鲜度（0~30）：6 个月半衰期的指数衰减，越新分越高
//   - 同分类加成（0 或 10）：同分类加 10 分
//
// 排序：有标签匹配的文章排在前面，无标签匹配的排在后面，组内均按总分降序，选满 N 篇为止。
func (s *PostServiceImpl) GetRelatedPosts(c context.Context, id int, limit int) ([]*model.PostRelatedResp, error) {
	current, err := s.client.Post.Query().
		Where(post.ID(id)).
		Select(post.FieldID, post.FieldTitle).
		WithCategories().
		WithTags().
		Only(c)
	if err != nil {
		return nil, err
	}

	candidates, err := s.client.Post.Query().
		Where(post.IDNEQ(id), post.StatusEQ("published"), post.IsVisible(true)).
		Select(postListFields...).
		WithCategories().
		WithTags().
		All(c)
	if err != nil {
		return nil, err
	}

	currentTagIDs := postTagIDSet(current)
	currentCategoryIDs := postCategoryIDSet(current)
	currentTitleTokens := titleTokenSet(current.Title)

	scored := make([]scoredCandidate, 0, len(candidates))
	for _, cand := range candidates {
		score, hasTagMatch := calcRelatedScore(cand, currentTagIDs, currentCategoryIDs, currentTitleTokens)
		scored = append(scored, scoredCandidate{
			resp: &model.PostRelatedResp{
				PostResp: toPostResp(cand),
				Score:    score,
			},
			hasTagMatch: hasTagMatch,
		})
	}

	// 有标签匹配的排前面，无标签匹配的排后面，组内均按总分降序
	sortRelatedCandidates(scored)

	if len(scored) > limit {
		scored = scored[:limit]
	}

	results := make([]*model.PostRelatedResp, 0, len(scored))
	for _, s := range scored {
		results = append(results, s.resp)
	}
	return results, nil
}

type scoredCandidate struct {
	resp        *model.PostRelatedResp
	hasTagMatch bool
}

// sortRelatedCandidates 相关文章排序：有标签匹配的排前面，无标签匹配的排后面，组内按总分降序
func sortRelatedCandidates(scored []scoredCandidate) {
	sort.Slice(scored, func(i, j int) bool {
		if scored[i].hasTagMatch != scored[j].hasTagMatch {
			return scored[i].hasTagMatch
		}
		if scored[i].resp.Score != scored[j].resp.Score {
			return scored[i].resp.Score > scored[j].resp.Score
		}
		return scored[i].resp.ID < scored[j].resp.ID
	})
}

// 中文分词器：使用 gse 嵌入词典，词典编译进二进制，运行时无需外部字典文件
var (
	segmenterOnce sync.Once
	segmenter     *gse.Segmenter
)

func getSegmenter() *gse.Segmenter {
	segmenterOnce.Do(func() {
		seg, err := gse.NewEmbed()
		if err != nil {
			slog.Error("加载中文分词词典失败", "error", err.Error())
			return
		}
		segmenter = &seg
	})
	return segmenter
}

// titleTokenSet 将标题按中文分词后转为 token 集合
func titleTokenSet(title string) map[string]struct{} {
	seg := getSegmenter()
	if seg == nil {
		return map[string]struct{}{}
	}
	tokens := seg.Cut(strings.TrimSpace(title), true)
	set := make(map[string]struct{}, len(tokens))
	for _, t := range tokens {
		t = strings.TrimSpace(t)
		if t == "" {
			continue
		}
		set[t] = struct{}{}
	}
	return set
}

// calcRelatedScore 计算候选文章的相关度总分，返回 (总分, 是否有标签匹配)
func calcRelatedScore(cand *ent.Post, currentTagIDs, currentCategoryIDs map[int]struct{}, currentTitleTokens map[string]struct{}) (float64, bool) {
	tagScore := 0.0
	hasTagMatch := false
	if tagSet := postTagIDSet(cand); len(tagSet) > 0 {
		tagScore = jaccard(currentTagIDs, tagSet) * 100
		hasTagMatch = tagScore > 0
	}

	titleScore := jaccard(currentTitleTokens, titleTokenSet(cand.Title)) * 100

	freshness := 0.0
	if cand.PublishedAt != nil {
		freshness = freshnessScore(*cand.PublishedAt, time.Now())
	}

	categoryBonus := 0.0
	if hasIntersection(currentCategoryIDs, postCategoryIDSet(cand)) {
		categoryBonus = 10
	}

	return tagScore + titleScore + freshness + categoryBonus, hasTagMatch
}

// freshnessScore 时间新鲜度：以 6 个月为半衰期的指数衰减，上限 30 分
func freshnessScore(publishedAt, now time.Time) float64 {
	if publishedAt.After(now) {
		publishedAt = now
	}
	ageDays := now.Sub(publishedAt).Hours() / 24
	halfLifeDays := 182.625 // 6 个月
	return 30 * math.Pow(0.5, ageDays/halfLifeDays)
}

// jaccard 计算两个集合的 Jaccard 相似度，空集合返回 0
func jaccard[T comparable](a, b map[T]struct{}) float64 {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	intersection := 0
	for k := range a {
		if _, ok := b[k]; ok {
			intersection++
		}
	}
	union := len(a) + len(b) - intersection
	if union == 0 {
		return 0
	}
	return float64(intersection) / float64(union)
}

func hasIntersection[T comparable](a, b map[T]struct{}) bool {
	for k := range a {
		if _, ok := b[k]; ok {
			return true
		}
	}
	return false
}

func postTagIDSet(p *ent.Post) map[int]struct{} {
	set := make(map[int]struct{}, len(p.Edges.Tags))
	for _, t := range p.Edges.Tags {
		set[t.ID] = struct{}{}
	}
	return set
}

func postCategoryIDSet(p *ent.Post) map[int]struct{} {
	set := make(map[int]struct{}, len(p.Edges.Categories))
	for _, c := range p.Edges.Categories {
		set[c.ID] = struct{}{}
	}
	return set
}

// toPostResp 将 ent.Post 转换为响应模型（列表场景，不包含正文内容）
func toPostResp(p *ent.Post) model.PostResp {
	return model.PostResp{
		ID:                    p.ID,
		Title:                 p.Title,
		Slug:                  p.Slug,
		ContentType:           string(p.ContentType),
		Status:                string(p.Status),
		IsAutogenSummary:      p.IsAutogenSummary,
		IsVisible:             p.IsVisible,
		IsPinToTop:            p.IsPinToTop,
		IsAllowComment:        p.IsAllowComment,
		IsVisibleAfterComment: p.IsVisibleAfterComment,
		IsVisibleAfterPay:     p.IsVisibleAfterPay,
		Price:                 float32(p.Price) / 100,
		PublishedAt:           (*model.LocalTime)(p.PublishedAt),
		ViewCount:             p.ViewCount,
		CommentCount:          p.CommentCount,
		Cover:                 p.Cover,
		Keywords:              p.Keywords,
		Copyright:             p.Copyright,
		Author:                p.Author,
		Summary:               p.Summary,
		CreatedAt:             (model.LocalTime)(p.CreatedAt),
		Categories:            p.Edges.Categories,
		CategoryIds: func() []int {
			ids := make([]int, len(p.Edges.Categories))
			for i, cat := range p.Edges.Categories {
				ids[i] = cat.ID
			}
			return ids
		}(),
		Tags: p.Edges.Tags,
		TagIds: func() []int {
			ids := make([]int, len(p.Edges.Tags))
			for i, tag := range p.Edges.Tags {
				ids[i] = tag.ID
			}
			return ids
		}(),
	}
}

func (s *PostServiceImpl) SearchPosts(c context.Context, req model.PostSearchReq) ([]*model.PostSearchResp, int, error) {
	cacheKey := fmt.Sprintf("post:search:%s:%d:%d", req.Keyword, req.Page, req.Size)

	if cached, found := cache.GetCache().Get(cacheKey); found {
		if result, ok := cached.([]*model.PostSearchResp); ok {
			return result, len(result), nil
		}
	}

	keyword := strings.ToLower(req.Keyword)

	posts, err := s.client.Post.Query().
		Where(
			post.Or(
				post.TitleContains(keyword),
				post.ContentContains(keyword),
				post.SummaryContains(keyword),
				post.KeywordsContains(keyword),
			),
			post.StatusEQ("published"),
			post.IsVisible(true),
		).
		Order(ent.Desc(post.FieldPublishedAt)).
		All(c)

	if err != nil {
		return nil, 0, err
	}

	var results []*model.PostSearchResp

	for _, p := range posts {
		relevance := s.calculateRelevance(p, keyword)
		if relevance > 0 {
			results = append(results, &model.PostSearchResp{
				ID:          p.ID,
				Title:       p.Title,
				Summary:     p.Summary,
				Content:     p.Content,
				Slug:        p.Slug,
				Cover:       p.Cover,
				Author:      p.Author,
				PublishedAt: (*model.LocalTime)(p.PublishedAt),
				ViewCount:   p.ViewCount,
				Relevance:   relevance,
			})
		}
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Relevance == results[j].Relevance {
			if results[i].PublishedAt != nil && results[j].PublishedAt != nil {
				return results[i].PublishedAt.Time().After(results[j].PublishedAt.Time())
			}
			return results[i].ID > results[j].ID
		}
		return results[i].Relevance > results[j].Relevance
	})

	total := len(results)

	start := (req.Page - 1) * req.Size
	end := start + req.Size

	if start >= total {
		return []*model.PostSearchResp{}, total, nil
	}
	if end > total {
		end = total
	}

	pagedResults := results[start:end]

	cache.GetCache().Set(cacheKey, pagedResults, 5*time.Minute)

	return pagedResults, total, nil
}

func (s *PostServiceImpl) calculateRelevance(p *ent.Post, keyword string) float64 {
	var relevance float64 = 0

	title := strings.ToLower(p.Title)
	content := strings.ToLower(p.Content)
	summary := strings.ToLower(p.Summary)
	keywords := strings.ToLower(p.Keywords)

	if strings.Contains(title, keyword) {
		if title == keyword {
			relevance += 10.0
		} else if strings.HasPrefix(title, keyword) {
			relevance += 8.0
		} else {
			relevance += 5.0
		}
	}

	if strings.Contains(summary, keyword) {
		relevance += 3.0
	}

	if strings.Contains(keywords, keyword) {
		relevance += 2.0
	}

	if strings.Contains(content, keyword) {
		count := strings.Count(content, keyword)
		relevance += float64(count) * 0.5
	}

	return relevance
}

func (s *PostServiceImpl) PublishPost(c context.Context, id int) (*ent.Post, error) {
	post, err := s.client.Post.UpdateOneID(id).
		SetStatus("published").
		SetPublishedAt(time.Now()).
		Save(c)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) UnpublishPost(c context.Context, id int) (*ent.Post, error) {
	post, err := s.client.Post.UpdateOneID(id).
		SetStatus("draft").
		Save(c)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) PostCountByCategory(c context.Context, categoryID int) (int, error) {
	count, err := s.client.Post.Query().
		Where(post.HasCategoriesWith(category.IDEQ(categoryID))).
		Count(c)
	if err != nil {
		return 0, err
	}
	return count, nil
}
