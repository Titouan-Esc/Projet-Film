import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Films from "../pages/Films";
// import HomePage from "../pages/HomePage";

const PublicRoutes = () => (
  <BrowserRouter>
    <Routes>
      {/* <Route path="/" element={<HomePage />} /> */}
      <Route path="/" element={<Films />} />
    </Routes>
  </BrowserRouter>
);

export default PublicRoutes;
