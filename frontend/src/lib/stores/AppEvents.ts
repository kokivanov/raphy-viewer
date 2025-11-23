import { writable, type Writable } from "svelte/store";
import { GetConnectionState } from "$lib/wailsjs/go/main/App";



export enum AppStateEnum {
    OFFLINE = 0,
    ONLINE = 1
}

export const AppConnection : Writable<AppStateEnum> = writable(AppStateEnum.OFFLINE)

GetConnectionState().then(state => {
    AppConnection.set(Number(state))
})
