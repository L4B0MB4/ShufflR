import { createSlice, Dispatch } from "@reduxjs/toolkit";
import { meApi, topTracksApi } from "../../api/meApi";
import { Profile, TrackItem } from "../../models/models";
import { RootState } from "../store";

// Define a type for the slice state
interface MeState {
  profile: Profile | null;
  tracks: TrackItem[] | null;
}

// Define the initial state using that type
const initialState: MeState = {
  profile: null,
  tracks: null,
};

export const counterSlice = createSlice({
  name: "me",
  // `createSlice` will infer the state type from the `initialState` argument
  initialState,
  reducers: {
    setMe: (state, action) => {
      state.profile = action.payload;
    },
    setTopTracks: (state, action) => {
      state.tracks = action.payload;
    },
    /*     incrementByAmount: (state, action: PayloadAction<number>) => {
      state.value += action.payload;
    }, */
  },
});

export const getMeAsync = () => (dispatch: Dispatch) => {
  meApi().then((res) => {
    console.log(res);
    if (res != null) dispatch(setMe(res));
  });
};

export const getTopTracks = () => (dispatch: Dispatch) => {
  topTracksApi().then((res) => {
    console.log(res);
    if (res != null) dispatch(setTopTracks(res));
  });
};

export const { setMe, setTopTracks } = counterSlice.actions;

export const isAuthorized = (state: RootState) => state.me.profile != null;

export default counterSlice.reducer;
