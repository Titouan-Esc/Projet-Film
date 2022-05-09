import React from "react";
import { IOneFilm, IFilms } from "../interfaces/App_interface";

export interface FilmProps {
  film: IOneFilm | undefined;
}

const Film = ({ film }: FilmProps) => {
  return (
    <>
      {film?.map((value: IFilms, index: number, array: IOneFilm) => {
        return (
          <main key={value.Id}>
            <h2>{value.Title}</h2>
            <img src={value.Image} alt="Image du film" />
            <p>{value.Description}</p>
          </main>
        );
      })}
    </>
  );
};

export default Film;
