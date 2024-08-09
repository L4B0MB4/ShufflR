export const Navbar = ({ isLoggedIn }: { isLoggedIn: boolean }) => {
  return (
    <nav className="grid navbar m-2">
      <div className="cell">
        <div className="navbar-brand">
          <a className="navbar-item" href="https://bulma.io">
            ShufflR
          </a>
        </div>
      </div>
      <div className="cell"></div>
      <div className="cell"></div>
      <div className="cell is-flex  is-justify-content-flex-end">
        <div className="buttons">
          {isLoggedIn ? <div /> : <a className="button is-primary">Log in</a>}
        </div>
      </div>
    </nav>
  );
};
