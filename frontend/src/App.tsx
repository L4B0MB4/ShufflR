import { useEffect } from "react";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { Navbar } from "./components/Navbar";
import { Root } from "./pages/Root";
import { useAppDispatch, useAppSelector } from "./redux/hook";
import { getMeAsync, isAuthorized } from "./redux/me/meSlice";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <div>not found</div>,
  },
]);

function App() {
  const dispatch = useAppDispatch();
  useEffect(() => {
    dispatch(getMeAsync());
  }, [dispatch]);
  const authenticated = useAppSelector(isAuthorized);
  return (
    <div>
      <Navbar />
      <div className="container">
        {authenticated ? <RouterProvider router={router} /> : <div />}
      </div>
    </div>
  );
}

export default App;
