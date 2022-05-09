import React, { useState } from "react";
import Film from "../components/Film";
import NavBar from "../components/NavBar";
import { IOneFilm } from "../interfaces/App_interface";

interface FilmsProps {
  film: IOneFilm | any;
}

const Films = () => {
  const [appFilm, setAppFilm] = useState<FilmsProps>();

  return (
    <main className="films">
      <h1>Tout les films</h1>
      <Film film={appFilm?.film} />
      <NavBar />
    </main>
  );
};

export default Films;
