import { useAppSelector } from "../redux/hook";

export const Navbar = () => {
  const me = useAppSelector((s) => s.me.profile);
  return (
    <nav className="grid navbar m-2">
      <div className="cell">
        <div className="navbar-brand">
          <a className="navbar-item" href="/">
            ShufflR
          </a>
        </div>
      </div>
      <div className="cell is-hidden-mobile"></div>
      <div className="cell is-hidden-mobile"></div>
      <div className="cell is-flex  is-justify-content-flex-end">
        <div className="buttons">
          {me != null ? (
            <figure className="image is-48x48">
              <img className="is-rounded" src={me.images[0].url} />
            </figure>
          ) : (
            <a className="button is-primary" href="http://localhost:8080/login">
              Log in
            </a>
          )}
        </div>
      </div>
    </nav>
  );
};
