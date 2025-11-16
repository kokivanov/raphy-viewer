import type { main } from "$lib/wailsjs/go/models";
import { writable, type Writable } from "svelte/store";

export const LocalSettings : Writable<main.LocalSettings> = writable()