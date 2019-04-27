// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Authors", testAuthors)
	t.Run("Chapters", testChapters)
	t.Run("ChapterQueues", testChapterQueues)
	t.Run("Covers", testCovers)
	t.Run("Genres", testGenres)
	t.Run("Languages", testLanguages)
	t.Run("Novels", testNovels)
	t.Run("NovelQueues", testNovelQueues)
	t.Run("Sources", testSources)
	t.Run("Tags", testTags)
}

func TestDelete(t *testing.T) {
	t.Run("Authors", testAuthorsDelete)
	t.Run("Chapters", testChaptersDelete)
	t.Run("ChapterQueues", testChapterQueuesDelete)
	t.Run("Covers", testCoversDelete)
	t.Run("Genres", testGenresDelete)
	t.Run("Languages", testLanguagesDelete)
	t.Run("Novels", testNovelsDelete)
	t.Run("NovelQueues", testNovelQueuesDelete)
	t.Run("Sources", testSourcesDelete)
	t.Run("Tags", testTagsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Authors", testAuthorsQueryDeleteAll)
	t.Run("Chapters", testChaptersQueryDeleteAll)
	t.Run("ChapterQueues", testChapterQueuesQueryDeleteAll)
	t.Run("Covers", testCoversQueryDeleteAll)
	t.Run("Genres", testGenresQueryDeleteAll)
	t.Run("Languages", testLanguagesQueryDeleteAll)
	t.Run("Novels", testNovelsQueryDeleteAll)
	t.Run("NovelQueues", testNovelQueuesQueryDeleteAll)
	t.Run("Sources", testSourcesQueryDeleteAll)
	t.Run("Tags", testTagsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Authors", testAuthorsSliceDeleteAll)
	t.Run("Chapters", testChaptersSliceDeleteAll)
	t.Run("ChapterQueues", testChapterQueuesSliceDeleteAll)
	t.Run("Covers", testCoversSliceDeleteAll)
	t.Run("Genres", testGenresSliceDeleteAll)
	t.Run("Languages", testLanguagesSliceDeleteAll)
	t.Run("Novels", testNovelsSliceDeleteAll)
	t.Run("NovelQueues", testNovelQueuesSliceDeleteAll)
	t.Run("Sources", testSourcesSliceDeleteAll)
	t.Run("Tags", testTagsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Authors", testAuthorsExists)
	t.Run("Chapters", testChaptersExists)
	t.Run("ChapterQueues", testChapterQueuesExists)
	t.Run("Covers", testCoversExists)
	t.Run("Genres", testGenresExists)
	t.Run("Languages", testLanguagesExists)
	t.Run("Novels", testNovelsExists)
	t.Run("NovelQueues", testNovelQueuesExists)
	t.Run("Sources", testSourcesExists)
	t.Run("Tags", testTagsExists)
}

func TestFind(t *testing.T) {
	t.Run("Authors", testAuthorsFind)
	t.Run("Chapters", testChaptersFind)
	t.Run("ChapterQueues", testChapterQueuesFind)
	t.Run("Covers", testCoversFind)
	t.Run("Genres", testGenresFind)
	t.Run("Languages", testLanguagesFind)
	t.Run("Novels", testNovelsFind)
	t.Run("NovelQueues", testNovelQueuesFind)
	t.Run("Sources", testSourcesFind)
	t.Run("Tags", testTagsFind)
}

func TestBind(t *testing.T) {
	t.Run("Authors", testAuthorsBind)
	t.Run("Chapters", testChaptersBind)
	t.Run("ChapterQueues", testChapterQueuesBind)
	t.Run("Covers", testCoversBind)
	t.Run("Genres", testGenresBind)
	t.Run("Languages", testLanguagesBind)
	t.Run("Novels", testNovelsBind)
	t.Run("NovelQueues", testNovelQueuesBind)
	t.Run("Sources", testSourcesBind)
	t.Run("Tags", testTagsBind)
}

func TestOne(t *testing.T) {
	t.Run("Authors", testAuthorsOne)
	t.Run("Chapters", testChaptersOne)
	t.Run("ChapterQueues", testChapterQueuesOne)
	t.Run("Covers", testCoversOne)
	t.Run("Genres", testGenresOne)
	t.Run("Languages", testLanguagesOne)
	t.Run("Novels", testNovelsOne)
	t.Run("NovelQueues", testNovelQueuesOne)
	t.Run("Sources", testSourcesOne)
	t.Run("Tags", testTagsOne)
}

func TestAll(t *testing.T) {
	t.Run("Authors", testAuthorsAll)
	t.Run("Chapters", testChaptersAll)
	t.Run("ChapterQueues", testChapterQueuesAll)
	t.Run("Covers", testCoversAll)
	t.Run("Genres", testGenresAll)
	t.Run("Languages", testLanguagesAll)
	t.Run("Novels", testNovelsAll)
	t.Run("NovelQueues", testNovelQueuesAll)
	t.Run("Sources", testSourcesAll)
	t.Run("Tags", testTagsAll)
}

func TestCount(t *testing.T) {
	t.Run("Authors", testAuthorsCount)
	t.Run("Chapters", testChaptersCount)
	t.Run("ChapterQueues", testChapterQueuesCount)
	t.Run("Covers", testCoversCount)
	t.Run("Genres", testGenresCount)
	t.Run("Languages", testLanguagesCount)
	t.Run("Novels", testNovelsCount)
	t.Run("NovelQueues", testNovelQueuesCount)
	t.Run("Sources", testSourcesCount)
	t.Run("Tags", testTagsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Authors", testAuthorsHooks)
	t.Run("Chapters", testChaptersHooks)
	t.Run("ChapterQueues", testChapterQueuesHooks)
	t.Run("Covers", testCoversHooks)
	t.Run("Genres", testGenresHooks)
	t.Run("Languages", testLanguagesHooks)
	t.Run("Novels", testNovelsHooks)
	t.Run("NovelQueues", testNovelQueuesHooks)
	t.Run("Sources", testSourcesHooks)
	t.Run("Tags", testTagsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Authors", testAuthorsInsert)
	t.Run("Authors", testAuthorsInsertWhitelist)
	t.Run("Chapters", testChaptersInsert)
	t.Run("Chapters", testChaptersInsertWhitelist)
	t.Run("ChapterQueues", testChapterQueuesInsert)
	t.Run("ChapterQueues", testChapterQueuesInsertWhitelist)
	t.Run("Covers", testCoversInsert)
	t.Run("Covers", testCoversInsertWhitelist)
	t.Run("Genres", testGenresInsert)
	t.Run("Genres", testGenresInsertWhitelist)
	t.Run("Languages", testLanguagesInsert)
	t.Run("Languages", testLanguagesInsertWhitelist)
	t.Run("Novels", testNovelsInsert)
	t.Run("Novels", testNovelsInsertWhitelist)
	t.Run("NovelQueues", testNovelQueuesInsert)
	t.Run("NovelQueues", testNovelQueuesInsertWhitelist)
	t.Run("Sources", testSourcesInsert)
	t.Run("Sources", testSourcesInsertWhitelist)
	t.Run("Tags", testTagsInsert)
	t.Run("Tags", testTagsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("ChapterQueueToChapterUsingChapter", testChapterQueueToOneChapterUsingChapter)
	t.Run("ChapterQueueToNovelUsingNovel", testChapterQueueToOneNovelUsingNovel)
	t.Run("NovelToCoverUsingCover", testNovelToOneCoverUsingCover)
	t.Run("NovelToSourceUsingSource", testNovelToOneSourceUsingSource)
	t.Run("NovelToLanguageUsingLanguage", testNovelToOneLanguageUsingLanguage)
	t.Run("NovelQueueToNovelUsingNovel", testNovelQueueToOneNovelUsingNovel)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("AuthorToNovels", testAuthorToManyNovels)
	t.Run("ChapterToNovels", testChapterToManyNovels)
	t.Run("ChapterToChapterQueues", testChapterToManyChapterQueues)
	t.Run("CoverToNovels", testCoverToManyNovels)
	t.Run("GenreToNovels", testGenreToManyNovels)
	t.Run("LanguageToNovels", testLanguageToManyNovels)
	t.Run("NovelToAuthors", testNovelToManyAuthors)
	t.Run("NovelToChapters", testNovelToManyChapters)
	t.Run("NovelToChapterQueues", testNovelToManyChapterQueues)
	t.Run("NovelToGenres", testNovelToManyGenres)
	t.Run("NovelToNovelQueues", testNovelToManyNovelQueues)
	t.Run("NovelToRecommendedNovelNovels", testNovelToManyRecommendedNovelNovels)
	t.Run("NovelToNovels", testNovelToManyNovels)
	t.Run("NovelToTags", testNovelToManyTags)
	t.Run("SourceToNovels", testSourceToManyNovels)
	t.Run("TagToNovels", testTagToManyNovels)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("ChapterQueueToChapterUsingChapterQueues", testChapterQueueToOneSetOpChapterUsingChapter)
	t.Run("ChapterQueueToNovelUsingChapterQueues", testChapterQueueToOneSetOpNovelUsingNovel)
	t.Run("NovelToCoverUsingNovels", testNovelToOneSetOpCoverUsingCover)
	t.Run("NovelToSourceUsingNovels", testNovelToOneSetOpSourceUsingSource)
	t.Run("NovelToLanguageUsingNovels", testNovelToOneSetOpLanguageUsingLanguage)
	t.Run("NovelQueueToNovelUsingNovelQueues", testNovelQueueToOneSetOpNovelUsingNovel)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("NovelToCoverUsingNovels", testNovelToOneRemoveOpCoverUsingCover)
	t.Run("NovelToLanguageUsingNovels", testNovelToOneRemoveOpLanguageUsingLanguage)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("AuthorToNovels", testAuthorToManyAddOpNovels)
	t.Run("ChapterToNovels", testChapterToManyAddOpNovels)
	t.Run("ChapterToChapterQueues", testChapterToManyAddOpChapterQueues)
	t.Run("CoverToNovels", testCoverToManyAddOpNovels)
	t.Run("GenreToNovels", testGenreToManyAddOpNovels)
	t.Run("LanguageToNovels", testLanguageToManyAddOpNovels)
	t.Run("NovelToAuthors", testNovelToManyAddOpAuthors)
	t.Run("NovelToChapters", testNovelToManyAddOpChapters)
	t.Run("NovelToChapterQueues", testNovelToManyAddOpChapterQueues)
	t.Run("NovelToGenres", testNovelToManyAddOpGenres)
	t.Run("NovelToNovelQueues", testNovelToManyAddOpNovelQueues)
	t.Run("NovelToRecommendedNovelNovels", testNovelToManyAddOpRecommendedNovelNovels)
	t.Run("NovelToNovels", testNovelToManyAddOpNovels)
	t.Run("NovelToTags", testNovelToManyAddOpTags)
	t.Run("SourceToNovels", testSourceToManyAddOpNovels)
	t.Run("TagToNovels", testTagToManyAddOpNovels)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("AuthorToNovels", testAuthorToManySetOpNovels)
	t.Run("ChapterToNovels", testChapterToManySetOpNovels)
	t.Run("CoverToNovels", testCoverToManySetOpNovels)
	t.Run("GenreToNovels", testGenreToManySetOpNovels)
	t.Run("LanguageToNovels", testLanguageToManySetOpNovels)
	t.Run("NovelToAuthors", testNovelToManySetOpAuthors)
	t.Run("NovelToChapters", testNovelToManySetOpChapters)
	t.Run("NovelToGenres", testNovelToManySetOpGenres)
	t.Run("NovelToRecommendedNovelNovels", testNovelToManySetOpRecommendedNovelNovels)
	t.Run("NovelToNovels", testNovelToManySetOpNovels)
	t.Run("NovelToTags", testNovelToManySetOpTags)
	t.Run("TagToNovels", testTagToManySetOpNovels)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("AuthorToNovels", testAuthorToManyRemoveOpNovels)
	t.Run("ChapterToNovels", testChapterToManyRemoveOpNovels)
	t.Run("CoverToNovels", testCoverToManyRemoveOpNovels)
	t.Run("GenreToNovels", testGenreToManyRemoveOpNovels)
	t.Run("LanguageToNovels", testLanguageToManyRemoveOpNovels)
	t.Run("NovelToAuthors", testNovelToManyRemoveOpAuthors)
	t.Run("NovelToChapters", testNovelToManyRemoveOpChapters)
	t.Run("NovelToGenres", testNovelToManyRemoveOpGenres)
	t.Run("NovelToRecommendedNovelNovels", testNovelToManyRemoveOpRecommendedNovelNovels)
	t.Run("NovelToNovels", testNovelToManyRemoveOpNovels)
	t.Run("NovelToTags", testNovelToManyRemoveOpTags)
	t.Run("TagToNovels", testTagToManyRemoveOpNovels)
}

func TestReload(t *testing.T) {
	t.Run("Authors", testAuthorsReload)
	t.Run("Chapters", testChaptersReload)
	t.Run("ChapterQueues", testChapterQueuesReload)
	t.Run("Covers", testCoversReload)
	t.Run("Genres", testGenresReload)
	t.Run("Languages", testLanguagesReload)
	t.Run("Novels", testNovelsReload)
	t.Run("NovelQueues", testNovelQueuesReload)
	t.Run("Sources", testSourcesReload)
	t.Run("Tags", testTagsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Authors", testAuthorsReloadAll)
	t.Run("Chapters", testChaptersReloadAll)
	t.Run("ChapterQueues", testChapterQueuesReloadAll)
	t.Run("Covers", testCoversReloadAll)
	t.Run("Genres", testGenresReloadAll)
	t.Run("Languages", testLanguagesReloadAll)
	t.Run("Novels", testNovelsReloadAll)
	t.Run("NovelQueues", testNovelQueuesReloadAll)
	t.Run("Sources", testSourcesReloadAll)
	t.Run("Tags", testTagsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Authors", testAuthorsSelect)
	t.Run("Chapters", testChaptersSelect)
	t.Run("ChapterQueues", testChapterQueuesSelect)
	t.Run("Covers", testCoversSelect)
	t.Run("Genres", testGenresSelect)
	t.Run("Languages", testLanguagesSelect)
	t.Run("Novels", testNovelsSelect)
	t.Run("NovelQueues", testNovelQueuesSelect)
	t.Run("Sources", testSourcesSelect)
	t.Run("Tags", testTagsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Authors", testAuthorsUpdate)
	t.Run("Chapters", testChaptersUpdate)
	t.Run("ChapterQueues", testChapterQueuesUpdate)
	t.Run("Covers", testCoversUpdate)
	t.Run("Genres", testGenresUpdate)
	t.Run("Languages", testLanguagesUpdate)
	t.Run("Novels", testNovelsUpdate)
	t.Run("NovelQueues", testNovelQueuesUpdate)
	t.Run("Sources", testSourcesUpdate)
	t.Run("Tags", testTagsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Authors", testAuthorsSliceUpdateAll)
	t.Run("Chapters", testChaptersSliceUpdateAll)
	t.Run("ChapterQueues", testChapterQueuesSliceUpdateAll)
	t.Run("Covers", testCoversSliceUpdateAll)
	t.Run("Genres", testGenresSliceUpdateAll)
	t.Run("Languages", testLanguagesSliceUpdateAll)
	t.Run("Novels", testNovelsSliceUpdateAll)
	t.Run("NovelQueues", testNovelQueuesSliceUpdateAll)
	t.Run("Sources", testSourcesSliceUpdateAll)
	t.Run("Tags", testTagsSliceUpdateAll)
}