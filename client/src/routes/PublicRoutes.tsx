import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "../pages/HomePage";
import OneFilm from "../pages/OneFilm";

interface PublicRoutesProps {}

const PublicRoutes: React.FunctionComponent<PublicRoutesProps> = (props) => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/one-film" element={<OneFilm />} />
      </Routes>
    </BrowserRouter>
  );
};

export default PublicRoutes;
