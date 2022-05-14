import { useState, useEffect } from "react";
import Film from "../components/Film";
import NavBar from "../components/NavBar";
import { IOneFilm } from "../interfaces/App_interface";
import axios from "axios";

const Films = () => {
  const [appFilm, setAppFilm] = useState<IOneFilm>();

  const getMovies = async (): Promise<IOneFilm> =>
    new Promise((resolve, reject) => {
      try {
        axios.get(`${process.env.REACT_APP_API_URL}/movies`).then((res) => {
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
    <>
      <main className="films">
        <Film film={appFilm} getMovies={getMovies}/>
      </main>
      <NavBar />
    </>
  );
};

export default Films;
