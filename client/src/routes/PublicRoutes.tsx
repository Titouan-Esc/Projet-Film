import { BrowserRouter, Routes, Route } from "react-router-dom";
import HomePage from "../pages/HomePage";
import Films from "../pages/Films";

interface PublicRoutesProps {}

const PublicRoutes: React.FunctionComponent<PublicRoutesProps> = (props) => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/films" element={<Films />} />
      </Routes>
    </BrowserRouter>
  );
};

export default PublicRoutes;
