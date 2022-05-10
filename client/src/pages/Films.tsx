import React, { useState, useEffect } from "react";
import Film from "../components/Film";
import NavBar from "../components/NavBar";
import { IOneFilm } from "../interfaces/App_interface";
import axios from "axios";

interface FilmsProps {
  film: IOneFilm | any;
}

const Films = () => {
  const [appFilm, setAppFilm] = useState<FilmsProps>();

  const getMovies = async (): Promise<IOneFilm> =>
    new Promise((resolve, reject) => {
      try {
        axios
          .post(
            `${process.env.REACT_APP_API_URL}/movies`,
            JSON.stringify({ page: 1 })
          )
          .then((res) => {
            resolve(res.data);
            setAppFilm(res.data);
          });
      } catch (error) {
        reject(error);
      }
    });

  useEffect(() => {
    getMovies();
  }, []);
  return (
    <main className="films">
      <h1>Tout les films</h1>
      <Film film={appFilm?.film} />
      <NavBar />
    </main>
  );
};

export default Films;
