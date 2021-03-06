import React from "react";
import ReactDOM from "react-dom/client";
import Routing from "./routes/PublicRoutes";
import "./scss/style.scss";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <Routing />
  </React.StrictMode>
);
