package defaultApi

const (
	QUERY_ALL_REQUEST = `
	query ExampleQuery($page: Int = 0, $perPage: Int = 10, $type: MediaType = ANIME, $sort: [MediaSort] = TRENDING_DESC, $statusIn: [MediaStatus], $genreIn: [String] = null, $search: String = null, $tagIn: [String] = null, $isAdult: Boolean = false) {
		Page(page: $page, perPage: $perPage) {
			media(type: $type, sort: $sort, status_in: $statusIn, genre_in: $genreIn, search: $search, tag_in: $tagIn, isAdult: $isAdult) {
			idMal
			title {
				english
				userPreferred
				native
				}
			status
			isAdult
			coverImage {
				large
				extraLarge
				medium
				color
				}
			rankings {
				id
				rank
				type
				format
				year
				season
				allTime
				context
				}
			}
		}
	}`

	QUERY_BY_ID_REQUEST = `query Query($mediaId: Int, $type: MediaType = ANIME, $reviewsPerPage2: Int = 20, $reviewsPage2: Int = 0, $asHtml: Boolean = true, $page: Int = 0, $perPage: Int = 5) {
		Media(idMal: $mediaId, type: $type) {
			idMal
			title {
			romaji
			english
			native
			userPreferred
			}
			siteUrl
			reviews(perPage: $reviewsPerPage2, page: $reviewsPage2) {
			pageInfo {
				total
				hasNextPage
			}
			}
			nextAiringEpisode {
			airingAt
			episode
			}
			isAdult
			studios {
			edges {
				node {
				name
				siteUrl
				isAnimationStudio
				}
			}
			}
			popularity
			averageScore
			genres
			synonyms
			bannerImage
			coverImage {
			extraLarge
			large
			color
			}
			rankings {
			id
			rank
			type
			format
			year
			season
			allTime
			context
			}
			tags {
			name
			isAdult
			}
			trailer {
			site
			thumbnail
			}
			hashtag
			isLicensed
			countryOfOrigin
			episodes
			seasonYear
			season
			startDate {
			year
			month
			day
			}
			format
			status
			description(asHtml: $asHtml)
			trending
			relations {
			edges {
				relationType
				node {
				idMal
				title {
					userPreferred
					english
				}
				coverImage {
					medium
					color
				}
				}
			}
			}
			endDate {
			year
			month
			day
			}
			recommendations(page: $page, perPage: $perPage) {
			pageInfo {
				total
				hasNextPage
			}
			edges {
				node {
				mediaRecommendation {
					idMal
					title {
					english
					userPreferred
					}
					coverImage {
					color
					medium
					}
				}
				}
			}
			}
		}
	}`
)
