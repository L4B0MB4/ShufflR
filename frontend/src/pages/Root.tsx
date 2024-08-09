import { useEffect } from "react";
import { useAppDispatch, useAppSelector } from "../redux/hook";
import { getTopTracks } from "../redux/me/meSlice";

export const Root = () => {
  const tracks = useAppSelector((s) => s.me.tracks);
  const profile = useAppSelector((s) => s.me.profile!);
  const dispatch = useAppDispatch();
  useEffect(() => {
    if (tracks == null) {
      dispatch(getTopTracks());
    }
  }, [tracks]);

  if (tracks == null) return null;
  return (
    <section className="section">
      <div className="container">
        <div className="columns is-vcentered">
          <div className="column is-narrow  is-hidden-mobile">
            <figure className="image is-128x128">
              <img
                src={profile.images[1].url}
                alt="Profile Picture"
                className="is-rounded"
              />
            </figure>
          </div>
          <div className="column">
            <h1 className="title is-3">{profile.display_name}</h1>
          </div>
        </div>
        <div className="fixed-grid has-4-cols has-2-cols-mobile">
          <div className="grid">
            {tracks.map((track) => {
              return (
                <div key={track.id} className="cell is-one-quarter">
                  <div className="card" style={{ height: "100%" }}>
                    <div className="card-image">
                      <figure className="image is-square">
                        <img
                          src={track.album.images[0].url}
                          alt="Song Cover 1"
                        />
                      </figure>
                    </div>
                    <div className="card-content">
                      <div className="media-content">
                        <p className="title is-6">{track.artists[0].name}</p>
                        <p className="subtitle is-6">{track.name}</p>
                      </div>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </section>
  );
};
