package post

import (
	"testing"
	"time"

	"github.com/shuTwT/hoshikuzu/ent"
	"github.com/shuTwT/hoshikuzu/pkg/domain/model"
)

func TestJaccard(t *testing.T) {
	tests := []struct {
		name string
		a, b map[string]struct{}
		want float64
	}{
		{"both empty", map[string]struct{}{}, map[string]struct{}{}, 0},
		{"one empty", map[string]struct{}{"a": {}}, map[string]struct{}{}, 0},
		{"identical", map[string]struct{}{"a": {}, "b": {}}, map[string]struct{}{"a": {}, "b": {}}, 1},
		{"no overlap", map[string]struct{}{"a": {}, "b": {}}, map[string]struct{}{"c": {}, "d": {}}, 0},
		{"partial", map[string]struct{}{"a": {}, "b": {}, "c": {}}, map[string]struct{}{"b": {}, "c": {}, "d": {}}, 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jaccard(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("jaccard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreshnessScore(t *testing.T) {
	now := time.Date(2026, 8, 21, 12, 0, 0, 0, time.UTC)
	halfLifeHours := 182.625 * 24

	tests := []struct {
		name  string
		delta time.Duration
		want  float64
	}{
		{"just published", 0, 30},
		{"half life (6 months)", time.Duration(halfLifeHours * float64(time.Hour)), 15},
		{"one year", time.Duration(2 * halfLifeHours * float64(time.Hour)), 7.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := freshnessScore(now.Add(-tt.delta), now)
			if diff := got - tt.want; diff > 0.01 || diff < -0.01 {
				t.Errorf("freshnessScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFreshnessScoreClampsFuture(t *testing.T) {
	now := time.Now()
	if got := freshnessScore(now.Add(24*time.Hour), now); got != 30 {
		t.Errorf("future published time should clamp to 30, got %v", got)
	}
}

// mockPost 构造带标签、分类、标题和发布时间的文章
func mockPost(id int, title string, publishedAt time.Time, tagIDs []int, categoryIDs []int) *ent.Post {
	p := &ent.Post{
		ID:        id,
		Title:     title,
		CreatedAt: publishedAt,
	}
	p.PublishedAt = &publishedAt
	for _, tid := range tagIDs {
		p.Edges.Tags = append(p.Edges.Tags, &ent.Tag{ID: tid})
	}
	for _, cid := range categoryIDs {
		p.Edges.Categories = append(p.Edges.Categories, &ent.Category{ID: cid})
	}
	return p
}

func TestCalcRelatedScore(t *testing.T) {
	now := time.Now()
	current := mockPost(1, "Go语言并发编程实战", now, []int{10, 20}, []int{5})

	tagIDs := postTagIDSet(current)
	catIDs := postCategoryIDSet(current)
	titleTokens := titleTokenSet(current.Title)

	// 完全无关：无共享标签、标题无重叠、不同分类，仍有时间新鲜度分
	unrelated := mockPost(2, "前端开发入门指南", now.Add(-365*24*time.Hour), []int{30}, []int{6})
	score, hasMatch := calcRelatedScore(unrelated, tagIDs, catIDs, titleTokens)
	if hasMatch {
		t.Errorf("unrelated post should have no tag match, got hasMatch=true")
	}
	if score <= 0 {
		t.Errorf("unrelated post should still have freshness score > 0, got %v", score)
	}

	// 同标签，应有标签匹配分，且总分高于完全无关的文章
	sameTag := mockPost(3, "前端开发入门指南", now.Add(-365*24*time.Hour), []int{10, 30}, []int{6})
	score2, hasMatch2 := calcRelatedScore(sameTag, tagIDs, catIDs, titleTokens)
	if !hasMatch2 {
		t.Errorf("same tag post should have tag match, got hasMatch=false")
	}
	if score2 <= score {
		t.Errorf("tag-matched post should outscore unrelated post, got %v vs %v", score2, score)
	}

	// 标签完全匹配 + 同分类，应比仅共享部分标签多 10 分
	best := mockPost(4, "前端开发入门指南", now.Add(-365*24*time.Hour), []int{10}, []int{5})
	score3, _ := calcRelatedScore(best, tagIDs, catIDs, titleTokens)
	if score3 < score2+10 {
		t.Errorf("same category bonus should add 10, got %v vs %v", score3, score2)
	}
}

func TestSortRelatedCandidates(t *testing.T) {
	now := time.Now()
	current := mockPost(1, "Go语言并发编程实战", now, []int{10}, []int{5})
	currentTagIDs := postTagIDSet(current)
	currentCatIDs := postCategoryIDSet(current)
	currentTitleTokens := titleTokenSet(current.Title)

	// 有标签匹配但分数低
	tagMatchLow := mockPost(2, "Go语言内存模型", now.Add(-300*24*time.Hour), []int{10}, []int{5})
	scoreTagMatch, _ := calcRelatedScore(tagMatchLow, currentTagIDs, currentCatIDs, currentTitleTokens)

	// 无标签匹配但分数高（更新鲜）
	freshNoTag := mockPost(3, "旅游摄影心得分享", now, []int{98}, []int{98})
	scoreNoTag, _ := calcRelatedScore(freshNoTag, currentTagIDs, currentCatIDs, currentTitleTokens)

	scored := []scoredCandidate{
		{resp: &model.PostRelatedResp{PostResp: model.PostResp{ID: 3}, Score: scoreNoTag}, hasTagMatch: false},
		{resp: &model.PostRelatedResp{PostResp: model.PostResp{ID: 2}, Score: scoreTagMatch}, hasTagMatch: true},
	}
	sortRelatedCandidates(scored)

	if scored[0].resp.ID != 2 {
		t.Errorf("tag-matched post should rank first regardless of score, got first ID=%d", scored[0].resp.ID)
	}
	if scored[1].resp.ID != 3 {
		t.Errorf("non-tag-matched post should rank last, got second ID=%d", scored[1].resp.ID)
	}
}
