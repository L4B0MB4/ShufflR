import { Profile, TrackItem } from "../models/models";
import { fetchUtil } from "./fetchhelper";

export const meApi = () => {
  return fetchUtil<Profile>("/api/me", "GET", null);
};

export const topTracksApi = () => {
  return fetchUtil<TrackItem[]>("/api/me/top/tracks", "GET", null);
};
