import type { SortTypesEnum } from "$lib";
import { QueryAllContent } from "$lib/wailsjs/go/api/ApiService";
import { GetFavoritesJoin } from "$lib/wailsjs/go/data/DataService";

export async function getMedia(sort: SortTypesEnum) {
		const rawRes = await QueryAllContent({
			perPage: 20,
			page: 1,
			sort: sort,
			statusIn: ['FINISHED', 'RELEASING']
		});

		const favs = await GetFavoritesJoin(rawRes.map((v) => v.idMal));

		console.log(favs);

		const favList = (favs ?? []).map((v) => v.ID);

		return rawRes.map((v) => {
			return {
				...v,
				isFav: favList.includes(v.idMal),
				convertValues: (a: any, classs: any, asMap: boolean = false) => {}
			};
		});
	}